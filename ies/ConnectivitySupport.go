package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ConnectivitySupportENDCSupportSupported    int64 = 0
	ConnectivitySupportENDCSupportNotSupported int64 = 1
)

var connectivitySupportENDCSupportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type ConnectivitySupportENDCSupport struct {
	Value int64
}

func (ie *ConnectivitySupportENDCSupport) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, connectivitySupportENDCSupportConstraints)
}

func (ie *ConnectivitySupportENDCSupport) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(connectivitySupportENDCSupportConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var connectivitySupportConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eNDC-Support"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ConnectivitySupport struct {
	ENDCSupport  ConnectivitySupportENDCSupport
	IEExtensions []byte
}

func (ie *ConnectivitySupport) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(connectivitySupportConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.ENDCSupport.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ConnectivitySupport) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(connectivitySupportConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ENDCSupport.Decode(d); err != nil {
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
