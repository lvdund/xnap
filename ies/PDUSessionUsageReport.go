package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDUSessionUsageReportRATTypeNr              int64 = 0
	PDUSessionUsageReportRATTypeEutra           int64 = 1
	PDUSessionUsageReportRATTypeNrUnlicensed    int64 = 2
	PDUSessionUsageReportRATTypeEUtraUnlicensed int64 = 3
)

var pDUSessionUsageReportRATTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2, 3},
}

type PDUSessionUsageReportRATType struct {
	Value int64
}

func (ie *PDUSessionUsageReportRATType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDUSessionUsageReportRATTypeConstraints)
}

func (ie *PDUSessionUsageReportRATType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDUSessionUsageReportRATTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var pDUSessionUsageReportConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rATType"},
		{Name: "pDUSessionTimedReportList"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionUsageReport struct {
	RATType                   PDUSessionUsageReportRATType
	PDUSessionTimedReportList VolumeTimedReportList
	IEExtensions              []byte
}

func (ie *PDUSessionUsageReport) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionUsageReportConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.RATType.Encode(e); err != nil {
		return err
	}
	if err := ie.PDUSessionTimedReportList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionUsageReport) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionUsageReportConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RATType.Decode(d); err != nil {
		return err
	}
	if err := ie.PDUSessionTimedReportList.Decode(d); err != nil {
		return err
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
