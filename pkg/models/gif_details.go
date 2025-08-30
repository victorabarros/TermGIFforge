package models

import (
	"sync"
	"time"
)

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

// TODO rename to GIFState
type GIFDetail struct {
	Status     GIFStatus
	LastAccess time.Time
}

type GIFDetails struct {
	GIF   map[string]GIFDetail
	Mutex *sync.Mutex
}

func (d *GIFDetails) Get(id string) (GIFDetail, bool) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	val, ok := d.GIF[id]
	return val, ok
}

func (d *GIFDetails) SetStatus(id string, val GIFStatus) {
	if d.GIF == nil {
		return
	}
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	detail, ok := d.GIF[id]
	if !ok {
		d.GIF[id] = GIFDetail{
			Status:     val,
			LastAccess: time.Now(),
		}
		return
	}

	d.GIF[id] = GIFDetail{
		Status:     val,
		LastAccess: detail.LastAccess,
	}
}

func (d *GIFDetails) SetLastAccess(id string, val time.Time) {
	if d.GIF == nil {
		return
	}
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	detail, ok := d.GIF[id]
	if !ok {
		return
	}

	d.GIF[id] = GIFDetail{
		Status:     detail.Status,
		LastAccess: val,
	}
}

func (d *GIFDetails) Del(id string) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	delete(d.GIF, id)
}

func NewGIFDetails() GIFDetails {
	return GIFDetails{
		GIF:   map[string]GIFDetail{},
		Mutex: &sync.Mutex{},
	}
}
