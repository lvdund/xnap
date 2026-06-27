package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBBStatusTransfer18bitsSNConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "receiveStatusofPDCPSDU", Optional: true},
		{Name: "cOUNTValue"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBBStatusTransfer18bitsSN struct {
	ReceiveStatusofPDCPSDU *per.BitString
	COUNTValue             COUNTPDCPSN18
	IEExtensions           []byte
}

func (ie *DRBBStatusTransfer18bitsSN) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBBStatusTransfer18bitsSNConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ReceiveStatusofPDCPSDU != nil, false}); err != nil {
		return err
	}
	if ie.ReceiveStatusofPDCPSDU != nil {
		if err := e.EncodeBitString(*ie.ReceiveStatusofPDCPSDU, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(131072)),
		}); err != nil {
			return err
		}
	}
	if err := ie.COUNTValue.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DRBBStatusTransfer18bitsSN) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBBStatusTransfer18bitsSNConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(131072)),
		})
		if err != nil {
			return err
		}
		ie.ReceiveStatusofPDCPSDU = &val
	}
	if err := ie.COUNTValue.Decode(d); err != nil {
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
