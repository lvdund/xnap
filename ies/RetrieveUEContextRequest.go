package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// RetrieveUEContextRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{RetrieveUEContextRequest-IEs}},
//	...
// }

type RetrieveUEContextRequest struct {
	NewNGRANnodeUEXnAPID                    NGRANnodeUEXnAPID
	UEContextID                             UEContextID
	MACI                                    MACI
	NewNGRANCellIdentity                    NGRANCellIdentity
	RRCResumeCause                          *RRCResumeCause
	SDTSupportRequest                       *SDTSupportRequest
	SRSPositioningConfigOrActivationRequest *SRSPositioningConfigOrActivationRequest
}

func (msg *RetrieveUEContextRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NewNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NewNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UEContextID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MACI)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MACI,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NewNGRANCellIdentity)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NewNGRANCellIdentity,
	})
	if msg.RRCResumeCause != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RRCResumeCause)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RRCResumeCause,
		})
	}
	if msg.SDTSupportRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SDTSupportRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SDTSupportRequest,
		})
	}
	if msg.SRSPositioningConfigOrActivationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SRSPositioningConfigOrActivationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SRSPositioningConfigOrActivationRequest,
		})
	}
	return ies
}

func (msg *RetrieveUEContextRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_RetrieveUEContext); err != nil {
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

func (msg *RetrieveUEContextRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode RetrieveUEContextRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode RetrieveUEContextRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode RetrieveUEContextRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode RetrieveUEContextRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode RetrieveUEContextRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NewNGRANnodeUEXnAPID:
			err = msg.NewNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_UEContextID:
			err = msg.UEContextID.Decode(ies)
		case common.ProtocolIEID_MACI:
			err = msg.MACI.Decode(ies)
		case common.ProtocolIEID_NewNGRANCellIdentity:
			err = msg.NewNGRANCellIdentity.Decode(ies)
		case common.ProtocolIEID_RRCResumeCause:
			v := &RRCResumeCause{}
			err = v.Decode(ies)
			msg.RRCResumeCause = v
		case common.ProtocolIEID_SDTSupportRequest:
			v := &SDTSupportRequest{}
			err = v.Decode(ies)
			msg.SDTSupportRequest = v
		case common.ProtocolIEID_SRSPositioningConfigOrActivationRequest:
			v := &SRSPositioningConfigOrActivationRequest{}
			err = v.Decode(ies)
			msg.SRSPositioningConfigOrActivationRequest = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported RetrieveUEContextRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode RetrieveUEContextRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode RetrieveUEContextRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NewNGRANnodeUEXnAPID,
		common.ProtocolIEID_UEContextID,
		common.ProtocolIEID_MACI,
		common.ProtocolIEID_NewNGRANCellIdentity,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory RetrieveUEContextRequest IE id=%d", id)
		}
	}

	return nil
}
