package http

import (
	"github.com/diianpro/template/internal/service"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"os"
)

type Server struct {
	tmpl *service.Template
}

func New(tmpl *service.Template) *Server {
	return &Server{
		tmpl: tmpl,
	}
}

func (s *Server) CreateTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ts, err := template.ParseFiles("./html/upload.html")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		err = ts.ExecuteTemplate(os.Stdout, "upload.html", nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

		}
		return c.JSON(http.StatusOK, "Successfully")
	}
}
