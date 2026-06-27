package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDCPSNLengthUlPDCPSNLengthV12bits int64 = 0
	PDCPSNLengthUlPDCPSNLengthV18bits int64 = 1
)

var pDCPSNLengthUlPDCPSNLengthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PDCPSNLengthUlPDCPSNLength struct {
	Value int64
}

func (ie *PDCPSNLengthUlPDCPSNLength) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDCPSNLengthUlPDCPSNLengthConstraints)
}

func (ie *PDCPSNLengthUlPDCPSNLength) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDCPSNLengthUlPDCPSNLengthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	PDCPSNLengthDlPDCPSNLengthV12bits int64 = 0
	PDCPSNLengthDlPDCPSNLengthV18bits int64 = 1
)

var pDCPSNLengthDlPDCPSNLengthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PDCPSNLengthDlPDCPSNLength struct {
	Value int64
}

func (ie *PDCPSNLengthDlPDCPSNLength) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDCPSNLengthDlPDCPSNLengthConstraints)
}

func (ie *PDCPSNLengthDlPDCPSNLength) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDCPSNLengthDlPDCPSNLengthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var pDCPSNLengthConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ulPDCPSNLength"},
		{Name: "dlPDCPSNLength"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDCPSNLength struct {
	UlPDCPSNLength PDCPSNLengthUlPDCPSNLength
	DlPDCPSNLength PDCPSNLengthDlPDCPSNLength
	IEExtensions   []byte
}

func (ie *PDCPSNLength) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCPSNLengthConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UlPDCPSNLength.Encode(e); err != nil {
		return err
	}
	if err := ie.DlPDCPSNLength.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDCPSNLength) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCPSNLengthConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UlPDCPSNLength.Decode(d); err != nil {
		return err
	}
	if err := ie.DlPDCPSNLength.Decode(d); err != nil {
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
