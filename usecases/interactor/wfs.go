package interactor

import (
	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/ports"
)

// DeviceInteractor is the interactor for MAO controller.
type DeviceInteractor struct {
	Config     ports.Config
	WFSHandler domain.WFSHandler
}

// Start starts MAO wavefront sensor.
func (i *DeviceInteractor) Start() error {
	if err := i.WFSHandler.Initialize(i.Config); err != nil {
		return err
	}
	if err := i.WFSHandler.Start(); err != nil {
		return err
	}
	return nil
}

// Halt stops MAO wavefront sensor.
func (i *DeviceInteractor) Halt() error {
	if err := i.WFSHandler.Halt(); err != nil {
		return err
	}
	if err := i.WFSHandler.Finalize(); err != nil {
		return err
	}
	return nil
}
