/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 21:27:32
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-26 18:34:04
 */
package module

import (
	"errors"
	"time"
)

//卖出手续费
type SaleFee struct {
	Level map[int32]float64
}

type FundInfo struct {
	Code string
	// Name        string
	PurchaseFee float64  //买入手续费
	SaleFee     *SaleFee //卖出手续费
	DaysInfo    *FundDaysInfo
}

type FundDaysInfo struct {
	DaysInfo map[time.Time]*FundDayInfo
	//cache
	StartDay time.Time
}

type FundDayInfo struct {
	Date           time.Time
	AccumulatedNet float64
}

//基金当日信息,如果不存在就前一天
func (days_info *FundDaysInfo) GetFundDayInfo(day time.Time) (day_info *FundDayInfo, err error) {
	day_info, ok := days_info.DaysInfo[day]
	for ; !ok; day_info, ok = days_info.DaysInfo[day] {
		if day.Before(days_info.StartDay) {
			err = errors.New("net find day info")
			return
		}
		// log.Println("can't find", day.Format("2006-1-2"))
		day = day.AddDate(0, 0, -1)
	}
	return
}

// type FundInfoGetter interface {
// 	GetFundDaysInfo() (days_info []*FundDaysInfo, err error)
// }
