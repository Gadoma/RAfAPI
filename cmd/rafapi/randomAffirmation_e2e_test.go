package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gadoma/rafapi/test"
)

func TestApiGetRandomAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryIds := []int{1, 2}

	response, err := http.Get(fmt.Sprintf("http://%s/random_affirmation?categoryIds=%d&categoryIds=%d", TestServerAddr, categoryIds[0], categoryIds[1]))

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

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
	} else if got, want := result.Status, "OK"; got != want {
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

	categoryIds := []int{1000, 2000}

	response, err := http.Get(fmt.Sprintf("http://%s/random_affirmation?categoryIds=%d&categoryIds=%d", TestServerAddr, categoryIds[0], categoryIds[1]))

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

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
	} else if got, want := result.Status, "OK"; got != want {
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
		response, err := http.Get(fmt.Sprintf("http://%s/random_affirmation?categoryIds=%s", TestServerAddr, testCase.input))

		if err != nil {
			t.Errorf("Could not send request because of %q", err)
		}

		data, err := io.ReadAll(response.Body)
		defer response.Body.Close()

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
		} else if got, want := result.Status, "ERROR"; got != want {
			t.Errorf("result.Status=%v, want %v", got, want)
		}
	}
}
