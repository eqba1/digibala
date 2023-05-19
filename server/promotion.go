package server

import (
	"digibala/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

// promotion routes comes here
func promotionHandler(c echo.Context) error {
	promotion, err := json.Marshal(&models.Promotion{})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, promotion)
}

func findPromotionHandler(c echo.Context) error {
	promotionID := c.Param("id")
	prid, _ := strconv.Atoi(promotionID)

	category := new(models.Category)
	promotion := models.Promotion{
		ID:          prid,
		Name:        "Promotion 1",
		Category:    *category,
		Description: "Promotion 1 description",
		Discount:    0.2,
		ExpireDate:  time.Now(),
	}

	return c.JSON(http.StatusOK, promotion)
}

func deletePromotionHandler(c echo.Context) error {
	promotionID := c.Param("id")
	prid, _ := strconv.Atoi(promotionID)

	msg, _ := json.Marshal(fmt.Sprintf("Delete the promotion %d from the database", prid))
	// Delete the promotion
	return c.JSON(http.StatusOK, msg)
}

func createPromotionHandler(c echo.Context) error {
	categoryName := c.Param("name")

	category := &models.Category{
		ID:   1,
		Name: categoryName,
	}

	promotion := models.Promotion{
		ID:          1,
		Name:        "Promotion 1",
		Category:    *category,
		Description: "Promotion 1 description",
		Discount:    0.2,
		ExpireDate:  time.Now(),
	}

	if err := c.Bind(promotion); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, promotion)
}

func updatePromotionHandler(c echo.Context) error {
	promotionID := c.Param("id")
	prid, _ := strconv.Atoi(promotionID)

	msg, _ := json.Marshal(fmt.Sprintf("Update the promotion %d in the database", prid))
	// Delete the promotion
	return c.JSON(http.StatusOK, msg)
}

func promotionRoutes(e *echo.Echo) {
	fmt.Println(models.Promotion{})
	e.GET("/promotion", promotionHandler)
	e.GET("/promotion/:id", findPromotionHandler)
	e.DELETE("/promotion/:id", deletePromotionHandler)
	e.PUT("/promotion/category/:name", createPromotionHandler)
	e.PATCH("/promotion/:id", updatePromotionHandler)
}