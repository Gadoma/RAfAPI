package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gadoma/rafapi/internal/category/test"
	"github.com/oklog/ulid/v2"
)

var categoryIdString = "01GEJ0CR9DWN7SA1QBSJE4DVKF"

func TestApiGetCategories(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	response, err := http.Get("http://" + testServerAddr + "/categories")

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

	var result test.GetCategoriesResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data[0].Id.String(), categoryIdString; got != want {
		t.Errorf("result.Data[0].Id=%v, want %v", got, want)
	} else if got, want := len(result.Data), 5; got != want {
		t.Errorf("len=%v, want %v", got, want)
	} else if got, want := result.Count, 5; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiGetCategory(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryId, _ := ulid.Parse(categoryIdString)

	response, err := http.Get(fmt.Sprintf("http://%s/categories/%s", testServerAddr, categoryId))

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

	var result test.GetCategoryResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Could not decode response because of %q", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("response.StatusCode=%v, want %v", got, want)
	} else if got, want := result.Status, statusOk; got != want {
		t.Errorf("result.Status=%v, want %v, message %v", got, want, result.Message)
	} else if got, want := result.Data.Id.String(), categoryIdString; got != want {
		t.Errorf("result.Data[0].Id=%v, want %v", got, want)
	} else if got, want := result.Count, 1; got != want {
		t.Errorf("n=%v, want %v", got, want)
	}
}

func TestApiGetCategoryError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	testCases := []struct {
		input    string
		expected int
	}{
		{"01HEH0CT8THFY2V8VYJJKAGP4N", http.StatusNotFound},
		{ulid.Make().String(), http.StatusNotFound},
	}

	for _, testCase := range testCases {
		response, err := http.Get(fmt.Sprintf("http://%s/categories/%s", testServerAddr, testCase.input))

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

		var result test.GetCategoryResponse
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

func TestApiCreateCategory(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	payload := []byte(`{"name": "Created"}`)
	bodyReader := bytes.NewReader(payload)

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/categories", testServerAddr), bodyReader)

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

	var result test.CreateCategoryResponse
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

func TestApiCreateCategoryError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	testCases := []struct {
		input    string
		expected int
	}{
		{`I am not JSON`, http.StatusBadRequest},
		{`{"name": 123}`, http.StatusBadRequest},
		{`{"name": ""}`, http.StatusUnprocessableEntity},
	}

	for _, testCase := range testCases {
		payload := []byte(testCase.input)
		bodyReader := bytes.NewReader(payload)
		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/categories", testServerAddr), bodyReader)

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

		var result test.CreateCategoryResponse
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

func TestApiUpdateCategory(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryId, _ := ulid.Parse(categoryIdString)
	payload := []byte(`{"name": "Updated"}`)
	bodyReader := bytes.NewReader(payload)

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/categories/%s", testServerAddr, categoryId), bodyReader)

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

	var result test.UpdateCategoryResponse
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

func TestApiUpdateCategoryError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryId, _ := ulid.Parse(categoryIdString)

	testCases := []struct {
		input    string
		expected int
	}{
		{`I am not JSON`, http.StatusBadRequest},
		{`{"name": 123}`, http.StatusBadRequest},
		{`{"name": ""}`, http.StatusUnprocessableEntity},
	}

	for _, testCase := range testCases {
		payload := []byte(testCase.input)
		bodyReader := bytes.NewReader(payload)
		request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/categories/%s", testServerAddr, categoryId), bodyReader)

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

		var result test.UpdateCategoryResponse
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

func TestApiDeleteCategory(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	categoryId, _ := ulid.Parse("01GEJ0CSKKVPY3PR8VXJMPDQAY")

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/categories/%s", testServerAddr, categoryId), nil)

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

	var result test.DeleteCategoryResponse
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

func TestApiDeleteCategoryError(t *testing.T) {
	app := MustRunMain(t)
	defer MustCloseMain(t, app)

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/categories/92233720368547758071", testServerAddr), nil)

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

	var result test.DeleteCategoryResponse
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
