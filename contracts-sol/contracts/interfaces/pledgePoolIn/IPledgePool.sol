// SPDX-License-Identifier:UNLICENED
pragma solidity ^0.8.0;

interface IPledgePool {
    function pledge(uint64 index, uint256 money, bytes memory sign) external payable;
    function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes memory sign) external;
    function addToken(address tAddr, uint32 tIndex) external;
    function getPledge(uint32 tIndex) external view returns (uint256);
    function getBalance(uint64 index, uint32 tIndex) external view returns (uint256);
}