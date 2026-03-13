package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode  int
	Code        string
	Description string
}

func (e APIError) Error() string {
	return fmt.Sprintf("scrapper api error: status=%d code=%s description=%s", e.StatusCode, e.Code, e.Description)
}

func parseErrorResponse(resp *http.Response) error {
	var dto ApiErrorResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return APIError{
		StatusCode:  resp.StatusCode,
		Code:        dto.Code,
		Description: dto.Description,
	}
}