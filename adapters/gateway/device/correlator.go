package device

import (
	"regexp"
	"strconv"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/domain"
)

// Result codes of the correlator operations.
const (
	resultOK int = iota
	resultReset
	resultNotExecutable
	resultInvalidArgs
	resultUnknownError
	_
	resultConflict
	resultInvalidKwd
)

// Error messages of the correlator operations.
const (
	msgErrNotExecutable = "cannot executed due to the differense of the operation mode"
	msgErrInvalidArgs   = "invalid arguments"
	msgErrUnknown       = "error while executing this command"
	msgErrConflict      = "cannot executed due to a contradiction or conflicting commands"
	msgErrInvalidKwd    = "invalid keyword"
)

// CorrelatorHandler is the handler of the correlator of MAO-WFS.
type CorrelatorHandler struct {
	Client
}

// NewCorrelatorHandler returns the handler of the correlator.
func NewCorrelatorHandler(clt Client) *CorrelatorHandler {
	return &CorrelatorHandler{Client: clt}
}

// Initialize initializes the correlator of MAO-WFS.
func (h *CorrelatorHandler) Initialize() error {
	if err := h.initialize(); err != nil {
		return xerrors.Errorf("error in Initialize(): %w", err)
	}
	return nil
}

// Finalize finalizes the correlator of MAO-WFS.
// This is the external method.
func (h *CorrelatorHandler) Finalize() error {
	if err := h.finalize(); err != nil {
		return xerrors.Errorf("error in Finalize(): %w", err)
	}
	return nil
}

// Start starts the correlator of MAO-WFS.
func (h *CorrelatorHandler) Start(t domain.SensingTime) error {
	return nil
}

// Halt halts the correlator of MAO-WFS.
func (h *CorrelatorHandler) Halt() error {
	msg := "corr_stop=2002001010000;"
	buf, err := h.Query(msg, defaultBufSize)
	if err != nil {
		return xerrors.Errorf("error in Halt(): %w", err)
	}
	res := string(buf)
	if err := h.checkResult(res); err != nil {
		return xerrors.Errorf("error in Halt(): %w", err)
	}
	return nil
}

// initialize initializes the correlator.
// This is the internal method.
func (h *CorrelatorHandler) initialize() error {
	return nil
}

// finalize finalizes the correlator of MAO-WFS.
// This is the internal method.
func (h *CorrelatorHandler) finalize() error {
	if err := h.reset(); err != nil {
		return err
	}
	return nil
}

// checkResult checks the result code of the correlator operations.
func (h *CorrelatorHandler) checkResult(res string) error {
	resCode := h.getResultCode(res)
	switch resCode {
	case resultNotExecutable:
		return xerrors.New(msgErrNotExecutable)
	case resultInvalidArgs:
		return xerrors.New(msgErrInvalidArgs)
	case resultUnknownError:
		return xerrors.New(msgErrUnknown)
	case resultConflict:
		return xerrors.New(msgErrConflict)
	case resultInvalidKwd:
		return xerrors.New(msgErrInvalidKwd)
	default:
		return nil
	}
}

// getResultCode gets the result code of the correlator operations.
func (h *CorrelatorHandler) getResultCode(msg string) int {
	re := regexp.MustCompile(`[0-9]`)
	resCode, _ := strconv.Atoi(re.FindString(msg))
	return resCode
}

// reset resets the correlator.
func (h *CorrelatorHandler) reset() error {
	msg := "reset=system;"
	buf, err := h.Query(msg, defaultBufSize)
	if err != nil {
		return err
	}
	res := string(buf)
	if err := h.checkResult(res); err != nil {
		return xerrors.Errorf("error in Halt(): %w", err)
	}
	return nil
}
