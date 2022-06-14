// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IERC20.sol";
import "../interfaces/IAccess.sol";

/**
 *@author MemoLabs
 *@title erc20 token in memo system
 */
contract ERC20 is IERC20 {

    uint16 public version = 2;
    
    address public access;

    uint256 public override totalSupply; // 上限6亿；初始发行3亿
    uint256 public constant initialSupply = 3*10**26;
    uint256 public constant maxSupply = 6*10**26;

    string public override name;
    string public override symbol;

    mapping(address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowances;

    /// @dev created by admin
    constructor(address _access, string memory _name, string memory _symbol) {
        name = _name;
        symbol = _symbol;
        access = _access;
        totalSupply = initialSupply;
        balances[msg.sender] += initialSupply;

        emit Transfer(address(0), msg.sender, initialSupply);
    }

    function balanceOf(address acc) external view virtual override returns (uint256) {
        return balances[acc];
    }

    function allowance(address owner, address spender) external view virtual override returns (uint256) {
        return allowances[owner][spender];
    }

    function transfer(address recipient, uint256 amount) external virtual override {
        _transfer(msg.sender, recipient, amount);
    }

    function approve(address spender, uint256 amount) external virtual override {
        _approve(msg.sender, spender, amount);
    }

    function transferFrom(address sender, address recipient, uint256 amount) external virtual override {
        uint256 ca = allowances[sender][msg.sender];
        require(ca >= amount, "AEB"); // transfer amount exceeds balance
        _transfer(sender, recipient, amount);

        unchecked{
            _approve(sender, msg.sender, ca - amount);
        }
    }


    function _transfer(address sender, address recipient, uint256 amount) internal virtual {
        require(sender != address(0), "IS"); // illegal sender
        require(recipient != address(0), "IR"); // illegal recipient; need?

        uint256 sb = balances[sender];
        require(sb >= amount, "AEB"); // amount exceeds balance

        unchecked {
            balances[sender] = sb - amount;
        }
        balances[recipient] += amount;
        emit Transfer(sender, recipient, amount);
    }

    function _approve(address owner, address spender, uint256 amount) internal virtual {
        require(owner != address(0), "IO"); // illegal owner
        require(spender != address(0), "IS"); // illegal spender; need?
        allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    // use a mint contract if need
    function mint(address target, uint256 mintedAmount) external virtual override {
        require(IAccess(access).can(msg.sender), "CNM"); // can not mint

        totalSupply += mintedAmount;
        require(totalSupply<=maxSupply, "EX"); // exceed the limit

        balances[target] += mintedAmount;
        emit Transfer(address(0), target, mintedAmount);
    }

    // everyone can burn its value; need it or add it in transfer
    function burn(uint256 burnAmount) external virtual override {
        uint256 accountBalance = balances[msg.sender];
        require(accountBalance >= burnAmount, "BAEB"); // burn amount exceeds balance
        unchecked{
            balances[msg.sender] = accountBalance - burnAmount;
        }
        totalSupply -= burnAmount;
        emit Transfer(msg.sender, address(0), burnAmount);
    }
}