/**
 * Name:
 * Comment:
 * Author: Rhinux
 * Web: http://www.rhinux.info./
 * Created: 2014-01-13 17:22:00
 * Last-Modified: 2014-01-17 16:45:52
 */
package controllers

import (
    "github.com/astaxie/beego"
    //不知道为什么beego.Context无法引入需要导入下面包
    "github.com/astaxie/beego/context"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    isExit := this.Input().Get("exit") == "true"
    if isExit {
        this.Ctx.SetCookie("uname", "", -1, "/")
        this.Ctx.SetCookie("pwd", "", -1, "/")
        this.Redirect("/", 301)
    }
    this.TplNames = "login.html"
}

func (this *LoginController) Post() {
    uname := this.Input().Get("uname")
    pwd := this.Input().Get("pwd")
    autoLogin := this.Input().Get("autologin") == "on"
    if beego.AppConfig.String("uname") == uname &&
        beego.AppConfig.String("pwd") == pwd {
        maxAge := 0
        if autoLogin {
            maxAge = 1<<31 - 1
        }
        this.Ctx.SetCookie("uname", uname, maxAge, "/")
        this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
    }
    this.Redirect("/", 301)
    return
}

func checkAccount(ctx *context.Context) bool {

    ck, err := ctx.Request.Cookie("uname")
    if err != nil {
        return false
    }
    uname := ck.Value

    ck, err = ctx.Request.Cookie("pwd")
    if err != nil {
        return false
    }
    pwd := ck.Value
    return beego.AppConfig.String("uname") == uname &&
        beego.AppConfig.String("pwd") == pwd
}
