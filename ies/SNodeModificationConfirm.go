package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeModificationConfirm ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeModificationConfirm-IEs}},
//	...
// }

type SNodeModificationConfirm struct {
	MNGRANnodeUEXnAPID                NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                NGRANnodeUEXnAPID
	PDUSessionAdmittedModSNModConfirm *PDUSessionAdmittedModSNModConfirm
	PDUSessionReleasedSNModConfirm    *PDUSessionReleasedSNModConfirm
	MNToSNContainer                   []byte
	AdditionalDRBIDs                  *DRBList
	CriticalityDiagnostics            *CriticalityDiagnostics
	MRDCResourceCoordinationInfo      *MRDCResourceCoordinationInfo
	QMCCoordinationResponse           *QMCCoordinationResponse
}

func (msg *SNodeModificationConfirm) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	if msg.PDUSessionAdmittedModSNModConfirm != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionAdmittedModSNModConfirm)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionAdmittedModSNModConfirm,
		})
	}
	if msg.PDUSessionReleasedSNModConfirm != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionReleasedSNModConfirm)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionReleasedSNModConfirm,
		})
	}
	if len(msg.MNToSNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNToSNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.MNToSNContainer},
		})
	}
	if msg.AdditionalDRBIDs != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AdditionalDRBIDs)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AdditionalDRBIDs,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.MRDCResourceCoordinationInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MRDCResourceCoordinationInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MRDCResourceCoordinationInfo,
		})
	}
	if msg.QMCCoordinationResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QMCCoordinationResponse)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QMCCoordinationResponse,
		})
	}
	return ies
}

func (msg *SNodeModificationConfirm) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeModificationPreparation); err != nil {
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

func (msg *SNodeModificationConfirm) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeModificationConfirm extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeModificationConfirm protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeModificationConfirm IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeModificationConfirm IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeModificationConfirm IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionAdmittedModSNModConfirm:
			v := &PDUSessionAdmittedModSNModConfirm{}
			err = v.Decode(ies)
			msg.PDUSessionAdmittedModSNModConfirm = v
		case common.ProtocolIEID_PDUSessionReleasedSNModConfirm:
			v := &PDUSessionReleasedSNModConfirm{}
			err = v.Decode(ies)
			msg.PDUSessionReleasedSNModConfirm = v
		case common.ProtocolIEID_MNToSNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.MNToSNContainer = v.Value
		case common.ProtocolIEID_AdditionalDRBIDs:
			v := &DRBList{}
			err = v.Decode(ies)
			msg.AdditionalDRBIDs = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_MRDCResourceCoordinationInfo:
			v := &MRDCResourceCoordinationInfo{}
			err = v.Decode(ies)
			msg.MRDCResourceCoordinationInfo = v
		case common.ProtocolIEID_QMCCoordinationResponse:
			v := &QMCCoordinationResponse{}
			err = v.Decode(ies)
			msg.QMCCoordinationResponse = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeModificationConfirm IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeModificationConfirm IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeModificationConfirm extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeModificationConfirm IE id=%d", id)
		}
	}

	return nil
}
