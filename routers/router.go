package routers

import (
	"beego02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/register", &controllers.UserController{})
}
