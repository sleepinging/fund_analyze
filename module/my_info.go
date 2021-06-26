/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-26 17:23:47
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-26 18:38:39
 */

package module

import "time"

type FundShare struct {
	PurchaseDay   time.Time
	PurchaseShare float64
}

type MyFund struct {
	Fund       *FundInfo    //基金信息
	FundShares []*FundShare //基金份额信息
	//cache
	TotalShare float64 //基金总份额
	TotolCost  float64 //基金总成本
}

type MyInfo struct {
	MoneyAvailable float64 //可用金额
	MyFunds        []*MyFund
	//cache
	TotolCosts float64 //基金总成本
}

//基金收益率
func (mf *MyFund) Yield(day time.Time) (yield float64, err error) {
	//(基金价值-成本) / 成本
	//当日净值
	day_info, err := mf.Fund.DaysInfo.GetFundDayInfo(day)
	if err != nil {
		return
	}
	yield = (day_info.AccumulatedNet*mf.TotalShare - mf.TotolCost) / mf.TotolCost
	return
}

//基金总金额
func (mi *MyInfo) CalcTotalAmount(day time.Time) (amount float64) {
	return
}

//买基金
func (mi *MyInfo) PurchaseFund(day time.Time, fund_info *FundInfo, cost float64) (err error) {
	return
}

//卖基金
func (mi *MyInfo) SaleFund(day time.Time, fund_info *FundInfo, share float64) (err error) {
	return
}
