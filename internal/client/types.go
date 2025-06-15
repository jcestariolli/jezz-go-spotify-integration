package client

import "jezz-go-spotify-integration/internal/model"

type HTTPApiClient interface {
	DoRequest(
		method model.HTTPMethod,
		url string,
		queryParams *model.QueryParams,
		contentType string,
		accessToken *model.AccessToken,
		responseTypedOutput any,
	) error
}
