package utils

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"zeropiece/common"
)

func GetLocation(ip string) interface{} {
	if ip == "127.0.0.1" || ip == "localhost" || ip == "::1" {
		return "内部IP地址"
	}
	resp, err := http.Get(common.CONFIG.BaiduMap.IpLocationUrl + common.CONFIG.BaiduMap.AK + "&ip=" + ip + "&coor=bd09ll")
	if err != nil {
		common.LOG.Warn(err.Error())

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)

	address := gjson.Get(string(s), "content.address").String()
	if err != nil {
		return "未知位置"
	}
	if address == "" {
		return "未知位置"
	}

	return address
}
