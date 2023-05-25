pragma solidity >=0.8.0;

contract EthereumDIDRegistry {
  function validDelegate(address identity, bytes32 delegateType, address delegate) public view returns(bool);
}