// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface ITokenSetter {
    function addT(address _t) external returns(uint8);
    function banT(address _t) external; 
}

interface ITokenGetter {
    function getTA(uint8 _ti) external view returns (address, bool);
    function getTI(address _t) external view returns (uint8, bool);
}

interface IToken is ITokenSetter, ITokenGetter {}