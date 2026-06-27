package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lTEV2XServicesAuthorizedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "vehicleUE", Optional: true},
		{Name: "pedestrianUE", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type LTEV2XServicesAuthorized struct {
	VehicleUE    *VehicleUE
	PedestrianUE *PedestrianUE
	IEExtensions []byte
}

func (ie *LTEV2XServicesAuthorized) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTEV2XServicesAuthorizedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.VehicleUE != nil, ie.PedestrianUE != nil, false}); err != nil {
		return err
	}
	if ie.VehicleUE != nil {
		if err := ie.VehicleUE.Encode(e); err != nil {
			return err
		}
	}
	if ie.PedestrianUE != nil {
		if err := ie.PedestrianUE.Encode(e); err != nil {
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

func (ie *LTEV2XServicesAuthorized) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTEV2XServicesAuthorizedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.VehicleUE = new(VehicleUE)
		if err := ie.VehicleUE.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.PedestrianUE = new(PedestrianUE)
		if err := ie.PedestrianUE.Decode(d); err != nil {
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
