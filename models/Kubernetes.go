package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置Kubernetes表名
func (a *Kubernetes) TableName() string {
	return KubernetesTBName()
}

// KubernetesQueryParam 用于查询的类
type KubernetesQueryParam struct {
	BaseQueryParam
	NameLike string //模糊查询
	TargetNodeLike string //模糊查询
	SearchStatus string //为空不查询，有值精确查询
}

// Kubernetes 实体类
type Kubernetes struct {
	Id                 int
	Name       string    `orm:"size(32)" form:"Name"  valid:"Required"`
	TargetNode string    `orm:"size(128)" form:"TargetNode" valid:"Required;"`
	ApiServer  string    `orm:"null;size(128)" form:"ApiServer" valid:"Required;"`
	KubeConfig string    `orm:"null;size(200)" form:"KubeConfig" valid:"Required;"`
	Status     int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
}

// KubernetesPageList 获取分页数据
func KubernetesPageList(params *KubernetesQueryParam) ([]*Kubernetes, int64) {
	query := orm.NewOrm().QueryTable(KubernetesTBName())

	data := make([]*Kubernetes, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("name__istartswith", params.NameLike)
	query = query.Filter("targetnode__istartswith", params.TargetNodeLike)
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// KubernetesOne 根据id获取单条
func KubernetesOne(id int) (*Kubernetes, error) {
	o := orm.NewOrm()
	m := Kubernetes{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// KubernetesOneByName 根据用户名密码获取单条
func KubernetesOneByName(name string) (*Kubernetes, error) {
	m := Kubernetes{}
	err := orm.NewOrm().QueryTable(KubernetesTBName()).Filter("name", name).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
