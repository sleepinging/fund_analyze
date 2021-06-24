/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:03:02
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-24 21:40:03
 */

package main

import (
	"fmt"
	"fund_analyze/ttjj"
)

func main() {
	fmt.Println("start")
	days_info, err := ttjj.GetFundDaysInfo("001593")
	if err != nil {
		fmt.Print(err)
	}
	for _, v := range days_info {
		fmt.Println(v.Date, v.AccumulatedNet)
	}
}
