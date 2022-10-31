package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strings"
)

type PtlistParams struct {
	Period    string `query:"period"`
	Timezone  string `query:"tz"`
	TimeStart string `query:"t1"`
	TimeEnd   string `query:"t2"`
}

type PtlistParamsDTO struct {
	Period    string
	Timezone  string
	TimeStart string
	TimeEnd   string
}

type ValidationError struct {
	Status      string
	Description string
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, randate("Europe/Athens"))
	})

	e.GET("/ptlist", func(c echo.Context) (err error) {
		ptParams := new(PtlistParams)
		if err = c.Bind(ptParams); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if !periodExists(ptParams.Period) {
			validationError := ValidationError{
				Status:      "error",
				Description: "Unsupported period",
			}
			return c.JSON(http.StatusNotFound, validationError)
		}

		// Load into separate struct for security
		ptlistDTO := PtlistParamsDTO{
			Period:    ptParams.Period,
			Timezone:  ptParams.Timezone,
			TimeStart: ptParams.TimeStart,
			TimeEnd:   ptParams.TimeEnd,
		}

		return c.JSON(http.StatusUnauthorized, findPeriodicTasks(ptlistDTO.Period, ptlistDTO.Timezone, ptlistDTO.TimeStart, ptlistDTO.TimeEnd))
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func periodExists(periodParam string) bool {
	supportedPeriods := []string{"h", "d", "m", "y"}
	periodValid := false
	for _, sp := range supportedPeriods {
		if strings.Contains(periodParam, sp) {
			periodValid = true
		}
	}
	return periodValid
}
