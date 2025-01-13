package pkg

import (
	"encoding/json"
)

type PhaseMonitorEntry struct {
	phaseState              int     // st4PhaseState(1)
	phaseStatus             int     // st4PhaseStatus(2)
	phaseVoltage            float64 // st4PhaseVoltage(3)
	phaseVoltageStatus      int     // st4PhaseVoltageStatus(4)
	phaseVoltageDeviation   float64 // st4PhaseVoltageDeviation(5)
	phaseCurrent            float64 // st4PhaseCurrent(6)
	phaseCurrentCrestFactor float64 // st4PhaseCurrentCrestFactor(7)
	phaseActivePower        float64 // st4PhaseActivePower(8)
	phaseApparentPower      float64 // st4PhaseApparentPower(9)
	phasePowerFactor        float64 // st4PhasePowerFactor(10)
	phasePowerFactorStatus  int     // st4PhasePowerFactorStatus(11)
	phaseReactance          float64 // st4PhaseReactance(12)
	phaseEnergy             float64 // st4PhaseEnergy(13)
}

func NewPhaseMonitorEntry() *PhaseMonitorEntry {
	return &PhaseMonitorEntry{
		phaseState:              0,
		phaseStatus:             0,
		phaseVoltage:            0.0,
		phaseVoltageStatus:      0,
		phaseVoltageDeviation:   0.0,
		phaseCurrent:            0.0,
		phaseCurrentCrestFactor: 0.0,
		phaseActivePower:        0.0,
		phaseApparentPower:      0.0,
		phasePowerFactor:        0.0,
		phasePowerFactorStatus:  0,
		phaseReactance:          0.0,
		phaseEnergy:             0.0,
	}
}

func (p *PhaseMonitorEntry) SetPhaseState(phaseState int) {
	p.phaseState = phaseState
}

func (p *PhaseMonitorEntry) SetPhaseStatus(phaseStatus int) {
	p.phaseStatus = phaseStatus
}

func (p *PhaseMonitorEntry) SetPhaseVoltage(phaseVoltage float64) {
	p.phaseVoltage = phaseVoltage
}

func (p *PhaseMonitorEntry) SetPhaseVoltageStatus(phaseVoltageStatus int) {
	p.phaseVoltageStatus = phaseVoltageStatus
}

func (p *PhaseMonitorEntry) SetPhaseVoltageDeviation(phaseVoltageDeviation float64) {
	p.phaseVoltageDeviation = phaseVoltageDeviation
}

func (p *PhaseMonitorEntry) SetPhaseCurrent(phaseCurrent float64) {
	p.phaseCurrent = phaseCurrent
}

func (p *PhaseMonitorEntry) SetPhaseCurrentCrestFactor(phaseCurrentCrestFactor float64) {
	p.phaseCurrentCrestFactor = phaseCurrentCrestFactor
}

func (p *PhaseMonitorEntry) SetPhaseActivePower(phaseActivePower float64) {
	p.phaseActivePower = phaseActivePower
}

func (p *PhaseMonitorEntry) SetPhaseApparentPower(phaseApparentPower float64) {
	p.phaseApparentPower = phaseApparentPower
}

func (p *PhaseMonitorEntry) SetPhasePowerFactor(phasePowerFactor float64) {
	p.phasePowerFactor = phasePowerFactor
}

func (p *PhaseMonitorEntry) SetPhasePowerFactorStatus(phasePowerFactorStatus int) {
	p.phasePowerFactorStatus = phasePowerFactorStatus
}

func (p *PhaseMonitorEntry) SetPhaseReactance(phaseReactance float64) {
	p.phaseReactance = phaseReactance
}

func (p *PhaseMonitorEntry) SetPhaseEnergy(phaseEnergy float64) {
	p.phaseEnergy = phaseEnergy
}

func (p *PhaseMonitorEntry) GetPhaseState() int {
	return p.phaseState
}

func (p *PhaseMonitorEntry) GetPhaseStatus() int {
	return p.phaseStatus
}

func (p *PhaseMonitorEntry) GetPhaseVoltage() float64 {
	return p.phaseVoltage
}

func (p *PhaseMonitorEntry) GetPhaseVoltageStatus() int {
	return p.phaseVoltageStatus
}

func (p *PhaseMonitorEntry) GetPhaseVoltageDeviation() float64 {
	return p.phaseVoltageDeviation
}

func (p *PhaseMonitorEntry) GetPhaseCurrent() float64 {
	return p.phaseCurrent
}

func (p *PhaseMonitorEntry) GetPhaseCurrentCrestFactor() float64 {
	return p.phaseCurrentCrestFactor
}

func (p *PhaseMonitorEntry) GetPhaseActivePower() float64 {
	return p.phaseActivePower
}

func (p *PhaseMonitorEntry) GetPhaseApparentPower() float64 {
	return p.phaseApparentPower
}

func (p *PhaseMonitorEntry) GetPhasePowerFactor() float64 {
	return p.phasePowerFactor
}

func (p *PhaseMonitorEntry) GetPhasePowerFactorStatus() int {
	return p.phasePowerFactorStatus
}

func (p *PhaseMonitorEntry) GetPhaseReactance() float64 {
	return p.phaseReactance
}

func (p *PhaseMonitorEntry) GetPhaseEnergy() float64 {
	return p.phaseEnergy
}

func (p PhaseMonitorEntry) ToJSON() (string, error) {

	data := struct {
		PhaseState              int     `json:"phase_state"`
		PhaseStatus             int     `json:"phase_status"`
		PhaseVoltage            float64 `json:"phase_voltage"`
		PhaseVoltageStatus      int     `json:"phase_voltage_status"`
		PhaseVoltageDeviation   float64 `json:"phase_voltage_deviation"`
		PhaseCurrent            float64 `json:"phase_current"`
		PhaseCurrentCrestFactor float64 `json:"phase_current_crest_factor"`
		PhaseActivePower        float64 `json:"phase_active_power"`
		PhaseApparentPower      float64 `json:"phase_apparent_power"`
		PhasePowerFactor        float64 `json:"phase_power_factor"`
		PhasePowerFactorStatus  int     `json:"phase_power_factor_status"`
		PhaseReactance          float64 `json:"phase_reactance"`
		PhaseEnergy             float64 `json:"phase_energy"`
	}{
		PhaseState:              p.phaseState,
		PhaseStatus:             p.phaseStatus,
		PhaseVoltage:            p.phaseVoltage,
		PhaseVoltageStatus:      p.phaseVoltageStatus,
		PhaseVoltageDeviation:   p.phaseVoltageDeviation,
		PhaseCurrent:            p.phaseCurrent,
		PhaseCurrentCrestFactor: p.phaseCurrentCrestFactor,
		PhaseActivePower:        p.phaseActivePower,
		PhaseApparentPower:      p.phaseApparentPower,
		PhasePowerFactor:        p.phasePowerFactor,
		PhasePowerFactorStatus:  p.phasePowerFactorStatus,
		PhaseReactance:          p.phaseReactance,
		PhaseEnergy:             p.phaseEnergy,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
