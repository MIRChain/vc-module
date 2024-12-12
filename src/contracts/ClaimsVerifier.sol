// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

pragma experimental ABIEncoderV2;

import "./lib/ECDSA.sol";
import "./openzeppelin-contracts/contracts/access/AccessControlEnumerable.sol";
import "./AbstractClaimsVerifier.sol";
import "./ClaimTypes.sol";

contract ClaimsVerifier is AbstractClaimsVerifier, ClaimTypes, AccessControlEnumerable {

    using ECDSA for bytes32;

    bytes32 public constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 public constant SIGNER_ROLE = keccak256("SIGNER_ROLE");

    constructor (address _registryAddress)
    AbstractClaimsVerifier(
        "EIP712Domain",
        "1",
        648529,
        address(this),
        _registryAddress
    ) public {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function verifyCredential(VerifiableCredential memory vc, uint8 v, bytes32 r, bytes32 s) public view returns (bool, bool, bool, bool, bool) {
        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                hashVerifiableCredential(vc)
            )
        );
        return (_exist(digest, vc.issuer), _verifyRevoked(digest, vc.issuer), _verifyIssuer(digest, vc.issuer, v, r, s), (_verifySigners(digest, vc.issuer) == getRoleMemberCount(keccak256("SIGNER_ROLE"))), _validPeriod(vc.validFrom, vc.validTo));
    }

    function verifySigner(VerifiableCredential memory vc, bytes calldata _signature) public view returns (bool){
        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                hashVerifiableCredential(vc)
            )
        );

        address signer = digest.recover(_signature);
        return hasRole(SIGNER_ROLE, signer) && _isSigner(digest, vc.issuer, _signature);
    }

    function registerCredential(address _subject, bytes32 _credentialHash, uint256 _from, uint256 _exp, bytes calldata _signature) public onlyIssuer returns (bool) {
        address signer = _credentialHash.recover(_signature);
        require(msg.sender == signer, "Sender hasn't signed the credential");
        return _registerCredential(msg.sender, _subject, _credentialHash, _from, _exp, _signature);
    }

    function registerSignature(bytes32 _credentialHash, address issuer, bytes calldata _signature) public onlySigner returns (bool){
        address signer = _credentialHash.recover(_signature);
        require(msg.sender == signer, "Sender hasn't signed the credential");
        return _registerSignature(_credentialHash, issuer, _signature);
    }

    function credentialHash(VerifiableCredential memory vc) public view returns (bytes32) {
        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                hashVerifiableCredential(vc)
            )
        );
        return digest;
    }

    modifier onlyAdmin(){
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "Caller is not Admin");
        _;
    }

    modifier onlySigner() {
        require(hasRole(SIGNER_ROLE, msg.sender), "Caller is not a signer");
        _;
    }

    modifier onlyIssuer() {
        string memory senderAddress = _bytes20ToLiteralString(
            bytes20(msg.sender)
        );
        string memory issuerRole = _bytes32ToLiteralString(
            bytes32(ISSUER_ROLE)
        );
        require(hasRole(ISSUER_ROLE, msg.sender), "Caller is not an issuer");
        _;
    }

    function _bytes20ToLiteralString(bytes20 data) 
        private
        pure
        returns (string memory result)
    {
        bytes memory temp = new bytes(41);
        uint256 count;

        for (uint256 i = 0; i < 20; i++) {
            bytes1 currentByte = bytes1(data << (i * 8));
            
            uint8 c1 = uint8(
                bytes1((currentByte << 4) >> 4)
            );
            
            uint8 c2 = uint8(
                bytes1((currentByte >> 4))
            );
        
            if (c2 >= 0 && c2 <= 9) temp[++count] = bytes1(c2 + 48);
            else temp[++count] = bytes1(c2 + 87);
            
            if (c1 >= 0 && c1 <= 9) temp[++count] = bytes1(c1 + 48);
            else temp[++count] = bytes1(c1 + 87);
        }
        
        result = string(temp);
    }
    function _bytes32ToLiteralString(bytes32 data) 
        private
        pure
        returns (string memory result)
    {
        bytes memory temp = new bytes(65);
        uint256 count;

        for (uint256 i = 0; i < 32; i++) {
            bytes1 currentByte = bytes1(data << (i * 8));
            
            uint8 c1 = uint8(
                bytes1((currentByte << 4) >> 4)
            );
            
            uint8 c2 = uint8(
                bytes1((currentByte >> 4))
            );
        
            if (c2 >= 0 && c2 <= 9) temp[++count] = bytes1(c2 + 48);
            else temp[++count] = bytes1(c2 + 87);
            
            if (c1 >= 0 && c1 <= 9) temp[++count] = bytes1(c1 + 48);
            else temp[++count] = bytes1(c1 + 87);
        }
        
        result = string(temp);
    }
}