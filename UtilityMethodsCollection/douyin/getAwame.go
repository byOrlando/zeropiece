package douyin

import (
	"fmt"
	"github.com/tidwall/gjson"
	"sync"
	"time"
	"zeropiece/UtilityMethodsCollection/push"
	"zeropiece/UtilityMethodsCollection/requests"
	"zeropiece/common"
	"zeropiece/dao"
)

var waitGroup sync.WaitGroup

// requestNetworkData Request data from network
func requestNetworkData(uid string, name string) gjson.Result {
	p1 := requests.ClientOption{}
	p1.Url = "https://www.iesdouyin.com/web/api/v2/aweme/post/"
	p1.Params = map[string]string{
		"sec_uid":    uid,
		"count":      "21",
		"max_cursor": "0",
	}
	p1.Headers = map[string]string{
		"DouyinUser-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
	}
	str, err := p1.Get()
	if err != nil {
		fmt.Println("获取Awame失败", err)
		WSError(fmt.Sprintf("获取【%s】数据失败:%s", name, err), err)
		// Error retry
		str, err = p1.Get()
	}
	// Get new work information list data
	return gjson.Parse(str)
}

func dataCleaning(data gjson.Result, uid string) (works []dao.DouyinWorks) {
	for _, v := range data.Get("aweme_list").Array() {
		works = append(works, dao.DouyinWorks{
			UID:       uid,
			Title:     v.Get("desc").String(),
			WorkID:    v.Get("aweme_id").String(),
			Name:      v.Get("author.nickname").String(),
			CommCount: int(v.Get("statistics.comment_count").Int()),
		})
	}
	return
}

// getAllOldData Get information about all old works
func getAllOldData(uid string) (worksIdList []interface{}) {
	var err error
	works := dao.DouyinWorks{UID: uid}
	worksIdList, err = works.FindWorksIdList()
	if err != nil {
		fmt.Println("获取数据库数据失败")
		WSError("获取数据库数据失败", err)
	}
	return
}

// CheckDataRepeat Returns new work data
func checkDataRepeat(works []dao.DouyinWorks, worksIdList []interface{}) (newWorks []dao.DouyinWorks) {
	work := dao.DouyinWorks{}
	for _, v := range works {
		err := work.CheckWorkIdInList(v.WorkID, worksIdList)
		if err != nil {
			newWorks = append(newWorks, v)
		}
	}
	return
}

// SaveAllNewData Save new work data
func saveAllNewData(works []dao.DouyinWorks) {
	for _, v := range works {
		err := v.Create()
		if err != nil {
			fmt.Println("创建失败")
			WSError("创建失败", err)
			break
		}
	}
}

// PushNewWorksToIphone Push data of new works
func pushNewWorksToIphone(works []dao.DouyinWorks) {
	for _, v := range works {
		if v.CommCount < 15 {
			var option push.PushOption
			option.Init()
			option.Text = v.Title
			option.AwameID = v.WorkID
			option.Title = v.Name
			option.Push()
			WSSuccess("发现新作品", nil)
		}
	}
}

// MainProcessingMethod Main processing method
func mainProcessingMethod(uid string, name string) {
	defer common.Error()()
	// Get new work information list data
	data := requestNetworkData(uid, name)
	works := dataCleaning(data, uid)
	// Get information about all old works
	worksIdList := getAllOldData(uid)
	// Returns new work data
	newWorks := checkDataRepeat(works, worksIdList)
	// Save new work data
	saveAllNewData(newWorks)
	// Push data of new works
	pushNewWorksToIphone(newWorks)
	// complete
	waitGroup.Done()
}

var NewWorkInspections = 0

// CheckAllUsersWorkCycle The entrance of the program
func CheckAllUsersWorkCycle() {
	defer common.Error()()
	for {
		user := dao.DouyinUser{}
		UserIdList, err := user.FindAllUserIdList()
		if err != nil {
			common.LOG.Error("获取用户id列表失败")
			WSError("获取用户id列表失败", err)
		}
		for _, UserData := range UserIdList {
			waitGroup.Add(1)
			go mainProcessingMethod(UserData.UID, UserData.UserName)
		}

		waitGroup.Wait()
		NewWorkInspections += 1
		WSSuccess(fmt.Sprintf("正在执行第-%d-次检查", NewWorkInspections), nil)
		time.Sleep(time.Second * 3)

	}

}
