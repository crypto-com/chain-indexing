package projection

type Base struct {
	projectionId string

	config interface{}
}

type Options struct {
	// Pointer to the configuration
	MaybeConfigPtr interface{}
}

func NewBase(projectionId string) Base {
	return NewBaseWithOptions(
		projectionId,
		Options{
			MaybeConfigPtr: nil,
		},
	)
}

func NewBaseWithOptions(projectionId string, options Options) Base {
	base := Base{
		projectionId: projectionId,

		config: nil,
	}

	if options.MaybeConfigPtr != nil {
		base.config = options.MaybeConfigPtr
	}

	return base
}

// Implements projection.Id()
func (base *Base) Id() string {
	return base.projectionId
}

func (base *Base) Config() interface{} {
	return base.config
}
