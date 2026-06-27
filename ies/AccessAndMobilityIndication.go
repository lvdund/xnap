package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// AccessAndMobilityIndication ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{AccessAndMobilityIndication-IEs}},
//	...
// }

type AccessAndMobilityIndication struct {
	RAReport                                *RAReport
	SuccessfulHOReportInformation           *SuccessfulHOReportInformation
	SuccessfulPSCellChangeReportInformation *SuccessfulPSCellChangeReportInformation
	DLLBTFailureInformationList             *DLLBTFailureInformationList
}

func (msg *AccessAndMobilityIndication) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	if msg.RAReport != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RAReport)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RAReport,
		})
	}
	if msg.SuccessfulHOReportInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SuccessfulHOReportInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SuccessfulHOReportInformation,
		})
	}
	if msg.SuccessfulPSCellChangeReportInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SuccessfulPSCellChangeReportInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SuccessfulPSCellChangeReportInformation,
		})
	}
	if msg.DLLBTFailureInformationList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DLLBTFailureInformationList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DLLBTFailureInformationList,
		})
	}
	return ies
}

func (msg *AccessAndMobilityIndication) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_AccessAndMobilityIndication); err != nil {
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

func (msg *AccessAndMobilityIndication) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode AccessAndMobilityIndication extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode AccessAndMobilityIndication protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode AccessAndMobilityIndication IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode AccessAndMobilityIndication IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode AccessAndMobilityIndication IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_RAReport:
			v := &RAReport{}
			err = v.Decode(ies)
			msg.RAReport = v
		case common.ProtocolIEID_SuccessfulHOReportInformation:
			v := &SuccessfulHOReportInformation{}
			err = v.Decode(ies)
			msg.SuccessfulHOReportInformation = v
		case common.ProtocolIEID_SuccessfulPSCellChangeReportInformation:
			v := &SuccessfulPSCellChangeReportInformation{}
			err = v.Decode(ies)
			msg.SuccessfulPSCellChangeReportInformation = v
		case common.ProtocolIEID_DLLBTFailureInformationList:
			v := &DLLBTFailureInformationList{}
			err = v.Decode(ies)
			msg.DLLBTFailureInformationList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported AccessAndMobilityIndication IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode AccessAndMobilityIndication IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode AccessAndMobilityIndication extensions: %w", err)
	}

	return nil
}
