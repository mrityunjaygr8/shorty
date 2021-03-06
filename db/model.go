package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

const CONFIG_EXPIRY_HOURS = 6

type Shortener struct {
	Token      string `gorm:"primaryKey;size:8"`
	Long_URL   string `gorm:"size:1000"`
	Created_at time.Time
	Expiry_at  time.Time
	Hits       uint
	ID         uint
}

func Create(long string, db gorm.DB) (Shortener, error) {
	var id int
	last := Shortener{}
	result := db.Limit(1).Order("id desc").First(&last)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			id = 1
		} else {
			return Shortener{}, result.Error
		}
	} else {
		id = int(last.ID + 1)
	}

	token, err := encodeHashid(id)
	if err != nil {
		return Shortener{}, err
	}
	now := time.Now().UTC()
	expiry := now.Add(time.Hour * CONFIG_EXPIRY_HOURS)
	short := Shortener{
		Token:      token,
		Long_URL:   long,
		Created_at: now,
		Expiry_at:  expiry,
		Hits:       0,
		ID:         uint(id),
	}
	if err := db.Create(&short).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// Case where the token already exists in the database
			// https://github.com/jackc/pgerrcode/blob/master/errcode.go#L142
			if pgErr.Code == "23505" {
				return Shortener{}, fmt.Errorf("URL with token: %s already exists", short.Token)
			}
		}
	}

	return short, nil
}

func LookupUsingURL(url string, db gorm.DB) (Shortener, bool, error) {
	value := Shortener{}
	result := db.Order("created_at desc").Find(&value, "long_url = ?", url)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Shortener{}, false, nil
		} else {
			return Shortener{}, false, result.Error
		}
	}
	return value, true, nil
}
func LookupUsingToken(token string, db gorm.DB) (Shortener, bool, error) {
	long := Shortener{}
	result := db.First(&long, "token = ?", token)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Shortener{}, false, nil
	}

	long.Hits += 1
	long.Expiry_at = time.Now().UTC().Add(time.Hour * CONFIG_EXPIRY_HOURS)
	result = db.Save(&long)
	if result.Error != nil {
		return Shortener{}, false, result.Error
	}

	return long, true, nil
}
