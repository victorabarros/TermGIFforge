package models

import "sync"

type GIFStatus string

var (
	GIFStatuses = struct {
		Fail       GIFStatus
		Processing GIFStatus
		Ready      GIFStatus
	}{
		Fail:       GIFStatus("Fail"),
		Processing: GIFStatus("Processing"),
		Ready:      GIFStatus("Ready"),
	}
)

type StatusDetails struct {
	Status map[string]GIFStatus
	mt     *sync.Mutex
}

func (sd *StatusDetails) Get(id string) (GIFStatus, bool) {
	sd.mt.Lock()
	defer sd.mt.Unlock()
	status, ok := sd.Status[id]
	return status, ok
}

func (sd *StatusDetails) Set(id string, status GIFStatus) {
	if sd.Status == nil {
		return
	}
	sd.mt.Lock()
	defer sd.mt.Unlock()
	sd.Status[id] = status
}

func (sd *StatusDetails) Del(id string) {
	sd.mt.Lock()
	defer sd.mt.Unlock()
	delete(sd.Status, id)
}

func NewStatusDetails() StatusDetails {
	return StatusDetails{
		Status: map[string]GIFStatus{},
		mt:     &sync.Mutex{},
	}
}
