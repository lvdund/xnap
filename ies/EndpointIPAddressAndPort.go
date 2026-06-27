package ies

import (
	"github.com/lvdund/asn1go/per"
)

var endpointIPAddressAndPortConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "endpointIPAddress"},
		{Name: "portNumber"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress
	PortNumber        PortNumber
	IEExtensions      []byte
}

func (ie *EndpointIPAddressAndPort) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(endpointIPAddressAndPortConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.EndpointIPAddress.Encode(e); err != nil {
		return err
	}
	if err := ie.PortNumber.Encode(e); err != nil {
		return err
	}
	return nil
}

func (ie *EndpointIPAddressAndPort) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(endpointIPAddressAndPortConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EndpointIPAddress.Decode(d); err != nil {
		return err
	}
	if err := ie.PortNumber.Decode(d); err != nil {
		return err
	}
	return nil
}
