package projection

type Base struct {
	projectionId string
}

func NewBase(projectionId string) Base {
	return Base{
		projectionId,
	}
}

// Implements projection.Id()
func (base *Base) Id() string {
	return base.projectionId
}
