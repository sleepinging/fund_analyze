/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 21:27:32
 * @LastEditors: taowentao
 * @LastEditTime: 2021-07-25 16:11:56
 */
package module

import (
	"errors"
	"fund_analyze/util"
	"time"

	mapset "github.com/deckarep/golang-set"
)

//卖出手续费
type SaleFee struct {
	Level mapset.Set
}

type FundInfo struct {
	Code string
	// Name        string
	PurchaseFee float64  //买入手续费
	SaleFee     *SaleFee //卖出手续费
	DaysInfo    *FundDaysInfo
}

type FundDaysInfo struct {
	DaysInfo map[uint32]*FundDayInfo
	//cache
	StartDay uint32
}

type FundDayInfo struct {
	Date           uint32
	AccumulatedNet float64
}

//基金当日信息,如果不存在就前一天
func (days_info *FundDaysInfo) GetFundDayInfoOrBefore(day time.Time) (day_info *FundDayInfo, err error) {
	day_info, ok := days_info.DaysInfo[util.ParseTimeToDays(day)]
	for ; !ok; day_info, ok = days_info.DaysInfo[util.ParseTimeToDays(day)] {
		if day.Before(util.ParseDaysToTime(days_info.StartDay)) {
			err = errors.New("net find day info")
			return
		}
		// log.Println("can't find", day.Format("2006-1-2"))
		day = day.AddDate(0, 0, -1)
	}
	return
}

//基金当日信息,如果不存在就后一天
func (days_info *FundDaysInfo) GetFundDayInfoOrAfter(day time.Time) (day_info *FundDayInfo, err error) {
	day_info, ok := days_info.DaysInfo[util.ParseTimeToDays(day)]
	for ; !ok; day_info, ok = days_info.DaysInfo[util.ParseTimeToDays(day)] {
		if day.Before(util.ParseDaysToTime(days_info.StartDay)) {
			err = errors.New("net find day info")
			return
		}
		// log.Println("can't find", day.Format("2006-1-2"))
		day = day.AddDate(0, 0, 1)
	}
	return
}

// type FundInfoGetter interface {
// 	GetFundDaysInfo() (days_info []*FundDaysInfo, err error)
// }
