package productdb

import (
	"time"

	"github.com/ardanlabs/service/business/core/product"
)

// dbProduct represents an individual product.
type dbProduct struct {
	ID          string    `db:"product_id"`   // Unique identifier.
	Name        string    `db:"name"`         // Display name of the product.
	Cost        int       `db:"cost"`         // Price for one item in cents.
	Quantity    int       `db:"quantity"`     // Original number of items available.
	Sold        int       `db:"sold"`         // Aggregate field showing number of items sold.
	Revenue     int       `db:"revenue"`      // Aggregate field showing total cost of sold items.
	UserID      string    `db:"user_id"`      // ID of the user who created the product.
	DateCreated time.Time `db:"date_created"` // When the product was added.
	DateUpdated time.Time `db:"date_updated"` // When the product record was last modified.
}

// =============================================================================

func toDBProduct(prd product.Product) dbProduct {
	return dbProduct(prd)
}

func toCoreProduct(dbPrd dbProduct) product.Product {
	return product.Product(dbPrd)
}

func toCoreProductSlice(dbProducts []dbProduct) []product.Product {
	prds := make([]product.Product, len(dbProducts))
	for i, dbPrd := range dbProducts {
		prds[i] = toCoreProduct(dbPrd)
	}
	return prds
}
