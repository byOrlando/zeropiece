package douyin_Data_Out

type DouYinOutData struct {
	Code     string      `json:"code"`
	Msg      string      `json:"msg"`
	DataType string      `json:"dataType"`
	Value    interface{} `json:"value"`
}
