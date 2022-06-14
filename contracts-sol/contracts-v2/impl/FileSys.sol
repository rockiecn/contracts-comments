// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IFileSys.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title manage fs in memo system
 */
contract FileSys is IFileSys, Owner {

    struct StoreInfo {
        uint64 start;  // last start 
        uint64 end;    // 什么时刻的状态，last end time
        uint64 size;   // 在该存储节点上的存储总量，byte
        uint256 sprice; 
    }

    struct AggOrder {
        uint64 nonce;    // 防止order重复提交
        uint64 subNonce; // 用于订单到期
        mapping(uint8 => StoreInfo) sInfo; // 不同代币的支付信息，tokenIndex => StoreInfo
    }

    struct FsInfo {
        uint64[] providers;             // provider索引
        mapping(uint64 => AggOrder) ao; // 该User对每个Provider的订单信息
    }

    // Settlement indicates billing information
    struct Settlement {
        uint64 time; // store状态改变或支付的时间, 与 epoch 对齐
        uint64 size; // 在该存储节点上的存储总量
        uint256 sprice; // 累积的sprice(即sizePrice)

        uint256 maxPay;  // 对此provider所有user聚合总额度； expected 加和
        uint256 hasPaid; // 已经支付
        uint256 canPay;  // 最近一次store/pay时刻，可以支付的金额
        uint256 lost;    // lost due to unable response to challenge
        uint256 lostPaid;// pay to repair

        uint256 managePay; // pay for group keepers >= endPaid+linearPaid
        uint256 endPaid;   // release 25% of manage when order expire; need?
        uint256 linearPaid;// release 75% of manage when pay for provider; need?
    }

    uint8 public constant taxRate = 1;

    uint16 public version = 2;

    mapping(uint64 => mapping(uint8 => uint256)) balances; // 账户可用的余额
    mapping(uint64 => mapping(uint8 => uint256)) lock;     // cannot withdraw 

    mapping(uint64 => FsInfo) fs; // user => FsInfo; user 0 is repair fs

    mapping(uint64 => mapping(uint8 => Settlement)) proInfo; // pro => token => income

    constructor(address _ctl, address _a) Owner(_ctl, _a) {
    }

    function _settlementAdd(uint64 _pIndex, uint8 _tokenIndex, uint64 start, uint64 size, uint256 sprice, uint256 pay, uint256 manage) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.time < start){ 
            if(se.time!=0){ // 非首次addOrder
                proInfo[_pIndex][_tokenIndex].canPay += (start-se.time) * se.sprice;
            }
            proInfo[_pIndex][_tokenIndex].time = start;
        }else if(se.time > start){
            proInfo[_pIndex][_tokenIndex].canPay += uint256(se.time - start)*sprice;
        }

        proInfo[_pIndex][_tokenIndex].sprice += sprice;
        proInfo[_pIndex][_tokenIndex].size += size;
        proInfo[_pIndex][_tokenIndex].maxPay += pay; // update maxPay; hardlimit
        proInfo[_pIndex][_tokenIndex].managePay += manage; // pay to keeper, 4% of pay
    }

    // roughly
    function _settlementSub(uint64 _pIndex, uint8 _tokenIndex, uint64 end, uint64 size, uint256 sprice) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.time < end){
            if (se.time != 0) {
                proInfo[_pIndex][_tokenIndex].canPay += (end - se.time) * se.sprice;
            }
            proInfo[_pIndex][_tokenIndex].time = end;
        } else if(se.time > end) {
            // should sub it
            uint256 hp = (se.time - end) * sprice;
            if (proInfo[_pIndex][_tokenIndex].canPay > hp) {
                proInfo[_pIndex][_tokenIndex].canPay -= hp;
            } else {
                proInfo[_pIndex][_tokenIndex].canPay = 0;
            }
        }

        // update size and price
        if (proInfo[_pIndex][_tokenIndex].sprice > sprice) {
            proInfo[_pIndex][_tokenIndex].sprice -= sprice;  
        } else {
            proInfo[_pIndex][_tokenIndex].sprice = 0; 
        }

        if (proInfo[_pIndex][_tokenIndex].size > size) {
            proInfo[_pIndex][_tokenIndex].size -= size;
        } else {
            proInfo[_pIndex][_tokenIndex].size = 0;
        }
    }

    function _settlementCal(uint64 _pIndex, uint8 _tokenIndex, uint256 pay, uint256 lost) internal returns (uint256) {
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.maxPay<pay){
            return 0;
        }
        
        // 'has paid', or 'lost' is not right
        if(se.hasPaid > pay || se.lost > lost){
            return 0;
        }
        proInfo[_pIndex][_tokenIndex].lost = lost;

        uint64 ntime = uint64(block.timestamp);
        if(se.time < ntime){
            proInfo[_pIndex][_tokenIndex].canPay += (ntime - se.time) * se.sprice;
            proInfo[_pIndex][_tokenIndex].time = ntime;
        }

        // can pay is not right
        if(proInfo[_pIndex][_tokenIndex].canPay<pay){
            return 0;
        }

        uint256 res = pay - se.hasPaid;
        proInfo[_pIndex][_tokenIndex].hasPaid = pay;
        return res;
    }

    function addOrder(OrderIn memory ps, uint256 mr) external override onlyOwner {
        uint256 pay = (ps.end-ps.start) * ps.sPrice;
        uint256 manage = pay / 100 * uint256(mr);
        uint256 tax = pay / 100 * uint256(taxRate);
        uint256 payAndTax = pay + manage + tax;
        require(balances[ps.uIndex][ps.tIndex] + lock[ps.uIndex][ps.tIndex] >= payAndTax, "BNE"); // balance not enough

        // 验证nonce
        require(fs[ps.uIndex].ao[ps.pIndex].nonce == ps.nonce, "NE"); // nonce error
        // start不减, end不减
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start <= ps.start, "SE"); // start error, start shouldn't less than last order's start
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end <= ps.end, "EE"); // end error, end shouldn't less than last order's end

        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].sprice += ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size += ps.size;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start = ps.start;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end = ps.end;

        _settlementAdd(ps.pIndex, ps.tIndex, ps.start, ps.size, ps.sPrice, pay, manage);

        fs[ps.uIndex].ao[ps.pIndex].nonce++;

        // add to foundation
        balances[0][ps.tIndex] += tax;

        // pay from lock first
        if (lock[ps.uIndex][ps.tIndex] >= payAndTax) {
            lock[ps.uIndex][ps.tIndex] -= payAndTax;
        } else {
            payAndTax -= lock[ps.uIndex][ps.tIndex]; 
            lock[ps.uIndex][ps.tIndex] = 0;
            balances[ps.uIndex][ps.tIndex] -= payAndTax;
        }
    }

    function subOrder(OrderIn memory ps, uint256 mr) external override onlyOwner returns(uint256) {
        require(fs[ps.uIndex].ao[ps.pIndex].subNonce == ps.nonce, "EN"); // nonce error

        // update size and sprice
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].sprice -= ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size -= ps.size;

        // update settlement
        _settlementSub(ps.pIndex, ps.tIndex, ps.end, ps.size, ps.sPrice);

        uint256 endPaid = ps.sPrice * uint256(ps.end-ps.start) * mr / 400;
        proInfo[ps.pIndex][ps.tIndex].endPaid += endPaid;

        fs[ps.uIndex].ao[ps.pIndex].subNonce++;
        
        return endPaid;
    }

    function recharge(uint64 _i, uint8 _ti, uint256 money, bool isLock) external override onlyOwner {
        if (isLock) {
            lock[_i][_ti] += money;
        } else {
            balances[_i][_ti] += money;
        }  
    }

    // provider withdraw money, called by owner
    function proWithdraw(PWIn memory ps, uint256 _mr) external override onlyOwner returns(uint256, uint256) {
        // pay to provider
        uint256 thisPay = _settlementCal(ps.pIndex, ps.tIndex, ps.pay, ps.lost);
        if(thisPay==0){
            return (0,0);
        }

        // linear pay to keepers
        uint256 lpay = thisPay * _mr * 3 / 400; // 75% managePay for linearPaid
        proInfo[ps.pIndex][ps.tIndex].linearPaid += lpay;
        
        return (thisPay, lpay);
    }

    function withdraw(uint64 index, uint8 tokenIndex, uint256 amount) external override onlyOwner returns(uint256){
        uint256 bal = balances[index][tokenIndex];
        if(amount>bal){
            amount=bal;
        }

        balances[index][tokenIndex] -= amount;
        return amount;
    }

    // ================get=============
    function getFsInfo(uint64 _ui, uint64 _pi) external view override returns (FsOut memory){
        FsOut memory f;
        f = FsOut(fs[_ui].ao[_pi].nonce,fs[_ui].ao[_pi].subNonce);
        return f;
    }

    // get storeInfo in fs
    function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) external view override returns (StoreOut memory){
        StoreOut memory s;
        s = StoreOut(fs[_ui].ao[_pi].sInfo[_ti].start,fs[_ui].ao[_pi].sInfo[_ti].end,fs[_ui].ao[_pi].sInfo[_ti].size,fs[_ui].ao[_pi].sInfo[_ti].sprice);
        return s;
    }

    // 获得支付计费相关的信息
    function getSettleInfo(uint64 _pi, uint8 _ti) external view override returns (SettleOut memory){
        SettleOut memory s;
        s = SettleOut(proInfo[_pi][_ti].time,proInfo[_pi][_ti].size,proInfo[_pi][_ti].sprice,proInfo[_pi][_ti].maxPay,proInfo[_pi][_ti].hasPaid,proInfo[_pi][_ti].canPay,proInfo[_pi][_ti].lost,proInfo[_pi][_ti].lostPaid,proInfo[_pi][_ti].managePay,proInfo[_pi][_ti].endPaid,proInfo[_pi][_ti].linearPaid);
        return s;
    }

    function balanceOf(uint64 _i, uint8 _ti) external view override returns(uint256, uint256){
        uint256 avail = balances[_i][_ti];
        uint256 _lock = lock[_i][_ti];
        Settlement memory se = proInfo[_i][_ti];
        uint256 canPay = se.canPay;
        if (se.time > 0) {
            if (block.timestamp > se.time) {
                canPay += uint256(block.timestamp - se.time) * se.sprice;
            }

            uint256 tmp = se.maxPay - se.lost;
            if(canPay > tmp){
                canPay = tmp;
            }

            _lock += (canPay - se.hasPaid);
        }
        
        return (avail, _lock);
    }
}