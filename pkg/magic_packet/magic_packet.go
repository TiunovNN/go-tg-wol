package magic_packet

import "net"

func Create(mac string) ([]byte, error) {
	buf := make([]byte, 0, 6*17)
	macBytes, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 6; i++ {
		buf = append(buf, 0xFF)
	}
	for i := 0; i < 16; i++ {
		for _, c := range macBytes {
			buf = append(buf, c)
		}
	}
	return buf, nil
}
