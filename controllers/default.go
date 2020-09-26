package controllers

import (
	"beego02/db_mysql"
	"beego02/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1.获取请求数据
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("pwd")
	//使用固定数据进行数据校验
	//数据校验失败执行此代码
	if userName != "admin" || password != "123456" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据错误"))
		return
	}
	//数据校验成功执行此代码
	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据查询成功"))

	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "2424624564@qq.com"
	c.TplName = "index.tpl"
}

//编写一个post方法
//func (c *MainController) Post() {
//	//1.接收post请求的参数
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	//使用固定数据进行数据校验
//	if name != "admin" && age != "20"{
//		c.Ctx.ResponseWriter.Write([]byte("数据校验失败"))
//		return
//	}
//	c.Ctx.WriteString("数据校验成功")
//}
//


//编写一个post方法
func (c *MainController) Post() {
	//1.解析JSON数据
	var user models.User
	dataBytes , err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误")
		return
	}
	err = json.Unmarshal(dataBytes,&user)
	if err != nil {
		c.Ctx.WriteString("数据解析错误")
		return
	}

	id, err := db_mysql.InsertUser(user)
	if err != nil {
		c.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)

	result := models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	c.Data["json"] = &result
	c.ServeJSON()
	//c.Ctx.WriteString("恭喜，用户保存成功")

}
//编写一个post方法
//func (c *MainController) Post() {
//	//1.解析JSON数据
//	var person models.Personer
//	dataBytes , err := ioutil.ReadAll(c.Ctx.Request.Body)
//	if err != nil {
//		c.Ctx.WriteString("数据接收错误")
//		return
//	}
//	err = json.Unmarshal(dataBytes,&person)
//	if err != nil {
//		c.Ctx.WriteString("数据解析错误")
//		return
//	}
//	fmt.Println("姓名：",person.Name)
//	fmt.Println("生日：",person.Birthday)
//	fmt.Println("住址：",person.Address)
//	fmt.Println("昵称：",person.Nick)
//	c.Ctx.WriteString("数据解析成功")
//}

////编写一个Delete方法
//func (c *MainController) Delete() {
//	//1.接收Delete请求的参数
//
//}