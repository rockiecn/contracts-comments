// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// for manage other contracts
interface IAuth {
    // change owner, control etc...
    function perm(bytes32 _h, bytes[] memory signs) external;
}