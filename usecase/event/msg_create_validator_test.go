package event_test

import (
	event_entity "github.com/crypto-com/chainindex/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/event"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	description := model.MsgValidatorDescription{
		Moniker:         "mymonicker",
		Identity:        "myidentity",
		Website:         "mywebsite",
		SecurityContact: "mysecuritycontact",
		Details:         "mydetails",
	}
	commission := model.MsgValidatorCommission{
		Rate:          "0.100000000000000000",
		MaxRate:       "0.200000000000000000",
		MaxChangeRate: "0.010000000000000000",
	}

	Describe("En/DecodeMsgCreateValidator", func() {
		It("should able to encode and decode to the same event", func() {
			event := event_usecase.NewMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: int64(503978),
					TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description:       description,
					Commission:        commission,
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					Pubkey:            "tcrocnclconspub1zcjduepqa5rksn4ds9u6jmmg4n86d9wct7wmj23pyqe6p7e252lffzqsgcvqxm5lc2",
					Amount:            coin.MustNewCoinFromString("10"),
				},
			)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_CREATE_VALIDATOR_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgCreateValidator)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_CREATE_VALIDATOR_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash()).To(Equal(event.TxHash()))
			Expect(typedEvent.MsgIndex).To(Equal(event.MsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(event.DelegatorAddress))
			Expect(typedEvent.ValidatorAddress).To(Equal(event.ValidatorAddress))
			Expect(typedEvent.CommissionRates).To(Equal(event.CommissionRates))
			Expect(typedEvent.MinSelfDelegation).To(Equal(event.MinSelfDelegation))
			Expect(typedEvent.Amount).To(Equal(event.Amount))
		})

		It("should able to encode and decode to failed event", func() {
			event := event_usecase.NewMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: int64(503978),
					TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
					TxSuccess:   false,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description:       description,
					Commission:        commission,
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					Pubkey:            "tcrocnclconspub1zcjduepqa5rksn4ds9u6jmmg4n86d9wct7wmj23pyqe6p7e252lffzqsgcvqxm5lc2",
					Amount:            coin.MustNewCoinFromString("10"),
				},
			)
			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_CREATE_VALIDATOR_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgCreateValidator)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_CREATE_VALIDATOR_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
