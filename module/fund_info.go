/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 21:27:32
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-26 17:22:04
 */
package module

//卖出手续费
type SaleFee struct {
	Level map[int32]float64
}

type FundInfo struct {
	Code string
	// Name        string
	PurchaseFee float64  //买入手续费
	SaleFee     *SaleFee //卖出手续费
	DaysInfo    *FundDayInfo
}

type FundDayInfo struct {
	Date           string
	AccumulatedNet float64
}

// type FundInfoGetter interface {
// 	GetFundDaysInfo() (days_info []*FundDaysInfo, err error)
// }
