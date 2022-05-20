package wsocket

type ClientMapHandlers struct {
	Name          string
	ClientMap     *Connection
	ClientMapList []*Connection
}

func (c *ClientMapHandlers) Set(control *Connection, name string) {
	if c.ClientMap == nil && c.ClientMapList == nil {
		c.ClientMap = control
		c.Name = name

	} else {
		if c.ClientMapList == nil {
			c.ClientMapList = []*Connection{
				c.ClientMap,
				control,
			}
			c.ClientMap = nil
		} else {
			c.ClientMapList = append(c.ClientMapList, control)
		}

	}
}
func (c *ClientMapHandlers) Write(message string) {
	if c.ClientMap != nil {
		c.ClientMap.OutChan <- []byte(message)
	}
	if c.ClientMapList != nil {
		for _, v := range c.ClientMapList {
			v.OutChan <- []byte(message)
		}
	}
}

func (c *ClientMapHandlers) Remove(control *Connection) {
	if c.ClientMap == control {
		c.ClientMap = nil
	}
	if c.ClientMapList != nil {
		for i, v := range c.ClientMapList {
			if v == control {
				c.ClientMapList = append(c.ClientMapList[:i], c.ClientMapList[i+1:]...)
				break
			}
		}
	}
	if c.ClientMapList == nil && c.ClientMap == nil {
		delete(ClientMapHandler, c.Name)
	}
}

// ClientMapHandler 映射关系表
var ClientMapHandler = make(map[string]*ClientMapHandlers)
