package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mBSDataForwardingResponseInfofromTargetItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mRB-ID"},
		{Name: "dlForwardingUPTNL"},
		{Name: "mRB-ProgressInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSDataForwardingResponseInfofromTargetItem struct {
	MRBID                  MRBID
	DlForwardingUPTNL      UPTransportLayerInformation
	MRBProgressInformation *MRBProgressInformation
	IEExtensions           []byte
}

func (ie *MBSDataForwardingResponseInfofromTargetItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSDataForwardingResponseInfofromTargetItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MRBProgressInformation != nil, false}); err != nil {
		return err
	}
	if err := ie.MRBID.Encode(e); err != nil {
		return err
	}
	if err := ie.DlForwardingUPTNL.Encode(e); err != nil {
		return err
	}
	if ie.MRBProgressInformation != nil {
		if err := ie.MRBProgressInformation.Encode(e); err != nil {
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

func (ie *MBSDataForwardingResponseInfofromTargetItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSDataForwardingResponseInfofromTargetItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MRBID.Decode(d); err != nil {
		return err
	}
	if err := ie.DlForwardingUPTNL.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.MRBProgressInformation = new(MRBProgressInformation)
		if err := ie.MRBProgressInformation.Decode(d); err != nil {
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
