package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置Hadoop表名
func (a *Hadoop) TableName() string {
	return HadoopTBName()
}

// HadoopQueryParam 用于查询的类
type HadoopQueryParam struct {
	BaseQueryParam
	HadoopNameLike string //模糊查询
	HadoopConfigLike string //模糊查询
	SearchStatus string //为空不查询，有值精确查询
}

// Hadoop 实体类
type Hadoop struct {
	Id                 int
	HadoopName       string    `orm:"size(32)" form:"HadoopName"  valid:"Required"`
	HadoopConfig string    `orm:"size(128)" form:"HadoopConfig"  valid:"Required"`
	Status     int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
}

// HadoopPageList 获取分页数据
func HadoopPageList(params *HadoopQueryParam) ([]*Hadoop, int64) {
	query := orm.NewOrm().QueryTable(HadoopTBName())

	data := make([]*Hadoop, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("hadoopname__istartswith", params.HadoopNameLike)
	query = query.Filter("hadoopconfig__icontains", params.HadoopConfigLike)
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// HadoopOne 根据id获取单条
func HadoopOne(id int) (*Hadoop, error) {
	o := orm.NewOrm()
	m := Hadoop{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// HadoopOneByHadoopName 根据用户名密码获取单条
func HadoopOneByHadoopName(hadoopname string) (*Hadoop, error) {
	m := Hadoop{}
	err := orm.NewOrm().QueryTable(HadoopTBName()).Filter("hadoopname", hadoopname).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
