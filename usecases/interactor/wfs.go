package interactor

import (
	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/ports"
)

// WFSInteractor is the interactor for MAO controller.
type WFSInteractor struct {
	Config  ports.Config
	Handler domain.WFSHandler
}

// NewWFSInteractor creates the new interactor for the WFS usecase.
func NewWFSInteractor(conf ports.Config, handler domain.WFSHandler) *WFSInteractor {
	return &WFSInteractor{
		Config:  conf,
		Handler: handler,
	}
}

// Start starts MAO wavefront sensor.
func (i *WFSInteractor) Start() error {
	if err := i.Handler.Initialize(i.Config); err != nil {
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
