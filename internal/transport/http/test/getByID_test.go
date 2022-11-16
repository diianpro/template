package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	cli := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/template/0e44db45-1098-4161-abce-42904f67cb5b", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "text/html")

	resp, err := cli.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	fmt.Println(string(body))

}
