package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var serviceAreaItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity"},
		{Name: "allowed-TACs-ServiceArea", Optional: true},
		{Name: "not-allowed-TACs-ServiceArea", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServiceAreaItem struct {
	PlmnIdentity              PLMNIdentity
	AllowedTACsServiceArea    []*TAC
	NotAllowedTACsServiceArea []*TAC
	IEExtensions              []byte
}

func (ie *ServiceAreaItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(serviceAreaItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.AllowedTACsServiceArea) > 0, len(ie.NotAllowedTACsServiceArea) > 0, false}); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Encode(e); err != nil {
		return err
	}
	if len(ie.AllowedTACsServiceArea) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofAllowedAreas)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.AllowedTACsServiceArea))); err != nil {
			return err
		}
		for _, item := range ie.AllowedTACsServiceArea {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	}
	if len(ie.NotAllowedTACsServiceArea) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofAllowedAreas)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.NotAllowedTACsServiceArea))); err != nil {
			return err
		}
		for _, item := range ie.NotAllowedTACsServiceArea {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ServiceAreaItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(serviceAreaItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofAllowedAreas)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.AllowedTACsServiceArea = make([]*TAC, n)
		for i := range ie.AllowedTACsServiceArea {
			ie.AllowedTACsServiceArea[i] = new(TAC)
			if err := ie.AllowedTACsServiceArea[i].Decode(d); err != nil {
				return err
			}
		}
	}
	if seq.IsComponentPresent(2) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofAllowedAreas)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.NotAllowedTACsServiceArea = make([]*TAC, n)
		for i := range ie.NotAllowedTACsServiceArea {
			ie.NotAllowedTACsServiceArea[i] = new(TAC)
			if err := ie.NotAllowedTACsServiceArea[i].Decode(d); err != nil {
				return err
			}
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
