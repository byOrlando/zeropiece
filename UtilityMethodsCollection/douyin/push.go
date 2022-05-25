package douyin

import (
	"encoding/json"
	"zeropiece/UtilityMethodsCollection/wsocket"
	"zeropiece/structData/douyin_Data_Out"
)

// WSError 抖音错误消息
func WSError(msg string, value interface{}) {

	if ws := wsocket.ClientMapHandler["DouYin"]; ws != nil {
		JsonData, _ := json.Marshal(douyin_Data_Out.DouYinOutData{
			Code:     "200",
			Msg:      msg,
			DataType: "error",
			Value:    value,
		})

		wsocket.WSSend("DouYin", JsonData)
	}

}

// WSSuccess 抖音成功消息
func WSSuccess(msg string, value interface{}) {

	if ws := wsocket.ClientMapHandler["DouYin"]; ws != nil {
		JsonData, _ := json.Marshal(douyin_Data_Out.DouYinOutData{
			Code:     "200",
			Msg:      msg,
			DataType: "ok",
			Value:    value,
		})
		wsocket.WSSend("DouYin", JsonData)
	}

}
