package dao

type CityCode struct {
	// 地区代码
	DistrictCode string `gorm:"primary_key" json:"districtCode"`
	// 省份
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
	// 区县
	District string `json:"district"`
}
