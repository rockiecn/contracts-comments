// SPDX-License-Identifier:UNLICENED
pragma solidity ^0.8.0;

interface IPledgeSetter {
    function pledge(uint64 _i, uint256 money) external;
    function withdraw(uint64 _i, uint8 _ti, uint256 money, uint256 lock) external returns (uint256);

    function addT(uint8 _ti) external;
}

interface IPledgeGetter {
    function balanceOf(uint64 _i, uint8 _to) external view returns (uint256);
}

interface IPledge is IPledgeSetter, IPledgeGetter {}