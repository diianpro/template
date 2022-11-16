package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	cli := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/template/92f2db56-5875-4cde-9892-1abdd3dddef2", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "text/html")

	resp, err := cli.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	fmt.Println(string(body))

}
