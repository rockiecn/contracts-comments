// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/fileSysIn/IFileSys.sol";
import "../interfaces/erc20In/IERC20.sol";

/// @dev This contract is related to data storage and data reading payment in the file system.
contract FileSys is IFileSys {

    // StoreInfo is at some time
    struct StoreInfo {
        uint64 start; // last start 
        uint64 end;   // 什么时刻的状态，last end time
        uint64 size;   // 在该存储节点上的存储总量，byte
        uint256 price; // 按周期计费; per cycle
    }

    // AggOrder is AggregatedOrder is user->provider order
    struct AggOrder {
        uint64 nonce;    // 防止order重复提交
        uint64 subNonce; // 用于订单到期
        mapping(uint32 => StoreInfo) sInfo; // 不同代币的支付信息，tokenIndex => StoreInfo
    }

    // FsInfo each user have at most one group
    struct FsInfo {
        bool isActive;
        uint64[] providers;             // provider索引
        mapping(uint64 => AggOrder) ao; // 该User对每个Provider的订单信息
    }

    // Settlement indicates billing information
    struct Settlement {
        uint64 time; // store状态改变或支付的时间, 与 epoch 对齐
        uint64 size; // 在该存储节点上的存储总量
        uint256 price; // 累积的sprice(即sizePrice)

        uint256 maxPay;  // 对此provider所有user聚合总额度； expected 加和
        uint256 hasPaid; // 已经支付
        uint256 canPay;  // 最近一次store/pay时刻，可以支付的金额
        uint256 lost;    // lost due to unable response to challenge
        uint256 lostPaid;// pay to repair

        uint256 managePay; // pay for group keepers >= endPaid+linearPaid
        uint256 endPaid;   // release when order expire
        uint256 linearPaid;// release when pay for provider
    }

    // fs合约状态变量, constant
    uint8 manageRate = 4; // group分得资金的百分比；4% for group, 其中3% linear and 1% at end;
    uint8 taxRate = 1;    // 基金会分得资金的百分比；1% for foundation;
    uint64 gIndex;        // 指代所属的group

    mapping(uint64 => mapping(uint32 => uint256)) balances; // 账户可用的余额
    mapping(uint64 => mapping(uint32 => uint256)) penalty;  // 由于没有回应挑战而受到的惩罚

    mapping(uint64 => FsInfo) fs; // user => FsInfo; user 0 is repair fs

    // keeper profit related 
    uint64[] keepers; // for profit
    uint64 period;    // keeper根据比例获取收益的时间间隔
    uint64 lastTime;  // 上次分利润时间
    mapping(uint32 => uint256) tAcc; // 记录分润值，每次分润后归0，tokenIndex=>num
    uint64 totalCount; // 记录所有keeper触发order相关函数的总次数
    mapping(uint64 => uint64) count; // 记录keeper触发Order相关函数的次数，用于分润
    uint32[] tokens; // user使用某token时候加进来

    // pro聚合信息
    mapping(uint64 => mapping(uint32 => Settlement)) proInfo;

    address public role;   // role合约地址
    address public rolefs; // roleFS合约地址; can be reset?

    /// @dev created by admin; 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address
    constructor(uint64 _gIndex, address r, address rfs, uint64[] memory _keepers) {
        role = r;
        rolefs = rfs;

        gIndex = _gIndex;

        fs[0].isActive = true;

        keepers = _keepers;
        period = 1;
        lastTime = uint64(block.timestamp);

        for(uint64 i=0; i<keepers.length; i++){
            count[keepers[i]] = 1;
        }
        totalCount = uint64(keepers.length);
        tokens.push(0);
    }

    receive() external payable {}

    //函数修饰符，判断是不是owner调用
    modifier onlyRole(){
        require(msg.sender == role, "N");
        _;
    }

    modifier onlyRoleFS(){
        require(msg.sender == rolefs, "N");
        _;
    }

    function _settlementAdd(uint64 _pIndex, uint32 _tokenIndex, uint64 start, uint64 size, uint256 sprice, uint256 pay, uint256 manage) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.time < start){
            if(se.time!=0){ // 非首次addOrder
                proInfo[_pIndex][_tokenIndex].canPay += (start-se.time) * se.price;
            }
            proInfo[_pIndex][_tokenIndex].time = start;
        }else if(se.time > start){
            proInfo[_pIndex][_tokenIndex].canPay += uint256(se.time - start)*sprice;
        }

        // update price and size
        proInfo[_pIndex][_tokenIndex].price += sprice;
        proInfo[_pIndex][_tokenIndex].size += size;

        // update maxPay; hardlimit
        proInfo[_pIndex][_tokenIndex].maxPay += pay;

        // pay to keeper, 4% of pay
        proInfo[_pIndex][_tokenIndex].managePay += manage;
    }

    // roughly
    function _settlementSub(uint64 _pIndex, uint32 _tokenIndex, uint64 end, uint64 size, uint256 sprice) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.time < end){
            if (se.time != 0) {
                proInfo[_pIndex][_tokenIndex].canPay += (end - se.time) * se.price;
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
        if (proInfo[_pIndex][_tokenIndex].price > sprice) {
            proInfo[_pIndex][_tokenIndex].price -= sprice;  
        } else {
            proInfo[_pIndex][_tokenIndex].price = 0; 
        }

        if (proInfo[_pIndex][_tokenIndex].size > size) {
            proInfo[_pIndex][_tokenIndex].size -= size;
        } else {
            proInfo[_pIndex][_tokenIndex].size = 0;
        }
    }

    // _settlementCal called by func withdraw
    function _settlementCal(uint64 _pIndex, uint32 _tokenIndex, uint256 pay, uint256 lost) internal returns (uint256) {
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(proInfo[_pIndex][_tokenIndex].maxPay<pay){
            return 0;
        }
        
        // 'has paid', or 'lost' is not right
        if(se.hasPaid > pay || se.lost > lost){
            return 0;
        }
        proInfo[_pIndex][_tokenIndex].lost = lost;

        uint64 ntime = uint64(block.timestamp);
        if(se.time < ntime){
            proInfo[_pIndex][_tokenIndex].canPay += (ntime - se.time) * se.price;
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

    // called by Role-contract，在user注册时被role合约调用
    function createFs(uint64 _uIndex) external override onlyRole {
        require(!fs[_uIndex].isActive, "E"); // the fs already exists
        fs[_uIndex].isActive = true;
    }

    // called by Role-contract; when add keeper to group
    function addKeeper(uint64 _kIndex) external override onlyRole {
        keepers.push(_kIndex);
        count[_kIndex] = 1;
        totalCount++;
    }

    // called by RoleFS-contract
    function addOrder(AOParams memory ps) external override onlyRoleFS {
        require(fs[ps.uIndex].isActive, "NE"); // the fs does not exist

        if(!_hasToken(ps.tIndex)){
            tokens.push(ps.tIndex);
        }

        // 验证金额是否足够
        uint256 pay = (ps._end-ps._start) * ps.sPrice;
        uint256 manage = pay / 100 * uint256(manageRate);
        uint256 tax = pay / 100 * uint256(taxRate);
        uint256 payAndTax = pay + manage + tax;
        require(balances[ps.uIndex][ps.tIndex] >= payAndTax, "BNE"); // balance not enough

        // 验证nonce
        require(fs[ps.uIndex].ao[ps.pIndex].nonce == ps.nonce, "NE"); // nonce error
        // start不减, end不减
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start <= ps._start, "SE"); // start error, start shouldn't less than last order's start
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end <= ps._end, "EE"); // end error, end shouldn't less than last order's end

        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].price += ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size += ps._size;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start = ps._start;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end = ps._end;

        _settlementAdd(ps.pIndex, ps.tIndex, ps._start, ps._size, ps.sPrice, pay, manage);

        fs[ps.uIndex].ao[ps.pIndex].nonce++;

        // add to foundation
        balances[0][ps.tIndex] += tax;
        balances[ps.uIndex][ps.tIndex] -= payAndTax;
    }

    // called by RoleFS-contract
    function subOrder(uint64 kIndex, AOParams memory ps) external override onlyRoleFS {
        if(!_hasToken(ps.tIndex)){
            tokens.push(ps.tIndex);
        }

        require(fs[ps.uIndex].ao[ps.pIndex].subNonce == ps.nonce, "EN"); // nonce error

        // update size and price
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].price -= ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size -= ps._size;

        // update settlement
        _settlementSub(ps.pIndex, ps.tIndex, ps._end, ps._size, ps.sPrice);

        // pay to keeper, 其中的1%在结束时才支付,存储在endPaid中
        uint256 endPaid = ps.sPrice * uint256(ps._end-ps._start) / 100;
        proInfo[ps.pIndex][ps.tIndex].endPaid += endPaid;
        tAcc[ps.tIndex] += endPaid;

        // pay to keeper, 3% for linear, do in proWithdraw

        fs[ps.uIndex].ao[ps.pIndex].subNonce++;
        if(kIndex!=ps.uIndex){
            count[kIndex]++;
        }
        totalCount++;
    }

    // 由Role合约调用
    // 调用前需要uAddr先approve本合约地址（FileSys.sol）账户指定金额
    // user往文件系统FileSys合约中充值
    function recharge(address sender, uint64 uIndex, uint32 tIndex, address tAddr, uint256 money) external override onlyRole {
        require(fs[uIndex].isActive, "NA"); // the fs with user is not active

        IERC20(tAddr).transferFrom(sender, address(this), money);
        balances[uIndex][tIndex] += money;

        // add tIndex
        if(!_hasToken(tIndex)){
            tokens.push(tIndex);
        }
    }

    // provider withdraw money, called by owner
    function proWithdraw(address pAddr,address tAddr ,PWParams memory ps) external override onlyRoleFS {
        // pay to provider
        uint256 thisPay = _settlementCal(ps.pIndex, ps.tIndex, ps.pay, ps.lost);
        if(thisPay==0){
            return;
        }

        // linear pay to keepers
        uint256 lpay = thisPay * 3 / 100; // 3% thisPay for linearPaid
        proInfo[ps.pIndex][ps.tIndex].linearPaid += lpay;
        tAcc[ps.tIndex] += lpay;

        IERC20(tAddr).transfer(pAddr, thisPay); // transfer fail, the previous operation will be rolled back 
    }

    // user、keeper、foundation取回余额, called by owner
    function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) external override onlyRole {
        if(roleType == 3){ // keeper
            uint64 ntime = uint64(block.timestamp);
            if(ntime-lastTime > period){
                if(totalCount != 0){
                    for(uint i = 0; i< tokens.length; i++){
                        uint256 value = tAcc[tokens[i]]; // 分润值
                        if(value==0){
                            continue;
                        }
                        uint256 per = value / uint256(totalCount);
                        for(uint j =0; j<keepers.length; j++){
                            uint64 kc = count[keepers[j]]; // 触发的次数
                            if(kc==0){
                                count[keepers[j]] = 1;
                            }else{
                                uint256 pro = per * kc;
                                balances[keepers[j]][tokens[i]] += pro;
                            }
                        }
                        tAcc[tokens[i]] = 0;
                    }
                }

                lastTime = ntime;
            }
        }
        uint256 bal = balances[index][tokenIndex];
        if(amount>bal){
            amount=bal;
        }

        balances[index][tokenIndex] -= amount;
        IERC20(tAddr).transfer(addr, amount);
    }

    // called by owner
    function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external override onlyRoleFS {  
    }

    // called by owner
    function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external override onlyRoleFS {
    }

    // ================get=============

    // get some parameters in FsInfo
    function getFsInfo(uint64 user) external view override returns (bool){
        return fs[user].isActive;
    }

    // get providers sum in fs info
    function getFsPSum(uint64 user) external view override returns (uint64) {
        return uint64(fs[user].providers.length);
    }

    // get provider in fs info by array index
    function getFsPro(uint64 user, uint64 i) external view override returns (uint64) {
        return fs[user].providers[i];
    }

    // get provider's aggregate order in FsInfo
    function getFsInfoAggOrder(uint64 user, uint64 provider) external view override returns (uint64, uint64) {
        return (fs[user].ao[provider].nonce,fs[user].ao[provider].subNonce);
    }

    // get storeInfo in fs
    function getStoreInfo(uint64 user, uint64 provider, uint32 token) external view override returns (uint64, uint64, uint64, uint256){
        return (fs[user].ao[provider].sInfo[token].start, fs[user].ao[provider].sInfo[token].end,fs[user].ao[provider].sInfo[token].size, fs[user].ao[provider].sInfo[token].price);
    }

    // judge if tokens has the _tokenIndex
    function _hasToken(uint32 _tokenIndex) internal view returns (bool) {
        for(uint32 i = 0; i<tokens.length; i++){
            if(tokens[i]==_tokenIndex){
                return true;
            }
        }
        return false;
    }

    // 获得支付计费相关的信息
    function getSettleInfo(uint64 pIndex, uint32 tIndex) external view override returns (uint64, uint64, uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256){
        Settlement memory se = proInfo[pIndex][tIndex];
        return (se.time,se.size,se.price,se.maxPay,se.hasPaid,se.canPay,se.lost,se.lostPaid,se.managePay,se.endPaid,se.linearPaid);
    }

    // 获得账户收益余额
    function getBalance(uint64 index, uint32 tIndex) external view override returns(uint256, uint256){
        uint256 avail = balances[index][tIndex];
        Settlement memory se = proInfo[index][tIndex];
        uint256 canPay = se.canPay;
        if (block.timestamp > se.time) {
            canPay += uint256(block.timestamp - se.time) * se.price;
        }
        
        uint256 tmp = 0; 
        tmp = se.maxPay - se.lost;
        if(canPay > tmp){
            canPay = tmp;
        }
        tmp = canPay - se.hasPaid;

        if(totalCount == 0){
            return (avail, tmp);
        }

        if(count[index]!=0){ // 如果index是keeper,就需要根据比例计算分润值
            uint256 sum = tAcc[tIndex] * count[index];
            uint256 pro = sum / totalCount;
            if((block.timestamp-lastTime)>=period){
                avail += pro;
            }else{
                tmp+=pro;
            }
        }

        return (avail, tmp);
    }
}