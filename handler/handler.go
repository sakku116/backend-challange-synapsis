package handler

import (
	"synapsis/domain/rest"
	"synapsis/service"
	"synapsis/utils/http_response"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	respWriter     http_response.IResponseWriter
	productService service.IProductService
}

type IProductHandler interface {
	GetList(ctx *gin.Context)
}

func NewProductHandler(respWriter http_response.IResponseWriter, productService service.IProductService) IProductHandler {
	return &ProductHandler{
		respWriter:     respWriter,
		productService: productService,
	}
}

// Get Product List
// @Summary get product list
// @Tags Product
// @Success 200 {object} rest.BaseJSONResp{data=[]rest.GetProductListResp}
// @Router /products [get]
// @param query  query  rest.GetProductListReq  false "query"
// @Security BearerAuth
func (slf *ProductHandler) GetList(ctx *gin.Context) {
	var query rest.GetProductListReq
	err := ctx.Bind(&query)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	products, err := slf.productService.GetList(query.Category, query.Search, query.Page, query.Limit, query.SortBy, query.SortOrder)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	var result []*rest.GetProductListResp
	dto := rest.GetProductListResp{}
	result = append(result, dto.ParseFromEntityList(&products)...)

	slf.respWriter.HTTPJson(ctx, result)
}
