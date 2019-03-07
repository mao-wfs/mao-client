package interactor

import (
	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/input"
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

// InitializeWFS initializes MAO-WFS.
func (i *WFSInteractor) InitializeWFS(req *input.InitRequest) error {
	conf := req.GetWFSConf()
	return i.Handler.Initialize(conf)
}

// FinalizeWFS finalizes MAO-WFS.
func (i *WFSInteractor) FinalizeWFS() error {
	return i.Handler.Finalize()
}

// StartWFS starts MAO-WFS.
func (i *WFSInteractor) StartWFS(req *input.StartRequest) error {
	t := req.GetWFSTime()
	return i.Handler.Start(t)
}

// HaltWFS halts MAO-WFS.
func (i *WFSInteractor) HaltWFS() error {
	return i.Handler.Halt()
}
