/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-26 17:23:47
 * @LastEditors: taowentao
 * @LastEditTime: 2021-07-25 16:58:53
 */

package module

import (
	"fmt"
	"log"
	"time"
)

type MyFund struct {
	FundInfo   *FundInfo          //基金信息
	FundShares map[uint32]float64 //基金份额信息 [确认时间]份额
	//cache
	TotalShare   float64 //基金总份额
	TotolCost    float64 //基金总花费
	TotolBenefit float64 // 基金总收益
}

type MyInfo struct {
	MoneyAvailable float64            //可用金额
	MyFunds        map[string]*MyFund //基金代码[持有基金信息]
	//cache
	TotolCosts float64 //基金总成本
}

func (mi *MyInfo) Init(money float64) {
	mi.MyFunds = make(map[string]*MyFund)
	mi.MoneyAvailable = money
}

//基金收益率
func (mf *MyFund) Yield(day time.Time) (yield float64, err error) {
	//(基金价值-成本) / 成本
	//当日净值
	day_info, err := mf.FundInfo.DaysInfo.GetFundDayInfoOrBefore(day)
	if err != nil {
		return
	}
	yield = (day_info.AccumulatedNet*mf.TotalShare - mf.TotolCost) / mf.TotolCost
	return
}

//基金总金额
func (mi *MyInfo) CalcTotalAmount(day time.Time) (amount float64, err error) {
	for _, my_fund := range mi.MyFunds {
		//当日净值*总份额
		var day_info *FundDayInfo
		day_info, err = my_fund.FundInfo.DaysInfo.GetFundDayInfoOrBefore(day)
		if err != nil {
			continue
		}
		amount += day_info.AccumulatedNet * my_fund.TotalShare
	}
	return
}

//买基金
func (mi *MyInfo) PurchaseFund(day time.Time, fund_info *FundInfo, cost float64) (err error) {
	//找到我当前购买的基金
	my_fund, ok := mi.MyFunds[fund_info.Code]
	if !ok {
		my_fund = new(MyFund)
		my_fund.FundInfo = fund_info
		my_fund.FundShares = make(map[uint32]float64)
		mi.MyFunds[fund_info.Code] = my_fund
	}
	//计算当日净值
	day_info, err := fund_info.DaysInfo.GetFundDayInfoOrAfter(day)
	if err != nil {
		return
	}
	// 获取确认日期
	confirm_day, err := fund_info.DaysInfo.GetFundDayInfoOrAfter(day)
	if err != nil {
		return
	}
	day_net := cost * (1 - fund_info.PurchaseFee) / day_info.AccumulatedNet
	my_fund.FundShares[confirm_day.Date] = day_net
	my_fund.TotalShare += day_net
	my_fund.TotolCost += cost

	mi.TotolCosts += cost
	return
}

//卖基金, share < 0 全部卖出
func (mi *MyInfo) SaleFund(day time.Time, fund_info *FundInfo, share float64) (money float64, err error) {
	my_fund, ok := mi.MyFunds[fund_info.Code]
	if !ok {
		err = fmt.Errorf("not find fund %s", fund_info.Code)
		return
	}
	if share < 0 {
		share = my_fund.TotalShare
	}
	if share > my_fund.TotalShare {
		log.Println("share not enough")
		share = my_fund.TotalShare
	}
	//计算当日净值
	day_info, err := fund_info.DaysInfo.GetFundDayInfoOrAfter(day)
	if err != nil {
		return
	}
	// TODO: 计算手续费

	money = day_info.AccumulatedNet * share

	my_fund.TotalShare -= share
	my_fund.TotolBenefit += money
	return
}
