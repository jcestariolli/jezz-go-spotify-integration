package commons

import (
	"errors"
	"testing"
)

func TestResourceError_Error_HappyPath(t *testing.T) {
	re := ResourceError{
		Status:  404,
		Message: "Resource Not Found",
	}

	expectedJSON := `{"status":404,"message":"Resource Not Found"}`
	actualError := re.Error()

	if actualError != expectedJSON {
		t.Errorf("Expected error string '%s', but got '%s'", expectedJSON, actualError)
	}
}

func TestResourceError_Error_MarshalErrorPath(t *testing.T) {
	originalJSONMarshal := jsonMarshal
	defer func() {
		jsonMarshal = originalJSONMarshal
	}()
	jsonMarshal = func(_ interface{}) ([]byte, error) {
		return nil, errors.New("mock marshal error")
	}

	re := ResourceError{
		Status:  400,
		Message: "Bad Request",
	}

	expectedErrorMessage := "resource error, no details provided"
	actualError := re.Error()
	if actualError != expectedErrorMessage {
		t.Errorf("Expected error string '%s' when json.Marshal fails, but got '%s'", expectedErrorMessage, actualError)
	}
}

func TestAppError_Error_HappyPath(t *testing.T) {
	ae := AppError{
		Code:    "1",
		Message: "Some message",
		Details: "Some details",
	}

	expectedJSON := `{"code":"1","message":"Some message","details":"Some details"}`
	actualError := ae.Error()

	if actualError != expectedJSON {
		t.Errorf("Expected error string '%s', but got '%s'", expectedJSON, actualError)
	}
}

func TestAppError_Error_MarshalErrorPath(t *testing.T) {
	originalJSONMarshal := jsonMarshal
	defer func() {
		jsonMarshal = originalJSONMarshal
	}()
	jsonMarshal = func(_ interface{}) ([]byte, error) {
		return nil, errors.New("mock marshal error")
	}

	ae := AppError{
		Code:    "1",
		Message: "Some message",
		Details: "Some details",
	}

	expectedErrorMessage := "app error, no details provided"
	actualError := ae.Error()
	if actualError != expectedErrorMessage {
		t.Errorf("Expected error string '%s' when json.Marshal fails, but got '%s'", expectedErrorMessage, actualError)
	}
}
