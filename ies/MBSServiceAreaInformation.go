package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mBSServiceAreaInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mBS-ServiceAreaCell-List", Optional: true},
		{Name: "mBS-ServiceAreaTAI-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSServiceAreaInformation struct {
	MBSServiceAreaCellList *MBSServiceAreaCellList
	MBSServiceAreaTAIList  *MBSServiceAreaTAIList
	IEExtensions           []byte
}

func (ie *MBSServiceAreaInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSServiceAreaInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MBSServiceAreaCellList != nil, ie.MBSServiceAreaTAIList != nil, false}); err != nil {
		return err
	}
	if ie.MBSServiceAreaCellList != nil {
		if err := ie.MBSServiceAreaCellList.Encode(e); err != nil {
			return err
		}
	}
	if ie.MBSServiceAreaTAIList != nil {
		if err := ie.MBSServiceAreaTAIList.Encode(e); err != nil {
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

func (ie *MBSServiceAreaInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSServiceAreaInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.MBSServiceAreaCellList = new(MBSServiceAreaCellList)
		if err := ie.MBSServiceAreaCellList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.MBSServiceAreaTAIList = new(MBSServiceAreaTAIList)
		if err := ie.MBSServiceAreaTAIList.Decode(d); err != nil {
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
