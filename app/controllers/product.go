package controllers

import (
	"aquaculture/models"
	"aquaculture/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service services.ProductService
}

func InitProductController() ProductController {
	return ProductController{
		service: services.InitProductService(),
	}
}

func (pc *ProductController) GetAll(c echo.Context) error {
	products, err := pc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching products",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Product]{
		Status:  "success",
		Message: "all products",
		Data:    products,
	})
}

func (pc *ProductController) GetByID(c echo.Context) error {
	productID := c.Param("id")

	product, err := pc.service.GetByID(productID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "product not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Product]{
		Status:  "success",
		Message: "product found",
		Data:    product,
	})
}

func (pc *ProductController) Create(c echo.Context) error {
	var productReq models.ProductRequest

	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := productReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	product, err := pc.service.Create(productReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a product",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Product]{
		Status:  "success",
		Message: "product created",
		Data:    product,
	})
}

func (pc *ProductController) Update(c echo.Context) error {
	var productReq models.ProductRequest

	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	productID := c.Param("id")

	err := productReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	product, err := pc.service.Update(productReq, productID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a product",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Product]{
		Status:  "success",
		Message: "product updated",
		Data:    product,
	})
}

func (pc *ProductController) Delete(c echo.Context) error {
	productID := c.Param("id")

	err := pc.service.Delete(productID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete a product",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "product deleted",
	})
}
