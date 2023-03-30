package main
import (
	"fmt"
	"syscall"
	"unsafe"
	"time"
	"strings"
	"strconv"
	"bytes"
)

//资金账号
type AccountData struct{
	//投资者帐号
	accountID string
	//可用资金
	available float64
	//期货结算准备金
	balance float64
	//经纪公司代码
	brokerID string
	//资金差额
	cashIn float64
	//平仓盈亏
	closeProfit float64
	//手续费
	commission float64
	//信用额度
	credit float64
	//币种代码
	currencyID string
	//当前保证金总额
	currMargin float64
	//投资者交割保证金
	deliveryMargin float64
	//入金金额
	deposit float64
	//动态权益(新增)
	dynamicBalance float64
	//交易所交割保证金
	exchangeDeliveryMargin float64
	//交易所保证金
	exchangeMargin float64
	//浮动盈亏
	floatProfit float64
	//冻结的资金
	frozenCash float64
	//冻结的手续费
	frozenCommission float64
	//冻结的保证金
	frozenMargin float64
	//货币质押余额
	fundMortgageAvailable float64
	//货币质入金额
	fundMortgageIn float64
	//货币质出金额
	fundMortgageOut float64
	//利息收入
	interest float64
	//利息基数
	interestBase float64
	//质押金额
	mortgage float64
	//可质押货币金额
	mortgageableFund float64
	//持仓盈亏
	positionProfit float64
	//上次结算准备金
	preBalance float64
	//上次信用额度
	preCredit float64
	//上次存款额
	preDeposit float64
	//上次货币质入金额
	preFundMortgageIn float64
	//上次货币质出金额
	preFundMortgageOut float64
	//上次占用的保证金
	preMargin float64
	//上次质押金额
	preMortgage float64
	//基本准备金
	reserve float64
	//保底期货结算准备金
	reserveBalance float64
	//风险度(新增)
	risk float64
	//结算编号
	settlementID float64
	//特殊产品平仓盈亏
	specProductCloseProfit float64
	//特殊产品手续费
	specProductCommission float64
	//特殊产品交易所保证金
	specProductExchangeMargin float64
	//特殊产品冻结手续费
	specProductFrozenCommission float64
	//特殊产品冻结保证金
	specProductFrozenMargin float64
	//特殊产品占用保证金
	specProductMargin float64
	//特殊产品持仓盈亏
	specProductPositionProfit float64
	//根据持仓盈亏算法计算的特殊产品持仓盈亏
	specProductPositionProfitByAlg float64
	//交易日
	tradingDay string
	//出金金额
	withdraw float64
	//可取资金
	withdrawQuota float64
}

//码表
type Security struct{
	//组合类型
	combinationType string
	//创建日
	createDate string
	//交割月
	deliveryMonth float64
	//交割年份
	deliveryYear float64
	//结束交割日
	endDelivDate string
	//交易所代码
	exchangeID string
	//合约在交易所的代码
	exchangeInstID string
	//到期日
	expireDate string
	//保留小数位数
	digit int
	//合约生命周期状态
	instLifePhase string
	//合约代码
	instrumentID string
	//合约名称
	instrumentName string
	//当前是否交易
	isTrading string
	//多头保证金率
	longMarginRatio float64
	//限价单最大下单量
	maxLimitOrderVolume float64
	//是否使用大额单边保证金算法
	maxMarginSideAlgorithm string
	//市价单最大下单量
	maxMarketOrderVolume float64
	//限价单最小下单量
	minLimitOrderVolume float64
	//市价单最小下单量
	minMarketOrderVolume float64
	//上市日
	openDate string
	//期权类型
	optionsType string
	//持仓日期类型
	positionDateType string
	//持仓类型
	positionType string
	//最小变动价位
	priceTick float64
	//产品类型
	productClass string
	//产品代码
	productID string
	//空头保证金率
	shortMarginRatio float64
	//开始交割日
	startDelivDate string
	//执行价
	strikePrice float64
	//基础商品代码
	underlyingInstrID string
	//基础商品名称
	underlyingInstrName string
	//合约基础商品乘数
	underlyingMultiple float64
	//合约数量乘数
	volumeMultiple float64
}

//持仓
type InvestorPosition struct{
	//放弃执行冻结
	abandonFrozen float64
	//经纪公司代码
	brokerID string
	//资金差额
	cashIn float64
	//平仓金额
	closeAmount float64
	//平仓盈亏
	closeProfit float64
	//逐日盯市平仓盈亏
	closeProfitByDate float64
	//逐笔对冲平仓盈亏
	closeProfitByTrade float64
	//平仓量
	closeVolume float64
	//合约代码
	code string
	//组合多头冻结
	combLongFrozen float64
	//组合成交形成的持仓
	combPosition float64
	//组合空头冻结
	combShortFrozen float64
	//手续费
	commission float64
	//交易所保证金
	exchangeMargin float64
	//浮动盈亏
	floatProfit float64
	//冻结的资金
	frozenCash float64
	//冻结的手续费
	frozenCommission float64
	//冻结的保证金
	frozenMargin float64
	//投机套保标志
	hedgeFlag string
	//投资者代码
	investorID string
	//多头冻结
	longFrozen float64
	//多头冻结金额
	longFrozenAmount float64
	//保证金
	margin float64
	//保证金率
	marginRateByMoney float64
	//保证金率(按手数)
	marginRateByVolume float64
	//开仓金额
	openAmount float64
	//开仓成本
	openCost float64
	//开仓价格
	openPrice float64
	//开仓量
	openVolume float64
	//今日持仓
	position float64
	//持仓日期
	positionDate string
	//持仓多空方向
	posiDirection string
	//持仓成本
	positionCost float64
	//持仓盈亏
	positionProfit float64
	//上次占用的保证金
	preMargin float64
	//上次结算价
	preSettlementPrice float64
	//结算编号
	settlementID float64
	//本次结算价
	settlementPrice float64
	//空头冻结
	shortFrozen float64
	//空头冻结金额
	shortFrozenAmount float64
	//执行冻结
	strikeFrozen float64
	//执行冻结金额
	strikeFrozenAmount float64
	//今日持仓
	todayPosition float64
	//交易日
	tradingDate string
	//占用的保证金
	useMargin float64
	//上日持仓
	ydPosition float64
}

//持仓明细
type InvestorPositionDetail struct{
	//经纪公司代码
	brokerID string
	//平仓金额
	closeAmount float64
	//平仓盈亏
	closeProfit float64
	//逐日盯市持仓盈亏
	closeProfitByDate float64
	//逐笔对冲持仓盈亏
	closeProfitByTrade float64
	//平仓量
	closeVolume float64
	//合约代码
	code string
	//组合合约代码
	combInstrumentID string
	//买卖
	direction string
	//交易所代码
	exchangeID string
	//交易所保证金
	exchMargin float64
	//浮动盈亏
	floatProfit float64
	//投机套保标志
	hedgeFlag string
	//投资者代码
	investorID string
	//昨收盘
	lastPrice float64
	//昨结算价
	lastSettlementPrice float64
	//投资者保证金
	margin float64
	//保证金率
	marginRateByMoney float64
	//保证金率(按手数)
	marginRateByVolume float64
	//内部编号
	orderRef string
	//开仓日期
	openDate string
	//开仓价
	openPrice float64
	//开仓量
	openVolume float64
	//本地持仓号，服务器编写
	positionNo string
	//持仓盈亏
	positionProfit float64
	//逐日盯市持仓利润
	positionProfitByDate float64
	//逐笔对冲持仓利润
	positionProfitByTrade float64
	//持仓流号
	positionStreamId float64
	//昨结算价
	preSettlementPrice float64
	//持仓盈亏基准价
	profitBasePrice float64
	//结算编号
	settlementID float64
	//结算价
	settlementPrice float64
	//成交日期
	tradingDay string
	//成交编号
	tradeID string
	//成交类型
	tradeType string
	//数量
	volume float64
}

//委托回报
type OrderInfo struct{
	//激活时间
	activeTime string
	//最后修改交易所交易员代码
	activeTraderID string
	//操作用户代码
	activeUserID string
	//经纪公司代码
	brokerID string
	//经纪公司报单编号
	brokerOrderSeq float64
	//业务单元
	businessUnit string
	//撤销时间
	cancelTime string
	//结算会员编号
	clearingPartID string
	//客户代码
	clientID string
	//合约代码
	code string
	//组合投机套保标志
	combHedgeFlag string
	//组合开平标志
	combOffsetFlag string
	//触发条件
	contingentCondition string
	//买卖方向
	direction string
	//交易所代码
	exchangeID string
	//合约在交易所的代码
	exchangeInstID string
	//强平原因
	forceCloseReason string
	//前置编号
	frontID float64
	//GTD日期
	gTDDate string
	//价格
	limitPrice float64
	//报单日期
	insertDate string
	//委托时间
	insertTime string
	//安装编号
	installID string
	//投资者代码
	investorID string
	//自动挂起标志
	isAutoSuspend float64
	//互换单标志
	isSwapOrder float64
	//最小成交量
	minVolume float64
	//报单提示序号
	notifySequence float64
	//本地报单编号
	orderLocalID string
	//报单价格条件
	orderPriceType string
	//报单引用
	orderRef string
	//报单状态
	orderStatus string
	//报单来源
	orderSource string
	//报单提交状态
	orderSubmitStatus string
	//报单编号
	orderSysID string
	//报单类型
	orderType string
	//会员代码
	participantID string
	//相关报单
	relativeOrderSysID string
	//请求编号
	requestID float64
	//序号
	sequenceNo float64
	//会话编号
	sessionID float64
	//结算编号
	settlementID float64
	//状态信息
	statusMsg string
	//止损价
	stopPrice float64
	//挂起时间
	suspendTime string
	//有效期类型
	timeCondition string
	//交易所交易员代码
	traderID string
	//交易日
	tradingDay string
	//最后修改时间
	updateTime string
	//用户强评标志
	userForceClose float64
	//用户代码
	userID string
	//用户端产品信息
	userProductInfo string
	//成交量类型
	volumeCondition string
	//剩余数量
	volumeTotal float64
	//数量
	volumeTotalOriginal float64
	//今成交数量
	volumeTraded float64
	//郑商所成交数量
	zCETotalTradedVolume float64
}

//最新数据
type SecurityLatestData struct{
	//触发日
	actionDay string
	//卖1价
	askPrice1 float64
	//卖2价
	askPrice2 float64
	//卖3价
	askPrice3 float64
	//卖4价
	askPrice4 float64
	//卖5价
	askPrice5 float64
	//卖1量
	askVolume1 int
	//卖2量
	askVolume2 int
	//卖3量
	askVolume3 int
	//卖4量
	askVolume4 int
	//卖5量
	askVolume5 int
	//平均价格
	averagePrice float64
	//买1价
	bidPrice1 float64
	//买2价
	bidPrice2 float64
	//买3价
	bidPrice3 float64
	//买4价
	bidPrice4 float64
	//买5价
	bidPrice5 float64
	//买1量
	bidVolume1 int
	//买2量
	bidVolume2 int
	//买3量
	bidVolume3 int
	//买4量
	bidVolume4 int
	//买5量
	bidVolume5 int
	//最新价
	close float64
	//合约代码
	code string
	//今虚实度
	currDelta float64
	//交易所ID
	exchangeID string
	//交易所InstID
	exchangeInstID string
	//最高价
	high float64
	//昨日收盘价
	lastClose float64
	//最低价
	low float64
	//跌停价
	lowerLimit float64
	//开盘价
	open float64
	//持仓量
	openInterest float64
	//昨收盘
	preClose float64
	//昨虚实度
	preDelta float64
	//昨持仓量
	preOpenInterest float64
	//上次结算价
	preSettlementPrice float64
	//本次结算价
	settlementPrice float64
	//交易日
	tradingDay string
	//成交金额
	turnover float64
	//最后修改毫秒
	updateMillisec float64
	//最后修改时间
	updateTime string
	//涨停价
	upperLimit float64
	//成交量
	volume float64
}

//成交回报
type TradeRecord struct{
	//经纪公司代码
	brokerID string
	///经纪公司报单编号
	brokerOrderSeq float64
	///业务单元
	businessUnit string
	///结算会员编号
	clearingPartID string
	///客户代码
	clientID string
	//合约代码
	code string
	//手续费
	commission float64
	//买卖方向
	direction string
	//市场代码
	exchangeID string
	//合约在交易所的代码
	exchangeInstID string
	//投机套保标志
	hedgeFlag string
	//投资者代码
	investorID string
	//开平标志
	offsetFlag string
	//本地报单编号
	orderLocalID string
	//报单引用
	orderRef string
	//报单编号
	orderSysID string
	//会员代码
	participantID string
	//价格
	price float64
	//成交价来源
	priceSource string
	//序号
	sequenceNo float64
	//结算编号
	settlementID float64
	//成交编号
	tradeID string
	//交易所交易员代码
	traderID string
	//成交时期
	tradeDate string
	//成交来源
	tradeSource string
	//成交时间
	tradeTime string
	//交易日
	tradingDay string
	//成交类型
	tradeType string
	//交易角色
	tradingRole string
	//数量
	volume float64
	//用户代码
	userID string
}

//合约手续费率
type CommissionRate struct{
	//经纪公司代码
	brokerID string
	//收费方式
	calculateMode string
	//平仓手续费率
	closeRatioByMoney float64
	//平仓手续费
	closeRatioByVolume float64
	//平今手续费率
	closeTodayRatioByMoney float64
	//平今手续费
	closeTodayRatioByVolume float64
	//平今费
	closeTodayFee float64
	//合约代码
	code string
	//代码
	commodityNo string
	//类型
	commodityType string
	//交易所编号
	exchangeNo string
	//投资者代码
	investorID string
	//投资者范围
	investorRange string
	//来源
	matchSource string
	//开平费
	openCloseFee float64
	//开仓手续费率
	openRatioByMoney float64
	//开仓手续费
	openRatioByVolume float64
}

//合约保证金率
type MarginRate struct{
	//经纪公司代码
	brokerID string
	//收费方式
	calculateMode string
	//看涨看跌标示
	callOrPutFlag string
	//合约代码
	code string
	//代码
	commodityNo string
	//类型
	commodityType string
	//合约
	contractNo string
	//投机套保标志
	hedgeFlag string
	initialMargin float64
	//投资者代码
	investorID string
	//多头保证金率
	longMarginRatioByMoney float64
	//多头保证金费
	longMarginRatioByVolume float64
	//投资者范围
	investorRange string
	//是否相对交易所收取
	isRelativel float64
	lockMargin float64
	maintenanceMargin float64
	sellInitialMargin float64
	sellMaintenanceMargin float64
	//空头保证金率
	shortMarginRatioByMoney float64
	//空头保证金费
	shortMarginRatioByVolume float64
	strikePrice string
}

//字符串转换成float64
func convertToFloat64(str string)float64{
	floatvar, err := strconv.ParseFloat(str, 64)
	if(err != nil){
		fmt.Println(err)
	}
	return floatvar
}

//字符串转换成int
func convertToInt(str string)int{
	intvar, err := strconv.ParseInt(str, 10, 32)
	if(err != nil){
		fmt.Println(err)
	}
	return int(intvar)
}

//转换成CTP手续费率
func convertToCTPCommissionRate(result string)(CommissionRate){
	cTPCommissionRate := CommissionRate{}
	results := strings.Split(result, ",")
	cTPCommissionRate.code = results[0]
	cTPCommissionRate.investorRange = results[1]
	cTPCommissionRate.brokerID = results[2]
	cTPCommissionRate.investorID = results[3]
	cTPCommissionRate.openRatioByMoney = convertToFloat64(results[4])
	cTPCommissionRate.openRatioByVolume = convertToFloat64(results[5])
	cTPCommissionRate.closeRatioByMoney = convertToFloat64(results[6])
	cTPCommissionRate.closeRatioByVolume = convertToFloat64(results[7])
	cTPCommissionRate.closeTodayRatioByMoney = convertToFloat64(results[8])
	cTPCommissionRate.closeTodayRatioByVolume = convertToFloat64(results[9])
	return cTPCommissionRate
}

//转换CTP保证金率
func convertToCTPMarginRate(result string)(MarginRate){
	cTPMarginRate := MarginRate{}
	results := strings.Split(result, ",")
	cTPMarginRate.code = results[0]
	cTPMarginRate.brokerID = results[1]
	cTPMarginRate.investorID = results[2]
	cTPMarginRate.hedgeFlag = results[3]
	cTPMarginRate.longMarginRatioByMoney = convertToFloat64(results[4])
	cTPMarginRate.longMarginRatioByVolume = convertToFloat64(results[5])
	cTPMarginRate.shortMarginRatioByMoney = convertToFloat64(results[6])
	cTPMarginRate.shortMarginRatioByVolume = convertToFloat64(results[7])
	cTPMarginRate.isRelativel = convertToFloat64(results[8])
	return cTPMarginRate
}

//转换资金账户结构
func convertToCTPAccountData(result string)(AccountData){
	cTPTradingAccount := AccountData{}
	results := strings.Split(result, ",")
	if(len(results) >= 46){
		cTPTradingAccount.brokerID = results[0]
		cTPTradingAccount.accountID = results[1]
		cTPTradingAccount.preMortgage = convertToFloat64(results[2])
		cTPTradingAccount.preCredit = convertToFloat64(results[3])
		cTPTradingAccount.preDeposit = convertToFloat64(results[4])
		cTPTradingAccount.preBalance = convertToFloat64(results[5])
		cTPTradingAccount.preMargin = convertToFloat64(results[6])
		cTPTradingAccount.interestBase = convertToFloat64(results[7])
		cTPTradingAccount.interest = convertToFloat64(results[8])
		cTPTradingAccount.deposit = convertToFloat64(results[9])
		cTPTradingAccount.withdraw = convertToFloat64(results[10])
		cTPTradingAccount.frozenMargin = convertToFloat64(results[11])
		cTPTradingAccount.frozenCash = convertToFloat64(results[12])
		cTPTradingAccount.frozenCommission = convertToFloat64(results[13])
		cTPTradingAccount.currMargin = convertToFloat64(results[14])
		cTPTradingAccount.cashIn = convertToFloat64(results[15])
		cTPTradingAccount.commission = convertToFloat64(results[16])
		cTPTradingAccount.closeProfit = convertToFloat64(results[17])
		cTPTradingAccount.positionProfit = convertToFloat64(results[18])
		cTPTradingAccount.balance = convertToFloat64(results[19])
		cTPTradingAccount.available = convertToFloat64(results[20])
		cTPTradingAccount.withdrawQuota = convertToFloat64(results[21])
		cTPTradingAccount.reserve = convertToFloat64(results[22])
		cTPTradingAccount.tradingDay = results[23]
		cTPTradingAccount.settlementID = convertToFloat64(results[24])
		cTPTradingAccount.credit = convertToFloat64(results[25])
		cTPTradingAccount.mortgage = convertToFloat64(results[26])
		cTPTradingAccount.exchangeMargin = convertToFloat64(results[27])
		cTPTradingAccount.deliveryMargin = convertToFloat64(results[28])
		cTPTradingAccount.exchangeDeliveryMargin = convertToFloat64(results[29])
		cTPTradingAccount.reserveBalance = convertToFloat64(results[30])
		cTPTradingAccount.currencyID = results[31]
		cTPTradingAccount.preFundMortgageIn = convertToFloat64(results[32])
		cTPTradingAccount.preFundMortgageOut = convertToFloat64(results[33])
		cTPTradingAccount.fundMortgageIn = convertToFloat64(results[34])
		cTPTradingAccount.fundMortgageOut = convertToFloat64(results[35])
		cTPTradingAccount.fundMortgageAvailable = convertToFloat64(results[36])
		cTPTradingAccount.mortgageableFund = convertToFloat64(results[37])
		cTPTradingAccount.specProductMargin = convertToFloat64(results[38])
		cTPTradingAccount.specProductFrozenMargin = convertToFloat64(results[39])
		cTPTradingAccount.specProductCommission = convertToFloat64(results[40])
		cTPTradingAccount.specProductFrozenCommission = convertToFloat64(results[41])
		cTPTradingAccount.specProductPositionProfit = convertToFloat64(results[42])
		cTPTradingAccount.specProductCloseProfit = convertToFloat64(results[43])
		cTPTradingAccount.specProductPositionProfitByAlg = convertToFloat64(results[44])
		cTPTradingAccount.specProductExchangeMargin = convertToFloat64(results[45])
		cTPTradingAccount.dynamicBalance = convertToFloat64(results[46])
		cTPTradingAccount.risk = convertToFloat64(results[47])
		cTPTradingAccount.floatProfit = convertToFloat64(results[48])
	}
	return cTPTradingAccount
}

//转换深度数据结构
func convertToCTPDepthMarketData(result string)([]SecurityLatestData){
	var list []SecurityLatestData
	strs := strings.Split(result, ",")
	length := len(strs)
	for i := 0; i < length; i++ {
		results := strings.Split(strs[i], ",")
		if (len(results) >= 43){
			cTPDepthMarketData := SecurityLatestData{}
			cTPDepthMarketData.tradingDay = results[0]
			cTPDepthMarketData.code = results[1]
			cTPDepthMarketData.exchangeID = results[2]
			cTPDepthMarketData.exchangeInstID = results[3]
			cTPDepthMarketData.close = convertToFloat64(results[4])
			cTPDepthMarketData.preSettlementPrice = convertToFloat64(results[5])
			cTPDepthMarketData.preClose = convertToFloat64(results[6])
			cTPDepthMarketData.preOpenInterest = convertToFloat64(results[7])
			cTPDepthMarketData.open = convertToFloat64(results[8])
			cTPDepthMarketData.high = convertToFloat64(results[9])
			cTPDepthMarketData.low = convertToFloat64(results[10])
			cTPDepthMarketData.volume = convertToFloat64(results[11])
			cTPDepthMarketData.turnover = convertToFloat64(results[12])
			cTPDepthMarketData.openInterest = convertToFloat64(results[13])
			cTPDepthMarketData.lastClose = convertToFloat64(results[14])
			cTPDepthMarketData.settlementPrice = convertToFloat64(results[15])
			cTPDepthMarketData.upperLimit = convertToFloat64(results[16])
			cTPDepthMarketData.lowerLimit = convertToFloat64(results[17])
			cTPDepthMarketData.preDelta = convertToFloat64(results[18])
			cTPDepthMarketData.currDelta = convertToFloat64(results[19])
			cTPDepthMarketData.updateTime = results[20]
			cTPDepthMarketData.updateMillisec = convertToFloat64(results[21])
			cTPDepthMarketData.bidPrice1 = convertToFloat64(results[22])
			cTPDepthMarketData.bidVolume1 = convertToInt(results[23])
			cTPDepthMarketData.askPrice1 = convertToFloat64(results[24])
			cTPDepthMarketData.askVolume1 = convertToInt(results[25])
			cTPDepthMarketData.bidPrice2 = convertToFloat64(results[26])
			cTPDepthMarketData.bidVolume2 = convertToInt(results[27])
			cTPDepthMarketData.askPrice2 = convertToFloat64(results[28])
			cTPDepthMarketData.askVolume2 = convertToInt(results[29])
			cTPDepthMarketData.bidPrice3 = convertToFloat64(results[30])
			cTPDepthMarketData.bidVolume3 = convertToInt(results[31])
			cTPDepthMarketData.askPrice3 = convertToFloat64(results[32])
			cTPDepthMarketData.askVolume3 = convertToInt(results[33])
			cTPDepthMarketData.bidPrice4 = convertToFloat64(results[34])
			cTPDepthMarketData.bidVolume4 = convertToInt(results[35])
			cTPDepthMarketData.askPrice4 = convertToFloat64(results[36])
			cTPDepthMarketData.askVolume4 = convertToInt(results[37])
			cTPDepthMarketData.bidPrice5 = convertToFloat64(results[38])
			cTPDepthMarketData.bidVolume5 = convertToInt(results[39])
			cTPDepthMarketData.askPrice5 = convertToFloat64(results[40])
			cTPDepthMarketData.askVolume5 = convertToInt(results[41])
			cTPDepthMarketData.averagePrice = convertToFloat64(results[42])
			cTPDepthMarketData.actionDay = results[43]
			list = append(list, cTPDepthMarketData)
		}
	}
	return list
}

//转换期货信息结构
func convertToCTPInstrumentDatas(result string)([]Security){
	var cTPInstrumentDatas []Security
	strs := strings.Split(result, ",")
	length := len(strs)
	for i := 0; i < length; i++ {
		results := strings.Split(strs[i], ",")
		if (len(results) >= 30){
			cTPInstrumentData := Security{}
			cTPInstrumentData.instrumentID = results[0]
			cTPInstrumentData.exchangeID = results[1]
			cTPInstrumentData.instrumentName = results[2]
			cTPInstrumentData.exchangeInstID = results[3]
			cTPInstrumentData.productID = results[4]
			cTPInstrumentData.productClass = results[5]
			cTPInstrumentData.deliveryYear = convertToFloat64(results[6])
			cTPInstrumentData.deliveryMonth = convertToFloat64(results[7])
			cTPInstrumentData.maxMarketOrderVolume = convertToFloat64(results[8])
			cTPInstrumentData.minMarketOrderVolume = convertToFloat64(results[9])
			cTPInstrumentData.maxLimitOrderVolume = convertToFloat64(results[10])
			cTPInstrumentData.minLimitOrderVolume = convertToFloat64(results[11])
			cTPInstrumentData.volumeMultiple = convertToFloat64(results[12])
			cTPInstrumentData.priceTick = convertToFloat64(results[13])
			cTPInstrumentData.createDate = results[14]
			cTPInstrumentData.openDate = results[15]
			cTPInstrumentData.expireDate = results[16]
			cTPInstrumentData.startDelivDate = results[17]
			cTPInstrumentData.endDelivDate = results[18]
			cTPInstrumentData.instLifePhase = results[19]
			cTPInstrumentData.isTrading = results[20]
			cTPInstrumentData.positionType = results[21]
			cTPInstrumentData.positionDateType = results[22]
			cTPInstrumentData.longMarginRatio = convertToFloat64(results[23])
			cTPInstrumentData.shortMarginRatio = convertToFloat64(results[24])
			cTPInstrumentData.maxMarginSideAlgorithm = results[25]
			cTPInstrumentData.underlyingInstrID = results[26]
			cTPInstrumentData.strikePrice = convertToFloat64(results[27])
			cTPInstrumentData.optionsType = results[28]
			cTPInstrumentData.underlyingMultiple = convertToFloat64(results[29])
			cTPInstrumentData.combinationType = results[30]
			cTPInstrumentDatas = append(cTPInstrumentDatas, cTPInstrumentData)
		}
	}
	return cTPInstrumentDatas
}

//转换持仓结构
func convertToCTPInvestorPosition(result string)([]InvestorPosition){
	var list []InvestorPosition
	strs := strings.Split(result, ",")
	length := len(strs)
	for i := 0; i < length; i++ {
		results := strings.Split(strs[i], ",")
		if (len(results) >= 42){
			cTPInvestorPosition := InvestorPosition{}
			cTPInvestorPosition.code = results[0]
			cTPInvestorPosition.brokerID = results[1]
			cTPInvestorPosition.investorID = results[2]
			cTPInvestorPosition.posiDirection = results[3]
			cTPInvestorPosition.hedgeFlag = results[4]
			cTPInvestorPosition.positionDate = results[5]
			cTPInvestorPosition.ydPosition = convertToFloat64(results[6])
			cTPInvestorPosition.position = convertToFloat64(results[7])
			cTPInvestorPosition.longFrozen = convertToFloat64(results[8])
			cTPInvestorPosition.shortFrozen = convertToFloat64(results[9])
			cTPInvestorPosition.longFrozenAmount = convertToFloat64(results[10])
			cTPInvestorPosition.shortFrozenAmount = convertToFloat64(results[11])
			cTPInvestorPosition.openVolume = convertToFloat64(results[12])
			cTPInvestorPosition.closeVolume = convertToFloat64(results[13])
			cTPInvestorPosition.openAmount = convertToFloat64(results[14])
			cTPInvestorPosition.closeAmount = convertToFloat64(results[15])
			cTPInvestorPosition.positionCost = convertToFloat64(results[16])
			cTPInvestorPosition.preMargin = convertToFloat64(results[17])
			cTPInvestorPosition.useMargin = convertToFloat64(results[18])
			cTPInvestorPosition.frozenMargin = convertToFloat64(results[19])
			cTPInvestorPosition.frozenCash = convertToFloat64(results[20])
			cTPInvestorPosition.frozenCommission = convertToFloat64(results[21])
			cTPInvestorPosition.cashIn = convertToFloat64(results[22])
			cTPInvestorPosition.margin = convertToFloat64(results[23])
			cTPInvestorPosition.floatProfit = convertToFloat64(results[24])
			cTPInvestorPosition.positionProfit = convertToFloat64(results[25])
			cTPInvestorPosition.preSettlementPrice = convertToFloat64(results[26])
			cTPInvestorPosition.settlementPrice = convertToFloat64(results[27])
			cTPInvestorPosition.tradingDate = results[28]
			cTPInvestorPosition.settlementID = convertToFloat64(results[29])
			cTPInvestorPosition.openCost = convertToFloat64(results[30])
			cTPInvestorPosition.exchangeMargin = convertToFloat64(results[31])
			cTPInvestorPosition.combPosition = convertToFloat64(results[32])
			cTPInvestorPosition.combLongFrozen = convertToFloat64(results[33])
			cTPInvestorPosition.combShortFrozen = convertToFloat64(results[34])
			cTPInvestorPosition.closeProfitByDate = convertToFloat64(results[35])
			cTPInvestorPosition.closeProfitByTrade = convertToFloat64(results[36])
			cTPInvestorPosition.todayPosition = convertToFloat64(results[37])
			cTPInvestorPosition.marginRateByMoney = convertToFloat64(results[38])
			cTPInvestorPosition.marginRateByVolume = convertToFloat64(results[39])
			cTPInvestorPosition.strikeFrozen = convertToFloat64(results[40])
			cTPInvestorPosition.strikeFrozenAmount = convertToFloat64(results[41])
			cTPInvestorPosition.abandonFrozen = convertToFloat64(results[42])
			cTPInvestorPosition.openPrice = convertToFloat64(results[43])
			list = append(list, cTPInvestorPosition)
		}
	}
	return list
}

//转换持仓明细结构
func convertToCTPInvestorPositionDetail(result string)([]InvestorPositionDetail){
	var list []InvestorPositionDetail
	strs := strings.Split(result, ",")
	length := len(strs)
	for i := 0; i < length; i++ {
		results := strings.Split(strs[i], ",")
		if (len(results) >= 25){
			cTPInvestorPositionDetail := InvestorPositionDetail{}
			cTPInvestorPositionDetail.code = results[0]
			cTPInvestorPositionDetail.brokerID = results[1]
			cTPInvestorPositionDetail.investorID = results[2]
			cTPInvestorPositionDetail.hedgeFlag = results[3]
			cTPInvestorPositionDetail.direction = results[4]
			cTPInvestorPositionDetail.openDate = results[5]
			cTPInvestorPositionDetail.tradeID = results[6]
			cTPInvestorPositionDetail.volume = convertToFloat64(results[7])
			cTPInvestorPositionDetail.openPrice = convertToFloat64(results[8])
			cTPInvestorPositionDetail.tradingDay = results[9]
			cTPInvestorPositionDetail.settlementID = convertToFloat64(results[10])
			cTPInvestorPositionDetail.tradeType = results[11]
			cTPInvestorPositionDetail.combInstrumentID = results[12]
			cTPInvestorPositionDetail.exchangeID = results[13]
			cTPInvestorPositionDetail.closeProfitByDate = convertToFloat64(results[14])
			cTPInvestorPositionDetail.closeProfitByTrade = convertToFloat64(results[15])
			cTPInvestorPositionDetail.positionProfitByDate = convertToFloat64(results[16])
			cTPInvestorPositionDetail.positionProfitByTrade = convertToFloat64(results[17])
			cTPInvestorPositionDetail.margin = convertToFloat64(results[18])
			cTPInvestorPositionDetail.exchMargin = convertToFloat64(results[19])
			cTPInvestorPositionDetail.marginRateByMoney = convertToFloat64(results[20])
			cTPInvestorPositionDetail.marginRateByVolume = convertToFloat64(results[21])
			cTPInvestorPositionDetail.lastSettlementPrice = convertToFloat64(results[22])
			cTPInvestorPositionDetail.lastSettlementPrice = convertToFloat64(results[23])
			cTPInvestorPositionDetail.settlementPrice = convertToFloat64(results[24])
			cTPInvestorPositionDetail.closeVolume = convertToFloat64(results[25])
			list = append(list, cTPInvestorPositionDetail)
		}
	}
	return list
}

//转换成委托回报结构
func convertToCTPOrder(result string)(OrderInfo){
	cTPOrder := OrderInfo{}
	results := strings.Split(result, ",")
	if(len(results) >= 56){
		cTPOrder.brokerID = results[0]
		cTPOrder.investorID = results[1]
		cTPOrder.code = results[2]
		cTPOrder.orderRef = results[3]
		cTPOrder.userID = results[4]
		cTPOrder.orderPriceType = results[5]
		cTPOrder.direction = results[6]
		cTPOrder.combOffsetFlag = results[7]
		cTPOrder.combHedgeFlag = results[8]
		cTPOrder.limitPrice = convertToFloat64(results[9])
		cTPOrder.volumeTotalOriginal = convertToFloat64(results[10])
		cTPOrder.timeCondition = results[11]
		cTPOrder.gTDDate = results[12]
		cTPOrder.volumeCondition = results[13]
		cTPOrder.minVolume = convertToFloat64(results[14])
		cTPOrder.contingentCondition = results[15]
		cTPOrder.stopPrice = convertToFloat64(results[16])
		cTPOrder.forceCloseReason = results[17]
		cTPOrder.isAutoSuspend = convertToFloat64(results[18])
		cTPOrder.businessUnit = results[19]
		cTPOrder.requestID = convertToFloat64(results[20])
		cTPOrder.orderLocalID = results[21]
		cTPOrder.exchangeID = results[22]
		cTPOrder.participantID = results[23]
		cTPOrder.clientID = results[24]
		cTPOrder.exchangeInstID = results[25]
		cTPOrder.traderID = results[26]
		cTPOrder.installID = results[27]
		cTPOrder.orderSubmitStatus = results[28]
		cTPOrder.notifySequence = convertToFloat64(results[29])
		cTPOrder.tradingDay = results[30]
		cTPOrder.settlementID = convertToFloat64(results[31])
		cTPOrder.orderSysID = results[32]
		cTPOrder.orderSource = results[33]
		cTPOrder.orderStatus = results[34]
		cTPOrder.orderType = results[35]
		cTPOrder.volumeTraded = convertToFloat64(results[36])
		cTPOrder.volumeTotal = convertToFloat64(results[37])
		cTPOrder.insertDate = results[38]
		cTPOrder.insertTime = results[39]
		cTPOrder.activeTime = results[40]
		cTPOrder.suspendTime = results[41]
		cTPOrder.updateTime = results[42]
		cTPOrder.cancelTime = results[43]
		cTPOrder.activeTraderID = results[44]
		cTPOrder.clearingPartID = results[45]
		cTPOrder.sequenceNo = convertToFloat64(results[46])
		cTPOrder.frontID = convertToFloat64(results[47])
		cTPOrder.sessionID = convertToFloat64(results[48])
		cTPOrder.userProductInfo = results[49]
		cTPOrder.statusMsg = results[50]
		cTPOrder.userForceClose = convertToFloat64(results[51])
		cTPOrder.activeUserID = results[52]
		cTPOrder.brokerOrderSeq = convertToFloat64(results[53])
		cTPOrder.relativeOrderSysID = results[54]
		cTPOrder.zCETotalTradedVolume = convertToFloat64(results[55])
		cTPOrder.isSwapOrder = convertToFloat64(results[56])
	}
	return cTPOrder
}

//转换持仓明细结构
func convertToCTPOrderList(result string)([]OrderInfo){
	var lst []OrderInfo
	strs := strings.Split(result, ",")
	length := len(strs)
	for i := 0; i < length; i++ {
		results := strings.Split(strs[i], ",")
		lst = append(lst, convertToCTPOrder(results[i]))
	}
	return lst
}

//转换成成交回报结构
func convertToCTPTrade(result string)(TradeRecord){
	cTPTrade := TradeRecord{}
	results := strings.Split(result, ",")
	if(len(results) > 30){
		cTPTrade.brokerID = results[0]
		cTPTrade.investorID = results[1]
		cTPTrade.code = results[2]
		cTPTrade.orderRef = results[3]
		cTPTrade.userID = results[4]
		cTPTrade.exchangeID = results[5]
		cTPTrade.tradeID = results[6]
		cTPTrade.direction = results[7]
		cTPTrade.orderSysID = results[8]
		cTPTrade.participantID = results[9]
		cTPTrade.clientID = results[10]
		cTPTrade.tradingRole = results[11]
		cTPTrade.exchangeInstID = results[12]
		cTPTrade.offsetFlag = results[13]
		cTPTrade.hedgeFlag = results[14]
		cTPTrade.price = convertToFloat64(results[15])
		cTPTrade.volume = convertToFloat64(results[16])
		cTPTrade.tradeDate = results[17]
		cTPTrade.tradeTime = results[18]
		cTPTrade.tradeType = results[19]
		cTPTrade.priceSource = results[20]
		cTPTrade.traderID = results[21]
		cTPTrade.orderLocalID = results[22]
		cTPTrade.clearingPartID = results[23]
		cTPTrade.businessUnit = results[24]
		cTPTrade.sequenceNo = convertToFloat64(results[25])
		cTPTrade.tradingDay = results[26]
		cTPTrade.settlementID = convertToFloat64(results[27])
		cTPTrade.brokerOrderSeq = convertToFloat64(results[28])
		cTPTrade.tradeSource = results[29]
		cTPTrade.commission = convertToFloat64(results[30])
	}
	return cTPTrade
}

//转换成成交回报列表
func convertToCTPTradeRecords(result string)([]TradeRecord){
	var lst []TradeRecord
	strs := strings.Split(result, ",")
	length := len(strs)
	for i := 0; i < length; i++ {
		results := strings.Split(strs[i], ",")
		lst = append(lst, convertToCTPTrade(results[i]))
	}
	return lst
}

//从byte转化成uintptr
func ByteToPtr(n byte) uintptr {
	return uintptr(n)
}

//从uintptr转换成int
func IntToPtr(n int) uintptr {
	return uintptr(n)
}

//从uintptr转换成float64
func FloatToPtr(n float64)uintptr{
	return uintptr(n)
}

//从string转换成uintptr
func StringToUintPtr(val string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(val)))
}

//转换成UTF-8
func u16To8(u16 []byte) string {
	if len(u16)%2 != 0 {
		return ""
	}
	var body bytes.Buffer
	for i := 0; i < len(u16)/2; i++ {
		v := int(u16[2*i]) + int(u16[2*i+1])<<8
		if v <= 127 {
			body.WriteByte(byte(v))
		} else if v <= 2047 {
			a1 := byte(v&63) + 128
			v = v >> 6
			a2 := byte(v&31) + 192
			body.WriteByte(a2)
			body.WriteByte(a1)
 
		} else if v <= 65535 {
			a1 := byte(v&63) + 128
			v = v >> 6
			a2 := byte(v&63) + 128
			v = v >> 6
			a3 := byte(v&15) + 224
			body.WriteByte(a3)
			body.WriteByte(a2)
			body.WriteByte(a1)
		}
	}
	return string(body.Bytes())
}

//从uintptr转换成String
func StringFromPtr(sptr uintptr) (res string) {
	if sptr <= 0 {
		return
	}
	buf := make([]byte, 0)
	for i := 0; *(*byte)(unsafe.Pointer(sptr + uintptr(i))) != 0; i++ {
		buf = append(buf, *(*byte)(unsafe.Pointer(sptr + uintptr(i))))
	}

	return u16To8(buf)
}

//从uintptr转换成Bool
func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

//卖平：多单平仓
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
//exchangeID 交易所ID
//price 价格
//qty 数量
//timeCondition 有效期
//orderRef 附加信息
func ctp_askClose(dll *syscall.DLL, ctpID int, requestID int, code string, exchangeID string, price float64, qty byte, timeCondition int, orderRef string)(int){
	proc,err := dll.FindProc("askClose")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code), StringToUintPtr(exchangeID), FloatToPtr(price), ByteToPtr(qty), IntToPtr(timeCondition), StringToUintPtr(orderRef))
	return int(result)
}

//卖平今仓：平今天的开仓的空单
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
//exchangeID 交易所ID
//price 价格
//qty 数量
//timeCondition 有效期
//orderRef 附加信息
func ctp_askCloseToday(dll *syscall.DLL, ctpID int, requestID int, code string, exchangeID string, price float64, qty byte, timeCondition int, orderRef string)(int){
	proc,err := dll.FindProc("askCloseToday")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code), StringToUintPtr(exchangeID), FloatToPtr(price), ByteToPtr(qty), IntToPtr(timeCondition), StringToUintPtr(orderRef))
	return int(result)
}

//卖开：空单开仓
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
//exchangeID 交易所ID
//price 价格
//qty 数量
//timeCondition 有效期
//orderRef 附加信息
func ctp_askOpen(dll *syscall.DLL, ctpID int, requestID int, code string, exchangeID string, price float64, qty byte, timeCondition int, orderRef string)(int){
	proc,err := dll.FindProc("askOpen")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code), StringToUintPtr(exchangeID), FloatToPtr(price), ByteToPtr(qty), IntToPtr(timeCondition), StringToUintPtr(orderRef))
	return int(result)
}

//买平：空单平仓
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
//exchangeID 交易所ID
//price 价格
//qty 数量
//timeCondition 有效期
//orderRef 附加信息
func ctp_bidClose(dll *syscall.DLL, ctpID int, requestID int, code string, exchangeID string, price float64, qty byte, timeCondition int, orderRef string)(int){
	proc,err := dll.FindProc("bidClose")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code), StringToUintPtr(exchangeID), FloatToPtr(price), ByteToPtr(qty), IntToPtr(timeCondition), StringToUintPtr(orderRef))
	return int(result)
}

//买平今仓：平今天的开仓的空单
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
//exchangeID 交易所ID
//price 价格
//qty 数量
//timeCondition 有效期
//orderRef 附加信息
func ctp_bidCloseToday(dll *syscall.DLL, ctpID int, requestID int, code string, exchangeID string, price float64, qty byte, timeCondition int, orderRef string)(int){
	proc,err := dll.FindProc("bidCloseToday")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code), StringToUintPtr(exchangeID), FloatToPtr(price), ByteToPtr(qty), IntToPtr(timeCondition), StringToUintPtr(orderRef))
	return int(result)
}

//买开：多单开仓
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
//exchangeID 交易所ID
//price 价格
//qty 数量
//timeCondition 有效期
//orderRef 附加信息
func ctp_bidOpen(dll *syscall.DLL, ctpID int, requestID int, code string, exchangeID string, price float64, qty byte, timeCondition int, orderRef string)(int){
	proc,err := dll.FindProc("bidOpen")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code), StringToUintPtr(exchangeID), FloatToPtr(price), ByteToPtr(qty), IntToPtr(timeCondition), StringToUintPtr(orderRef))
	return int(result)
}

//撤单
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//exchangeID 交易所ID
//orderSysID 委托编号
//orderRef 附加信息
func ctp_cancelOrder(dll *syscall.DLL, ctpID int, requestID int, exchangeID string, orderSysID string, orderRef string)(int){
	proc,err := dll.FindProc("cancelOrder")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(exchangeID), StringToUintPtr(orderSysID), StringToUintPtr(orderRef))
	return int(result)
}

//创建交易
//dll 动态链接库
func ctp_create(dll *syscall.DLL)(int){
	proc,err := dll.FindProc("create")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call()
	return int(result)
}

//生成ctp请求编号
//dll 动态链接库
//ctpID ctp编号
func ctp_generateReqID(dll *syscall.DLL, ctpID int)(int){
	proc,err := dll.FindProc("generateReqID")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return int(result)
}

//是否有新的数据
//dll 动态链接库
//ctpID ctp编号
func ctp_hasNewDatas(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("hasNewDatas")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//获取资金账户信息
//dll 动态链接库
//ctpID ctp编号
func ctp_getAccountData(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getAccountData")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取经纪公司ID
//dll 动态链接库
//ctpID ctp编号
func ctp_getBrokerID(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getBrokerID")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取手续费率
//dll 动态链接库
//ctpID ctp编号
//code 代码
func ctp_getCommissionRate(dll *syscall.DLL, ctpID int, code string)(string){
	proc,err := dll.FindProc("getCommissionRate")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), StringToUintPtr(code), ptr)
	return StringFromPtr(ptr)
}

//获取深度市场行情
//dll 动态链接库
//ctpID ctp编号
func ctp_getDepthMarketData(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getDepthMarketData")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取合约数据
//dll 动态链接库
//ctpID ctp编号
func ctp_getInstrumentsData(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getInstrumentsData")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取投资者ID
//dll 动态链接库
//ctpID ctp编号
func ctp_getInvestorID(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getInvestorID")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取保证金率
//dll 动态链接库
//ctpID ctp编号
//code 代码
func ctp_getMarginRate(dll *syscall.DLL, ctpID int, code string)(string){
	proc,err := dll.FindProc("getMarginRate")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), StringToUintPtr(code), ptr)
	return StringFromPtr(ptr)
}

//获取投资者持仓数据
//dll 动态链接库
//ctpID ctp编号
func ctp_getPositionData(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getPositionData")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取投资者持仓明细数据
//dll 动态链接库
//ctpID ctp编号
func ctp_getPositionDetailData(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getPositionDetailData")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取最新的委托回报（上一条）
//dll 动态链接库
//ctpID ctp编号
func ctp_getOrderInfo(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getOrderInfo")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取所有的最新委托回报（今天的所有委托）
//dll 动态链接库
//ctpID ctp编号
func ctp_getOrderInfos(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getOrderInfos")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取sessionID
//dll 动态链接库
//ctpID ctp编号
func ctp_getSessionID(dll *syscall.DLL, ctpID int)(int){
	proc,err := dll.FindProc("getSessionID")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return int(result)
}

//获取结算单信息
//dll 动态链接库
//ctpID ctp编号
func ctp_getSettlementInfo(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getSettlementInfo")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取最新成交记录（上一条） 
//dll 动态链接库
//ctpID ctp编号
func ctp_getTradeRecord(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getTradeRecord")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取最新交易记录（今天的所有交易）
//dll 动态链接库
//ctpID ctp编号
func ctp_getTradeRecords(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getTradeRecords")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取交易日期
//dll 动态链接库
//ctpID ctp编号
func ctp_getTradingDate(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getTradingDate")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//获取交易时间
//dll 动态链接库
//ctpID ctp编号
func ctp_getTradingTime(dll *syscall.DLL, ctpID int)(string){
	proc,err := dll.FindProc("getTradingTime")
	if(err != nil){
		fmt.Println(err)
	}
	var bytes [1024000]byte
	var c *byte = &bytes[0]
	ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))
	proc.Call(IntToPtr(ctpID), ptr)
	return StringFromPtr(ptr)
}

//当天是否已经结算
//dll 动态链接库
//ctpID ctp编号
func ctp_isClearanced(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("isClearanced")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//是否是结算时间
//dll 动态链接库
//ctpID ctp编号
func ctp_isClearanceTime(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("isClearanceTime")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//数据是否准备好
//dll 动态链接库
//ctpID ctp编号
func ctp_isDataOk(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("isDataOk")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//行情数据服务器是否已经登录
//dll 动态链接库
//ctpID ctp编号
func ctp_isMdLogined(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("isMdLogined")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//交易是否已经登录
//dll 动态链接库
//ctpID ctp编号
func ctp_isTdLogined(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("isTdLogined")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//是否是交易时间
//dll 动态链接库
//ctpID ctp编号
func ctp_isTradingTime(dll *syscall.DLL, ctpID int)(bool){
	proc,err := dll.FindProc("isTradingTime")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID))
	return GoBool(result)
}

//是否是精确交易时间(去掉集合竞价时间和休息时间)
//dll 动态链接库
//ctpID ctp编号
//code 代码
func ctp_isTradingTimeExact(dll *syscall.DLL, ctpID int, code string)(bool){
	proc,err := dll.FindProc("isTradingTimeExact")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), StringToUintPtr(code))
	return GoBool(result)
}

//请求手续费率
//dll 动态链接库
//ctpID ctp编号
//code 代码
//requestID 请求ID
func ctp_reqCommissionRate(dll *syscall.DLL, ctpID int, code string, requestID int)(int){
	proc,err := dll.FindProc("reqCommissionRate")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), StringToUintPtr(code), IntToPtr(requestID))
	return int(result)
}

//请求确认结算信息
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
func ctp_reqQrySettlementInfoConfirm(dll *syscall.DLL, ctpID int, requestID int)(int){
	proc,err := dll.FindProc("reqQrySettlementInfoConfirm")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID))
	return int(result)
}

//请求确认结算信息
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//tradingDay 交易日
func ctp_reqQrySettlementInfo(dll *syscall.DLL, ctpID int, requestID int, tradingDay string)(int){
	proc,err := dll.FindProc("reqQrySettlementInfo")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(tradingDay))
	return int(result)
}

//请求保证金率
//dll 动态链接库
//ctpID ctp编号
//code 代码
//requestID 请求ID
func ctp_reqMarginRate(dll *syscall.DLL, ctpID int, code string, requestID int)(int){
	proc,err := dll.FindProc("reqMarginRate")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), StringToUintPtr(code), IntToPtr(requestID))
	return int(result)
}

//启动创建的连接(在create后执行)
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//appID APPID
//authCode 用户ID
//mdServer 行情地址
//tdServer 交易地址
//brokerID 机构号
//investorID 投资者账号
//password 密码
func ctp_start(dll *syscall.DLL, ctpID int, requestID int, appID string, authCode string, mdServer string, tdServer string, brokerID string, investorID string, password string)(int){
	proc,err := dll.FindProc("start")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(appID), StringToUintPtr(authCode), StringToUintPtr(mdServer), StringToUintPtr(tdServer), StringToUintPtr(brokerID), StringToUintPtr(investorID), StringToUintPtr(password))
	return int(result)
}

//订阅多个合约的行情数据
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
func ctp_subMarketDatas(dll *syscall.DLL, ctpID int, requestID int, code string)(int){
	proc,err := dll.FindProc("subMarketDatas")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code))
	return int(result)
}

//取消订阅多个合约的行情数据
//dll 动态链接库
//ctpID ctp编号
//requestID 请求ID
//code 代码
func ctp_unSubMarketDatas(dll *syscall.DLL, ctpID int, requestID int, code string)(int){
	proc,err := dll.FindProc("unSubMarketDatas")
	if(err != nil){
		fmt.Println(err)
	}
	result,_,_ := proc.Call(IntToPtr(ctpID), IntToPtr(requestID), StringToUintPtr(code))
	return int(result)
}

func main() {
	dll,err := syscall.LoadDLL("iCTP.dll")
	fmt.Println(dll)
	fmt.Println(err)
    ctpID := ctp_create(dll)
	reqID := ctp_generateReqID(dll, ctpID)
	ctp_start(dll, ctpID, reqID, "simnow_client_test", "0000000000000000", "180.168.146.187:10212", "180.168.146.187:10202", "9999", "021739", "123456")
	for{
		if(ctp_isDataOk(dll, ctpID)){
			break
		}else{
			time.Sleep(1 * time.Second)
		}
	}
	//获取码表
	recvData := ctp_getInstrumentsData(dll, ctpID)
	instrumentsData := convertToCTPInstrumentDatas(recvData)
	fmt.Println(instrumentsData)
	//获取委托回报历史
	recvData = ctp_getOrderInfos(dll, ctpID)
	orderInfos := convertToCTPOrderList(recvData)
	fmt.Println(orderInfos)
	//获取成交回报历史
	recvData = ctp_getTradeRecords(dll, ctpID)
	tradeRecords := convertToCTPTradeRecords(recvData)
	fmt.Println(tradeRecords)
	ctp_subMarketDatas(dll, ctpID, reqID, "cu2304,rb2304")
	for{
		if(ctp_hasNewDatas(dll, ctpID)){
			//获取最新数据
			recvData = ctp_getDepthMarketData(dll, ctpID)
			if(len(recvData) > 0){
				depthMarketData := convertToCTPDepthMarketData(recvData)
				fmt.Println(depthMarketData)
				continue
			}
			//获取资金账号
			recvData = ctp_getAccountData(dll, ctpID)
			if(len(recvData) > 0){
				accountData := convertToCTPAccountData(recvData)
				fmt.Println(accountData)
				continue
			}
			//获取委托回报
			recvData = ctp_getOrderInfo(dll, ctpID)
			if(len(recvData) > 0){
				orderInfo := convertToCTPOrder(recvData)
				fmt.Println(orderInfo)
				continue
			}
			//获取成交回报
			recvData = ctp_getTradeRecord(dll, ctpID)
			if(len(recvData) > 0){
				tradeRecord := convertToCTPTrade(recvData)
				fmt.Println(tradeRecord)
				continue
			}
			//获取持仓
			recvData = ctp_getPositionData(dll, ctpID)
			if(len(recvData) > 0){
				positionData := convertToCTPInvestorPosition(recvData)
				fmt.Println(positionData)
				continue
			}
			//获取持仓明细
			recvData = ctp_getPositionDetailData(dll, ctpID)
			if(len(recvData) > 0){
				positionDetailData := convertToCTPInvestorPositionDetail(recvData)
				fmt.Println(positionDetailData)
				continue
			}
		}
		time.Sleep(1 * time.Second)
	}
}