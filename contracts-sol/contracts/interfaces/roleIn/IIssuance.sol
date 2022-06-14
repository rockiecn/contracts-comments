// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct IssuParams {
        uint64 _start;
        uint64 _end;
        uint64 _size;
        uint256 sPrice;
    }

interface IIssuance {
    function issu(IssuParams memory ps) external returns (uint256);
    function setTP(uint256 _add, uint256 _sub) external;
}