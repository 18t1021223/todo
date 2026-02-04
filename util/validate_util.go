package util

import (
	"encoding/json"
	"net/http"
	"todo_list_roadmap/config"
	"todo_list_roadmap/handle/response"
)

func BindAndValidate(r *http.Request, out any) error {
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		return response.ErrInvalidReq
	}

	if err := config.ValidatorG.Struct(out); err != nil {
		return err
	}
	return nil
}
