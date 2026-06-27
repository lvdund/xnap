package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mBSMappingandDataForwardingRequestInfofromSourceItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mRB-ID"},
		{Name: "mBS-QoSFlow-List"},
		{Name: "mRB-ProgressInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSMappingandDataForwardingRequestInfofromSourceItem struct {
	MRBID                  MRBID
	MBSQoSFlowList         MBSQoSFlowList
	MRBProgressInformation *MRBProgressInformation
	IEExtensions           []byte
}

func (ie *MBSMappingandDataForwardingRequestInfofromSourceItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSMappingandDataForwardingRequestInfofromSourceItemConstraints)
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
	if err := ie.MBSQoSFlowList.Encode(e); err != nil {
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

func (ie *MBSMappingandDataForwardingRequestInfofromSourceItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSMappingandDataForwardingRequestInfofromSourceItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MRBID.Decode(d); err != nil {
		return err
	}
	if err := ie.MBSQoSFlowList.Decode(d); err != nil {
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
