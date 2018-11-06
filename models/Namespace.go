package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置Namespace表名
func (a *Namespace) TableName() string {
	return NamespaceTBName()
}

// NamespaceQueryParam 用于查询的类
type NamespaceQueryParam struct {
	BaseQueryParam
	NameSpaceLike string //模糊查询
	K8sNameLike string //模糊查询
	SearchStatus string //为空不查询，有值精确查询
}

// Namespace 实体类
type Namespace struct {
	Id                 int
	NameSpace      string    `orm:"unique;size(32)" form:"NameSpace"  valid:"Required;MaxSize(32);MinSize(4)"`
	HadoopName     string       `orm:"size(32)" form:"HadoopName" valid:"Required;"`
	K8sName        string       `orm:"size(32)" form:"K8sName" valid:"Required;"`
	Subsystem      string    `orm:"size(32)" form:"Subsystem" valid:"Required;"`
	ReleaseVersion string    `orm:"size(32)" form:"ReleaseVersion" valid:"Required;"`
	Remark         string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status         int       `orm:"default(4)" form:"Status" valid:"Range(0,1,2,3)"`
	StartName      string    `orm:"size(32)" form:"StartName"  valid:"Required;"`
	Creater        string    `orm:"size(32)" form:"Creater";"`
	Updater        string    `orm:"size(32)" form:"Updater";"`
	Createtime     time.Time `orm:"type(datetime);auto_now_add" `
	Updatetime     time.Time `orm:"null;type(datetime)" form:"-"`
}

// NamespacePageList 获取分页数据
func NamespacePageList(params *NamespaceQueryParam) ([]*Namespace, int64) {
	query := orm.NewOrm().QueryTable(NamespaceTBName())

	data := make([]*Namespace, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("namespace__istartswith", params.NameSpaceLike)
	query = query.Filter("k8sname__icontains", params.K8sNameLike)
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// NamespaceOne 根据id获取单条
func NamespaceOne(id int) (*Namespace, error) {
	o := orm.NewOrm()
	m := Namespace{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// NamespaceOneByNamespaceName 根据用户名密码获取单条
func NamespaceOneByNamespaceName(namespacename string) (*Namespace, error) {
	m := Namespace{}
	err := orm.NewOrm().QueryTable(NamespaceTBName()).Filter("namespacename", namespacename).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
