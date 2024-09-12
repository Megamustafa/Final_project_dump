package controllers

import (
	"aquaculture/models"
	"aquaculture/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionDetailController struct {
	service services.TransactionDetailService
}

func InitTransactionDetailController() TransactionDetailController {
	return TransactionDetailController{
		service: services.InitTransactionDetailService(),
	}
}

func (tdc *TransactionDetailController) GetAll(c echo.Context) error {
	transactionDetails, err := tdc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetdching transaction details",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.TransactionDetail]{
		Status:  "success",
		Message: "all transaction details",
		Data:    transactionDetails,
	})
}

func (tdc *TransactionDetailController) GetByID(c echo.Context) error {
	transactionDetailID := c.Param("id")

	transactionDetail, err := tdc.service.GetByID(transactionDetailID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "transaction detail not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.TransactionDetail]{
		Status:  "success",
		Message: "transaction detail found",
		Data:    transactionDetail,
	})
}

func (tdc *TransactionDetailController) Create(c echo.Context) error {
	var tdReq models.TransactionDetailRequest

	if err := c.Bind(&tdReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := tdReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	transactionDetail, err := tdc.service.Create(tdReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a transaction detail",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.TransactionDetail]{
		Status:  "success",
		Message: "transaction detail created",
		Data:    transactionDetail,
	})
}

func (tdc *TransactionDetailController) Update(c echo.Context) error {
	var tdReq models.TransactionDetailRequest

	if err := c.Bind(&tdReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	transactionDetailID := c.Param("id")

	err := tdReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	transactionDetail, err := tdc.service.Update(tdReq, transactionDetailID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a transaction detail",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.TransactionDetail]{
		Status:  "success",
		Message: "transaction detail updated",
		Data:    transactionDetail,
	})
}
