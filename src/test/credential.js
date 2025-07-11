const crypto = require("crypto");
const moment = require("moment");
const web3Abi = require("web3-eth-abi");
const web3Utils = require("web3-utils");
const ethUtil = require("ethereumjs-util");

const CredentialRegistry = artifacts.require("CredentialRegistry");
const ClaimsVerifier = artifacts.require("ClaimsVerifier");

const VERIFIABLE_CREDENTIAL_TYPEHASH = web3Utils.soliditySha3(
  "VerifiableCredential(address issuer,address subject,bytes32 data,uint256 validFrom,uint256 validTo)"
);
const EIP712DOMAIN_TYPEHASH = web3Utils.soliditySha3(
  "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
);

const sleep = (seconds) =>
  new Promise((resolve) => setTimeout(resolve, seconds * 1e3));

function sha256(data) {
  const hashFn = crypto.createHash("sha256");
  hashFn.update(data);
  return hashFn.digest("hex");
}

function getCredentialHash(vc, issuer, claimsVerifierContractAddress) {
  const hashDiplomaHex = `0x${sha256(JSON.stringify(vc.credentialSubject))}`;

  const encodeEIP712Domain = web3Abi.encodeParameters(
    ["bytes32", "bytes32", "bytes32", "uint256", "address"],
    [
      EIP712DOMAIN_TYPEHASH,
      web3Utils.sha3("EIP712Domain"),
      web3Utils.sha3("1"),
      648529,
      claimsVerifierContractAddress,
    ]
  );
  const hashEIP712Domain = web3Utils.soliditySha3(encodeEIP712Domain);

  const validFrom = new Date(vc.issuanceDate).getTime();
  const validTo = new Date(vc.expirationDate).getTime();
  const subjectAddress = vc.credentialSubject.id.split(":").slice(-1)[0];
  const encodeHashCredential = web3Abi.encodeParameters(
    ["bytes32", "address", "address", "bytes32", "uint256", "uint256"],
    [
      VERIFIABLE_CREDENTIAL_TYPEHASH,
      issuer.address,
      subjectAddress,
      hashDiplomaHex,
      Math.round(validFrom / 1000),
      Math.round(validTo / 1000),
    ]
  );
  const hashCredential = web3Utils.soliditySha3(encodeHashCredential);

  const encodedCredentialHash = web3Abi.encodeParameters(
    ["bytes32", "bytes32"],
    [hashEIP712Domain, hashCredential.toString(16)]
  );
  return web3Utils.soliditySha3(
    "0x1901".toString(16) + encodedCredentialHash.substring(2, 131)
  );
}

function signCredential(credentialHash, issuer) {
  const rsv = ethUtil.ecsign(
    Buffer.from(credentialHash.substring(2, 67), "hex"),
    Buffer.from(issuer.privateKey, "hex")
  );
  return ethUtil.toRpcSig(rsv.v, rsv.r, rsv.s);
}

contract("Verifiable Credentials", (accounts) => {
  const subject = accounts[1];
  const issuer = {
    address: accounts[0], // eg '0x47adc0faa4f6eb42b499187317949ed99e77ee85'
    // dont forget to update from ganache
    privateKey:
      "89f6fd874041aa38800b29c04ec6097bd5eb76c10cdbaafe746b1ff1c91ea782",
  };
  const signers = [
    {
      address: accounts[2], // eg '0x4a5a6460d00c4d8c2835a3067f53fb42021d5bb9'
      // dont forget to update from ganache
      privateKey:
        "ee5cd6a2cd74e20de5fa04ea156b25b3e55dffae6170873e3e4c94982111e00a",
    },
    {
      address: accounts[3], // eg '0x4222ec932c5a68b80e71f4ddebb069fa02518b8a'
      // dont forget to update from ganache
      privateKey:
        "d0e8b2d78981815f378f0c806850b351855d32596084e0d63774482867e1dd2d",
    },
  ];

  const vc = {
    "@context": "https://www.w3.org/2018/credentials/v1",
    id: "73bde252-cb3e-44ab-94f9-eba6a8a2f28d",
    type: "VerifiableCredential",
    issuer: `did:mir:main:${issuer.address}`,
    issuanceDate: moment().toISOString(),
    expirationDate: moment().add(1, "years").toISOString(),
    credentialSubject: {
      id: `did:mir:main:${subject}`,
      data: "test",
    },
    proof: [],
  };

  before(async () => {
    const instance = await ClaimsVerifier.deployed();
    await instance.grantRole(await instance.ISSUER_ROLE(), issuer.address);
    await instance.grantRole(await instance.SIGNER_ROLE(), signers[0].address);
    await instance.grantRole(await instance.SIGNER_ROLE(), signers[1].address);
  });

  it("should register a VC", async () => {
    const instance = await ClaimsVerifier.deployed();

    const credentialHash = getCredentialHash(vc, issuer, instance.address);
    const signature = await signCredential(credentialHash, issuer);
    const sigObject = await ethUtil.fromRpcSig(signature);
    console.log("Sig: ", signature);
    console.log("Sig: v", sigObject.v);
    console.log("Sig: r", sigObject.r);
    console.log("Sig: s", sigObject.s);
    const pubKeyRecovered = await ethUtil.ecrecover(
      Buffer.from(credentialHash.substring(2, 67), "hex"),
      sigObject.v - 27,
      sigObject.r,
      sigObject.s
    );
    addrRecovered = ethUtil.Address.fromPublicKey(pubKeyRecovered);
    console.log("Recovered address:", addrRecovered.toString());
    console.log("Issuer address:", issuer.address);
    const tx = await instance.registerCredential(
      subject,
      credentialHash,
      Math.round(moment(vc.issuanceDate).valueOf() / 1000),
      Math.round(moment(vc.expirationDate).valueOf() / 1000),
      signature,
      { from: issuer.address }
    );

    vc.proof.push({
      id: vc.issuer,
      type: "EcdsaSecp256k1Signature2019",
      proofPurpose: "assertionMethod",
      verificationMethod: `${vc.issuer}#vm-0`,
      domain: instance.address,
      proofValue: signature,
    });

    await sleep(1);

    return assert.equal(tx.receipt.status, true);
  });

  it("should fail verify additional signers", async () => {
    const instance = await ClaimsVerifier.deployed();

    const data = `0x${sha256(JSON.stringify(vc.credentialSubject))}`;
    const rsv = ethUtil.fromRpcSig(vc.proof[0].proofValue);
    const result = await instance.verifyCredential(
      [
        vc.issuer.replace("did:mir:main:", ""),
        vc.credentialSubject.id.replace("did:mir:main:", ""),
        data,
        Math.round(moment(vc.issuanceDate).valueOf() / 1000),
        Math.round(moment(vc.expirationDate).valueOf() / 1000),
      ],
      rsv.v,
      rsv.r,
      rsv.s
    );

    const additionalSigners = result[3];

    assert.equal(additionalSigners, false);
  });

  it("should register additional signatures to the VC", async () => {
    const instance = await ClaimsVerifier.deployed();

    const credentialHash = getCredentialHash(vc, issuer, instance.address);
    const signature1 = await signCredential(credentialHash, signers[0]);

    const tx1 = await instance.registerSignature(
      credentialHash,
      issuer.address,
      signature1,
      { from: signers[0].address }
    );

    vc.proof.push({
      id: `did:mir:main:${signers[0]}`,
      type: "EcdsaSecp256k1Signature2019",
      proofPurpose: "assertionMethod",
      verificationMethod: `did:mir:main:${signers[0]}#vm-0`,
      domain: instance.address,
      proofValue: signature1,
    });

    assert.equal(tx1.receipt.status, true);

    const signature2 = await signCredential(credentialHash, signers[1]);
    const tx2 = await instance.registerSignature(
      credentialHash,
      issuer.address,
      signature2,
      { from: signers[1].address }
    );

    vc.proof.push({
      id: `did:mir:main:${signers[1]}`,
      type: "EcdsaSecp256k1Signature2019",
      proofPurpose: "assertionMethod",
      verificationMethod: `did:mir:main:${signers[1]}#vm-0`,
      domain: instance.address,
      proofValue: signature2,
    });

    await sleep(1);

    return assert.equal(tx2.receipt.status, true);
  });

  it("should verify a VC", async () => {
    const instance = await ClaimsVerifier.deployed();
    // console.log( vc );

    const data = `0x${sha256(JSON.stringify(vc.credentialSubject))}`;
    const rsv = ethUtil.fromRpcSig(vc.proof[0].proofValue);
    const result = await instance.verifyCredential(
      [
        vc.issuer.replace("did:mir:main:", ""),
        vc.credentialSubject.id.replace("did:mir:main:", ""),
        data,
        Math.round(moment(vc.issuanceDate).valueOf() / 1000),
        Math.round(moment(vc.expirationDate).valueOf() / 1000),
      ],
      rsv.v,
      rsv.r,
      rsv.s
    );

    const credentialExists = result[0];
    const isNotRevoked = result[1];
    const issuerSignatureValid = result[2];
    const additionalSigners = result[3];
    const isNotExpired = result[4];

    assert.equal(credentialExists, true);
    assert.equal(isNotRevoked, true);
    assert.equal(issuerSignatureValid, true);
    assert.equal(additionalSigners, true);
    assert.equal(isNotExpired, true);
  });

  it("should verify additional signatures", async () => {
    const instance = await ClaimsVerifier.deployed();

    const data = `0x${sha256(JSON.stringify(vc.credentialSubject))}`;

    const sign1 = await instance.verifySigner(
      [
        vc.issuer.replace("did:mir:main:", ""),
        vc.credentialSubject.id.replace("did:mir:main:", ""),
        data,
        Math.round(moment(vc.issuanceDate).valueOf() / 1000),
        Math.round(moment(vc.expirationDate).valueOf() / 1000),
      ],
      vc.proof[1].proofValue
    );

    assert.equal(sign1, true);

    const sign2 = await instance.verifySigner(
      [
        vc.issuer.replace("did:mir:main:", ""),
        vc.credentialSubject.id.replace("did:mir:main:", ""),
        data,
        Math.round(moment(vc.issuanceDate).valueOf() / 1000),
        Math.round(moment(vc.expirationDate).valueOf() / 1000),
      ],
      vc.proof[2].proofValue
    );

    assert.equal(sign2, true);
  });

  it("should revoke the credential", async () => {
    const instance = await ClaimsVerifier.deployed();
    const registry = await CredentialRegistry.deployed();

    const credentialHash = getCredentialHash(vc, issuer, instance.address);

    const tx = await registry.revokeCredential(credentialHash);

    assert.equal(tx.receipt.status, true);
  });

  it("should fail the verification process due credential status", async () => {
    const instance = await ClaimsVerifier.deployed();

    const data = `0x${sha256(JSON.stringify(vc.credentialSubject))}`;
    const rsv = ethUtil.fromRpcSig(vc.proof[0].proofValue);
    const result = await instance.verifyCredential(
      [
        vc.issuer.replace("did:mir:main:", ""),
        vc.credentialSubject.id.replace("did:mir:main:", ""),
        data,
        Math.round(moment(vc.issuanceDate).valueOf() / 1000),
        Math.round(moment(vc.expirationDate).valueOf() / 1000),
      ],
      rsv.v,
      rsv.r,
      rsv.s
    );

    const isNotRevoked = result[1];

    assert.equal(isNotRevoked, false);
  });

  it("should verify credential status using the CredentialRegistry", async () => {
    const instance = await ClaimsVerifier.deployed();
    const registry = await CredentialRegistry.deployed();

    const credentialHash = getCredentialHash(vc, issuer, instance.address);

    const result = await registry.status(issuer.address, credentialHash);

    assert.equal(result, false);
  });
});
