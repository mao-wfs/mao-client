package interactor

import (
	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/ports"
)

// WFSInteractor is the interactor for MAO controller.
type WFSInteractor struct {
	ConfigPort ports.ConfigPort
	Handler    domain.WFSHandler
}

// NewWFSInteractor creates the new interactor for the WFS usecase.
func NewWFSInteractor(confPort ports.ConfigPort, handler domain.WFSHandler) *WFSInteractor {
	return &WFSInteractor{
		ConfigPort: confPort,
		Handler:    handler,
	}
}

// Start starts MAO wavefront sensor.
func (i *WFSInteractor) Start() error {
	if err := i.Handler.Initialize(i.ConfigPort); err != nil {
		return err
	}
	if err := i.Handler.Start(); err != nil {
		return err
	}
	return nil
}

// Halt stops MAO wavefront sensor.
func (i *WFSInteractor) Halt() error {
	if err := i.Handler.Halt(); err != nil {
		return err
	}
	if err := i.Handler.Finalize(); err != nil {
		return err
	}
	return nil
}
