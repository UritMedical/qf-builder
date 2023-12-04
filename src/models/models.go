package models

type Setting struct {
	Project string `yaml:"project"`
	//生成类型
	Type string `yaml:"type"`
	//语言
	Lang string `yaml:"lang"`
	//定义文件路径
	Define string `yaml:"define"`
	//输出路径
	Output string `yaml:"output"`
	//是否覆盖
	Cover bool `yaml:"cover"`
	//资源文件路径
	Src []string `yaml:"src"`
	//样式
	Style string `yaml:"style"`
	//初始化执行的脚本
	Init string `yaml:"init"`
}

type Model struct {
	Name    string
	Comment string
	Props   []Prop
}

// Prop 属性
type Prop struct {
	//Name 名称
	Name string
	//Comment 注释
	Comment string
	//TypeName 类型
	TypeName string
	//IsArray 是否数组
	IsArray bool
	//IsPoint 是否指针
	IsPoint bool
}

// Module 模块
type Module struct {
	//Name 名称
	Name string
	//Comment 注释
	Comment string
	//Route 路由
	Route string
	//Groups 分组
	Groups []Group
	//Notices 通知
	Notices []Notice
}

// Group 分组
type Group struct {
	//Name 名称
	Name string
	//Comment 注释
	Comment string
	//Heads 头
	Heads []Prop
	//Apis api
	Apis []Api
}

// Api api
type Api struct {
	//Name 名称
	Name string
	//Comment 注释
	Comment string
	//Route 路由
	Route string
	//ReqType 请求方式
	ReqType string
	//头参数
	Heads []Prop
	//Query query 参数
	Query []Prop
	//json参数
	Json []Prop
	//返回
	Return Prop
}

// Param 参数
type Param struct {
	//Name 名称
	Name string
	//Comment 注释
	Comment string
	//TypeName 类型
	TypeName string
}

// Notice 通知
type Notice struct {
	//Topic 主题
	Topic string
	//Comment 注释
	Comment string
	//Query 参数
	Params []Prop
}
