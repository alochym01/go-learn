package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/alochym01/ftp-api/api"
	"github.com/alochym01/ftp-api/api/controllers"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := api.SetupRouter()
	t.Run("Method Get", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/get", nil)
		router.ServeHTTP(w, req)
		w := httptest.NewRecorder()

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "pong", w.Body.String())

	})
	t.Run("Method Post w JSON value", func(t *testing.T) {
		// https://medium.com/@craigchilds94/testing-gin-json-responses-1f258ce3b0b1

		// create json variable
		var response map[string]string

		// body := gin.H{
		// 	"ping": "pong",
		// }

		// set router to url to serve request
		router.POST("/pongjson", controllers.PostPingJson)

		// prepare a request
		params := []byte(`{"ping":"pong"}`)
		// We use http.NewRequest to create a request
		req, _ := http.NewRequest("POST", "/pingjson", bytes.NewBuffer(params))
		req.Header.Add("Content-Type", "application/json")

		// net/http/httptest has a spy already made for us called ResponseRecorder so we can use that.
		// It has many helpful methods to inspect what has been written as a response.
		w := httptest.NewRecorder()

		// Running server to serve request
		router.ServeHTTP(w, req)

		// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to
		// Convert the JSON response to a map
		err := json.Unmarshal([]byte(w.Body.String()), &response)

		// check key exist in response reply
		value, exists := response["ping"]

		// compare response status code
		assert.Equal(t, http.StatusOK, w.Code)
		// compare an error with nil
		assert.Nil(t, err)
		// compare a key is existed
		assert.True(t, exists)
		// compare value
		assert.Equal(t, "pong", value)
	})
	t.Run("Method Post w Form value", func(t *testing.T) {
		router.POST("/pongform", controllers.PostPingForm)
		w := httptest.NewRecorder()
		params := url.Values{}
		params.Add("ping", "pong")
		req, _ := http.NewRequest("POST", "/pingform", strings.NewReader(params.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})
}
