package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		error  string
		status int
		result int
	}{
		{name: "double of 2", input: "2", error: "", status: 200, result: 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, "localhost:4050/double?v="+testCase.input, nil)
			if err != nil {
				t.Fatalf("cannot get the request %v", err)
			}

			rec := httptest.NewRecorder()
			doubleHandler(rec, request)
			res := rec.Result()

			if res.StatusCode != testCase.status {
				t.Errorf("received status code ")
				return
			}

			respBytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("cannot read the body")
				return
			}
			defer res.Body.Close()

			trimErr := strings.TrimSpace(string(respBytes))

			if res.StatusCode != http.StatusOK {
				//check the error message
				if trimErr != testCase.error {
					t.Errorf("received status code ")
				}
				return
			}

			result, err := strconv.Atoi(strings.TrimSpace(string(respBytes)))
			if err != nil {
				t.Fatalf("cannot convert")
				return
			}

			if result != testCase.result {
				t.Errorf("received result error")
			}
		})
	}

}
