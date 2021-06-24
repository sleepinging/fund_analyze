/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2021-06-24 20:04:34
 * @LastEditors: taowentao
 * @LastEditTime: 2021-06-24 20:43:45
 */
package util

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, headers map[string]string) (res string, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	for k, v := range headers {
		reqest.Header.Add(k, v)
	}

	resp, err := client.Do(reqest)
	// resp, err := http.Get(url.QueryEscape(downURL))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	res = string(data)
	return
}
