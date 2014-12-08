package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func loadPage(title string) string {
	filename := "wikiData/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(body)
}

func savePage(title string, body string) error {
	filename := "wikiData/" + title + ".txt"
	return ioutil.WriteFile(filename, []byte(body), 0600)
}

func (c App) View(title string) revel.Result {
	body := loadPage(title)
	if body == "" {
		return c.Redirect("/edit/" + title)
	}
	return c.Render(title, body)
}

func (c App) Edit(title string) revel.Result {
	body := loadPage(title)
	return c.Render(title, body)
}

func (c App) Save(title, body string) revel.Result {
	err := savePage(title, body)
	if err != nil {
		fmt.Println("Save Error : ", err)
		return c.RenderError(err)
	}
	return c.Redirect("/view/" + title)
}
