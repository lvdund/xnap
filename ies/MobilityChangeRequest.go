package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// MobilityChangeRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{MobilityChangeRequest-IEs}},
//	...
// }

type MobilityChangeRequest struct {
	NGRANnode1CellID                     GlobalNGRANCellID
	NGRANnode2CellID                     GlobalNGRANCellID
	NGRANnode1MobilityParameters         *MobilityParametersInformation
	NGRANnode2ProposedMobilityParameters MobilityParametersInformation
	Cause                                Cause
	SSBOffsetsList                       *SSBOffsetsList
}

func (msg *MobilityChangeRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode1CellID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANnode1CellID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode2CellID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANnode2CellID,
	})
	if msg.NGRANnode1MobilityParameters != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode1MobilityParameters)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NGRANnode1MobilityParameters,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode2ProposedMobilityParameters)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANnode2ProposedMobilityParameters,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.SSBOffsetsList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SSBOffsetsList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SSBOffsetsList,
		})
	}
	return ies
}

func (msg *MobilityChangeRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_MobilitySettingsChange); err != nil {
		return err
	}
	if err := criticalityXnMsg(outer, CriticalityReject); err != nil {
		return err
	}

	if err := outer.EncodeOpenType(inner.Bytes()); err != nil {
		return err
	}

	_, err := w.Write(outer.Bytes())
	return err
}

func (msg *MobilityChangeRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode MobilityChangeRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode MobilityChangeRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode MobilityChangeRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode MobilityChangeRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode MobilityChangeRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANnode1CellID:
			err = msg.NGRANnode1CellID.Decode(ies)
		case common.ProtocolIEID_NGRANnode2CellID:
			err = msg.NGRANnode2CellID.Decode(ies)
		case common.ProtocolIEID_NGRANnode1MobilityParameters:
			v := &MobilityParametersInformation{}
			err = v.Decode(ies)
			msg.NGRANnode1MobilityParameters = v
		case common.ProtocolIEID_NGRANnode2ProposedMobilityParameters:
			err = msg.NGRANnode2ProposedMobilityParameters.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_SSBOffsetsList:
			v := &SSBOffsetsList{}
			err = v.Decode(ies)
			msg.SSBOffsetsList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported MobilityChangeRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode MobilityChangeRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode MobilityChangeRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANnode1CellID,
		common.ProtocolIEID_NGRANnode2CellID,
		common.ProtocolIEID_NGRANnode2ProposedMobilityParameters,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory MobilityChangeRequest IE id=%d", id)
		}
	}

	return nil
}
