package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestGetByList(t *testing.T) {
	cli := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/templates?limit=1&offset=0", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	fmt.Println(string(body))

}
