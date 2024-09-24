package Routes

import (
	"kursus-online/Controller"

	"github.com/labstack/echo"
)

type Routes struct {
	Controller Controller.Controller
}

func (app *Routes) CollectRoutes(e *echo.Echo) {
	appRoutes := e

	// Mengelompokkan route untuk kursus
	course := e.Group("/courses")
	course.POST("/", app.Controller.CreateCourse)
	//course.GET("/", app.Controller.ListCourses)
	//course.PUT("/:id", app.Controller.UpdateCourse)
	//course.DELETE("/:id", app.Controller.DeleteCourse)
	//
	//user := e.Group("/user")
	//user.POST("/register", app.Controller.Register)

	appRoutes.Start(":3000")
}
