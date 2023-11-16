package handler

import (
	"synapsis/domain/rest"
	"synapsis/service"
	"synapsis/utils/http_response"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	respWriter  http_response.IResponseWriter
	cartService service.ICartService
}

type ICartHandler interface {
	GetCartItems(ctx *gin.Context)
	// RemoveItemFromCart(ctx *gin.Context)
}

func NewCartHandler(respWriter http_response.IResponseWriter, cartService service.ICartService) ICartHandler {
	return &CartHandler{
		respWriter:  respWriter,
		cartService: cartService,
	}
}

// @Summary get product list from current cart
// @Tags Cart
// @Router /cart/items [get]
// @Security BearerAuth
// @Success 200 {object} rest.BaseJSONResp{}
func (slf *CartHandler) GetCartItems(ctx *gin.Context) {
	user_id := ctx.GetString("user_id")
	orders, err := slf.cartService.GetCartItems(user_id)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	dto := rest.GetCartItemsResp{}
	resp := []rest.GetCartItemsResp{}
	resp = append(resp, dto.ParseFromEntityList(orders)...)

	slf.respWriter.HTTPJson(ctx, resp)
}

// @Summary remove item from cart
// @Tags Cart
// @Router /cart/items/{id} [delete]
// @Security BearerAuth
// @Success 200 {object} rest.BaseJSONResp{}
// func (slf *CartHandler) RemoveItemFromCart(ctx *gin.Context) {
// 	product_id := ctx.Param("id")

// }
