// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tGorpController struct {}
var GorpController tGorpController


func (_ tGorpController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Begin", args).URL
}

func (_ tGorpController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Commit", args).URL
}

func (_ tGorpController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Rollback", args).URL
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


type tMonitor struct {}
var Monitor tMonitor


func (_ tMonitor) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Monitor.Index", args).URL
}

func (_ tMonitor) AddiLOForm(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Monitor.AddiLOForm", args).URL
}

func (_ tMonitor) AddiLO(
		ilo interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ilo", ilo)
	return revel.MainRouter.Reverse("Monitor.AddiLO", args).URL
}

func (_ tMonitor) Overview(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Monitor.Overview", args).URL
}

func (_ tMonitor) Fans(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Monitor.Fans", args).URL
}

func (_ tMonitor) Powers(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Monitor.Powers", args).URL
}

func (_ tMonitor) Temperatures(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Monitor.Temperatures", args).URL
}

func (_ tMonitor) EventLog(
		id int,
		pageNumber int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "pageNumber", pageNumber)
	return revel.MainRouter.Reverse("Monitor.EventLog", args).URL
}

func (_ tMonitor) SystemLog(
		id int,
		pageNumber int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "pageNumber", pageNumber)
	return revel.MainRouter.Reverse("Monitor.SystemLog", args).URL
}


