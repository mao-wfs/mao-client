package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/mao-wfs/mao-ctrl/adapters/controller"
	"github.com/mao-wfs/mao-ctrl/usecases/input"
)

// InitializeWFS initializes MAO-WFS.
// It is implemented to use as a handler function of "Echo".
func InitializeWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(input.InitRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := ctrl.InitializeWFS(req); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, req)
	}
}

// FinalizeWFS finalizes MAO-WFS.
// It is implemented to use as a handler function of "Echo".
func FinalizeWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ctrl.FinalizeWFS(); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}

// StartWFS starts MAO-WFS.
// It is implemented to use as a handler function of "Echo".
func StartWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(input.StartRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := ctrl.StartWFS(req); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, req)
	}
}

// HaltWFS halts MAO-WFS.
// It is implemented to use as a handler function of "Echo".
func HaltWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ctrl.HaltWFS(); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
