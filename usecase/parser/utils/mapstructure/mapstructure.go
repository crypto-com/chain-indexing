package mapstructure

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

type MapstructureDecoder struct {
	weeklyTypedInput bool
	decodeHookFuncFn mapstructure.DecodeHookFunc
}

var DefaultMapstructureDecoder = (func() *MapstructureDecoder {
	weaklyTypedInput := true
	decodeHook := mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToTimeHookFunc(time.RFC3339),
		StringToJsonUint64HookFunc(),
		StringToJsonNumericStringHookFunc(),
		StringToDurationHookFunc(),
		StringToByteSliceHookFunc(),
	)
	return NewMapstructureDecoder(
		weaklyTypedInput,
		decodeHook,
	)
})()

func NewMapstructureDecoder(
	weeklyTypedInput bool,
	decodeHookFuncFn mapstructure.DecodeHookFunc,
) *MapstructureDecoder {
	return &MapstructureDecoder{
		weeklyTypedInput: weeklyTypedInput,
		decodeHookFuncFn: decodeHookFuncFn,
	}
}

func (generator *MapstructureDecoder) MustDecode(
	input interface{},
	resultPtr interface{},
) {
	if err := generator.Decode(input, resultPtr); err != nil {
		panic(err)
	}
}

func (generator *MapstructureDecoder) Decode(
	input interface{},
	resultPtr interface{},
) error {
	decoder, decoderErr := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook:       generator.decodeHookFuncFn,
		Result:           resultPtr,
	})
	if decoderErr != nil {
		return fmt.Errorf("error creating decoder: %v", decoderErr)
	}

	if decodeErr := decoder.Decode(input); decodeErr != nil {
		return fmt.Errorf("error decoding input: %v", decodeErr)
	}

	return nil
}
