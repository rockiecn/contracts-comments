// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Owner {
    //合约主人，即创建合约的人,设为private，只能在当前合约调用
    address private owner;

    event AlterOwner(address from, address to);

    constructor() {
        owner = msg.sender;
        emit AlterOwner(address(0), owner);
    }

    //函数修饰符，判断是不是owner调用
    modifier onlyOwner(){
        require(msg.sender == owner, "N");
        _;
    }

    function alterOwner(address newOwner) public onlyOwner{
        emit AlterOwner(owner, newOwner);
        owner = newOwner;
    }

    function getOwner() public view returns(address){
        return owner;
    }
}