package main

import (
	"context"
	"github.com/diianpro/template/internal/config"
	"github.com/diianpro/template/internal/service"
	"github.com/diianpro/template/internal/storage/mongo"
	"github.com/diianpro/template/internal/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	ctx := context.Background()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	repos := mongo.New(cfg.MgConnection)
	services := service.New(repos)
	handler := http.New(services)

	e := echo.New()

	e.POST("/template", handler.AddTemplate())
	e.GET("/template/:id", handler.GetByIDTemplate())
	e.GET("/templates", handler.GetListsTemplate())
	e.DELETE("/template/:id", handler.DeleteTemplate())

	e.Logger.Fatal(e.Start(":8080"))

}
