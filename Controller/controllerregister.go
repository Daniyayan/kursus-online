package Controller

import (
	"database/sql"
	"errors"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
	"kursus-online/Repository"
	"net/http"

	"github.com/labstack/echo"
)

type RegisterInterface interface {
	CreateRegister(ctx echo.Context) (err error)
	UpdateRegister(ctx echo.Context) (err error)
	DeleteRegister(ctx echo.Context) (err error)
	ListRegister(ctx echo.Context) (err error)
	DetailRegister(ctx echo.Context) (err error)
}

func (c *Controller) CreateRegister(ctx echo.Context) (err error) {
	var req Request.CreateRegister
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

	exists, err := Repository.ApplicationRepository.Register.CheckExistsRegisterEmail(ctx.Request().Context(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if exists {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Email sudah terdaftar").Error(),
		})
	}

	req.IdRegis = 1 // Simulate assigning a unique ID
	if err = Repository.ApplicationRepository.Register.CreateRegister(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    "",
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) UpdateRegister(ctx echo.Context) (err error) {
	var req Request.UpdateRegister
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

	existsRegisterId, err := Repository.ApplicationRepository.Register.CheckExistsRegisterId(ctx.Request().Context(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if !existsRegisterId {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Register ID tidak ditemukan").Error(),
		})
	}

	if err = Repository.ApplicationRepository.Register.UpdateRegister(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    "",
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) DeleteRegister(ctx echo.Context) (err error) {
	registerId := ctx.Param("id")
	existsRegisterId, err := Repository.ApplicationRepository.Register.CheckExistsRegisterId(ctx.Request().Context(), registerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if !existsRegisterId {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Register ID tidak ditemukan").Error(),
		})
	}

	if err = Repository.ApplicationRepository.Register.DeleteRegister(ctx.Request().Context(), registerId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    "",
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) ListRegister(ctx echo.Context) (err error) {
	sortBy := ctx.QueryParams().Get("sortBy")
	list, err := Repository.ApplicationRepository.Register.ListRegister(ctx.Request().Context(), sortBy)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    list,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    list,
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) DetailRegister(ctx echo.Context) (err error) {
	registerId := ctx.Param("id")
	detail, err := Repository.ApplicationRepository.Register.DetailRegister(ctx.Request().Context(), registerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusBadRequest, &Response.Responses{
				Data:    nil,
				Message: errors.New("Data tidak ditemukan").Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    detail,
		Message: http.StatusText(http.StatusOK),
	})
}
