// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IFileSys.sol";

interface IProxy {
    function activate(uint64 _i, bool _active, bytes[] memory signs) external;
    function ban(uint64 _i, bool _ban, bytes[] memory signs) external;
    function addT(address _t, bool _ban, bytes[] memory signs) external;
    function banG(uint64 _gi, bool _ban, bytes[] memory signs) external;

    function createGroup(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external;
    // register self to get index
    function registerAccount() external; 
    function registerRole(uint8 _rtype, bytes memory _extra) external;
    // add a user/keeper/provider to group
    function addToGroup(uint64 _gi) external;
    
    function pledge(uint64 _i, uint256 _money) external;
    function unpledge(uint64 _i, uint8 _ti, uint256 _money) external;

    function addOrder( OrderIn memory _oi) external;
    function subOrder(OrderIn memory _oi) external;

    function recharge(uint64 _i, uint8 _ti, uint256 _money, bool isLock) external;
    function withdraw(uint64 _i, uint8 _ti, uint256 _money) external;
    function proWithdraw(PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;

    function get(uint8 _type) external view returns(address);
}

   