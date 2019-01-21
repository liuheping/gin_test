package models

import (
	orm "gin_test/context"
	"time"
)

// BaseCategory 商品分类
type BaseCategory struct {
	CategoryID    int64
	CompanyID     int64
	ParentID      int64
	OrderNum      int64
	CategoryPnum  string
	CategoryNum   string
	CategoryName  string
	LevelNum      int64
	CategoryCount int64
	GoodsCount    int64
	IsDefault     string
	UpdateDate    time.Time
}

// Site 输出分类
type Site struct {
	SiteID   int64
	SiteName string
	ErpID    string
}

// GetCategories 获取分类列表
func (c *BaseCategory) GetCategories() (categories []BaseCategory, err error) {
	if err = orm.DataDB.Where("company_id = 1253").Find(&categories).Error; err != nil {
		return
	}
	return
}

// GetSite 获取分类列表
func (c *BaseCategory) GetSite() ([]Site, error) {
	var cat []BaseCategory
	var erpmap []BaseErpMap

	err := orm.DataDB.Where("company_id = 1253").Find(&cat).Error

	if err != nil {
		return nil, err
	}

	if len(cat) == 0 {
		return nil, nil
	}

	err = orm.DataDB.Where("company_id = 1253 and map_type = 'goodscategory'").Select("dhb_id,erp_id").Find(&erpmap).Error

	if err != nil {
		return nil, err
	}

	sh := make(map[int64]string)
	for _, n := range erpmap {
		sh[n.DhbID] = n.ErpID
	}

	s := make([]Site, len(cat))

	for x, y := range cat {
		s[x].SiteID = y.CategoryID
		s[x].SiteName = y.CategoryName
		s[x].ErpID = sh[y.CategoryID]
	}

	return s, nil
}

// GetSite 获取分类列表
func (c *BaseCategory) GetSite2() ([]Site, error) {
	var s []Site

	err := orm.DataDB.Table("base_category c").Select("c.category_id site_id, c.category_name site_name,m.erp_id").Joins("left join base_erp_map m on m.dhb_id = c.category_id").Where("m.map_type = 'goodcategory'").Scan(&s).Error

	if err != nil {
		return nil, err
	}

	return s, nil
}
