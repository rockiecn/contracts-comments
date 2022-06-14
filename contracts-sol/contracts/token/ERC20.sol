// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;
import "../interfaces/erc20In/IERC20.sol";
import "./AccessControl.sol";
import "../Recover.sol";

/// @dev This contract is about memo token.
contract ERC20 is IERC20, AccessControl {
    mapping(address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowances;

    uint256 private totalSupply; // 上限6亿；初始发行3亿
    uint256 public constant maxSupply = 6*10**26;

    string private name;
    string private symbol;

    uint16 public version;

    using Recover for bytes32;

    /// @dev created by admin
    constructor(string memory _name, string memory _symbol, uint16 _version, address[5] memory _addrs) {
        name = _name;
        symbol = _symbol;
        addrs = _addrs;
        version = _version;
        // 初始发行3亿
        // need public?
        uint256 initialSupply = 300000000000000000000000000;
        totalSupply = initialSupply;
        balances[msg.sender] += initialSupply;
        emit Transfer(address(0), msg.sender, initialSupply);
    }

    // 添加virtual表示函数可以被覆盖
    function getName() public view virtual override returns (string memory) {
        return name;
    }

    function getSymbol() public view virtual override returns (string memory) {
        return symbol;
    }

    // 表示代币最小单位与token之间的换算，比如 1 ether = 1*10^18 wei. 这里我们也规定为18
    // need ?
    function getDecimals() public view virtual override returns (uint8) {
        return 18;
    }

    function getTotalSupply() public view override virtual returns (uint256) {
        return totalSupply;
    }

    function balanceOf(address acc) public view virtual override returns (uint256) {
        return balances[acc];
    }

    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return allowances[owner][spender];
    }

    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(address sender, address recipient, uint256 amount) public virtual override returns (bool) {
        uint256 currentAllowance = allowances[sender][msg.sender];
        require(currentAllowance >= amount, "AEB"); // transfer amount exceeds balance
        _transfer(sender, recipient, amount);

        unchecked{
            _approve(sender, msg.sender, currentAllowance - amount);
        }
        return true;
    }

    // increaseAllowance and decreaseAllowance need? can remove it
    function increaseAllowance(address spender, uint256 addedValue) public virtual override returns (bool) {
        _approve(msg.sender, spender, allowances[msg.sender][spender]+addedValue);
        return true;
    }

    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual override returns (bool) {
        uint256 currentAllowance = allowances[msg.sender][spender];
        require(currentAllowance >= subtractedValue, "ABZ"); // decreased allowance below zero
        unchecked {
            _approve(msg.sender, spender, currentAllowance - subtractedValue);
        }
        return true;
    }

    function _transfer(address sender, address recipient, uint256 amount) internal virtual {
        require(!getPaused(), "TF"); // transfer forbidden
        require(sender != address(0), "IS"); // illegal sender
        require(recipient != address(0), "IR"); // illegal recipient; need?

        uint256 senderBalance = balances[sender];
        require(senderBalance >= amount, "AEB"); // transfer amount exceeds balance

        unchecked {
            balances[sender] = senderBalance - amount;
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
    function mint(address target, uint256 mintedAmount) external virtual override returns(bool) {
        require(!getPaused(), "TF");
        require(hasRole(MINTER_ROLE, c), "CNM"); // can not mint

        uint size;
        assembly {
            size := extcodesize(sender)
        }
        require(size != 0; "nca") // not contract address

        totalSupply += mintedAmount;
        require(totalSupply<=maxSupply, "EX"); // exceed the limit

        balances[target] += mintedAmount;
        emit Transfer(address(0), target, mintedAmount);

        return true;
    }

    // everyone can burn its value; need it or add it in transfer
    function burn(uint256 burnAmount) external virtual override returns (bool) {
        uint256 accountBalance = balances[msg.sender];
        require(accountBalance >= burnAmount, "BAEB"); // burn amount exceeds balance
        unchecked{
            balances[msg.sender] = accountBalance - burnAmount;
        }
        totalSupply -= burnAmount;
        emit Transfer(msg.sender, address(0), burnAmount);
        return true;
    }
}