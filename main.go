package main

import "task-order/cmd"

//go:generate swag init --parseDependency --parseDepth=6

// @title go-admin API
// @version 1.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
