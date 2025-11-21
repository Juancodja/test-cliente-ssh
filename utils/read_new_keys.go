package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type KeyExchangeReply struct {
	KeyType       []byte
	EdDSApub      []byte
	Q_c           []byte
	SignatureType []byte
	Signature     []byte
}

func ReadKeyExchangeReply(b *bytes.Buffer) (*KeyExchangeReply, error) {
	messageType, err := b.ReadByte()
	if err != nil {
		return nil, err
	}
	if messageType != 31 {
		return nil, errors.New("codigo mensaje erroneo")
	}

	var hostKeyLen uint32
	if err := binary.Read(b, binary.BigEndian, &hostKeyLen); err != nil {
		return nil, err
	}

	_, hostKeyType, err := ReadSshString(b)
	if err != nil {
		return nil, err
	}

	_, edDSApub, err := ReadSshString(b)
	if err != nil {
		return nil, err
	}

	_, ecdhServerEphemeral, err := ReadSshString(b)
	if err != nil {
		return nil, err
	}

	var signatureLen uint32
	if err := binary.Read(b, binary.BigEndian, &signatureLen); err != nil {
		return nil, err
	}

	_, signatureType, err := ReadSshString(b)
	if err != nil {
		return nil, err
	}

	signDataLen := int(signatureLen) - len(signatureType) - 4
	signatureData := make([]byte, signDataLen)
	_, err = b.Read(signatureData)
	if err != nil {
		return nil, err
	}

	return &KeyExchangeReply{
		hostKeyType,
		edDSApub,
		ecdhServerEphemeral,
		signatureType,
		signatureData,
	}, nil
}

func ReadNewKeys(b *bytes.Buffer) error {
	messageType, err := b.ReadByte()
	if err != nil {
		return err
	}
	if messageType != 21 {
		fmt.Println(messageType)
		return errors.New("codigo mensaje erroneo")
	}
	return nil
}
