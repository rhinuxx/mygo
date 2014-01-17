/**
 * Name:
 * Comment:
 * Author: Rhinux
 * Web: http://www.rhinux.info./
 * Created: 2014-01-17 18:44:21
 * Last-Modified: 2014-01-17 18:48:06
 */
package controllers

import (
    "github.com/astaxie/beego"
)

type CategoryController struct {
    beego.Controller
}

func (this *CategoryController) Get() {
    this.Data["IsCategory"] = true
    this.TplNames = "category.html"
}
