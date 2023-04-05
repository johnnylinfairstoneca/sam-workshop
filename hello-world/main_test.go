package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("verifies successful response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, "127.0.0.1")
		}))
		defer ts.Close()

		DefaultHTTPGetAddress = ts.URL

		result, err := handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if result.StatusCode != 200 {
			t.Fatal("should be HTTP 200")
		}
		response := make(map[string]interface{})
		err = json.Unmarshal([]byte(result.Body), &response)
		if err != nil {
			t.Fatal("handler must return json of obj")
		}

		var message interface{}
		var ok bool
		if message, ok = response["message"]; !ok {
			t.Fatal("response does not have message key")
		}

		if reflect.TypeOf(message).Kind() != reflect.String {
			t.Fatal("response['message'] should be string but is", reflect.TypeOf(message).Kind())
		}

		if message.(string) != "hello my friend" {
			t.Fatal("response['message'] is not 'hello my friend'")
		}

	})
}
