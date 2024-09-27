package Routes

import (
	"kursus-online/Controller"

	"net/http"

	"github.com/labstack/echo"
)

type Routes struct {
	Controller Controller.Controller
}

func (app *Routes) CollectRoutes(e *echo.Echo) {
	appRoutes := e

	// Mengelompokkan route untuk kursus
	Mapel := appRoutes.Group("/Mapel")
	Mapel.POST("/", app.Controller.CreateMapel)
	appRoutes.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "ping")
	})
	//course.GET("/", app.Controller.ListCourses)
	//course.PUT("/:id", app.Controller.UpdateCourse)
	//course.DELETE("/:id", app.Controller.DeleteCourse)
	//
	//user := e.Group("/user")
	//user.POST("/register", app.Controller.Register)

	appRoutes.Start(":8000")
}
