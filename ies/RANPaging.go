package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// RANPaging ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{RANPaging-IEs}},
//	...
// }

type RANPaging struct {
	UEIdentityIndexValue                      UEIdentityIndexValue
	UERANPagingIdentity                       UERANPagingIdentity
	PagingDRX                                 PagingDRX
	RANPagingArea                             RANPagingArea
	PagingPriority                            *PagingPriority
	AssistanceDataForRANPaging                *AssistanceDataForRANPaging
	UERadioCapabilityForPaging                *UERadioCapabilityForPaging
	ExtendedUEIdentityIndexValue              *ExtendedUEIdentityIndexValue
	EUTRAPagingeDRXInformation                *EUTRAPagingeDRXInformation
	UESpecificDRX                             *UESpecificDRX
	NRPagingeDRXInformation                   *NRPagingeDRXInformation
	NRPagingeDRXInformationforRRCINACTIVE     *NRPagingeDRXInformationforRRCINACTIVE
	PagingCause                               *PagingCause
	PEIPSassistanceInformation                *PEIPSassistanceInformation
	HashedUEIdentityIndexValue                *HashedUEIdentityIndexValue
	MTSDTInformation                          *MTSDTInformation
	NRPagingLongeDRXInformationforRRCINACTIVE *NRPagingLongeDRXInformationforRRCINACTIVE
}

func (msg *RANPaging) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEIdentityIndexValue)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UEIdentityIndexValue,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UERANPagingIdentity)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.UERANPagingIdentity,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PagingDRX)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.PagingDRX,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RANPagingArea)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.RANPagingArea,
	})
	if msg.PagingPriority != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PagingPriority)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PagingPriority,
		})
	}
	if msg.AssistanceDataForRANPaging != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AssistanceDataForRANPaging)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AssistanceDataForRANPaging,
		})
	}
	if msg.UERadioCapabilityForPaging != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UERadioCapabilityForPaging)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UERadioCapabilityForPaging,
		})
	}
	if msg.ExtendedUEIdentityIndexValue != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ExtendedUEIdentityIndexValue)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ExtendedUEIdentityIndexValue,
		})
	}
	if msg.EUTRAPagingeDRXInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_EUTRAPagingeDRXInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.EUTRAPagingeDRXInformation,
		})
	}
	if msg.UESpecificDRX != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UESpecificDRX)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UESpecificDRX,
		})
	}
	if msg.NRPagingeDRXInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NRPagingeDRXInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NRPagingeDRXInformation,
		})
	}
	if msg.NRPagingeDRXInformationforRRCINACTIVE != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NRPagingeDRXInformationforRRCINACTIVE)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NRPagingeDRXInformationforRRCINACTIVE,
		})
	}
	if msg.PagingCause != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PagingCause)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PagingCause,
		})
	}
	if msg.PEIPSassistanceInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PEIPSassistanceInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PEIPSassistanceInformation,
		})
	}
	if msg.HashedUEIdentityIndexValue != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_HashedUEIdentityIndexValue)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.HashedUEIdentityIndexValue,
		})
	}
	if msg.MTSDTInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MTSDTInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MTSDTInformation,
		})
	}
	if msg.NRPagingLongeDRXInformationforRRCINACTIVE != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NRPagingLongeDRXInformationforRRCINACTIVE)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NRPagingLongeDRXInformationforRRCINACTIVE,
		})
	}
	return ies
}

func (msg *RANPaging) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_RANPaging); err != nil {
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

func (msg *RANPaging) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode RANPaging extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode RANPaging protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode RANPaging IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode RANPaging IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode RANPaging IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_UEIdentityIndexValue:
			err = msg.UEIdentityIndexValue.Decode(ies)
		case common.ProtocolIEID_UERANPagingIdentity:
			err = msg.UERANPagingIdentity.Decode(ies)
		case common.ProtocolIEID_PagingDRX:
			err = msg.PagingDRX.Decode(ies)
		case common.ProtocolIEID_RANPagingArea:
			err = msg.RANPagingArea.Decode(ies)
		case common.ProtocolIEID_PagingPriority:
			v := &PagingPriority{}
			err = v.Decode(ies)
			msg.PagingPriority = v
		case common.ProtocolIEID_AssistanceDataForRANPaging:
			v := &AssistanceDataForRANPaging{}
			err = v.Decode(ies)
			msg.AssistanceDataForRANPaging = v
		case common.ProtocolIEID_UERadioCapabilityForPaging:
			v := &UERadioCapabilityForPaging{}
			err = v.Decode(ies)
			msg.UERadioCapabilityForPaging = v
		case common.ProtocolIEID_ExtendedUEIdentityIndexValue:
			v := &ExtendedUEIdentityIndexValue{}
			err = v.Decode(ies)
			msg.ExtendedUEIdentityIndexValue = v
		case common.ProtocolIEID_EUTRAPagingeDRXInformation:
			v := &EUTRAPagingeDRXInformation{}
			err = v.Decode(ies)
			msg.EUTRAPagingeDRXInformation = v
		case common.ProtocolIEID_UESpecificDRX:
			v := &UESpecificDRX{}
			err = v.Decode(ies)
			msg.UESpecificDRX = v
		case common.ProtocolIEID_NRPagingeDRXInformation:
			v := &NRPagingeDRXInformation{}
			err = v.Decode(ies)
			msg.NRPagingeDRXInformation = v
		case common.ProtocolIEID_NRPagingeDRXInformationforRRCINACTIVE:
			v := &NRPagingeDRXInformationforRRCINACTIVE{}
			err = v.Decode(ies)
			msg.NRPagingeDRXInformationforRRCINACTIVE = v
		case common.ProtocolIEID_PagingCause:
			v := &PagingCause{}
			err = v.Decode(ies)
			msg.PagingCause = v
		case common.ProtocolIEID_PEIPSassistanceInformation:
			v := &PEIPSassistanceInformation{}
			err = v.Decode(ies)
			msg.PEIPSassistanceInformation = v
		case common.ProtocolIEID_HashedUEIdentityIndexValue:
			v := &HashedUEIdentityIndexValue{}
			err = v.Decode(ies)
			msg.HashedUEIdentityIndexValue = v
		case common.ProtocolIEID_MTSDTInformation:
			v := &MTSDTInformation{}
			err = v.Decode(ies)
			msg.MTSDTInformation = v
		case common.ProtocolIEID_NRPagingLongeDRXInformationforRRCINACTIVE:
			v := &NRPagingLongeDRXInformationforRRCINACTIVE{}
			err = v.Decode(ies)
			msg.NRPagingLongeDRXInformationforRRCINACTIVE = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported RANPaging IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode RANPaging IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode RANPaging extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_UEIdentityIndexValue,
		common.ProtocolIEID_UERANPagingIdentity,
		common.ProtocolIEID_PagingDRX,
		common.ProtocolIEID_RANPagingArea,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory RANPaging IE id=%d", id)
		}
	}

	return nil
}
