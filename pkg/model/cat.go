package model

import (
	"net/http"
)

type CatRequest struct {
	Name   string `json:"cat_name"`
	Age    int    `json:"cat_age"`
	Race   string `json:"cat_race"`
	Weight int    `json:"cat_weight"`
}

func (a *CatRequest) Bind(r *http.Request) error {

	return nil
}

type CatResponse struct {
	Name   string `json:"cat_name"`
	Age    int    `json:"cat_age"`
	Race   string `json:"cat_race"`
	Weight int    `json:"cat_weight"`
}
