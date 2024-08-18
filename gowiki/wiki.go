package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")

	router.GET("/view/:title", viewHandler)
	router.GET("/edit/:title", editHandler)
	router.POST("/save/:title", saveHandler)

	router.Run("localhost:8080")
}

func viewHandler(c *gin.Context) {
	title := c.Param("title")
	p, err := loadPage(title)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/edit/%s", title))
		return
	}
	c.HTML(http.StatusOK, "view.html", p)
}

func editHandler(c *gin.Context) {
	title := c.Param("title")
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	c.HTML(http.StatusOK, "edit.html", p)
}

func saveHandler(c *gin.Context) {
	title := c.Param("title")
	body := c.Request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/view/%s", title))
}

// This is a method named save that takes as its receiver p, a pointer to Page .
// It takes no parameters, and returns a value of type error.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
