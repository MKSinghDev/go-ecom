package product

import (
	"fmt"
	"net/http"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/MKSinghDev/go-ecom/src/utils"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload interfaces.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.repo.CreateProduct(interfaces.CreateProductPayload{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, products)
}
