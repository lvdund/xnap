package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var resourceCoordRequestGNBInitiatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataTrafficResourceIndication"},
		{Name: "listofE-UTRACells", Optional: true},
		{Name: "spectrumSharingGroupID"},
		{Name: "listofNRCells", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ResourceCoordRequestGNBInitiated struct {
	DataTrafficResourceIndication DataTrafficResourceIndication
	ListofEUTRACells              []*EUTRACGI
	SpectrumSharingGroupID        SpectrumSharingGroupID
	ListofNRCells                 []*NRCGI
	IEExtensions                  []byte
}

func (ie *ResourceCoordRequestGNBInitiated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(resourceCoordRequestGNBInitiatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.ListofEUTRACells) > 0, len(ie.ListofNRCells) > 0, false}); err != nil {
		return err
	}
	if err := ie.DataTrafficResourceIndication.Encode(e); err != nil {
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
	if err := ie.SpectrumSharingGroupID.Encode(e); err != nil {
		return err
	}
	if len(ie.ListofNRCells) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.ListofNRCells))); err != nil {
			return err
		}
		for _, item := range ie.ListofNRCells {
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

func (ie *ResourceCoordRequestGNBInitiated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(resourceCoordRequestGNBInitiatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DataTrafficResourceIndication.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
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
	if err := ie.SpectrumSharingGroupID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.ListofNRCells = make([]*NRCGI, n)
		for i := range ie.ListofNRCells {
			ie.ListofNRCells[i] = new(NRCGI)
			if err := ie.ListofNRCells[i].Decode(d); err != nil {
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
