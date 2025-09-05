package cart

import (
	"github.com/MKSinghDev/go-ecom/src/interfaces"
)

func (h *Handler) createOrder(ps []interfaces.Product, items []interfaces.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]interfaces.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// Check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}

	// Calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// Reduce the quantity of the products in our db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productRepo.UpdateProduct(product)
	}

	// Create the order
	orderID, err := h.orderRepo.CreateOrder(interfaces.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "Some address",
	})
	if err != nil {
		return 0, 0, err
	}

	// Create order items
	for _, item := range items {
		h.orderRepo.CreateOrderItem(interfaces.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}
