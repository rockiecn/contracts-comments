// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IRToken {
    function addT(address t) external returns (uint32 index);
    function isValid(uint32 tIndex) external view returns (bool);
    function getTA(uint32 tIndex) external view returns (address);
    function getTI(address t) external view returns (uint32, bool);
}