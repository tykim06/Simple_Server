package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
	revel.OnAppStart(InitHpDB)
	revel.OnAppStart(InitHttp)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
