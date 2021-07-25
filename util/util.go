/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-07-25 16:06:33
 * @LastEditors: taowentao
 * @LastEditTime: 2021-07-25 16:11:29
 */

package util

import "time"

var tSince time.Time

func Init() {
	tSince, _ = time.Parse("2006-1-2", "1900-01-01")
}
