package http

import (
	"fmt"
	"github.com/diianpro/template/internal/domain"
	"github.com/diianpro/template/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	tmpl service.Template
}

func New(tmpl service.Template) *Server {
	return &Server{
		tmpl: tmpl,
	}
}

func (s *Server) AddTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := c.Request().ParseMultipartForm(10 << 20)
		if err != nil {
			log.Errorf("Parse file error: %v", err)

		}
		// Get handler for filename, size and headers
		file, _, err := c.Request().FormFile("template")
		if err != nil {
			log.Errorf("Get file error: %v", err)

		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Errorf("Read file error: %v", err)

		}

		fl := domain.Template{
			Data: fileBytes,
		}

		id, err := s.tmpl.CreateTemplate(c.Request().Context(), &fl)
		if err != nil {
			log.Errorf("Create file error: %v", err)

		}

		return c.JSON(http.StatusOK, id)
	}
}

func (s *Server) GetByIDTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := c.Param("id")
		id, err := uuid.Parse(resp)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "parse getById error")
		}
		file, err := s.tmpl.GetByID(c.Request().Context(), id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNoContent, "find getById error")
		}

		return c.JSON(http.StatusOK, file)
	}
}

func (s *Server) DeleteTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := c.Param("id")
		ID, err := uuid.Parse(resp)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "parse getById error")
		}
		err = s.tmpl.Delete(c.Request().Context(), ID.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "delete error")
		}
		return c.JSON(http.StatusOK, "Document delete")
	}
}

func (s *Server) GetListsTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		cnvLimit, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return fmt.Errorf("parse limit: error : %w", err)
		}
		offset := c.QueryParam("offset")
		cnvOffset, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return fmt.Errorf("parse offser: error : %w", err)
		}
		result, err := s.tmpl.GetAll(c.Request().Context(), cnvLimit, cnvOffset)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "get by list error")
		}
		return c.JSON(http.StatusOK, result)
	}
}
