package wsocket

import (
	"encoding/json"
)

type JsonDataIn struct {
	DataType string      `json:"dataType"`
	Value    interface{} `json:"value"`
}
type JsonDataOut struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (conn *Connection) MessageHandler(data []byte) {
	var JsonMessage JsonDataIn
	err := json.Unmarshal(data, &JsonMessage)
	conn.errorHandler(err)
	// todo:处理消息

}
func (conn *Connection) WriteMessageHandler(data interface{}) {
	jsonData := JsonDataOut{Code: "200", Msg: "ok", Value: data}
	jsonDataByte, err := json.Marshal(jsonData)
	conn.errorHandler(err)
	conn.OutChan <- jsonDataByte
}

func (conn *Connection) errorHandler(err error) {
	if err != nil {
		jsonData := JsonDataOut{Code: "201", Msg: "数据不正确", Value: nil}
		jsonDataByte, _ := json.Marshal(jsonData)
		conn.OutChan <- jsonDataByte
	}
}
