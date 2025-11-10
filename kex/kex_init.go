package kex

import (
	"bytes"
	"sushi/utils"
)

type KexInit struct {
	MessageCode              byte
	Cookie                   [16]byte
	KexAlgos                 NameList
	ServerHostKeyAlgos       NameList
	EncryptionClientToServer NameList
	EncryptionServerToClient NameList
	MacClientToServer        NameList
	MacServerToClient        NameList
	LanguagesClientToServer  NameList
	LanguagesServerToClient  NameList
	FirstKexPacketFollows    bool
	EmptyField               uint32
}

func (m *KexInit) Marshall() []byte {
	buf := new(bytes.Buffer)

	buf.WriteByte(m.MessageCode)

	buf.Write(m.Cookie[:])

	buf.Write(m.KexAlgos.Marshall())

	return buf.Bytes()
}
