package cart

import (
	"fmt"
	"net/http"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/MKSinghDev/go-ecom/src/service/auth"
	"github.com/MKSinghDev/go-ecom/src/utils"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	var cart interfaces.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	ids, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ps, err := h.productRepo.GetProductsByIDs(ids)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
