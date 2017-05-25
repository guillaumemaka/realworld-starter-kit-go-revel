// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tGormController struct {}
var GormController tGormController


func (_ tGormController) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Init", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tApplicationController struct {}
var ApplicationController tApplicationController


func (_ tApplicationController) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApplicationController.Init", args).URL
}

func (_ tApplicationController) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApplicationController.AddUser", args).URL
}


type tUserController struct {}
var UserController tUserController


func (_ tUserController) GetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.GetUser", args).URL
}

func (_ tUserController) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Register", args).URL
}

func (_ tUserController) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Login", args).URL
}

func (_ tUserController) UpdateUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.UpdateUser", args).URL
}


