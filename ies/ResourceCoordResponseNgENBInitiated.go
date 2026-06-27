package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var resourceCoordResponseNgENBInitiatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataTrafficResourceIndication"},
		{Name: "spectrumSharingGroupID"},
		{Name: "listofE-UTRACells", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ResourceCoordResponseNgENBInitiated struct {
	DataTrafficResourceIndication DataTrafficResourceIndication
	SpectrumSharingGroupID        SpectrumSharingGroupID
	ListofEUTRACells              []*EUTRACGI
	IEExtensions                  []byte
}

func (ie *ResourceCoordResponseNgENBInitiated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(resourceCoordResponseNgENBInitiatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.ListofEUTRACells) > 0, false}); err != nil {
		return err
	}
	if err := ie.DataTrafficResourceIndication.Encode(e); err != nil {
		return err
	}
	if err := ie.SpectrumSharingGroupID.Encode(e); err != nil {
		return err
	}
	if len(ie.ListofEUTRACells) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.ListofEUTRACells))); err != nil {
			return err
		}
		for _, item := range ie.ListofEUTRACells {
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

func (ie *ResourceCoordResponseNgENBInitiated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(resourceCoordResponseNgENBInitiatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DataTrafficResourceIndication.Decode(d); err != nil {
		return err
	}
	if err := ie.SpectrumSharingGroupID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.ListofEUTRACells = make([]*EUTRACGI, n)
		for i := range ie.ListofEUTRACells {
			ie.ListofEUTRACells[i] = new(EUTRACGI)
			if err := ie.ListofEUTRACells[i].Decode(d); err != nil {
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
