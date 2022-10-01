package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gadoma/rafapi/test"
)

func TestApiGetAffirmations(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	response, err := http.Get("http://" + TestServerAddr + "/affirmations")

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.GetAffirmationsResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, "OK"; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data[0].Id, 1; got != want {
		t.Errorf("result.Data[0].Id=%v, want %v", got, want)
	} else if got, want := len(result.Data), 8; got != want {
		t.Errorf("len=%v, want %v", got, want)
	} else if got, want := result.Count, 8; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiGetAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	affirmationId := 1

	response, err := http.Get(fmt.Sprintf("http://%s/affirmations/%d", TestServerAddr, affirmationId))

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.GetAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, "OK"; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data.Id, 1; got != want {
		t.Errorf("result.Data[0].Id=%v, want %v", got, want)
	} else if got, want := result.Count, 1; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiGetAffirmationError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	testCases := []struct {
		input    string
		expected int
	}{
		{"92233720368547758071", http.StatusBadRequest},
		{"10000", http.StatusNotFound},
	}

	for _, testCase := range testCases {
		response, err := http.Get(fmt.Sprintf("http://%s/affirmations/%s", TestServerAddr, testCase.input))

		if err != nil {
			t.Errorf("Could not send request because of %q", err)
		}

		data, err := io.ReadAll(response.Body)
		defer response.Body.Close()

		if err != nil {
			t.Errorf("Could not read response because of %q", err)
		}

		var result test.GetAffirmationResponse
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

func TestApiCreateAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	nextAffirmationId := 9

	payload := []byte(`{"text": "I am created.", "categoryId": 1}`)
	bodyReader := bytes.NewReader(payload)

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/affirmations", TestServerAddr), bodyReader)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.CreateAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, "OK"; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data, nextAffirmationId; got != want {
		t.Errorf("result.Data.Text=%v, want %v", got, want)
	} else if got, want := result.Count, 1; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiCreateAffirmationError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	testCases := []struct {
		input    string
		expected int
	}{
		{`I am not JSON`, http.StatusBadRequest},
		{`{"text": 123, "categoryId": 1}`, http.StatusBadRequest},
		{`{"text": "", "categoryId": 1}`, http.StatusUnprocessableEntity},
		{`{"text": "I am cool.", "categoryId": "category"}`, http.StatusBadRequest},
		{`{"text": "I am cool.", "categoryId": 0}`, http.StatusUnprocessableEntity},
	}

	for _, testCase := range testCases {
		payload := []byte(testCase.input)
		bodyReader := bytes.NewReader(payload)
		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/affirmations", TestServerAddr), bodyReader)

		if err != nil {
			t.Errorf("Could not prepare request because of %q", err)
		}

		response, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Errorf("Could not send request because of %q", err)
		}

		data, err := io.ReadAll(response.Body)
		defer response.Body.Close()

		if err != nil {
			t.Errorf("Could not read response because of %q", err)
		}

		var result test.CreateAffirmationResponse
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

func TestApiUpdateAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	affirmationId := 1
	payload := []byte(`{"text": "I am updated.", "categoryId": 1}`)
	bodyReader := bytes.NewReader(payload)

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/affirmations/%d", TestServerAddr, affirmationId), bodyReader)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.UpdateAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, "OK"; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := len(result.Data), 0; got != want {
		t.Errorf("len(result.Data)=%v, want %v", got, want)
	} else if got, want := result.Count, 0; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiUpdateAffirmationError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	affirmationId := 1

	testCases := []struct {
		input    string
		expected int
	}{
		{`I am not JSON`, http.StatusBadRequest},
		{`{"text": 123, "categoryId": 1}`, http.StatusBadRequest},
		{`{"text": "", "categoryId": 1}`, http.StatusUnprocessableEntity},
		{`{"text": "I am cool.", "categoryId": "category"}`, http.StatusBadRequest},
		{`{"text": "I am cool.", "categoryId": 0}`, http.StatusUnprocessableEntity},
	}

	for _, testCase := range testCases {
		payload := []byte(testCase.input)
		bodyReader := bytes.NewReader(payload)
		request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/affirmations/%d", TestServerAddr, affirmationId), bodyReader)

		if err != nil {
			t.Errorf("Could not prepare request because of %q", err)
		}

		response, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Errorf("Could not send request because of %q", err)
		}

		data, err := io.ReadAll(response.Body)
		defer response.Body.Close()

		if err != nil {
			t.Errorf("Could not read response because of %q", err)
		}

		var result test.UpdateAffirmationResponse
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

func TestApiDeleteAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	affirmationId := 1

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/affirmations/%d", TestServerAddr, affirmationId), nil)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.DeleteAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, "OK"; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := len(result.Data), 0; got != want {
		t.Errorf("len(result.Data)=%v, want %v", got, want)
	} else if got, want := result.Count, 0; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiDeleteAffirmationError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/affirmations/92233720368547758071", TestServerAddr), nil)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Errorf("Could not send request because of %q", err)
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Could not read response because of %q", err)
	}

	var result test.DeleteAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusBadRequest; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, "ERROR"; got != want {
		t.Errorf("result.Status=%v, want %v", got, want)
	}
}
