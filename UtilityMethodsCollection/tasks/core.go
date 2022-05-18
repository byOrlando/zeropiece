package tasks

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

type Tasks struct {
	Server     *gocron.Scheduler
	ServerType chan bool
}

func (t *Tasks) Init() {
	{
		t.Server = gocron.NewScheduler()
	}
	{
		t.Server.Every(6).Seconds().Do(SuperWang)
	}

}
func InitAndStart() {
	s := Tasks{}
	s.Init()
	s.ServerType = s.Server.Start()
	// 正确关闭
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()
}
func (t *Tasks) Stop() {
	close(t.ServerType)
}
func (t *Tasks) Remove(funcName interface{}) {
	t.Server.Remove(funcName)
}

func (t *Tasks) Clear() {
	t.Server.Clear()
}
