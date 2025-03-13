package http

import (
	"api-order/src/Product/application"
	"api-order/src/Product/domain/ports"
	"api-order/src/Product/infraestructure/adapters"
	"api-order/src/Product/infraestructure/http/controllers"
	"log"
)

var ProductRepository ports.IProduct

func init(){
	var err error
	ProductRepository, err = adapters.NewProductRepositoryMysql()
	if err != nil {
		log.Fatalf("Error al inicializar el repositorio de productos: %s", err)
	} 
}


func SetUpCreateProduct() *controllers.CreateProductController {
	productService := application.NewCreateProductUseCase(ProductRepository)
	return controllers.NewCreateProductController(productService)
}

func SetUpGetAllProducts() *controllers.GetAllProductsController {
	productService := application.NewGetAllProductsUseCase(ProductRepository)
	return controllers.NewGetAllProductsController(productService)
}