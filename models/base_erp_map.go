package models

import (
	orm "gin_test/context"
	"time"
)

// BaseErpMap base_erp_map关系
type BaseErpMap struct {
	MapID      int64
	MapType    string
	MapUser    string
	CompanyID  int64
	DhbID      int64
	ErpID      string
	UpdateTime time.Time
	CreateTime time.Time
}

// BaseErpMaps 分类列表
func (c *BaseErpMap) BaseErpMaps() (maps []BaseErpMap, err error) {
	if err = orm.DataDB.Where("company_id = 1253").Find(&maps).Error; err != nil {
		return
	}
	return
}
