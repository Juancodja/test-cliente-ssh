package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
)

const MIN_PAD_LEN = 4

type SSHMessage struct {
	PacketLength  uint32
	PaddingLength byte
	Payload       []byte
	Padding       []byte
	MAC           []byte
}

func (m *SSHMessage) Marshal() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, m.PacketLength)

	buf.WriteByte(m.PaddingLength)

	buf.Write(m.Payload)
	buf.Write(m.Padding)
	buf.Write(m.MAC)

	return buf.Bytes()
}

func NewSSHMessage(payload, mac []byte, blockSize int) *SSHMessage {
	packlen := len(payload) + 1 + MIN_PAD_LEN
	r := packlen % blockSize

	extrapad := blockSize - r
	packlen += extrapad

	padlen := MIN_PAD_LEN + extrapad

	pad := make([]byte, padlen)
	rand.Read(pad[:])

	return &SSHMessage{
		uint32(packlen),
		byte(padlen),
		payload,
		pad,
		mac,
	}
}

func ReadNextMessage(conn io.Reader, maclen int) *SSHMessage {
	var packlen uint32
	binary.Read(conn, binary.BigEndian, &packlen)

	var padlen byte
	binary.Read(conn, binary.BigEndian, &padlen)

	payloadlen := int(packlen) - int(padlen) - 1
	payload := make([]byte, payloadlen)
	_, err := io.ReadFull(conn, payload)
	if err != nil {
		panic(err)
	}
	padding := make([]byte, padlen)
	_, err = io.ReadFull(conn, padding)
	if err != nil {
		panic(err)
	}
	mac := make([]byte, maclen)
	_, err = io.ReadFull(conn, mac)
	if err != nil {
		panic(err)
	}

	return &SSHMessage{
		PacketLength:  packlen,
		PaddingLength: padlen,
		Payload:       payload,
		Padding:       padding,
		MAC:           mac,
	}
}
