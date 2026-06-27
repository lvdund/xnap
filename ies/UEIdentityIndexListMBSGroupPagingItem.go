package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEIdentityIndexListMBSGroupPagingItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ueIdentityIndexList-MBSGroupPagingValue"},
		{Name: "pagingDRX", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEIdentityIndexListMBSGroupPagingItem struct {
	UeIdentityIndexListMBSGroupPagingValue UEIdentityIndexListMBSGroupPagingValue
	PagingDRX                              *UESpecificDRX
	IEExtensions                           []byte
}

func (ie *UEIdentityIndexListMBSGroupPagingItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEIdentityIndexListMBSGroupPagingItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PagingDRX != nil, false}); err != nil {
		return err
	}
	if err := ie.UeIdentityIndexListMBSGroupPagingValue.Encode(e); err != nil {
		return err
	}
	if ie.PagingDRX != nil {
		if err := ie.PagingDRX.Encode(e); err != nil {
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

func (ie *UEIdentityIndexListMBSGroupPagingItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEIdentityIndexListMBSGroupPagingItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UeIdentityIndexListMBSGroupPagingValue.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PagingDRX = new(UESpecificDRX)
		if err := ie.PagingDRX.Decode(d); err != nil {
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
