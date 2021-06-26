/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:43:29
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-26 18:24:20
 */
package ttjj

import (
	"fmt"
	"fund_analyze/module"
	"fund_analyze/util"
	"log"
	"strconv"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

const refer = "http://fundf10.eastmoney.com/"
const url_fmt = "http://api.fund.eastmoney.com/f10/lsjz?callback=jQuery1830829753053641421_1624535738293&fundCode=%s&pageIndex=1&pageSize=9999999&startDate=&endDate=&_=%d238"

func get_fund_json(code string) (res string, err error) {
	headers := map[string]string{
		"Referer": refer,
	}
	ts := time.Now().Unix()
	url := fmt.Sprintf(url_fmt, code, ts)
	res, err = util.HttpGet(url, headers)
	return
}

func GetFundDaysInfo(code string) (days_info *module.FundDaysInfo, err error) {
	res, err := get_fund_json(code)
	if err != nil {
		return
	}
	start := len("jQuery1830829753053641421_1624535738293(")
	end := len(res) - 1
	js, err := simplejson.NewJson([]byte(res[start:end]))
	if err != nil {
		return
	}
	vs, err := js.Get("Data").Get("LSJZList").Array()
	if err != nil {
		return
	}
	if len(vs) == 0 {
		return
	}
	days_info = new(module.FundDaysInfo)
	days_info.DaysInfo = make(map[time.Time]*module.FundDayInfo, len(vs))
	for i, day_info_obj := range vs {
		if day_info, ok := day_info_obj.(map[string]interface{}); ok {
			info := new(module.FundDayInfo)

			if accumulated_net, ok := day_info["LJJZ"].(string); ok {
				info.AccumulatedNet, err = strconv.ParseFloat(accumulated_net, 64)
				if err != nil {
					log.Panicln(err)
					continue
				}
			}

			if date, ok := day_info["FSRQ"].(string); ok {
				info.Date, err = time.ParseInLocation("2006-1-2", date, time.Local)
				if err != nil {
					log.Panicln(err)
					continue
				}
				if i == 0 || days_info.StartDay.After(info.Date) {
					days_info.StartDay = info.Date
				}
			}

			days_info.DaysInfo[info.Date] = info
		}
	}
	return
}
