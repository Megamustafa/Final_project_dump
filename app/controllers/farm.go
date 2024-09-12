package controllers

import (
	"aquaculture/models"
	"aquaculture/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FarmController struct {
	service services.FarmService
}

func InitFarmController() FarmController {
	return FarmController{
		service: services.InitFarmService(),
	}
}

func (pc *FarmController) GetAll(c echo.Context) error {
	farms, err := pc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching farms",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Farm]{
		Status:  "success",
		Message: "all farms",
		Data:    farms,
	})
}

func (pc *FarmController) GetByID(c echo.Context) error {
	farmID := c.Param("id")

	farm, err := pc.service.GetByID(farmID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "farm not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Farm]{
		Status:  "success",
		Message: "farm found",
		Data:    farm,
	})
}

func (pc *FarmController) Create(c echo.Context) error {
	var farmReq models.FarmRequest

	if err := c.Bind(&farmReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := farmReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farm, err := pc.service.Create(farmReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a farm",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Farm]{
		Status:  "success",
		Message: "farm created",
		Data:    farm,
	})
}

func (pc *FarmController) Update(c echo.Context) error {
	var farmReq models.FarmRequest

	if err := c.Bind(&farmReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	farmID := c.Param("id")

	err := farmReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farm, err := pc.service.Update(farmReq, farmID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a farm",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Farm]{
		Status:  "success",
		Message: "farm updated",
		Data:    farm,
	})
}

func (pc *FarmController) Delete(c echo.Context) error {
	farmID := c.Param("id")

	err := pc.service.Delete(farmID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete a farm",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "farm deleted",
	})
}
