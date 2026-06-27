package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sSBOffsetsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nG-RANnode1SSBOffsets", Optional: true},
		{Name: "nG-RANnode2ProposedSSBOffsets"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SSBOffsetsItem struct {
	NGRANnode1SSBOffsets         *SSBOffsetInformation
	NGRANnode2ProposedSSBOffsets SSBOffsetInformation
	IEExtensions                 []byte
}

func (ie *SSBOffsetsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBOffsetsItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NGRANnode1SSBOffsets != nil, false}); err != nil {
		return err
	}
	if ie.NGRANnode1SSBOffsets != nil {
		if err := ie.NGRANnode1SSBOffsets.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.NGRANnode2ProposedSSBOffsets.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SSBOffsetsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBOffsetsItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.NGRANnode1SSBOffsets = new(SSBOffsetInformation)
		if err := ie.NGRANnode1SSBOffsets.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.NGRANnode2ProposedSSBOffsets.Decode(d); err != nil {
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
