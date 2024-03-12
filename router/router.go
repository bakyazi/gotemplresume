package router

import (
	"encoding/json"
	"github.com/bakyazi/gotemplresume/dto"
	"github.com/bakyazi/gotemplresume/pages"
	"github.com/labstack/echo/v4"
	"os"
)

type Router struct {
}

func Init(e *echo.Echo) {
	router := &Router{}
	e.Static("/resources", "resources")
	e.GET("/", router.Home)
	e.GET("/resume", router.Resume)
}

func (r *Router) Home(c echo.Context) error {
	return pages.Home().Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) Resume(c echo.Context) error {

	resumeFile, err := os.Open("resources/resume.json")
	if err != nil {
		return err
	}
	defer resumeFile.Close()

	resumeDecoder := json.NewDecoder(resumeFile)

	var resume dto.Resume

	err = resumeDecoder.Decode(&resume)
	if err != nil {
		return err
	}
	return pages.Resume(resume).Render(c.Request().Context(), c.Response().Writer)
}
