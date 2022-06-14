// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// for manage erc20 
interface IAccess {
    function set(address _a, bool _set, bytes[] memory signs) external;
    function can(address _a) external returns (bool);
}