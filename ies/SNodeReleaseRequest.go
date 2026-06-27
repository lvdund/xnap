package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeReleaseRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeReleaseRequest-IEs}},
//	...
// }

type SNodeReleaseRequest struct {
	MNGRANnodeUEXnAPID           NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID           *NGRANnodeUEXnAPID
	Cause                        Cause
	PDUSessionToBeReleasedRelReq *PDUSessionListWithCause
	UEContextKeptIndicator       *UEContextKeptIndicator
	MNToSNContainer              []byte
	DRBsTransferredToMN          *DRBList
}

func (msg *SNodeReleaseRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	if msg.SNGRANnodeUEXnAPID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeUEXnAPID,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.PDUSessionToBeReleasedRelReq != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionToBeReleasedRelReq)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionToBeReleasedRelReq,
		})
	}
	if msg.UEContextKeptIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextKeptIndicator)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEContextKeptIndicator,
		})
	}
	if len(msg.MNToSNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNToSNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.MNToSNContainer},
		})
	}
	if msg.DRBsTransferredToMN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DRBsTransferredToMN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DRBsTransferredToMN,
		})
	}
	return ies
}

func (msg *SNodeReleaseRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeRelease); err != nil {
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

func (msg *SNodeReleaseRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeReleaseRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeReleaseRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeReleaseRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			v := &NGRANnodeUEXnAPID{}
			err = v.Decode(ies)
			msg.SNGRANnodeUEXnAPID = v
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_PDUSessionToBeReleasedRelReq:
			v := &PDUSessionListWithCause{}
			err = v.Decode(ies)
			msg.PDUSessionToBeReleasedRelReq = v
		case common.ProtocolIEID_UEContextKeptIndicator:
			v := &UEContextKeptIndicator{}
			err = v.Decode(ies)
			msg.UEContextKeptIndicator = v
		case common.ProtocolIEID_MNToSNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.MNToSNContainer = v.Value
		case common.ProtocolIEID_DRBsTransferredToMN:
			v := &DRBList{}
			err = v.Decode(ies)
			msg.DRBsTransferredToMN = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeReleaseRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeReleaseRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeReleaseRequest IE id=%d", id)
		}
	}

	return nil
}
