// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct IssuIn {
    uint64 start;
    uint64 end;
    uint64 size;
    uint256 sPrice;
}

interface IIssuance {
    function issu(IssuIn memory ps) external returns (uint256);
}