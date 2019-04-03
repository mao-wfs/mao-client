package input

import (
	"github.com/mao-wfs/mao-ctrl/domain"
)

// WFSInputPort is the input port for MAO-WFS.
type WFSInputPort interface {
	InitializeWFS(req *InitRequest) error
	FinalizeWFS() error
	StartWFS(req *StartRequest) error
	HaltWFS() error
}

// InitRequest represents the request to initialize MAO-WFS.
type InitRequest struct {
	WFSConf *domain.WFSConf
}

// GetWFSConf gets "WFSConf".
func (req *InitRequest) GetWFSConf() *domain.WFSConf { return req.WFSConf }

// StartRequest represents the request to start MAO-WFS.
type StartRequest struct {
	WFSTime *domain.WFSTime
}

// GetWFSTime gets "WFSTime".
func (req *StartRequest) GetWFSTime() *domain.WFSTime { return req.WFSTime }
