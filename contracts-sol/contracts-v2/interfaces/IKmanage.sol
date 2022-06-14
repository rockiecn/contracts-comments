// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IPool.sol";

interface IKmanageSetter {
    function addKeeper(uint64 _ki) external;
    function addCnt(uint64 _ki, uint64 _cnt) external;
    function addProfit(uint8 _ti, uint256 _money) external;
    function addSP(uint8 _ti, uint64 _size, uint256 _price, bool isAdded) external;
    function withdraw(uint64 _ki, uint8 _ti, uint256 money) external returns (uint256);
}

interface IKmanageGetter {
    function getRate() external view returns (uint8);
    function getSP(uint8 _ti) external view returns (uint64, uint256); 
    function balanceOf(uint64 _ki, uint8 _ti) external view returns(uint256, uint256);
}

interface IKmanage is IKmanageSetter, IKmanageGetter {}