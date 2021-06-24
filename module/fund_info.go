/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 21:27:32
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-24 21:38:33
 */
package module

type FundInfo struct {
	Code string
	Name string
}

type FundDaysInfo struct {
	Date           string
	AccumulatedNet float64
}

// type FundInfoGetter interface {
// 	GetFundDaysInfo() (days_info []*FundDaysInfo, err error)
// }
