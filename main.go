/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:03:02
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-26 18:36:20
 */

package main

import (
	"fmt"
	"fund_analyze/ttjj"
	"time"
)

//每天的操作
func day_op(day time.Time, acc_net float64) {
	//建仓|买入
	//卖出
}

func main() {
	fmt.Println("day,value")
	days_info, err := ttjj.GetFundDaysInfo("001593")
	if err != nil {
		fmt.Print(err)
	}
	// day, _ := time.ParseInLocation("2006-1-2", "2020-07-01", time.Local)
	// info, err := days_info.GetFundDayInfo(day)
	// fmt.Println(info)
	for k, v := range days_info.DaysInfo {
		fmt.Println(k.Format("2006-1-2"), ",", v.AccumulatedNet)
		day_op(v.Date, v.AccumulatedNet)
	}
}
