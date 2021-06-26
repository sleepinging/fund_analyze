/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:03:02
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-26 17:13:17
 */

package main

import (
	"fmt"
	"fund_analyze/ttjj"
	"log"
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
	for _, v := range days_info {
		fmt.Println(v.Date, ",", v.AccumulatedNet)
		t, err := time.ParseInLocation("2006-1-2", v.Date, time.Local)
		if err != nil {
			log.Println(err)
			continue
		}
		day_op(t, v.AccumulatedNet)
	}
}
