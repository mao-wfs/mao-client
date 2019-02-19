package interactor

import (
	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/ports"
)

// WFSInteractor is the interactor for MAO controller.
type WFSInteractor struct {
	Handler domain.WFSHandler
}

// NewWFSInteractor creates the new interactor for the WFS usecase.
func NewWFSInteractor(handler domain.WFSHandler) *WFSInteractor {
	return &WFSInteractor{
		Handler: handler,
	}
}

// Initialize initialize MAO wavefront sensor.
func (i *WFSInteractor) Initialize(conf *ports.WFSConfig) error {
	return i.Handler.Initialize(conf)
}

// Finalize finalize MAO wavefront sensor.
func (i *WFSInteractor) Finalize() error {
	return i.Handler.Finalize()
}

// Start starts MAO wavefront sensor.
func (i *WFSInteractor) Start(t *ports.WFSTime) error {
	return i.Handler.Start(t)
}

// Halt stops MAO wavefront sensor.
func (i *WFSInteractor) Halt() error {
	return i.Handler.Halt()
}
