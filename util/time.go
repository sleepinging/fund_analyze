/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-07-25 15:52:13
 * @LastEditors: taowentao
 * @LastEditTime: 2021-07-25 16:11:37
 */
package util

import "time"

func ParseTimeToDays(t time.Time) uint32 {
	d := t.Sub(tSince)
	return uint32(d.Hours()) / 24
}

func ParseDaysToTime(day uint32) time.Time {
	t := tSince.AddDate(0, 0, int(day))
	return t
}
