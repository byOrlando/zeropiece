package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"zeropiece/middleware"
)

func GetWeather(c *gin.Context) {
	// 获取天气
	resp, err := http.Get("http://www.weather.com.cn/data/cityinfo/101060505.html")
	if err != nil {
		middleware.ResponseFail(c, 201, "获取天气失败")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var DATA interface{}
	err = json.Unmarshal(body, &DATA)
	if err != nil {
		middleware.ResponseFail(c, 201, "获取天气失败")
	}

	middleware.ResponseSucc(c, "获取天气成功", DATA)
}
