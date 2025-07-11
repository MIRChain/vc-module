// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;
pragma experimental ABIEncoderV2;

contract ClaimTypes {

    struct VerifiableCredential {
        address issuer;
        address subject;
        bytes32 data;
        uint256 validFrom;
        uint256 validTo;
    }

    bytes32 constant internal VERIFIABLE_CREDENTIAL_TYPEHASH = keccak256(
        "VerifiableCredential(address issuer,address subject,bytes32 data,uint256 validFrom,uint256 validTo)"
    );

    function hashVerifiableCredential(VerifiableCredential memory vc) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                VERIFIABLE_CREDENTIAL_TYPEHASH,
                vc.issuer,
                vc.subject,
                vc.data,
                vc.validFrom,
                vc.validTo
            )
        );
    }
}