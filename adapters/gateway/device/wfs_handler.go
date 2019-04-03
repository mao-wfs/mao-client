package device

import (
	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/domain"
)

// WFSHandler is the handler of MAO-WFS.
type WFSHandler struct {
	Correlator *CorrelatorHandler
	Switch     *SwitchHandler
}

// NewWFSHandler returns the handler of MAO-WFS.
func NewWFSHandler(corrClt, swClt Client) *WFSHandler {
	return &WFSHandler{
		Correlator: NewCorrelatorHandler(corrClt),
		Switch:     NewSwitchHandler(swClt),
	}
}

// Initialize initializes MAO-WFS.
func (h *WFSHandler) Initialize(conf *domain.WFSConf) error {
	if err := h.Correlator.Initialize(); err != nil {
		return xerrors.Errorf("error in correlator: %w", err)
	}
	swOrder := conf.GetSwitchOrder()
	if err := h.Switch.Initialize(swOrder); err != nil {
		return xerrors.Errorf("error in switch: %w", err)
	}
	return nil
}

// Finalize finalizes MAO-WFS.
func (h *WFSHandler) Finalize() error {
	if err := h.Correlator.Finalize(); err != nil {
		return xerrors.Errorf("error in correlator: %w", err)
	}
	if err := h.Switch.Finalize(); err != nil {
		return xerrors.Errorf("error in switch: %w", err)
	}
	return nil
}

// Start starts MAO-WFS.
func (h *WFSHandler) Start(t *domain.WFSTime) error {
	return nil
}

// Halt halts MAO-WFS.
func (h *WFSHandler) Halt() error {
	if err := h.Correlator.Halt(); err != nil {
		return xerrors.Errorf("error in correlator: %w", err)
	}
	if err := h.Switch.Finalize(); err != nil {
		return xerrors.Errorf("error in switch: %w", err)
	}
	return nil
}
