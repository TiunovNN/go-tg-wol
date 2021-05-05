package wol

import (
	"net"

	"github.com/TiunovNN/go-tg-wol/pkg/magic_packet"
)

const broadcastAddress string = "255.255.255.255:9"

func Send(mac string) error {
	conn, err := net.Dial("udp4", broadcastAddress)
	if err != nil {
		return err
	}
	message, err := magic_packet.Create(mac)
	if err != nil {
		return err
	}
	_, err = conn.Write(message)

	return err
}
