package xn

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

func XnDecode(wire []byte) (pdu XnApPDU, err error) {
	r := per.NewDecoder(wire, per.APER)

	choice := r.NewChoiceDecoder(per.ChoiceConstraints{
		Extensible: true,
		RootAlternatives: []per.AlternativeInfo{
			{Name: "initiatingMessage", Tag: int(common.InitiatingMessage)},
			{Name: "successfulOutcome", Tag: int(common.SuccessfulOutcome)},
			{Name: "unsuccessfulOutcome", Tag: int(common.UnsuccessfulOutcome)},
		},
	})

	choiceIdx, isExt, _, err := choice.DecodeChoice()
	if err != nil {
		return pdu, err
	}
	if isExt {
		return pdu, fmt.Errorf("unsupported XnAP-PDU extension choice")
	}

	pdu.Choice = uint8(choiceIdx)

	if err := pdu.ProcedureCode.Decode(r); err != nil {
		return pdu, err
	}
	if err := pdu.Criticality.Decode(r); err != nil {
		return pdu, err
	}

	body, err := r.DecodeOpenType()
	if err != nil {
		return pdu, err
	}

	msg, err := decodeMessage(pdu.Choice, pdu.ProcedureCode.Value, body)
	if err != nil {
		return pdu, err
	}
	pdu.Message = msg
	return pdu, nil
}
