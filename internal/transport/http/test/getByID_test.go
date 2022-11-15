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

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/template/752a70e7-dea7-422a-ae7a-470819aff1e2", nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "text/html")

	resp, err := cli.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	fmt.Println(string(body))

}
