package controllers

import (
	"fmt"
	"github.com/revel/revel"
	//"io/ioutil"
	"regexp"
	"revelApp/revelwiki/app/routes"
	//"runtime"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

var page = make(map[string]string)

func loadPage(title string) string {
	return page[title]
	/*
		filename := title + ".txt"
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("title:", title)
			fmt.Println("Load Error:", err)
			return ""
		}

		return string(body)
	*/
}

func savePage(title string, body string) error {
	page[title] = body
	return nil
	/*
		filename := title + ".txt"
		return ioutil.WriteFile(filename, []byte(body), 0600)
	*/
}

var titleValid = regexp.MustCompile("^([a-zA-Z0-9]+)$")

func (c App) View(title string) revel.Result {
	c.Validation.Match(title, titleValid)
	if c.Validation.HasErrors() {
		return c.NotFound(title)
	}

	body := loadPage(title)
	if body == "" {
		return c.Redirect(routes.App.Edit(title))
	}
	return c.Render(title, body)
}

func (c App) Edit(title string) revel.Result {
	c.Validation.Match(title, titleValid)
	if c.Validation.HasErrors() {
		return c.NotFound(title)
	}

	body := loadPage(title)
	return c.Render(title, body)
}

func (c App) Save(title, body string) revel.Result {
	c.Validation.Match(title, titleValid)
	if c.Validation.HasErrors() {
		return c.NotFound(title)
	}

	err := savePage(title, body)
	if err != nil {
		fmt.Println("title:", title)
		fmt.Println("Save Error:", err)
		return c.RenderError(err)
	}

	return c.Redirect(routes.App.View(title))
	//return c.Redirect("/view/" + title)
}
