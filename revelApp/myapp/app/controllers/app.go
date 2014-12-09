package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

type Hello struct {
	Hello1 string
	Hello2 string
}

func (c App) Index() revel.Result {
	var i Hello
	i.Hello1 = "hello"
	i.Hello2 = "world"
	return c.Render(i)
}
