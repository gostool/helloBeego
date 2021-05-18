package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

func (c *MainController) Prepare() {
	fmt.Println("hi I am in prepare")
}

func (c *MainController) Finish() {
	fmt.Println("hi I am in Finish")
}
func (c *MainController) Get() {
	mystruct := map[string]string{
		"Function": "Get",
		"Message":  "I am in Get",
	}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}

func (c *MainController) Post() {
	c.Ctx.Output.Body([]byte("Post"))
}

func (c *MainController) Put() {
	fmt.Println("hi I am in Put")
	c.Ctx.Output.Body([]byte("Put"))
}

func (c *MainController) Delete() {
	c.Ctx.Output.Body([]byte("Delete"))
}

func (c *MainController) Hello() {
	time.Sleep(time.Second)
	c.Ctx.ResponseWriter.Write([]byte("Hello, world"))
}
