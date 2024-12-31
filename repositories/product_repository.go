package repositories

import (
	"fmt"
	"log"
	"pencarian_elektronik/models"
)

func FindProductsByName(name string) ([]models.Product, error) {
	query := "SELECT id, name, description, image_url FROM products WHERE name LIKE ?"

	rows, err := DB.Query(query, "%"+name+"%")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return products, nil
}
