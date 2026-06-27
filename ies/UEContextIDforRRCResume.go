package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEContextIDforRRCResumeConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "i-rnti"},
		{Name: "allocated-c-rnti"},
		{Name: "accessPCI"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEContextIDforRRCResume struct {
	IRnti          IRNTI
	AllocatedCRnti CRNTI
	AccessPCI      NGRANCellPCI
	IEExtensions   []byte
}

func (ie *UEContextIDforRRCResume) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEContextIDforRRCResumeConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IRnti.Encode(e); err != nil {
		return err
	}
	if err := ie.AllocatedCRnti.Encode(e); err != nil {
		return err
	}
	if err := ie.AccessPCI.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UEContextIDforRRCResume) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEContextIDforRRCResumeConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IRnti.Decode(d); err != nil {
		return err
	}
	if err := ie.AllocatedCRnti.Decode(d); err != nil {
		return err
	}
	if err := ie.AccessPCI.Decode(d); err != nil {
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
