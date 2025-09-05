package cart

import (
	"fmt"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
)

func getCartItemsIDs(items []interfaces.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIds[i] = item.ProductID
	}

	return productIds, nil
}

func checkIfCartIsInStock(items []interfaces.CartItem, products map[int]interfaces.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(items []interfaces.CartItem, products map[int]interfaces.Product) float64 {
	var total float64

	for _, item := range items {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
