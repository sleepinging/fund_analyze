/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:03:02
 * @LastEditors: taowentao
 * @LastEditTime: 2021-07-25 16:08:36
 */

package main

import (
	"fmt"
	"fund_analyze/ttjj"
	"fund_analyze/util"
	"sort"
	"time"
)

//每天的操作
func day_op(day time.Time, acc_net float64) {
	//建仓|买入
	//卖出
}

func main() {
	util.Init()
	fmt.Println("day,value")
	days_info, err := ttjj.GetFundDaysInfo("001593")
	if err != nil {
		fmt.Print(err)
	}
	// day, _ := time.ParseInLocation("2006-1-2", "2020-07-01", time.Local)
	// info, err := days_info.GetFundDayInfo(day)
	// fmt.Println(info)

	//排序
	var times []int
	for k := range days_info.DaysInfo {
		times = append(times, int(k))
	}
	sort.Ints(times)
	for _, t := range times {
		v := days_info.DaysInfo[uint32(t)]
		day := util.ParseDaysToTime(v.Date)
		fmt.Println(day.Format("2006-1-2"), ",", v.AccumulatedNet)
		day_op(util.ParseDaysToTime(v.Date), v.AccumulatedNet)
	}
}
