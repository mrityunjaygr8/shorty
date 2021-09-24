package app

import "github.com/mrityunjaygr8/shorty/db"

func (a *App) Lookup(token string) (string, bool, error) {
	value, found, err := db.LookupUsingToken(token, *a.DB)
	if err != nil {
		return "", false, err
	}

	return value.Long_URL, found, nil
}

func (a *App) Create(url string) (string, error) {
	value, err := db.Create(url, *a.DB)
	if err != nil {
		return "", err
	}

	return value.Token, nil
}
