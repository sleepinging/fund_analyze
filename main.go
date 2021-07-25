/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:03:02
 * @LastEditors: taowentao
 * @LastEditTime: 2021-07-25 17:04:09
 */

package main

import (
	"fmt"
	"fund_analyze/module"
	"fund_analyze/ttjj"
	"fund_analyze/util"
	"sort"
	"time"
)

//我的信息
var my_info *module.MyInfo

//005913
var f_fund *module.FundInfo

func init() {
	util.Init()
	my_info = new(module.MyInfo)
	my_info.Init(300000)

	f_fund = new(module.FundInfo)
	saleFee := []module.SaleFee{
		{Day: 7, Fee: 1.5 / 100},
	}
	f_fund.Init("001593", 0.15/100, saleFee)
}

//每天的操作
func day_op(day time.Time, acc_net float64) {
	//计算收益率
	//2015.07.28买入
	bt, _ := time.Parse("2006-01-02", "2015-07-28")
	if day.Before(bt) {
		return
	}
	if len(my_info.MyFunds) == 0 {
		my_info.PurchaseFund(day, f_fund, 100)
	}
	fund := my_info.MyFunds["001593"]
	y, _ := fund.Yield(day)
	fmt.Printf("date:%s, acc: %f, yield: %.2f%%\n", day.Format("2006-1-2"), acc_net, y*100)
}

func main() {
	fmt.Println("day,value")
	days_info, err := ttjj.GetFundDaysInfo("001593")
	if err != nil {
		fmt.Print(err)
	}
	f_fund.DaysInfo = days_info
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
		// fmt.Println(day.Format("2006-1-2"), ",", v.AccumulatedNet)
		day_op(day, v.AccumulatedNet)
	}
}
