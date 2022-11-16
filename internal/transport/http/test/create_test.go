package test

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

func TestServer_CreateTemplate(t *testing.T) {
	req, err := newFileUploadRequest("http://localhost:8080/create", "template", "../../../../html/template.html")
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error")
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	log.Info(string(data))
}

// Creates a new file upload http request with optional extra params
func newFileUploadRequest(uri, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func TestRead(t *testing.T) {
	var template template.Template
	err := client.Database("templates").Collection("template").FindOne(context.Background(), bson.M{})
	if err != nil {
		log.Errorf("Read file error: %v", err)
	}
	fmt.Println(template)
}
