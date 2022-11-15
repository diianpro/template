package test

import (
	"context"
	"github.com/diianpro/template/internal/service"
	mongo2 "github.com/diianpro/template/internal/storage/mongo"
	"github.com/diianpro/template/internal/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var client *mongo.Client

func TestMain(m *testing.M) {
	ctx := context.Background()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	repos := mongo2.New(client)
	services := service.New(repos)
	handler := http.New(services)

	e := echo.New()
	e.POST("/create", handler.AddTemplate())
	e.GET("/template/:id", handler.GetByIDTemplate())
	e.DELETE("/template/:id", handler.DeleteTemplate())
	e.GET("/templates", handler.GetListsTemplate())

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-quit
		err := client.Disconnect(ctx)
		if err != nil {
			log.Errorf("disconnect error: %v", err)
		}
		err = e.Shutdown(ctx)
		if err != nil {
			log.Errorf("Error: %v", err)
		}
	}()

	m.Run()
}
