package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "index.tpl"

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

    this.Data["TrueCond"] = "true"
    this.Data["FalseCond"] = "false"

    type u struct {
        Name string
        Age  int
        Sex  string
    }
    user := &u{
        Name: "Joe",
        Age:  20,
        Sex:  "Male",
    }
    this.Data["User"] = user
    nums := []int{1,2,3,4,5,6}
    this.Data["Nums"] = nums
}
