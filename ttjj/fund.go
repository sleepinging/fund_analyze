/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:43:29
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-24 21:29:08
 */
package ttjj

import (
	"fmt"
	"fund_analyze/module"
	"fund_analyze/util"
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

func GetFundDaysInfo(code string) (days_info []*module.FundDaysInfo, err error) {
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
	for _, day_info_obj := range vs {
		if day_info, ok := day_info_obj.(map[string]interface{}); ok {
			info := new(module.FundDaysInfo)

			if accumulated_net, ok := day_info["LJJZ"].(string); ok {
				info.AccumulatedNet, err = strconv.ParseFloat(accumulated_net, 64)
				if err != nil {
					continue
				}
			}

			if date, ok := day_info["FSRQ"].(string); ok {
				info.Date = date
			}

			days_info = append(days_info, info)
		}
	}
	return
}
