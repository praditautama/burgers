package payloads

import "github.com/opam22/burgers/models"

type SuccessGetMenu struct {
	Code     int         `json:"code"`
	Response models.Menu `json:"response"`
}

type SuccessGetAllMenu struct {
	Code     int           `json:"code"`
	Response []models.Menu `json:"response"`
}

type SuccessGetMenuReceipt struct {
	Code     int         `json:"code"`
	Response MenuReceipt `json:"response"`
}

type MenuReceipt struct {
	Menu    models.Menu             `json:"menu"`
	Receipt []models.MenuIngredient `json:"menu_receipt"`
}

type SuccessOrder struct {
	Code     int          `json:"code"`
	Response models.Order `json:"response"`
}
