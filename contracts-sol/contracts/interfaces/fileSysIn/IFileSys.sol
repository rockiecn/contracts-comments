// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// struct for addOrder params
struct AOParams {
        uint64 uIndex;
        uint64 pIndex;
        uint64 _start;
        uint64 _end;
        uint64 _size;
        uint64 nonce;
        uint32 tIndex;
        uint256 sPrice;
        bytes usign;
        bytes psign;
    }

// struct for proWithdraw params
struct PWParams {
        uint64 pIndex;
        uint32 tIndex;
        uint256 pay;
        uint256 lost;
    }

interface IFileSys {
    function createFs(uint64 uIndex) external;
    function addKeeper(uint64 kIndex) external;
    function addOrder(AOParams memory ps) external;
    //function subOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external;
    function subOrder(AOParams memory ps) external;
    function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) external;
    //function proWithdraw(uint64 pIndex, uint32 tokenIndex, address pAddr, address tAddr, uint256 pay, uint256 lost) external;
    function proWithdraw(PWParams memory ps) external;
    function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) external;
    function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external;
    function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external;
    // =====get=====
    function getFsInfo(uint64 user) external view returns (bool);
    function getFsPSum(uint64 user) external view returns (uint64);
    function getFsPro(uint64 user, uint64 i) external view returns (uint64);
    function getFsInfoAggOrder(uint64 user, uint64 provider) external view returns (uint64, uint64);
    function getStoreInfo(uint64 user, uint64 provider, uint32 token) external view returns (uint64,uint64, uint64, uint256);
    function getChannelInfo(uint64 user, uint64 provider, uint32 token) external view returns (uint256, uint64, uint64);
    function getSettleInfo(uint64 index, uint32 tIndex) external view returns (uint64, uint64, uint256, uint256, uint256,uint256,uint256,uint256,uint256,uint256,uint256);
    function getBalance(uint64 index, uint32 tIndex) external view returns (uint256, uint256);
}