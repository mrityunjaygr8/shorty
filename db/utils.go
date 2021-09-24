package db

import (
	"github.com/speps/go-hashids/v2"
)

func encodeHashid(id int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = "Mera Salt"
	hd.MinLength = 8
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	token, _ := h.Encode([]int{id})
	return token, nil
}
