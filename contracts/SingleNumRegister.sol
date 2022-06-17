//Contract based on [https://docs.openzeppelin.com/contracts/3.x/erc721](https://docs.openzeppelin.com/contracts/3.x/erc721)
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SingleNumRegister {
  uint storedData;
  function set(uint x) public {
    storedData = x;
  }
  function get() public view returns (uint retVal){
    return storedData;
  }
}
