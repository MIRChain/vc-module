### Truffle testing

1. Install Truffle

https://trufflesuite.com/docs/truffle/how-to/install/

2. Install Ganache-cli

```sh
npm install -g ganache-cli
```

3. Run ganache-cli

```sh
ganache-cli --gasLimit=0x1fffffffffffff --allowUnlimitedContractSize -e 1000000000 -p 7545 --account="0x89f6fd874041aa38800b29c04ec6097bd5eb76c10cdbaafe746b1ff1c91ea782,1000000000000000000000" --account="0x2eb859d6120dbc59638cacac709af7fcb7e60311e6eaa90ad447e949bff54ee4,1000000000000000000000" --account="0xee5cd6a2cd74e20de5fa04ea156b25b3e55dffae6170873e3e4c94982111e00a,1000000000000000000000" --account="0xd0e8b2d78981815f378f0c806850b351855d32596084e0d63774482867e1dd2d,1000000000000000000000"
```

4. Run tests `test`

```sh
npm test
```
