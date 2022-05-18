package wsocket

import (
	"encoding/json"
	"zeropiece/structData/douyin_Data_Out"
)

// WSDouYinError 抖音错误消息
func (conn *Connection) WSDouYinError(msg string, value interface{}) {
	JsonData, _ := json.Marshal(douyin_Data_Out.DouyinOutData{
		Code:     "200",
		Msg:      msg,
		DataType: "error",
		Value:    value,
	})
	conn.OutChan <- JsonData
}

// WSDouYinSuccess 抖音成功消息
func (conn *Connection) WSDouYinSuccess(msg string, value interface{}) {
	JsonData, _ := json.Marshal(douyin_Data_Out.DouyinOutData{
		Code:     "200",
		Msg:      msg,
		DataType: "ok",
		Value:    value,
	})
	conn.OutChan <- JsonData

}

func CycleToSendSuccess(message string) {
	var ws = &ClientMap
	if len(*ws) > 0 {
		for _, v := range *ws {
			v.WSDouYinSuccess(message, nil)
		}
	}
}

func CycleToSendError(message string) {
	var ws = &ClientMap
	if len(*ws) > 0 {
		for _, v := range *ws {
			v.WSDouYinError(message, nil)
		}
	}
}
