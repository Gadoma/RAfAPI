package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gadoma/rafapi/internal/affirmation/test"
	"github.com/oklog/ulid/v2"
)

var affirmationIdString = "01GEJ0CNNA3VXV1HMJCKFNCYJV"

func TestApiGetAffirmations(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	response, err := http.Get("http://" + testServerAddr + "/affirmations")

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

	var result test.GetAffirmationsResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data[0].Id.String(), affirmationIdString; got != want {
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

	affirmationId, _ := ulid.Parse(affirmationIdString)

	response, err := http.Get(fmt.Sprintf("http://%s/affirmations/%s", testServerAddr, affirmationId))

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

	var result test.GetAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data.Id.String(), affirmationIdString; got != want {
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
		{"92233720368547758071", http.StatusNotFound},
		{ulid.Make().String(), http.StatusNotFound},
	}

	for _, testCase := range testCases {
		response, err := http.Get(fmt.Sprintf("http://%s/affirmations/%s", testServerAddr, testCase.input))

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

		var result test.GetAffirmationResponse
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

func TestApiCreateAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	payload := []byte(`{"text": "I am created.", "categoryId": "01GEJ0CR9DWN7SA1QBSJE4DVKF"}`)
	bodyReader := bytes.NewReader(payload)

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/affirmations", testServerAddr), bodyReader)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

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

	var result test.CreateAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
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
		{`{"text": 123, "categoryId": "01GEJ0CR9DWN7SA1QBSJE4DVKF"}`, http.StatusBadRequest},
		{`{"text": "", "categoryId": "01GEJ0CR9DWN7SA1QBSJE4DVKF"}`, http.StatusUnprocessableEntity},
		{`{"text": "I am cool.", "categoryId": "12345678"}`, http.StatusBadRequest},
	}

	for _, testCase := range testCases {
		payload := []byte(testCase.input)
		bodyReader := bytes.NewReader(payload)
		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/affirmations", testServerAddr), bodyReader)

		if err != nil {
			t.Errorf("Could not prepare request because of %q", err)
		}

		response, err := http.DefaultClient.Do(request)

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

		var result test.CreateAffirmationResponse
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

func TestApiUpdateAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	affirmationId, _ := ulid.Parse(affirmationIdString)
	payload := []byte(`{"text": "I am updated.", "categoryId": "01GEJ0CR9DWN7SA1QBSJE4DVKF"}`)
	bodyReader := bytes.NewReader(payload)

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/affirmations/%s", testServerAddr, affirmationId), bodyReader)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

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

	var result test.UpdateAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
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

	affirmationId, _ := ulid.Parse(affirmationIdString)

	testCases := []struct {
		input    string
		expected int
	}{
		{`I am not JSON`, http.StatusBadRequest},
		{`{"text": 123, "categoryId": "01GEJ0CR9DWN7SA1QBSJE4DVKF"}`, http.StatusBadRequest},
		{`{"text": "", "categoryId": "01GEJ0CR9DWN7SA1QBSJE4DVKF"}`, http.StatusUnprocessableEntity},
		{`{"text": "I am cool.", "categoryId": "1234567"}`, http.StatusBadRequest},
	}

	for _, testCase := range testCases {
		payload := []byte(testCase.input)
		bodyReader := bytes.NewReader(payload)
		request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/affirmations/%s", testServerAddr, affirmationId), bodyReader)

		if err != nil {
			t.Errorf("Could not prepare request because of %q", err)
		}

		response, err := http.DefaultClient.Do(request)

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

		var result test.UpdateAffirmationResponse
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

func TestApiDeleteAffirmation(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	affirmationId, _ := ulid.Parse(affirmationIdString)

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/affirmations/%s", testServerAddr, affirmationId), nil)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

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

	var result test.DeleteAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
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

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/affirmations/92233720368547758071", testServerAddr), nil)

	if err != nil {
		t.Errorf("Could not prepare request because of %q", err)
	}

	response, err := http.DefaultClient.Do(request)

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

	var result test.DeleteAffirmationResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusNotFound; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusError; got != want {
		t.Errorf("result.Status=%v, want %v", got, want)
	}
}
