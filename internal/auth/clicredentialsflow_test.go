package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type MockRoundTripper func(*http.Request) (*http.Response, error)

func (m MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m(req)
}

func newMockClient(rt MockRoundTripper) *http.Client {
	return &http.Client{Transport: rt}
}

type errorReader struct {
	err error
}

func (er *errorReader) Read(_ []byte) (n int, err error) {
	return 0, er.err
}

func (er *errorReader) Close() error {
	return nil
}

func TestNewCliCredentialsFlow(t *testing.T) {
	got := NewCliCredentialsFlow(
		"http://dummy.url",
		"client-id-mock",
		"client-secret-mock",
	)
	want := CliCredentialsFlow{
		accountURL:   "http://dummy.url",
		clientID:     "client-id-mock",
		clientSecret: "client-secret-mock",
		httpClient:   http.Client{},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewCliCredentialsFlow() = %v, want %v", got, want)
	}
}

func TestCliCredentialsFlow_Authenticate(t *testing.T) {
	accountURL := "http://dummy.url"
	clientID := "dummy-client-id"
	clientSecret := "dummy-client-secret"
	type want struct {
		auth *model.Authentication
		err  bool
	}
	tests := []struct {
		name               string
		mockHTTPNewRequest func(method, url string, body io.Reader) (*http.Request, error)
		mockRoundTripper   MockRoundTripper
		want               want
	}{
		{
			name:               "Successful Authentication",
			mockHTTPNewRequest: nil,
			mockRoundTripper: func(req *http.Request) (*http.Response, error) {
				if req.URL.String() != accountURL+"/api/token" {
					t.Errorf("Unexpected request URL: got %s, want %s", req.URL.String(), accountURL+"/api/token")
				}
				if req.Method != "POST" {
					t.Errorf("Unexpected request method: got %s, want POST", req.Method)
				}
				authResponse := model.Authentication{
					AccessToken: "mock_access_token",
					TokenType:   "Bearer",
					ExpiresIn:   3600,
				}
				respBody, _ := json.Marshal(authResponse)
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBuffer(respBody)),
					Header:     make(http.Header),
					Request:    req,
				}, nil
			},
			want: want{
				auth: &model.Authentication{
					AccessToken: "mock_access_token",
					TokenType:   "Bearer",
					ExpiresIn:   3600,
				},
				err: false,
			},
		},
		{
			name: "Error creating request",
			mockHTTPNewRequest: func(_, _ string, _ io.Reader) (*http.Request, error) {
				return nil, fmt.Errorf("mock request creation error")
			},
			mockRoundTripper: func(_ *http.Request) (*http.Response, error) {
				t.Fatal("MockRoundTripper should not be called if request creation fails")
				return nil, nil
			},
			want: want{
				auth: nil,
				err:  true,
			},
		},
		{
			name:               "Error connecting to client",
			mockHTTPNewRequest: nil,
			mockRoundTripper: func(_ *http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("mock connection error")
			},
			want: want{
				auth: nil,
				err:  true,
			},
		},
		{
			name:               "Error when validating response status",
			mockHTTPNewRequest: nil,
			mockRoundTripper: func(req *http.Request) (*http.Response, error) {
				authError := commons.AuthenticationError{
					Err:            "invalid_client",
					ErrDescription: "Invalid client credentials",
				}
				respBody, _ := json.Marshal(authError)
				return &http.Response{
					StatusCode: 400,
					Status:     "400 Bad Request",
					Body:       io.NopCloser(bytes.NewBuffer(respBody)),
					Header:     make(http.Header),
					Request:    req,
				}, nil
			},
			want: want{
				auth: nil,
				err:  true,
			},
		},
		{
			name:               "Error parsing response",
			mockHTTPNewRequest: nil,
			mockRoundTripper: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(`{"access_token": "token", "malformed": `)), // Malformed JSON
					Header:     make(http.Header),
					Request:    req,
				}, nil
			},
			want: want{
				auth: nil,
				err:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockHTTPNewRequest != nil {
				originalHTTPNewRequest := httpNewRequest
				defer func() {
					httpNewRequest = originalHTTPNewRequest
				}()
				httpNewRequest = tt.mockHTTPNewRequest
			}
			c := CliCredentialsFlow{
				accountURL:   accountURL,
				clientID:     clientID,
				clientSecret: clientSecret,
				httpClient:   *newMockClient(tt.mockRoundTripper), // Inject the mock client here!
			}
			authResp, err := c.Authenticate()

			if tt.want.err {
				if err == nil {
					t.Errorf("TestAuthenticate(%s): expected error, got nil", tt.name)
				}
			} else {
				if authResp == nil {
					t.Errorf("TestAuthenticate(%s): expected non-nil auth response, got nil", tt.name)
				} else if *authResp != *tt.want.auth {
					t.Errorf("TestAuthenticate(%s): expected auth %+v, got %+v", tt.name, tt.want.auth, authResp)
				}
			}
		})
	}
}

func TestCliCredentialsFlow_createRequest(t *testing.T) {
	type want struct {
		err           bool
		method        string
		url           string
		contentType   string
		authorization string
		requestBody   string
	}
	type config struct {
		deferHTTPNewRequestToError bool
	}
	tests := []struct {
		name           string
		cliCredentials CliCredentialsFlow
		config         config
		want           want
	}{
		{
			name: "Successful Request Creation",
			cliCredentials: CliCredentialsFlow{
				clientID:     "client-id-mock",
				clientSecret: "client-secret-mock",
				accountURL:   "http://dummy.url",
			},
			want: want{
				err:           false,
				method:        "POST",
				url:           "http://dummy.url/api/token",
				contentType:   "application/x-www-form-urlencoded",
				authorization: "Basic " + base64.StdEncoding.EncodeToString([]byte("client-id-mock:client-secret-mock")),
				requestBody:   "grant_type=client_credentials",
			},
		},
		{
			name: "Empty Client ID and Secret",
			cliCredentials: CliCredentialsFlow{
				clientID:     "",
				clientSecret: "",
				accountURL:   "http://dummy.url",
			},
			want: want{
				err:           false,
				method:        "POST",
				url:           "http://dummy.url/api/token",
				contentType:   "application/x-www-form-urlencoded",
				authorization: "Basic " + base64.StdEncoding.EncodeToString([]byte(":")),
				requestBody:   "grant_type=client_credentials",
			},
		},
		{
			name: "HTTP New Request causes error",
			cliCredentials: CliCredentialsFlow{
				clientID:     "client-id-mock",
				clientSecret: "client-secret-mock",
				accountURL:   "http://dummy.url",
			},
			config: config{
				deferHTTPNewRequestToError: true,
			},
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.config.deferHTTPNewRequestToError {
				originalHTTPNewRequest := httpNewRequest
				defer func() {
					httpNewRequest = originalHTTPNewRequest
				}()
				httpNewRequest = func(_ string, _ string, _ io.Reader) (*http.Request, error) {
					return nil, fmt.Errorf("mock error")
				}
			}
			req, err := tt.cliCredentials.createRequest()

			if tt.want.err {
				if err == nil {
					t.Errorf("createRequest() error = %v, wantErr %v", err, tt.want.err)
					return
				}
				return
			}

			if req == nil {
				t.Fatal("createRequest() returned nil request on success")
			}

			if req.Method != tt.want.method {
				t.Errorf("request Method = %q, want %q", req.Method, tt.want.method)
			}
			if req.URL.String() != tt.want.url {
				t.Errorf("request URL = %q, want %q", req.URL.String(), tt.want.url)
			}
			if gotContentType := req.Header.Get("Content-Type"); gotContentType != tt.want.contentType {
				t.Errorf("request Content-Type header = %q, want %q", gotContentType, tt.want.contentType)
			}
			if gotAuth := req.Header.Get("Authorization"); gotAuth != tt.want.authorization {
				t.Errorf("request Authorization header = %q, want %q", gotAuth, tt.want.authorization)
			}

			if req.Body == nil {
				t.Fatal("request Body is nil, wanted content")
			}
			bodyBytes, readErr := io.ReadAll(req.Body)
			if readErr != nil {
				t.Fatalf("failed to read request body: %v", readErr)
			}
			_ = req.Body.Close()
			reqBodyString := string(bodyBytes)
			if reqBodyString != tt.want.requestBody {
				t.Errorf("request Body = %q, want %q", reqBodyString, tt.want.requestBody)
			}
		})
	}
}

func TestCliCredentialsFlow_parseResponse(t *testing.T) {
	c := CliCredentialsFlow{}
	tests := []struct {
		name       string
		resp       *http.Response
		wantAuth   *model.Authentication
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "Successful Parsing",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(bytes.NewBufferString(`{"access_token": "some_token", "expires_in": 3600, "token_type": "Bearer"}`)),
			},
			wantAuth: &model.Authentication{
				AccessToken: "some_token",
				ExpiresIn:   3600,
				TokenType:   "Bearer",
			},
			wantErr: false,
		},
		{
			name: "Invalid JSON",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(bytes.NewBufferString(`{"access_token": "some_token", "expires_in": 3600, "token_type": "Bearer`)),
			},
			wantAuth:   nil,
			wantErr:    true,
			wantErrMsg: "error obtaining auth response",
		},
		{
			name: "Missing Access Token",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(bytes.NewBufferString(`{"expires_in": 3600, "token_type": "Bearer"}`)),
			},
			wantAuth:   nil,
			wantErr:    true,
			wantErrMsg: "error obtaining auth response",
		},
		{
			name: "Empty Response Body",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(bytes.NewBufferString(``)),
			},
			wantAuth:   nil,
			wantErr:    true,
			wantErrMsg: "error obtaining auth response",
		},
		{
			name: "IO Read Error",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(&errorReader{err: errors.New("simulated read error")}),
			},
			wantAuth:   nil,
			wantErr:    true,
			wantErrMsg: "simulated read error",
		},
		{
			name: "HTTP Error Status with Valid Body",
			resp: &http.Response{
				StatusCode: http.StatusBadRequest,
				Status:     "400 Bad Request",
				Body:       io.NopCloser(bytes.NewBufferString(`{"access_token": "some_token"}`)),
			},
			wantAuth: &model.Authentication{
				AccessToken: "some_token",
			},
			wantErr: false,
		},
		{
			name: "HTTP Error Status with Invalid Body",
			resp: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Status:     "401 Unauthorized",
				Body:       io.NopCloser(bytes.NewBufferString(`{"error": "invalid_client"}`)),
			},
			wantAuth:   nil,
			wantErr:    true,
			wantErrMsg: "error obtaining auth response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAuth, err := c.parseResponse(tt.resp)

			if tt.wantErr {
				if err == nil {
					t.Errorf("parseResponse() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantErrMsg != "" {
					if err.Error() != tt.wantErrMsg && !strings.Contains(err.Error(), tt.wantErrMsg) {
						t.Errorf("parseResponse() error message = %q, want %q", err.Error(), tt.wantErrMsg)
					}
				}
			}

			if !reflect.DeepEqual(gotAuth, tt.wantAuth) {
				t.Errorf("parseResponse() gotAuth = %v, want %v", gotAuth, tt.wantAuth)
			}
		})
	}
}

func TestCliCredentialsFlow_validateRespStatus(t *testing.T) {
	c := CliCredentialsFlow{}
	tests := []struct {
		name         string
		resp         *http.Response
		wantErr      bool
		wantAppError commons.AppError
	}{
		{
			name: "Status Code 200 - Success",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(bytes.NewBufferString("")),
			},
			wantErr: false,
		},
		{
			name: "Status Code not 200 - Default AppError (Empty Body)",
			resp: &http.Response{
				StatusCode: http.StatusBadRequest,
				Status:     "400 Bad Request",
				Body:       io.NopCloser(bytes.NewBufferString("")),
			},
			wantErr: true,
			wantAppError: commons.AppError{
				Code:    "400 Bad Request",
				Message: "error authenticating",
				Details: "no details were provided",
			},
		},
		{
			name: "Status Code not 200 - Default AppError (Malformed JSON)",
			resp: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Status:     "401 Unauthorized",
				Body:       io.NopCloser(bytes.NewBufferString(`{"error": "invalid_token", "error_description": "token expired`)), // Malformed JSON
			},
			wantErr: true,
			wantAppError: commons.AppError{
				Code:    "401 Unauthorized",
				Message: "error authenticating",
				Details: "no details were provided",
			},
		},
		{
			name: "Status Code not 200 - Custom AppError from AuthenticationError",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Status:     "403 Forbidden",
				Body:       io.NopCloser(bytes.NewBufferString(`{"error": "access_denied", "error_description": "User is not authorized for this action."}`)),
			},
			wantErr: true,
			wantAppError: commons.AppError{
				Code:    "403 Forbidden",
				Message: "access_denied",
				Details: "User is not authorized for this action.",
			},
		},
		{
			name: "Status Code not 200 - IO ReadAll Error",
			resp: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Status:     "500 Internal Server Error",
				Body:       io.NopCloser(&errorReader{err: errors.New("simulated read error")}),
			},
			wantErr: true,
			wantAppError: commons.AppError{
				Code:    "500 Internal Server Error",
				Message: "error authenticating",
				Details: "no details were provided",
			},
		},
		{
			name: "Status Code not 200 - JSON without Err field",
			resp: &http.Response{
				StatusCode: http.StatusBadRequest,
				Status:     "400 Bad Request",
				Body:       io.NopCloser(bytes.NewBufferString(`{"some_other_field": "value"}`)),
			},
			wantErr: true,
			wantAppError: commons.AppError{
				Code:    "400 Bad Request",
				Message: "error authenticating",
				Details: "no details were provided",
			},
		},
		{
			name: "Status Code not 200 - JSON with empty Err field",
			resp: &http.Response{
				StatusCode: http.StatusBadRequest,
				Status:     "400 Bad Request",
				Body:       io.NopCloser(bytes.NewBufferString(`{"error": "", "error_description": "Some description"}`)),
			},
			wantErr: true,
			wantAppError: commons.AppError{
				Code:    "400 Bad Request",
				Message: "error authenticating",
				Details: "no details were provided",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.validateRespStatus(tt.resp)

			if tt.wantErr {
				if err == nil {
					t.Errorf("validateRespStatus() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				var appErr commons.AppError
				ok := errors.As(err, &appErr)
				if !ok {
					t.Fatalf("validateRespStatus() returned error of unexpected type %T, want commons.AppError", err)
				}

				if appErr.Code != tt.wantAppError.Code {
					t.Errorf("validateRespStatus() AppError Code = %q, want %q", appErr.Code, tt.wantAppError.Code)
				}
				if appErr.Message != tt.wantAppError.Message {
					t.Errorf("validateRespStatus() AppError Message = %q, want %q", appErr.Message, tt.wantAppError.Message)
				}
				if appErr.Details != tt.wantAppError.Details {
					t.Errorf("validateRespStatus() AppError Details = %q, want %q", appErr.Details, tt.wantAppError.Details)
				}
			}
		})
	}
}
