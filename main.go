package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "github.com/pavelkrolevets/vc-module/src/bind"
)

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


	vc :=  map[string]interface{}{
		"@context": "https://www.w3.org/2018/credentials/v1",
		"id": "73bde252-cb3e-44ab-94f9-eba6a8a2f28d",
		"type": "VerifiableCredential",
		"issuer": "did:mir:main:"+issuer.Address.Hex(),
		"issuanceDate": time.Now().String(),
		"expirationDate": time.Now().Add(time.Hour*24),
		"credentialSubject": map[string]interface{} {
			"id": "did:mir:main:"+subject.Address.Hex(),
			"data": "test",
		},
		"proof":"[]",
	}
	jsonData, err := json.Marshal(vc)
	if err != nil {
		log.Printf("could not marshal json: %s\n", err)
		return
	}
	log.Println("json : ", jsonData)
}