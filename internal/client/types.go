package client

import "jezz-go-spotify-integration/internal/model"

type HTTPClient interface {
	DoRequest(
		method model.HTTPMethod,
		url string,
		queryParams *model.QueryParams,
		accessToken model.AccessToken,
		responseTypedOutput any,
	) error
}
