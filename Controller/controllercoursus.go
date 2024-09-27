package Controller

import (
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
	"kursus-online/Repository"
	"net/http"

	"github.com/labstack/echo"
)

// Coursus Interface
type CoursusInterface interface {
	CreateCoursus(ctx echo.Context) error
	UpdateCoursus(ctx echo.Context) error
	DeleteCoursus(ctx echo.Context) error
	ListCoursus(ctx echo.Context) error
	DetailCoursus(ctx echo.Context) error
}

// CreateCoursus
func (c *Controller) CreateCoursus(ctx echo.Context) (err error) {
	var req Request.CreateCoursus
	if err = ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if err = ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	req.IdCoursus = 0 // Use auto-increment
	//exists, err := Repository.ApplicationRepository.Coursus.CheckExistsCoursusTitle(ctx.Request().Context(), req.Description)
	//if err != nil {
	//	return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
	//		Data:    nil,
	//		Message: err.Error(),
	//	})
	//}

	//if exists {
	//	return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
	//		Data:    nil,
	//		Message: err.Error(),
	//	})
	//}

	req.IdCoursus = 1 // Simulate assigning a unique ID, you can change this as per your requirement
	if err = Repository.ApplicationRepository.Coursus.CreateCourse(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    req,
		Message: http.StatusText(http.StatusOK),
	})
}

// UpdateCoursus
//func (c *Controller) UpdateCoursus(ctx echo.Context) error {
//	var req Request.UpdateCoursus
//	if err := ctx.Bind(&req); err != nil {
//		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	if err := ctx.Validate(req); err != nil {
//		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	exists, err := Repository.ApplicationRepository.Coursus.CheckExistsCoursusTitle(ctx.Request().Context(), req.Description)
//	if err != nil {
//		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	if !exists {
//		return ctx.JSON(http.StatusNotFound, &Response.Responses{
//			Data:    nil,
//			Message: "Coursus Id Not Found",
//		})
//	}
//
//	if err := Repository.ApplicationRepository.Coursus.UpdateCoursus(ctx.Request().Context(), req); err != nil {
//		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	return ctx.JSON(http.StatusOK, &Response.Responses{
//		Data:    req,
//		Message: http.StatusText(http.StatusOK),
//	})
//}
//
//// DeleteCoursus
//func (c *Controller) DeleteCoursus(ctx echo.Context) error {
//	courseId := ctx.Param("id")
//	exists, err := Repository.ApplicationRepository.Coursus.CheckExistsCoursusId(ctx.Request().Context(), courseId)
//	if err != nil {
//		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	if !exists {
//		return ctx.JSON(http.StatusNotFound, &Response.Responses{
//			Data:    nil,
//			Message: "Coursus Id Not Found",
//		})
//	}
//
//	if err := Repository.ApplicationRepository.Coursus.DeleteCoursus(ctx.Request().Context(), courseId); err != nil {
//		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	return ctx.JSON(http.StatusOK, &Response.Responses{
//		Data:    nil,
//		Message: http.StatusText(http.StatusOK),
//	})
//}
//
//// ListCoursuss
//func (c *Controller) ListCoursuss(ctx echo.Context) error {
//	sortBy := ctx.QueryParams().Get("sortBy")
//	list, err := Repository.ApplicationRepository.Coursus.ListCoursus(ctx.Request().Context(), sortBy)
//	if err != nil {
//		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
//			Data:    list,
//			Message: err.Error(),
//		})
//	}
//
//	return ctx.JSON(http.StatusOK, &Response.Responses{
//		Data:    list,
//		Message: http.StatusText(http.StatusOK),
//	})
//}
//
//// DetailCoursus
//func (c *Controller) DetailCoursus(ctx echo.Context) error {
//	courseId := ctx.Param("id")
//	detail, err := Repository.ApplicationRepository.Coursus.DetailCoursus(ctx.Request().Context(), courseId)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return ctx.JSON(http.StatusNotFound, &Response.Responses{
//				Data:    nil,
//				Message: "Data Not Found",
//			})
//		}
//
//		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
//			Data:    nil,
//			Message: err.Error(),
//		})
//	}
//
//	return ctx.JSON(http.StatusOK, &Response.Responses{
//		Data:    detail,
//		Message: http.StatusText(http.StatusOK),
//	})
//}

// Similar implementations can be created for `Mapel`, `Register`, and `User`
