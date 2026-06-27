package utils

import (
	"net"

	"github.com/lvdund/asn1go/per"
)

func IPAddressToString(ipAddr per.BitString) (ipv4Addr, ipv6Addr string) {

	// Described in 38.414
	switch ipAddr.Length {
	case 32: // ipv4
		netIP := net.IPv4(ipAddr.Bits[0], ipAddr.Bits[1], ipAddr.Bits[2], ipAddr.Bits[3])
		ipv4Addr = netIP.String()
	case 128: // ipv6
		netIP := net.IP{}
		for i := range ipAddr.Bits {
			netIP = append(netIP, ipAddr.Bits[i])
		}
		ipv6Addr = netIP.String()
	case 160: // ipv4 + ipv6, and ipv4 is contained in the first 32 bits
		netIPv4 := net.IPv4(ipAddr.Bits[0], ipAddr.Bits[1], ipAddr.Bits[2], ipAddr.Bits[3])
		netIPv6 := net.IP{}
		for i := range ipAddr.Bits {
			netIPv6 = append(netIPv6, ipAddr.Bits[i+4])
		}
		ipv4Addr = netIPv4.String()
		ipv6Addr = netIPv6.String()
	}
	return
}

func IPAddressToNgap(ipv4Addr, ipv6Addr string) (ipAddr per.BitString) {

	if ipv4Addr == "" && ipv6Addr == "" {
		return
	}

	if ipv4Addr != "" && ipv6Addr != "" { // Both ipv4 & ipv6
		ipv4NetIP := net.ParseIP(ipv4Addr).To4()
		ipv6NetIP := net.ParseIP(ipv6Addr).To16()

		ipBits := []byte{ipv4NetIP[0], ipv4NetIP[1], ipv4NetIP[2], ipv4NetIP[3]}
		for i := range 16 {
			ipBits = append(ipBits, ipv6NetIP[i])
		}

		ipAddr = per.BitString{
			Bits: ipBits,
			Length: 160,
		}

	} else if ipv4Addr != "" && ipv6Addr == "" { // ipv4
		ipv4NetIP := net.ParseIP(ipv4Addr).To4()

		ipBits := []byte{ipv4NetIP[0], ipv4NetIP[1], ipv4NetIP[2], ipv4NetIP[3]}

		ipAddr = per.BitString{
			Bits: ipBits,
			Length: 32,
		}

	} else { // ipv6
		ipv6NetIP := net.ParseIP(ipv6Addr).To16()

		ipBits := []byte{}
		for i := range 16 {
			ipBits = append(ipBits, ipv6NetIP[i])
		}

		ipAddr = per.BitString{
			Bits: ipBits,
			Length: 128,
		}

	}

	return ipAddr
}
