package app

import (
	"time"

	"github.com/mrityunjaygr8/shorty/db"
)

func (a *App) Lookup(token string) (string, bool, error) {
	value, found, err := db.LookupUsingToken(token, *a.DB)
	if err != nil {
		return "", false, err
	}

	return value.Long_URL, found, nil
}

func (a *App) Create(url string) (string, error) {
	value, present, err := db.LookupUsingURL(url, *a.DB)
	if err != nil {
		return "", err
	}
	if present && value.Expiry_at.After(time.Now().UTC()) {
		return value.Token, nil

	} else {
		value, err := db.Create(url, *a.DB)
		if err != nil {
			return "", err
		}
		return value.Token, nil
	}
}
