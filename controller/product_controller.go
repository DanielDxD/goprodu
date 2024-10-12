package controller

import (
	"net/http"
	"strconv"

	"github.com/DanielDxD/goprodu/model"
	"github.com/DanielDxD/goprodu/usecases"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecases.ProductUseCase
}

func NewProductController(usecase usecases.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Ocorreu um erro",
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Parâmetros inválidos.",
		})
		return
	}

	newProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Ocorreu um erro",
		})
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)
}
func (p *productController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Nenhum ID foi enviado",
		})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "ID inválido",
		})
		return
	}

	product, err := p.productUseCase.GetProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Ocorreu um erro",
		})
		return
	}
	if product == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Produto não encontrado.",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
