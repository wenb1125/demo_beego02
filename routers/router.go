package routers

import (
	"beego02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.MainController{})
}
