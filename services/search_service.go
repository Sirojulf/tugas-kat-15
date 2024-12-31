package services

import (
	"errors"
	"pencarian_elektronik/models"
	"pencarian_elektronik/repositories"
)

var ErrQueryTooShort = errors.New("query too short, must be at least 3 characters")

func SearchProducts(query string) ([]models.Product, error) {
	if len(query) < 3 {
		return nil, ErrQueryTooShort
	}

	products, err := repositories.FindProductsByName(query)
	if err != nil {
		return nil, err
	}

	return products, nil
}
