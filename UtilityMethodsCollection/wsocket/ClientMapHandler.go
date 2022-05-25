package wsocket

func WSSET(name string, conn *Connection) {
	if ClientMapHandler[name] == nil {
		ClientMapHandler[name] = []*Connection{
			conn,
		}
	} else {
		ClientMapHandler[name] = append(ClientMapHandler[name], conn)
	}
}
func DELWS(name string, conn *Connection) {
	if ClientMapHandler[name] == nil {
		return
	}
	for i, v := range ClientMapHandler[name] {
		if v == conn {
			ClientMapHandler[name] = append(ClientMapHandler[name][:i], ClientMapHandler[name][i+1:]...)
			return
		}
	}
}
func WSSend(name string, data []byte) {
	if ClientMapHandler[name] == nil {
		return
	}
	for _, v := range ClientMapHandler[name] {
		v.OutChan <- data
	}
}

// ClientMapHandler 映射关系表
var ClientMapHandler = make(map[string][]*Connection)
