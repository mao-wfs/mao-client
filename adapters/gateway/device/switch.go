package device

import (
	"strings"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/domain"
)

// SwitchHandler is the handler of the switch of MAO-WFS.
type SwitchHandler struct {
	Client
}

// NewSwitchHandler returns the handler of the switch.
func NewSwitchHandler(clt Client) *SwitchHandler {
	return &SwitchHandler{Client: clt}
}

// Initialize initializes the switch of MAO-WFS.
func (h *SwitchHandler) Initialize(swOrder domain.SwitchOrder) error {
	if err := h.enableDigPatt(); err != nil {
		return xerrors.Errorf("error in Initialize(): %w", err)
	}
	return nil
}

// Finalize finalizes the switch of MAO-WFS.
func (h *SwitchHandler) Finalize() error {
	if err := h.reset(); err != nil {
		return xerrors.Errorf("error in Finalize(): %w", err)
	}
	if err := h.clearStatus(); err != nil {
		return xerrors.Errorf("error in Finalize(): %w", err)
	}
	return nil
}

// Start starts the switch of MAO-WFS.
func (h *SwitchHandler) Start() error {
	if err := h.start(); err != nil {
		return xerrors.Errorf("error in Start(): %w", err)
	}
	return nil
}

// Halt halts the switch of MAO-WFS.
func (h *SwitchHandler) Halt() error {
	if err := h.halt(); err != nil {
		return xerrors.Errorf("error in Start(): %w", err)
	}
	return nil
}

// checkResult checks the result of a switch operation.
func (h *SwitchHandler) checkResult() error {
	msg := "SYST:ERR?\n"
	buf, err := h.Query(msg, defaultBufSize)
	if err != nil {
		return err
	}
	res := string(buf)
	if res != "+0,\"No error\"\n" {
		errMsg := strings.TrimRight(res, "\n")
		return xerrors.New(errMsg)
	}
	return nil
}

func (h *SwitchHandler) reset() error {
	msg := "*RST\n"
	if err := h.Write(msg); err != nil {
		return err
	}
	return h.checkResult()
}

func (h *SwitchHandler) clearStatus() error {
	msg := "*CLS\n"
	if err := h.Write(msg); err != nil {
		return err
	}
	return h.checkResult()
}

func (h *SwitchHandler) start() error {
	msg := "OUTP ON\n"
	if err := h.Write(msg); err != nil {
		return err
	}
	return h.checkResult()
}

func (h *SwitchHandler) halt() error {
	msg := "OUTP OFF\n"
	if err := h.Write(msg); err != nil {
		return err
	}
	return h.checkResult()
}

func (h *SwitchHandler) enableDigPatt() error {
	msg := "DIG:PATT ON\n"
	if err := h.Write(msg); err != nil {
		return err
	}
	return h.checkResult()
}
