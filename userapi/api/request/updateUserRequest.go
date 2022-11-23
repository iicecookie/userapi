package request

import "net/http"

type UpdateUserRequest struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
}

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }
