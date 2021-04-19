package firefly_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/randaalex/finance_bot/pkg/firefly"
)

const (
	baseURL = "http://firefly/api/"
)

func TestFirefly_NewClient(t *testing.T) {
	httpClient := &http.Client{}

	t.Run("http client is nil", func(t *testing.T) {
		client := firefly.NewClient(nil, baseURL)

		assert.Equal(t, "/api/", client.BaseURL.Path)
		assert.NotNil(t, client.HTTPClient)
	})

	t.Run("http client is present", func(t *testing.T) {
		client := firefly.NewClient(httpClient, baseURL)

		assert.Equal(t, "/api/", client.BaseURL.Path)
		assert.Equal(t, httpClient, client.HTTPClient)
	})

	t.Run("http client with invalid baseUrl", func(t *testing.T) {
		client := firefly.NewClient(nil, "http://firefly/api")
		req, err := client.NewRequest("GET", "api", "")

		assert.Error(t, err, "BaseURL must have a trailing slash, but \"http://firefly/api\" does not")
		assert.Nil(t, req)

		assert.NotNil(t, client.HTTPClient)
	})
}
