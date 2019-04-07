package controller

import (
	"github.com/mao-wfs/mao-client/adapters/gateway/device"
	"github.com/mao-wfs/mao-client/usecases/input"
	"github.com/mao-wfs/mao-client/usecases/interactor"
)

// WFSController is the controller of MAO-WFS.
type WFSController struct {
	InputPort input.WFSInputPort
}

// NewWFSController returns "WFSController".
func NewWFSController(handler *device.WFSHandler) *WFSController {
	return &WFSController{
		InputPort: interactor.NewWFSInteractor(handler),
	}
}

// InitializeWFS initializes MAO-WFS.
func (c *WFSController) InitializeWFS(req *input.InitRequest) error {
	return c.InputPort.InitializeWFS(req)
}

// FinalizeWFS finalizes MAO-WFS.
func (c *WFSController) FinalizeWFS() error {
	return c.InputPort.FinalizeWFS()
}

// StartWFS starts MAO-WFS.
func (c *WFSController) StartWFS(req *input.StartRequest) error {
	return c.InputPort.StartWFS(req)
}

// HaltWFS halts MAO-WFS.
func (c *WFSController) HaltWFS() error {
	return c.InputPort.HaltWFS()
}
