package controller

import "net/http"

type OKDto struct {
	OK bool `json:"ok"`
}

func NewOkDto() *OKDto {
	return &OKDto{
		OK: true,
	}
}

func (hr *OKDto) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}
