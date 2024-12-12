// Quorum
//
// this is to generate go binding for smart contracts used in permissioning
//
// Require:
// 1. solc 0.8.14
// 2. abigen (make all from root)

//go:generate ./solc --abi --bin -o . --overwrite ../AbstractClaimsVerifier.sol
//go:generate ./solc --abi --bin -o . --overwrite ../ClaimsVerifier.sol
//go:generate ./solc --abi --bin -o . --overwrite ../ClaimTypes.sol
//go:generate ./solc --abi --bin -o . --overwrite ../CredentialRegistry.sol
//go:generate ./solc --abi --bin -o . --overwrite ../ICredentialRegistry.sol
//go:generate ./solc --abi --bin -o . --overwrite ../Migrations.sol

//go:generate ../../../bin/abigen --crypto gost --gostcurve id-GostR3410-2001-CryptoPro-A-ParamSet -pkg bind -abi  ./AbstractClaimsVerifier.abi -bin  ./AbstractClaimsVerifier.bin            -type AbstractClaimsVerifier   -out ../../bind/AbstractClaimsVerifier.go
//go:generate ../../../bin/abigen --crypto gost --gostcurve id-GostR3410-2001-CryptoPro-A-ParamSet -pkg bind -abi  ./ClaimsVerifier.abi            	 -bin  ./ClaimsVerifier.bin               -type ClaimsVerifier   -out ../../bind/ClaimsVerifier.go
//go:generate ../../../bin/abigen --crypto gost --gostcurve id-GostR3410-2001-CryptoPro-A-ParamSet -pkg bind -abi  ./ClaimTypes.abi               -bin  ./ClaimTypes.bin               -type ClaimTypes   -out ../../bind/ClaimTypes.go
//go:generate ../../../bin/abigen --crypto gost --gostcurve id-GostR3410-2001-CryptoPro-A-ParamSet -pkg bind -abi  ./CredentialRegistry.abi                -bin  ./CredentialRegistry.bin                -type CredentialRegistry    -out ../../bind/CredentialRegistry.go
//go:generate ../../../bin/abigen --crypto gost --gostcurve id-GostR3410-2001-CryptoPro-A-ParamSet -pkg bind -abi  ./ICredentialRegistry.abi -bin  ./ICredentialRegistry.bin -type PermImpl      -out ../../bind/ICredentialRegistry.go
//go:generate ../../../bin/abigen --crypto gost --gostcurve id-GostR3410-2001-CryptoPro-A-ParamSet -pkg bind -abi  ./Migrations.abi      -bin  ./Migrations.bin      -type Migrations -out ../../bind/Migrations.go

package gen
