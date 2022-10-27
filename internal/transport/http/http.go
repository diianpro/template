package http

import (
	"github.com/diianpro/template/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

type Server struct {
	tmpl *service.Template
}

func New(tmpl *service.Template) *Server {
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

		id, err := s.tmpl.CreateTemplate(c.Request().Context(), fileBytes)
		if err != nil {
			log.Errorf("Create file error: %v", err)

		}

		return c.JSON(http.StatusOK, id)
	}
}

func (s *Server) GetByIDTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := uuid.Parse("_id")
		if err != nil {
			log.Errorf("Parse getById error: %v", err)

		}
		file, err := s.tmpl.GetByID(c.Request().Context(), ID)
		if err != nil {
			log.Errorf("Find getById error: %v", err)
		}
		return c.JSON(http.StatusOK, file)
	}
}

func (s *Server) DeleteTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := uuid.Parse("_id")
		if err != nil {
			log.Errorf("Parse getById error: %v", err)

		}
		err = s.tmpl.Delete(c.Request().Context(), ID)
		if err != nil {
			log.Errorf("Delete error: %v", err)
		}
		return c.JSON(http.StatusOK, "Document delete")
	}
}
