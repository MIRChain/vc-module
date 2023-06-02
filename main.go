package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "github.com/pavelkrolevets/vc-module/src/bind"
	"golang.org/x/crypto/sha3"
)

var VERIFIABLE_CREDENTIAL_TYPEHASH = Sha256([]byte("VerifiableCredential(address issuer,address subject,bytes32 data,uint256 validFrom,uint256 validTo)"))
var EIP712DOMAIN_TYPEHASH = Sha256([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))

var (
	Uint256, _    = abi.NewType("uint256", "", nil)
	Uint32, _     = abi.NewType("uint32", "", nil)
	Uint16, _     = abi.NewType("uint16", "", nil)
	String, _     = abi.NewType("string", "", nil)
	Bool, _       = abi.NewType("bool", "", nil)
	Bytes, _      = abi.NewType("bytes", "", nil)
	Bytes32, _    = abi.NewType("bytes32", "", nil)
	Address, _    = abi.NewType("address", "", nil)
	Uint64Arr, _  = abi.NewType("uint64[]", "", nil)
	AddressArr, _ = abi.NewType("address[]", "", nil)
	Int8, _       = abi.NewType("int8", "", nil)
	// Special types for testing
	Uint32Arr2, _       = abi.NewType("uint32[2]", "", nil)
	Uint64Arr2, _       = abi.NewType("uint64[2]", "", nil)
	Uint256Arr, _       = abi.NewType("uint256[]", "", nil)
	Uint256Arr2, _      = abi.NewType("uint256[2]", "", nil)
	Uint256Arr3, _      = abi.NewType("uint256[3]", "", nil)
	Uint256ArrNested, _ = abi.NewType("uint256[2][2]", "", nil)
	Uint8ArrNested, _   = abi.NewType("uint8[][2]", "", nil)
	Uint8SliceNested, _ = abi.NewType("uint8[][]", "", nil)
)


type Credential struct {
	Context string `json:"@context"`
	Id string `json:"id"`
	Type string `json:"type"`
	Issuer string `json:"issuer"`
	IssuanceDate time.Time `json:"issuanceDate"`
	ExpirationDate time.Time `json:"expirationDate"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Proof []Proof `json:"proof"`

}

type CredentialSubject struct {
	Id string `json:"id"`
	Data string `json:"data"`
}

type Proof struct {
	Id string `json:"id"`
	Type string `json:"type"`
	ProofPurpose string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
	Domain common.Address `json:"domain"`
	ProofValue string `json:"proofValue"`

}

func main() {
	DeployRegistry()
}


func DeployRegistry(){
	// gost3410.GostCurve = gost3410.CurveIdGostR34102001CryptoProAParamSet()
	back, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	
	// Accounts
	jsonBytes, err := ioutil.ReadFile("./keys/keystore/UTC--2023-04-03T08-15-20.263797183Z--c6494ec224c2e9e21108005ca4682a81ed4786a2")
    if err != nil {
        log.Fatal(err)
    }
	subject, err := keystore.DecryptKey(jsonBytes, "12345678")
	if err != nil {
		panic(err)
	}

	jsonBytes, err = ioutil.ReadFile("./keys/keystore/UTC--2022-08-17T13-11-15.885898644Z--25face3782641d4ad4c7eaf2ee5eb8a7dcfab465")
    if err != nil {
        log.Fatal(err)
    }
	issuer, err := keystore.DecryptKey(jsonBytes, "Gfdtk81,")
	if err != nil {
		panic(err)
	}

	jsonBytes, err = ioutil.ReadFile("./keys/keystore/UTC--2018-06-24T06-42-15.344831400Z--01665a4eb869efbf3af991e0b791d5347718a49d")
    if err != nil {
        log.Fatal(err)
    }
	signer_1, err := keystore.DecryptKey(jsonBytes, "extreme8811")
	if err != nil {
		panic(err)
	}

	jsonBytes, err = ioutil.ReadFile("./keys/keystore/UTC--2018-07-18T10-22-47.145427470Z--3833067356d624e36fa8cfaf208e97263f3e0703")
    if err != nil {
        log.Fatal(err)
    }
	signer_2, err := keystore.DecryptKey(jsonBytes, "password")
	if err != nil {
		panic(err)
	}


	auth, err := bind.NewKeyedTransactorWithChainID(issuer.PrivateKey, big.NewInt(123456789))
	if err != nil {
		panic(err)
	}
	// Deploying Credential Registry
	registryAddress, tx, _, err := contract.DeployCredentialRegistry(auth, back)
	if err != nil {
		panic(err)
	}

	log.Println("Contract Credential Registry deployed at addr: ", registryAddress.Hex())
	log.Println("Tx hash ", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	log.Println("Contract receipt block num : ", receipt.BlockNumber.String())
	
	credentialRegistry, err := contract.NewCredentialRegistry(registryAddress, back)
	if err != nil {
		panic(err)
	}

	// Deploying Claims Verifier
	claimsVerifierAddress, tx, _, err := contract.DeployClaimsVerifier(auth, back, registryAddress)
	if err != nil {
		panic(err)
	}

	log.Println("Contract Credential Registry deployed at addr: ", claimsVerifierAddress.Hex())
	log.Println("Tx hash ", tx.Hash().Hex())
	
	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	log.Println("Contract receipt block num : ", receipt.BlockNumber.String())
	
	claimsVerifier, err := contract.NewClaimsVerifier(claimsVerifierAddress, back)
	if err != nil {
		panic(err)
	}

	// Assign roles
	issuerRole, err := claimsVerifier.ISSUERROLE(&bind.CallOpts{Pending: true, Context: ctx})
	if err != nil {
		panic(err)
	}
	log.Println("Issuer Role : ", hexutil.Encode(issuerRole[:]))

	signerRole, err := claimsVerifier.SIGNERROLE(&bind.CallOpts{Pending: true, Context: ctx})
	if err != nil {
		panic(err)
	}
	log.Println("Signer Role : ", hexutil.Encode(signerRole[:]))

	// Grant role for ClaimsVerifier
	tx, err = credentialRegistry.GrantRole(auth, issuerRole, claimsVerifierAddress)
	if err != nil {
		panic(err)
	}
	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	tx, err = claimsVerifier.GrantRole(auth, issuerRole, issuer.Address)
	if err != nil {
		panic(err)
	}
	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	log.Println("receipt block num : ", receipt.BlockNumber.String())
	tx, err = claimsVerifier.GrantRole(auth, signerRole, signer_1.Address)
	if err != nil {
		panic(err)
	}
	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	log.Println("receipt block num : ", receipt.BlockNumber.String())
	tx, err = claimsVerifier.GrantRole(auth, signerRole, signer_2.Address)
	if err != nil {
		panic(err)
	}
	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	log.Println("receipt block num : ", receipt.BlockNumber.String())


	vc :=  Credential { 
		Context: "https://www.w3.org/2018/credentials/v1",
		Id: "73bde252-cb3e-44ab-94f9-eba6a8a2f28d",
		Type: "VerifiableCredential",
		Issuer: `did:mir:main:` + issuer.Address.String(),
		IssuanceDate: time.Now(),
		ExpirationDate: time.Now().Add(time.Hour * 2),
		CredentialSubject: CredentialSubject{
			Id: `did:mir:main:` + subject.Address.Hex(),
			Data: "some_data",
		},
		Proof: []Proof{},
	}
	credentialHash, err := getCredentialHash( vc, issuer.Address, claimsVerifierAddress )
	
	sig, err := crypto.Sign(credentialHash, issuer.PrivateKey)
	if err != nil {
		panic(err)
	}

	_, _, v := decodeSignature(sig)

	sig[64] = v.Bytes()[0]

	// register a VC
	tx, err = claimsVerifier.RegisterCredential(auth, subject.Address, [32]byte(credentialHash), big.NewInt(vc.IssuanceDate.Unix()/1000), big.NewInt(vc.ExpirationDate.Unix()/1000), sig)
	if err != nil {
		panic(err)
	}

	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	log.Println("RegisterCredential receipt status : ", receipt.Status)

	vc.Proof = append(vc.Proof, Proof{
		Id: issuer.Address.Hex(),
		Type: "EcdsaSecp256k1Signature2019",
		ProofPurpose: "assertionMethod",
		VerificationMethod: vc.Issuer + `#vm-0`,
		Domain: claimsVerifierAddress,
		ProofValue: "0x"+hex.EncodeToString(sig),
	})

	credentialAtContract, err  := credentialRegistry.Credentials(&bind.CallOpts{Context: ctx}, [32]byte(credentialHash), subject.Address)
	if err != nil {
		panic(err)
	}
	log.Printf("%v", credentialAtContract)

	// verify a VC
	json, err := json.Marshal( vc.CredentialSubject )
	if err != nil {
		panic(err)
	}
	data := Sha256(json)
	log.Println(vc.Proof[0].ProofValue)
	sigToVerify := hexutil.MustDecode(vc.Proof[0].ProofValue)
	vcToVerify:= contract.ClaimTypesVerifiableCredential {
		Issuer: common.HexToAddress(strings.Split(vc.Issuer, ":")[3]),
		Subject: common.HexToAddress(strings.Split(vc.CredentialSubject.Id, ":")[3]),
		Data: [32]byte(data),
		ValidFrom: big.NewInt(vc.IssuanceDate.Unix()/1000),
		ValidTo: big.NewInt(vc.IssuanceDate.Unix()/1000),
	}

	credentialExists, isNotRevoked, issuerSignatureValid, additionalSigners, isNotExpired, err := claimsVerifier.VerifyCredential(&bind.CallOpts{Pending: true, Context: ctx}, vcToVerify, uint8(sigToVerify[64]), [32]byte(sigToVerify[:32]), [32]byte(sigToVerify[32:64]))
	if err != nil {
		panic(err)
	}

	log.Println(credentialExists, isNotRevoked, issuerSignatureValid, additionalSigners, isNotExpired)
}

func getCredentialHash(vc Credential, issuer common.Address, claimsVerifierContractAddress common.Address) ([]byte, error) {
	json, err := json.Marshal( vc.CredentialSubject )
	if err != nil {
		panic(err)
	}
	hashDiplomaHex := hexutil.Encode(Sha256(json))
	
	argumentsEIP712Domain := abi.Arguments{
        {
            Type: Bytes32,
        },
        {
            Type: Bytes32,
        },
        {
            Type: Bytes32,
        },
		{
            Type: Uint256,
        },
		{
            Type: Address,
        },
    }

	encodeEIP712Domain, err := argumentsEIP712Domain.Pack(
        EIP712DOMAIN_TYPEHASH,
        Sha256([]byte("EIP712Domain")),
        Sha256([]byte("1")),
		648529,
		claimsVerifierContractAddress,
    )

	hashEIP712Domain := Sha256(encodeEIP712Domain)

	argumentsHashCredential := abi.Arguments{
        {
            Type: Bytes32,
        },
        {
            Type: Address,
        },
        {
            Type: Address,
        },
		{
            Type: Bytes32,
        },
		{
            Type: Uint256,
        },
		{
            Type: Uint256,
        },
    }

	subjectAddress := strings.Split(vc.CredentialSubject.Id, ":")[3]

	encodeHashCredential, err := argumentsHashCredential.Pack(
		VERIFIABLE_CREDENTIAL_TYPEHASH,
		issuer,
		common.HexToAddress(subjectAddress),
		hashDiplomaHex,
		vc.IssuanceDate.Unix(),
		vc.ExpirationDate.Unix(),
	)
	hashCredential := Sha256(encodeHashCredential)


	arguments :=  abi.Arguments{
        {
            Type: Bytes32,
        },
		{
            Type: Bytes32,
        },
    }
	
	encodedCredentialHash, err := arguments.Pack(hashEIP712Domain, hashCredential)

	return Sha256(append([]byte("0x1901"), encodedCredentialHash...)), err

}

func Sha256(d []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(d)
	return hash.Sum(nil)
}

func decodeSignature(sig []byte) (r, s, v *big.Int) {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v
}