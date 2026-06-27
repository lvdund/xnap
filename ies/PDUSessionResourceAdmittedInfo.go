package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchangedTrue int64 = 0
)

var pDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchangedConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchanged struct {
	Value int64
}

func (ie *PDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchanged) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchangedConstraints)
}

func (ie *PDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchanged) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchangedConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var pDUSessionResourceAdmittedInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dL-NG-U-TNL-Information-Unchanged", Optional: true},
		{Name: "qosFlowsAdmitted-List"},
		{Name: "qosFlowsNotAdmitted-List", Optional: true},
		{Name: "dataForwardingInfoFromTarget", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceAdmittedInfo struct {
	DLNGUTNLInformationUnchanged *PDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchanged
	QosFlowsAdmittedList         QoSFlowsAdmittedList
	QosFlowsNotAdmittedList      *QoSFlowsListWithCause
	DataForwardingInfoFromTarget *DataForwardingInfoFromTargetNGRANnode
	IEExtensions                 []byte
}

func (ie *PDUSessionResourceAdmittedInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceAdmittedInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DLNGUTNLInformationUnchanged != nil, ie.QosFlowsNotAdmittedList != nil, ie.DataForwardingInfoFromTarget != nil, false}); err != nil {
		return err
	}
	if ie.DLNGUTNLInformationUnchanged != nil {
		if err := ie.DLNGUTNLInformationUnchanged.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.QosFlowsAdmittedList.Encode(e); err != nil {
		return err
	}
	if ie.QosFlowsNotAdmittedList != nil {
		if err := ie.QosFlowsNotAdmittedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataForwardingInfoFromTarget != nil {
		if err := ie.DataForwardingInfoFromTarget.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionResourceAdmittedInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceAdmittedInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DLNGUTNLInformationUnchanged = new(PDUSessionResourceAdmittedInfoDLNGUTNLInformationUnchanged)
		if err := ie.DLNGUTNLInformationUnchanged.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.QosFlowsAdmittedList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.QosFlowsNotAdmittedList = new(QoSFlowsListWithCause)
		if err := ie.QosFlowsNotAdmittedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DataForwardingInfoFromTarget = new(DataForwardingInfoFromTargetNGRANnode)
		if err := ie.DataForwardingInfoFromTarget.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
