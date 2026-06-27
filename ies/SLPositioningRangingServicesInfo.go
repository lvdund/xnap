package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sLPositioningRangingServicesInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sLPositioning-Ranging-Authorized"},
		{Name: "rSPP-transport-QoS-parameters", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type SLPositioningRangingServicesInfo struct {
	SLPositioningRangingAuthorized SLPositioningRangingAuthorized
	RSPPTransportQoSParameters     *RSPPTransportQoSParameters
	IEExtensions                   []byte
}

func (ie *SLPositioningRangingServicesInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPositioningRangingServicesInfoConstraints)
	if err := seq.EncodePreamble([]bool{ie.RSPPTransportQoSParameters != nil, false}); err != nil {
		return err
	}
	if err := ie.SLPositioningRangingAuthorized.Encode(e); err != nil {
		return err
	}
	if ie.RSPPTransportQoSParameters != nil {
		if err := ie.RSPPTransportQoSParameters.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SLPositioningRangingServicesInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPositioningRangingServicesInfoConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SLPositioningRangingAuthorized.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.RSPPTransportQoSParameters = new(RSPPTransportQoSParameters)
		if err := ie.RSPPTransportQoSParameters.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
