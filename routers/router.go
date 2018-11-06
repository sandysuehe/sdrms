package routers

import (
	"sdrms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")

	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	//Hadoop配置路由
	beego.Router("/hadoop/index", &controllers.HadoopController{}, "*:Index")
	beego.Router("/hadoop/datagrid", &controllers.HadoopController{}, "POST:DataGrid")
	beego.Router("/hadoop/edit/?:id", &controllers.HadoopController{}, "Get,Post:Edit")
	beego.Router("/hadoop/delete", &controllers.HadoopController{}, "Post:Delete")

	//Kubernetes配置路由
	beego.Router("/kubernetes/index", &controllers.KubernetesController{}, "*:Index")
	beego.Router("/kubernetes/datagrid", &controllers.KubernetesController{}, "POST:DataGrid")
	beego.Router("/kubernetes/edit/?:id", &controllers.KubernetesController{}, "Get,Post:Edit")
	beego.Router("/kubernetes/delete", &controllers.KubernetesController{}, "Post:Delete")

	//Namespace配置路由
	beego.Router("/namespace/index", &controllers.NamespaceController{}, "*:Index")
	beego.Router("/namespace/datagrid", &controllers.NamespaceController{}, "POST:DataGrid")
	beego.Router("/namespace/edit/?:id", &controllers.NamespaceController{}, "Get,Post:Edit")
	beego.Router("/namespace/delete", &controllers.NamespaceController{}, "Post:Delete")

	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	beego.Router("/", &controllers.HomeController{}, "*:Index")

}
