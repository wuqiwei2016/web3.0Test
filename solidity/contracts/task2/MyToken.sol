pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyToken is IERC20,Ownable{
    string private _name = "MyToken";
    string private _symbol="MTK";
    string private _decimals = "18";   
    uint256 private _totalSupply;

    //账户余额
    mapping(address => uint256) private _balances;
    //授权映射
    mapping(address => mapping(address => uint256)) private _allowances;

    constructor(uint256 initialSupply) Ownable(msg.sender){
           _mint(msg.sender, initialSupply * 10**18);  // 乘以 10^18 处理 decimals
    }

    // ERC20标准
    // 币的总额
    function totalSupply() public view override returns(uint256){
        return _totalSupply;
    }

    // 账户余额
    function balanceOf(address account) public view override returns(uint256){
        return _balances[account];
    }

    // 转账
    function transfer(address to,uint256 amount) public override returns(bool){
        _transfer(msg.sender,to,amount);
        return true;
    }
    // 银行卡授权支付 
    function approve(address spender,uint256 amount) public override returns(bool){
        _approve(msg.sender,spender,amount);
        return true;
    }

    // 代付或者自动扣款
    function transferForm(address from,address to,uint256 amount) public override returns(uint256){
        _spendAllowance(from, msg.sender, amount);
        _transfer(from, to, amount);
        return true;
    }

    // 查询授权支付的剩余限额
     function allowance(address owner, address spender) public view override returns (uint256) {
        return _allowances[owner][spender];
    }

    // --- 内部方法 ---
    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "ERC20: approve from zero address");
        require(spender != address(0), "ERC20: approve to zero address");
        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    function _spendAllowance(address owner, address spender, uint256 amount) internal {
        uint256 currentAllowance = allowance(owner, spender);
        require(currentAllowance >= amount, "ERC20: insufficient allowance");
        unchecked {
            _allowances[owner][spender] = currentAllowance - amount;
        }
    }

    function _transfer(address from, address to, uint256 amount) internal {
        require(from != address(0), "ERC20: transfer from zero address");
        require(to != address(0), "ERC20: transfer to zero address");
        require(_balances[from] >= amount, "ERC20: transfer amount exceeds balance");
        unchecked {
            _balances[from] -= amount;
            _balances[to] += amount;
        }
        emit Transfer(from, to, amount);
    }

    // 铸币（仅限所有者调用）
    function mint(address to, uint256 amount) public onlyOwner {
        require(to != address(0), "ERC20: mint to zero address");
        _mint(to, amount);
    }

    function _mint(address to,uint256 amount) internal {
        _totalSupply += amount;
        unchecked {
            _balances[to] += amount;
        }
        emit Transfer(address(0), to, amount);
    }
    
}