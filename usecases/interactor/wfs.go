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

// Initialize initialize MAO wavefront sensor.
func (i *WFSInteractor) Initialize(corrConf *ports.CorrelatorConfig, swConf *ports.SwitchConfig) error {
	wfsConf, err := i.ConfigPort.FormatConfig(corrConf, swConf)
	if err != nil {
		return err
	}
	return i.Handler.Initialize(wfsConf)
}

// Start starts MAO wavefront sensor.
func (i *WFSInteractor) Start() error {
	return i.Handler.Start()
}

// Halt stops MAO wavefront sensor.
func (i *WFSInteractor) Halt() error {
	return i.Handler.Halt()
}

// Finalize finalize MAO wavefront sensor.
func (i *WFSInteractor) Finalize() error {
	return i.Handler.Finalize()
}
