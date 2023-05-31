### Truffle testing

1. Install Truffle 

https://trufflesuite.com/docs/truffle/how-to/install/

2. Install Ganache-cli

```sh
npm install -g ganache-cli
```
3. Run ganache-cli

```sh
ganache-cli --gasLimit=0x1fffffffffffff --allowUnlimitedContractSize -e 1000000000 -p 7545
```

4. Run tests `test`

```sh
npm test
```