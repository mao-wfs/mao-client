package domain

import (
	"time"
)

// WFSHandler represents the domain handler for MAO-WFS.
type WFSHandler interface {
	Initialize(conf *WFSConf) error
	Finalize() error
	Start(t *WFSTime) error
	Halt() error
}

// WFSConf represents the configure of MAO-WFS.
type WFSConf struct {
	SwitchOrder SwitchOrder
}

// SwitchOrder is the order of the switch to enable.
type SwitchOrder []SwitchPort

// SwitchPort is the port of the switch to enable.
type SwitchPort string

// GetSwitchOrder gets "SwitchOrder".
func (conf *WFSConf) GetSwitchOrder() SwitchOrder { return conf.SwitchOrder }

// WFSTime represents the time of MAO-WFS.
type WFSTime struct {
	SensingTime SensingTime
}

// SensingTime is the sensing time of MAO-WFS.
type SensingTime time.Duration

// GetSensingTime gets "SensingTime".
func (t *WFSTime) GetSensingTime() SensingTime { return t.SensingTime }
