package mba

import "sync"

// mba defines the internal MBA structure.
type mba struct {
	// keep a list of initilialized drivers.
	drivers map[string]Driver

	mutex sync.Mutex
}

// Keep the mba context at the package global level.
var gmba mba

func init() {
	gmba.drivers = make(map[string]Driver)
}

// MustInitDriver add the driver to the list of drivers.
// If the name is already used, it panicks.
func MustInitDriver(name string, driver Driver) {
	gmba.mutex.Lock()
	defer gmba.mutex.Unlock()

	if _, ok := gmba.drivers[name]; ok {
		panic("The driver name is already used.")
	}

	gmba.drivers[name] = driver
}

// GetDriver returns the configured driver
func GetDriver(name string) Driver {
	gmba.mutex.Lock()
	defer gmba.mutex.Unlock()

	return gmba.drivers[name]
}
