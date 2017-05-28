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

func (_ tApplicationController) ExtractArticle(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApplicationController.ExtractArticle", args).URL
}


type tArticleController struct {}
var ArticleController tArticleController


func (_ tArticleController) Index(
		tag string,
		favorited string,
		author string,
		offset int,
		limit int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tag", tag)
	revel.Unbind(args, "favorited", favorited)
	revel.Unbind(args, "author", author)
	revel.Unbind(args, "offset", offset)
	revel.Unbind(args, "limit", limit)
	return revel.MainRouter.Reverse("ArticleController.Index", args).URL
}

func (_ tArticleController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleController.Create", args).URL
}

func (_ tArticleController) Read(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleController.Read", args).URL
}

func (_ tArticleController) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleController.Update", args).URL
}

func (_ tArticleController) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ArticleController.Delete", args).URL
}


type tFavoriteController struct {}
var FavoriteController tFavoriteController


func (_ tFavoriteController) Post(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("FavoriteController.Post", args).URL
}

func (_ tFavoriteController) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("FavoriteController.Delete", args).URL
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


