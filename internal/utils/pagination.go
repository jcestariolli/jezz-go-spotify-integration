package utils

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
)

func ValidatePaginationParams(
	limit *model.Limit,
	offset *model.Offset,
) (*model.Limit, *model.Offset, error) {
	if limit != nil && ((*limit).Int() < 0 || (*limit).Int() < 0) {
		err := fmt.Errorf("limit is invalid - must be between 0 and 50")
		return nil, nil, err
	}

	if offset != nil && ((*offset).Int() < 0) {
		err := fmt.Errorf("offset is invalid - must be above 0")
		return nil, nil, err
	}

	return limit, offset, nil
}
