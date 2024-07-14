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
	"golang.org/x/crypto/sha3"

	contract "github.com/MIRChain/vc-module/src/bind"
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
	Context           string            `json:"@context"`
	Id                string            `json:"id"`
	Type              string            `json:"type"`
	Issuer            string            `json:"issuer"`
	IssuanceDate      time.Time         `json:"issuanceDate"`
	ExpirationDate    time.Time         `json:"expirationDate"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Proof             []Proof           `json:"proof"`
}

type CredentialSubject struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

type Proof struct {
	Id                 string         `json:"id"`
	Type               string         `json:"type"`
	ProofPurpose       string         `json:"proofPurpose"`
	VerificationMethod string         `json:"verificationMethod"`
	Domain             common.Address `json:"domain"`
	ProofValue         string         `json:"proofValue"`
}

func main() {
	DeployRegistry()
}

func DeployRegistry() {
	// gost3410.GostCurve = gost3410.CurveIdGostR34102001CryptoProAParamSet()
	back, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// Accounts
	jsonBytes, err := ioutil.ReadFile("./keys/keystore/UTC--2024-06-30T09-13-21.083288244Z--014c0bb53c88290976012ea871c95a8d6dd03d0e")
	if err != nil {
		log.Fatal(err)
	}
	subject, err := keystore.DecryptKey(jsonBytes, "12345678")
	if err != nil {
		panic(err)
	}

	jsonBytes, err = ioutil.ReadFile("./keys/keystore/UTC--2024-06-30T09-40-17.290930840Z--95a51c875f2764c46b4424e1fc361d2300d043f4")
	if err != nil {
		log.Fatal(err)
	}
	issuer, err := keystore.DecryptKey(jsonBytes, "12345678")
	if err != nil {
		panic(err)
	}

	jsonBytes, err = ioutil.ReadFile("./keys/keystore/UTC--2024-06-30T09-40-39.641539706Z--a32f367e95b038664e99851e4355947a5d2cdbda")
	if err != nil {
		log.Fatal(err)
	}
	signer_1, err := keystore.DecryptKey(jsonBytes, "12345678")
	if err != nil {
		panic(err)
	}

	jsonBytes, err = ioutil.ReadFile("./keys/keystore/UTC--2024-06-30T09-40-53.521102494Z--5246b17c9c99fcb0a2e8c79e3f151e45ca70b9f9")
	if err != nil {
		log.Fatal(err)
	}
	signer_2, err := keystore.DecryptKey(jsonBytes, "12345678")
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(issuer.PrivateKey, big.NewInt(648529))
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
	log.Println("receipt block num : ", receipt.BlockNumber.String())
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

	vc := Credential{
		Context:        "https://www.w3.org/2018/credentials/v1",
		Id:             "73bde252-cb3e-44ab-94f9-eba6a8a2f28d",
		Type:           "VerifiableCredential",
		Issuer:         `did:mir:main:` + issuer.Address.String(),
		IssuanceDate:   time.Now(),
		ExpirationDate: time.Now().Add(time.Hour * 2),
		CredentialSubject: CredentialSubject{
			Id:   `did:mir:main:` + subject.Address.Hex(),
			Data: "diploma_magna_cum_laude",
		},
		Proof: []Proof{},
	}
	credentialHash, err := getCredentialHash(vc, issuer.Address, claimsVerifierAddress)
	if err != nil {
		panic(err)
	}
	sig, err := crypto.Sign(credentialHash[:], issuer.PrivateKey)
	if err != nil {
		panic(err)
	}

	_, _, v := decodeSignature(sig)

	sig[64] = v.Bytes()[0]

	// register a VC
	tx, err = claimsVerifier.RegisterCredential(auth, subject.Address, credentialHash, big.NewInt(vc.IssuanceDate.Unix()/1000), big.NewInt(vc.ExpirationDate.Unix()/1000), sig)
	if err != nil {
		panic(err)
	}

	receipt, err = bind.WaitMined(ctx, back, tx)
	if err != nil {
		panic(err)
	}
	if receipt.Status != 1 {
		panic("credential is not registered")
	}
	log.Println("RegisterCredential receipt status : ", receipt.Status)

	// read event
	// query := ethereum.FilterQuery{
	// 	FromBlock: big.NewInt(0),
	// 	ToBlock:   big.NewInt(1000),
	// 	Addresses: []common.Address{
	// 		claimsVerifierAddress,
	// 	},
	// }
	// logs, err := back.FilterLogs(context.Background(), query)
	// if err != nil {
	// 	panic(err)
	// }
	// contractAbi, err := abi.JSON(strings.NewReader(contract.CredentialRegistryABI))
	// if err != nil {
	// 	panic(err)
	// }
	// for _, vLog := range logs {
	// 	fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
	// 	fmt.Println(vLog.BlockNumber)     // 2394201
	// 	fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

	// event := make(map[string]interface{})
	// err := contractAbi.UnpackIntoMap(event, "CredentialRegistered", vLog.Data)
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	continue
	// }
	// fmt.Printf("Event %v", event)
	// 	var topics [4]string
	// 	for i := range vLog.Topics {
	// 		topics[i] = vLog.Topics[i].Hex()
	// 		fmt.Printf("Topic %s", vLog.Topics[i].Hex())
	// 	}
	// }

	// Check if registered
	credentialAtContract, err := credentialRegistry.Credentials(&bind.CallOpts{Context: ctx}, credentialHash, issuer.Address)
	if err != nil {
		panic(err)
	}
	log.Printf("Credential at contract %v", credentialAtContract)
	log.Printf("Credential hash %x", credentialHash)

	vc.Proof = append(vc.Proof, Proof{
		Id:                 issuer.Address.Hex(),
		Type:               "EcdsaSecp256k1Signature2019",
		ProofPurpose:       "assertionMethod",
		VerificationMethod: vc.Issuer + `#vm-0`,
		Domain:             claimsVerifierAddress,
		ProofValue:         "0x" + hex.EncodeToString(sig),
	})

	// verify a VC
	json, err := json.Marshal(vc.CredentialSubject)
	if err != nil {
		panic(err)
	}
	data := Sha256(json)

	log.Printf("Proof value %s", vc.Proof[0].ProofValue)

	sigToVerify := hexutil.MustDecode(vc.Proof[0].ProofValue)

	vcToVerify := contract.ClaimTypesVerifiableCredential{
		Issuer:    common.HexToAddress(strings.Split(vc.Issuer, ":")[3]),
		Subject:   common.HexToAddress(strings.Split(vc.CredentialSubject.Id, ":")[3]),
		Data:      [32]byte(data),
		ValidFrom: big.NewInt(vc.IssuanceDate.Unix() / 1000),
		ValidTo:   big.NewInt(vc.IssuanceDate.Unix() / 1000),
	}

	credentialHashAtContract, err := claimsVerifier.CredentialHash(&bind.CallOpts{Context: ctx}, vcToVerify)
	if err != nil {
		panic(err)
	}
	log.Printf("Credential hash at contract %x", credentialHashAtContract)

	// tx, err = claimsVerifier.RegisterSignature(auth, credentialHash, signer_1.Address, sig)
	// if err != nil {
	// 	panic(err)
	// }
	// receipt, err = bind.WaitMined(ctx, back, tx)
	// if err != nil {
	// 	panic(err)
	// }
	// if receipt.Status != 1 {
	// 	panic("signature is not registered")
	// }

	// log.Println("RegisterSignature receipt status : ", receipt.Status)

	verifySigner, err := claimsVerifier.VerifySigner(&bind.CallOpts{Pending: true, Context: ctx}, vcToVerify, sigToVerify)
	if err != nil {
		panic(err)
	}
	log.Printf("Verify signer %t", verifySigner)

	credentialExists, isNotRevoked, issuerSignatureValid, additionalSigners, isNotExpired, err := claimsVerifier.VerifyCredential(&bind.CallOpts{Pending: true, Context: ctx}, vcToVerify, uint8(sigToVerify[64]), [32]byte(sigToVerify[:32]), [32]byte(sigToVerify[32:64]))
	if err != nil {
		panic(err)
	}

	log.Println(credentialExists, isNotRevoked, issuerSignatureValid, additionalSigners, isNotExpired)
}

func getCredentialHash(vc Credential, issuer common.Address, claimsVerifierContractAddress common.Address) ([32]byte, error) {
	json, err := json.Marshal(vc.CredentialSubject)
	if err != nil {
		return [32]byte{}, fmt.Errorf("cant marshall %w", err)
	}
	data := Sha256(json)
	// hashDiplomaHex := hexutil.Encode(jsonHash[:])

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
		big.NewInt(648529),
		claimsVerifierContractAddress,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("cant abi pack argumentsEIP712Domain %w", err)
	}

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
		data,
		big.NewInt(vc.IssuanceDate.Unix()),
		big.NewInt(vc.ExpirationDate.Unix()),
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("cant abi pack argumentsHashCredential %w", err)
	}
	hashCredential := Sha256(encodeHashCredential)

	arguments := abi.Arguments{
		{
			Type: Bytes32,
		},
		{
			Type: Bytes32,
		},
	}

	encodedCredentialHash, err := arguments.Pack(hashEIP712Domain, hashCredential)
	if err != nil {
		return [32]byte{}, fmt.Errorf("cant abi pack arguments %w", err)
	}
	return Sha256(append([]byte("0x1901"), encodedCredentialHash...)), nil

}

func Sha256(d []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(d)
	var res [32]byte
	copy(res[:], hash.Sum(nil))
	return res
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
