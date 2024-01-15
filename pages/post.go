package pages

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"isabelroses.com/lib"
)

func Post(c echo.Context) error {
	var (
		finalPost lib.Post
		posts     = lib.GetBlogPosts()
		slug      = c.Param("slug")
	)

	parts := strings.Split(slug, "-")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		log.Print(err)
	}

	for _, post := range posts {
		if id == post.ID {
			// if the slug is not the same as the post's slug
			if post.Slug != slug {
				c.Redirect(http.StatusSeeOther, "/blog/"+post.Slug)
			}

			finalPost = post
		}
	}

	props := lib.Post{
		ID:          finalPost.ID,
		Title:       finalPost.Title,
		Description: finalPost.Description,
		Content:     finalPost.Content,
		Date:        finalPost.Date,
		Tags:        finalPost.Tags,
		Slug:        finalPost.Slug,
	}

	templates := []string{
		lib.GetPath("templates/layouts/base.html"),
		lib.GetPath("templates/components/header.html"),
		lib.GetPath("templates/pages/post.html"),
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return ts.ExecuteTemplate(c.Response().Writer, "base", props)
}