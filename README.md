## SDRMS

# 特点
1. 分页列表页面的搜索条件、搜索面板、PageSize、当前页数、显示/隐藏列在变化时自动保存，页面刷新后、重新进入时，这些状态依然保持；
2. TreeTabe列表节点展开/收缩状态、滚动条位置时自动保存，页面刷新后、重新进入时，这些状态依然保持；
3. 编辑分页列表、TreeTabe列表中数据后，当前数据行背景闪烁，如果当前数据行由于顺序变化跳出可视区域，则滚动条自动滚动，将当前数据行移动至可视区域；
4. 精确至Action的轻量级功能权限控制，后台用户与角色、角色与资源（菜单、按钮）都是多对多关系，可以灵活配置用户可访问的资源。
# 后端框架
1. 基于Beego 1.9.1（官方的版本已经升级到v1.10.1，本项目也可以支持），使用官方的orm、cache、session、logs等模块，感谢原作者提供了如此简单易用的框架（<a href="https://beego.me/">更多信息</a>）;
2. 代码风格源自笔者本人其他语言的风格，但参考了多个开源系统的代码风格如 PPGo_ApiAdmin（<a href="https://github.com/george518/PPGo_ApiAdmin">更多信息</a>）、ERP系统（<a href="https://github.com/hexiaoyun128/ERP">更多信息</a>）等。
# 前端框架
1. 基于AdminLTE2（<a href="https://adminlte.io/themes/AdminLTE/index2.html">更多信息</a>）；
2. 弹出层插件使用了Layer（<a href="http://layer.layui.com/">更多信息</a>）；
3. 分页列表使用Bootstrap-table（<a href="http://bootstrap-table.wenzhixin.net.cn/zh-cn/getting-started/">更多信息</a>），并集成cookie、x-editable等辅助插件（<a href="http://bootstrap-table.wenzhixin.net.cn/zh-cn/extensions/">更多信息</a>）实现状态保持、快速编辑等功能。对Bootstrap-table进行了扩展使分页导航可以通过下拉迅速定位。修复Bootstrap-table-cookie的若干Bug；
4. TreeTable列表使用jQuery-treetable插件，按照Boostrap风格进行了样式调整（<a href="http://ludo.cubicphuse.nl/jquery-treetable/">更多信息</a>）；
5. 下拉框使用Bootstrap-select（<a href="http://silviomoreto.github.io/bootstrap-select/">更多信息</a>）；
6. 高亮显示使用的是笔者自已开发的插件；
7. 分页列表里搜索条件、搜索面板状态自动保存使用的是笔者自已开发的插件。

# 安装方法

本系统基于beego开发，默认使用mysql数据库，缓存redis 


1. 安装golang环境（ 略）

2. 安装本系统
```
go get github.com/lhtzbj12/sdrms
```
3. 将根目录下的sdrms.sql导入mysql

4. 修改配置文件 conf/app.conf
 需要配置mysql和redis的参数
5. 运行
在 sdrms 目录使用beego官方提供的命令运行
```
bee run
```
在浏览器里打开 http://localhost:8080 进行访问

# 补充说明
 beego升级到v1.10.1后，启动本项目时报错
 ```
 cannot find package "github.com/gomodule/redigo/redis"
 ```
 解决方法很简单，只需要在终端运行下面命令，下载需要的包即可
 ```
 go get github.com/gomodule/redigo/redis
 ```
