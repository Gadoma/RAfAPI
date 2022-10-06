package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gadoma/rafapi/internal/randomAffirmation/test"
	"github.com/oklog/ulid/v2"
)

func TestApiGetRandomAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryId1, _ := ulid.Parse("01GEJ0CR9DWN7SA1QBSJE4DVKF")
	categoryId2, _ := ulid.Parse("01GEJ0CRM2JW0KY2Z4R5CH4349")

	response, err := http.Get(fmt.Sprintf("http://%s/random_affirmation?categoryIds=%s&categoryIds=%s", testServerAddr, categoryId1, categoryId2))

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.GetRandomAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if result.Data.Text == "" {
		t.Error("result.Data.Text should not be empty")
	} else if got, want := result.Count, 1; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiGetRandomAffirmationEmpty(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryId1 := ulid.Make()
	categoryId2 := ulid.Make()

	response, err := http.Get(fmt.Sprintf("http://%s/random_affirmation?categoryIds=%s&categoryIds=%s", testServerAddr, categoryId1, categoryId2))

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.GetRandomAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if result.Data.Text != "" {
		t.Error("result.Data.Text should be empty")
	} else if got, want := result.Count, 0; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiGetRandomAffirmationError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	testCases := []struct {
		input    string
		expected int
	}{
		{"92233720368547758071", http.StatusBadRequest},
		{"abc", http.StatusBadRequest},
	}

	for _, testCase := range testCases {
		response, err := http.Get(fmt.Sprintf("http://%s/random_affirmation?categoryIds=%s", testServerAddr, testCase.input))

		if err != nil {
			t.Errorf("Could not send request because of %q", err)
		}

		data, err := io.ReadAll(response.Body)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(response.Body)

		if err != nil {
			t.Errorf("Could not read response because of %q", err)
		}

		var result test.GetRandomAffirmationResponse
		err = json.Unmarshal(data, &result)

		if err != nil {
			t.Errorf("Could not decode response because of %q", err)
		}

		if got, want := response.StatusCode, testCase.expected; got != want {
			t.Errorf("response.StatusCode=%v, want %v", got, want)
		} else if got, want := result.Status, statusError; got != want {
			t.Errorf("result.Status=%v, want %v", got, want)
		}
	}
}
