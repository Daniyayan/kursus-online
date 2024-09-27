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

type MapelInterface interface {
	CreateMapel(ctx echo.Context) (err error)
	UpdateMapel(ctx echo.Context) (err error)
	DeleteMapel(ctx echo.Context) (err error)
	ListMapel(ctx echo.Context) (err error)
	DetailMapel(ctx echo.Context) (err error)
}

func (c *Controller) CreateMapel(ctx echo.Context) (err error) {
	var req Request.CreateMapel
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
	req.IdMapel = 0
	exists, err := Repository.ApplicationRepository.Mapel.CheckExistsMapelTitle(ctx.Request().Context(), req.NamaMapel)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if exists {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Nama Mata Pelajaran Duplicate").Error(),
		})
	}

	req.IdMapel = 1 // Simulate assigning a unique ID, you can change this as per your requirement
	if err = Repository.ApplicationRepository.Mapel.CreateMapel(ctx.Request().Context(), req); err != nil {
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

func (c *Controller) UpdateMapel(ctx echo.Context) (err error) {
	var req Request.UpdateMapel
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

	existsMapelId, err := Repository.ApplicationRepository.Mapel.CheckExistsMapelId(ctx.Request().Context(), req.NamaDosen)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if !existsMapelId {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Mapel ID Not Found").Error(),
		})
	}

	exists, err := Repository.ApplicationRepository.Mapel.CheckExistsMapelTitle(ctx.Request().Context(), req.NamaMapel)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if exists {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Nama Mata Pelajaran Duplicate").Error(),
		})
	}

	if err = Repository.ApplicationRepository.Mapel.UpdateMapel(ctx.Request().Context(), req); err != nil {
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

func (c *Controller) DeleteMapel(ctx echo.Context) (err error) {
	mapelId := ctx.Param("id")
	existsMapelId, err := Repository.ApplicationRepository.Mapel.CheckExistsMapelId(ctx.Request().Context(), mapelId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if !existsMapelId {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Mapel ID Not Found").Error(),
		})
	}

	if err = Repository.ApplicationRepository.Mapel.DeleteMapel(ctx.Request().Context(), mapelId); err != nil {
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

func (c *Controller) ListMapel(ctx echo.Context) (err error) {
	sortBy := ctx.QueryParams().Get("sortBy")
	list, err := Repository.ApplicationRepository.Mapel.ListMapel(ctx.Request().Context(), sortBy)
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

func (c *Controller) DetailMapel(ctx echo.Context) (err error) {
	mapelId := ctx.Param("id")
	detail, err := Repository.ApplicationRepository.Mapel.DetailMapel(ctx.Request().Context(), mapelId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusBadRequest, &Response.Responses{
				Data:    nil,
				Message: errors.New("Data Not Found").Error(),
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
