package controllers

import (
	"github.com/astaxie/beego"
	"TestBlog/models"
	"fmt"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")

	println("op-----", len(op))
	println("Param-----", c.Ctx.Input.Param(":op"))
	println("GetBody-----", c.Ctx.Request.GetBody)
	println("GetString-----", c.GetString("op"))
	println("Get-----", c.Input().Get("op"))
	println("RequestBody-----", c.Ctx.Input.RequestBody)

	println("URI-----", c.Ctx.Input.URI())
	println("URL-----", c.Ctx.Input.URL())
	println("ParseForm-----", c.Ctx.Request.ParseForm())

	switch op {
	case "add":
		name := c.Input().Get("name")

		println("name-----",name)

		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			println("err",err)
		}
		c.Redirect("/category",301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}

	c.Data["IsCategory"] = true
	c.TplName = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		println("err",err)
	}
}

func (c *CategoryController) Post() {
	c.Ctx.WriteString(fmt.Sprint(c.Input()))

	op := c.Input().Get("op")
	println("op-----", len(op))

	switch op {
	case "add":
		name := c.Input().Get("name")

		println("name-----",name)

		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			println("err",err)
		}
		c.Redirect("/category",301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}

	c.Data["IsCategory"] = true
	c.TplName = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		println("err",err)
	}
}