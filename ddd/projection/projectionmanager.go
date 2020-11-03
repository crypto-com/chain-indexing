package projection

import (
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/rdb"
)

type Manager struct {
	projections []Projection
}

func NewManager(conn rdb.Conn) *Manager {
	return &Manager{
		projections: make([]Projection, 0),
	}
}

func (manager *Manager) RegisterProjection(projection Projection) error {
	if manager.IsProjectionRegistered(projection) {
		return fmt.Errorf("Projection `%s` already registered", projection.Id())
	}
	manager.projections = append(manager.projections, projection)
	return nil
}

func (manager *Manager) IsProjectionRegistered(projection Projection) bool {
	for _, registeredProjection := range manager.projections {
		if projection.Id() == registeredProjection.Id() {
			return true
		}
	}
	return false
}

// Starts ProjectionManager by subscribing to events feed and invoke projection accordingly.
func (manager *Manager) Run() {}
