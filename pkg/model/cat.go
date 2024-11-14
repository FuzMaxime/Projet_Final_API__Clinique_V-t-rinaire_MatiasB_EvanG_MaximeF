package model

import (
	"errors"
	"net/http"
)

type CatRequest struct {
    Name string `json:"cat_name"`
	Age int `json:"cat_age"`
	Race string `json:"cat_race"`
	Weight int `json:"cat_weight"`
}

func (a *CatRequest) Bind(r *http.Request) error {
    if a.Age < 0 {
        return errors.New("human_age must be a positive integer")
    }
    return nil
}

type CatResponse struct {
    CatAge int `json:"cat_age"`
}