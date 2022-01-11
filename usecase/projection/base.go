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
		//decoderConfig := &mapstructure.DecoderConfig{
		//	WeaklyTypedInput: true,
		//	DecodeHook: mapstructure.ComposeDecodeHookFunc(
		//		mapstructure.StringToTimeDurationHookFunc(),
		//		mapstructure.StringToTimeHookFunc(time.RFC3339),
		//	),
		//	Result: options.MaybeConfigPtr,
		//}
		//decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
		//if decoderErr != nil {
		//	panic(fmt.Errorf("error creating projection config decoder: %v", decoderErr))
		//}
		//
		//projectionConfigId := strcase.ToSnake(projectionId)
		//projectionConfig, projectionConfigFound := GlobalConfig2.Projection[projectionConfigId]
		//if !projectionConfigFound {
		//	panic(fmt.Sprintf("no projection config found for projection id: %s", projectionConfigId))
		//}
		//if err := decoder.Decode(projectionConfig); err != nil {
		//	panic(fmt.Errorf("error decoding projection %s config: %v", projectionId, err))
		//}

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
