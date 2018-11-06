package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"sdrms/enums"
	"sdrms/models"
	"github.com/astaxie/beego/orm"
)

type HadoopController struct {
	BaseController
}

func (c *HadoopController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()

}
func (c *HadoopController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "hadoop/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "hadoop/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("HadoopController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("HadoopController", "Delete")
}
func (c *HadoopController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.HadoopQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.HadoopPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *HadoopController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.Hadoop{}
	var err error
	if Id > 0 {
		m, err = models.HadoopOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	} else {
		//添加用户时默认状态为启用
		m.Status = enums.Enabled
	}
	c.Data["m"] = m
	c.setTpl("hadoop/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "hadoop/edit_footerjs.html"
}
func (c *HadoopController) Save() {
	m := models.Hadoop{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	if m.Id == 0 {
		if _, err := o.Insert(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if oM, err := models.HadoopOne(m.Id); err != nil {
			c.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.HadoopName = oM.HadoopName
			m.HadoopConfig = oM.HadoopConfig
		}
		if _, err := o.Update(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
	c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)

}
func (c *HadoopController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(models.HadoopTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
