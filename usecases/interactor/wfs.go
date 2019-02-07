package interactor

import (
	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/ports"
)

// WFSInteractor is the interactor for MAO controller.
type WFSInteractor struct {
	Config     ports.Config
	WFSHandler domain.WFSHandler
}

// NewWFSInteractor creates the new interactor for the WFS usecase.
func NewWFSInteractor(conf ports.Config, handler domain.WFSHandler) *WFSInteractor {
	return &WFSInteractor{
		Config:     conf,
		WFSHandler: handler,
	}
}

// Start starts MAO wavefront sensor.
func (i *WFSInteractor) Start() error {
	if err := i.WFSHandler.Initialize(i.Config); err != nil {
		return err
	}
	if err := i.WFSHandler.Start(); err != nil {
		return err
	}
	return nil
}

// Halt stops MAO wavefront sensor.
func (i *WFSInteractor) Halt() error {
	if err := i.WFSHandler.Halt(); err != nil {
		return err
	}
	if err := i.WFSHandler.Finalize(); err != nil {
		return err
	}
	return nil
}
