package routes

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	database "github.com/utiiz/go/templ/internal/db"
	"github.com/utiiz/go/templ/internal/models"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var db database.DB

func NewRouter(d database.DB) *echo.Echo {
	db = d
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "internal/static")

	t := &Template{
		templates: template.Must(template.ParseGlob("internal/templates/**/*.html")),
	}
	e.Renderer = t

	e.GET("/", indexHandler)
	e.PUT("/task/:id", taskHandler)
	return e
}

type IndexHandlerData struct {
	Title         string
	TasksByStatus map[models.Status][]models.Task
}

func indexHandler(c echo.Context) error {
	data := IndexHandlerData{
		Title:         "Tasks",
		TasksByStatus: db.GetTasksByStatus(),
	}
	return c.Render(http.StatusOK, "indexPage", data)
}

type TaskHandlerData struct {
	TasksByStatus map[models.Status][]models.Task
}

func taskHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	db.UpdateTaskToDone(id)

	data := TaskHandlerData{
		TasksByStatus: db.GetTasksByStatus(),
	}

	return c.Render(http.StatusOK, "tasksLayout", data)
}
