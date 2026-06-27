package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// CellTrafficTrace ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{CellTrafficTraceIEs}},
//	...
// }

type CellTrafficTrace struct {
	MNGRANnodeUEXnAPID             NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID             NGRANnodeUEXnAPID
	NGRANTraceID                   NGRANTraceID
	TraceCollectionEntityIPAddress TransportLayerAddress
	PrivacyIndicator               *PrivacyIndicator
	TraceCollectionEntityURI       *URIaddress
}

func (msg *CellTrafficTrace) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANTraceID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.NGRANTraceID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TraceCollectionEntityIPAddress)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.TraceCollectionEntityIPAddress,
	})
	if msg.PrivacyIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PrivacyIndicator)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PrivacyIndicator,
		})
	}
	if msg.TraceCollectionEntityURI != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TraceCollectionEntityURI)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TraceCollectionEntityURI,
		})
	}
	return ies
}

func (msg *CellTrafficTrace) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_CellTrafficTrace); err != nil {
		return err
	}
	if err := criticalityXnMsg(outer, CriticalityIgnore); err != nil {
		return err
	}

	if err := outer.EncodeOpenType(inner.Bytes()); err != nil {
		return err
	}

	_, err := w.Write(outer.Bytes())
	return err
}

func (msg *CellTrafficTrace) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode CellTrafficTrace extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode CellTrafficTrace protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode CellTrafficTrace IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode CellTrafficTrace IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode CellTrafficTrace IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_NGRANTraceID:
			err = msg.NGRANTraceID.Decode(ies)
		case common.ProtocolIEID_TraceCollectionEntityIPAddress:
			err = msg.TraceCollectionEntityIPAddress.Decode(ies)
		case common.ProtocolIEID_PrivacyIndicator:
			v := &PrivacyIndicator{}
			err = v.Decode(ies)
			msg.PrivacyIndicator = v
		case common.ProtocolIEID_TraceCollectionEntityURI:
			v := &URIaddress{}
			err = v.Decode(ies)
			msg.TraceCollectionEntityURI = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported CellTrafficTrace IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode CellTrafficTrace IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode CellTrafficTrace extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_NGRANTraceID,
		common.ProtocolIEID_TraceCollectionEntityIPAddress,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory CellTrafficTrace IE id=%d", id)
		}
	}

	return nil
}
