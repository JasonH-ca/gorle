package gorle

import (
	"bytes"
	"encoding/binary"
)

func encodeLength(length uint16) []byte {
	buf := make([]byte,0)
	if length < 80 {
		buf = append(buf, byte(length & 0x7f) )
	} else if length < 0x8000 {
		buf = append(buf, byte((length & 0x7f00) >> 8 + 0x80))
		buf = append(buf, byte(length & 0x00ff))
	} else {
		return nil
	}
	return buf
}

func Encode(data []byte) []byte {
	count := uint16(1)
	var encoded bytes.Buffer
	var i int

	for i = 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			count++
		} else {
			itemdata := encodeLength(count)
			if itemdata == nil {
				return nil
			}
			itemdata = append(itemdata, data[i])
			encoded.Write(itemdata)
			count = 1
		}
	}
	itemdata := encodeLength(count)
	itemdata = append(itemdata, data[i])
	encoded.Write(itemdata)
	return encoded.Bytes()
}

func Decode(data []byte) []byte {
	var decoded bytes.Buffer
	length := 0
	offset := 0
	for {
		if data[offset] < 0x80 {
			length = int(data[offset])
			offset++
		} else if data[offset] > 0x80 {
			length = int(binary.LittleEndian.Uint16(data[offset:offset+2]))
			length = length & 0x7fff
			offset += 2
		}
		for {
			decoded.WriteByte(data[offset])
			length--
			if length == 0 {
				break
			}
		}
		offset++
		if offset >= len(data) {
			break
		}
	}
	return decoded.Bytes()
}