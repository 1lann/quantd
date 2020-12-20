package quantdimage

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ImageRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Image":
			z.Image, err = dc.ReadBytes(z.Image)
			if err != nil {
				err = msgp.WrapError(err, "Image")
				return
			}
		case "DX":
			z.DX, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "DX")
				return
			}
		case "DY":
			z.DY, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "DY")
				return
			}
		case "Speed":
			z.Speed, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Speed")
				return
			}
		case "MaxColors":
			z.MaxColors, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "MaxColors")
				return
			}
		case "Dither":
			z.Dither, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Dither")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ImageRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Image"
	err = en.Append(0x86, 0xa5, 0x49, 0x6d, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Image)
	if err != nil {
		err = msgp.WrapError(err, "Image")
		return
	}
	// write "DX"
	err = en.Append(0xa2, 0x44, 0x58)
	if err != nil {
		return
	}
	err = en.WriteInt(z.DX)
	if err != nil {
		err = msgp.WrapError(err, "DX")
		return
	}
	// write "DY"
	err = en.Append(0xa2, 0x44, 0x59)
	if err != nil {
		return
	}
	err = en.WriteInt(z.DY)
	if err != nil {
		err = msgp.WrapError(err, "DY")
		return
	}
	// write "Speed"
	err = en.Append(0xa5, 0x53, 0x70, 0x65, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Speed)
	if err != nil {
		err = msgp.WrapError(err, "Speed")
		return
	}
	// write "MaxColors"
	err = en.Append(0xa9, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.MaxColors)
	if err != nil {
		err = msgp.WrapError(err, "MaxColors")
		return
	}
	// write "Dither"
	err = en.Append(0xa6, 0x44, 0x69, 0x74, 0x68, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Dither)
	if err != nil {
		err = msgp.WrapError(err, "Dither")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ImageRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Image"
	o = append(o, 0x86, 0xa5, 0x49, 0x6d, 0x61, 0x67, 0x65)
	o = msgp.AppendBytes(o, z.Image)
	// string "DX"
	o = append(o, 0xa2, 0x44, 0x58)
	o = msgp.AppendInt(o, z.DX)
	// string "DY"
	o = append(o, 0xa2, 0x44, 0x59)
	o = msgp.AppendInt(o, z.DY)
	// string "Speed"
	o = append(o, 0xa5, 0x53, 0x70, 0x65, 0x65, 0x64)
	o = msgp.AppendInt(o, z.Speed)
	// string "MaxColors"
	o = append(o, 0xa9, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73)
	o = msgp.AppendInt(o, z.MaxColors)
	// string "Dither"
	o = append(o, 0xa6, 0x44, 0x69, 0x74, 0x68, 0x65, 0x72)
	o = msgp.AppendFloat64(o, z.Dither)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ImageRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Image":
			z.Image, bts, err = msgp.ReadBytesBytes(bts, z.Image)
			if err != nil {
				err = msgp.WrapError(err, "Image")
				return
			}
		case "DX":
			z.DX, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DX")
				return
			}
		case "DY":
			z.DY, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DY")
				return
			}
		case "Speed":
			z.Speed, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Speed")
				return
			}
		case "MaxColors":
			z.MaxColors, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxColors")
				return
			}
		case "Dither":
			z.Dither, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Dither")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ImageRequest) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.Image) + 3 + msgp.IntSize + 3 + msgp.IntSize + 6 + msgp.IntSize + 10 + msgp.IntSize + 7 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ImageResponse) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Image":
			z.Image, err = dc.ReadBytes(z.Image)
			if err != nil {
				err = msgp.WrapError(err, "Image")
				return
			}
		case "DX":
			z.DX, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "DX")
				return
			}
		case "DY":
			z.DY, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "DY")
				return
			}
		case "Palette":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Palette")
				return
			}
			if cap(z.Palette) >= int(zb0002) {
				z.Palette = (z.Palette)[:zb0002]
			} else {
				z.Palette = make([]RGBA, zb0002)
			}
			for za0001 := range z.Palette {
				err = z.Palette[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Palette", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ImageResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Image"
	err = en.Append(0x84, 0xa5, 0x49, 0x6d, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Image)
	if err != nil {
		err = msgp.WrapError(err, "Image")
		return
	}
	// write "DX"
	err = en.Append(0xa2, 0x44, 0x58)
	if err != nil {
		return
	}
	err = en.WriteInt(z.DX)
	if err != nil {
		err = msgp.WrapError(err, "DX")
		return
	}
	// write "DY"
	err = en.Append(0xa2, 0x44, 0x59)
	if err != nil {
		return
	}
	err = en.WriteInt(z.DY)
	if err != nil {
		err = msgp.WrapError(err, "DY")
		return
	}
	// write "Palette"
	err = en.Append(0xa7, 0x50, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Palette)))
	if err != nil {
		err = msgp.WrapError(err, "Palette")
		return
	}
	for za0001 := range z.Palette {
		err = z.Palette[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Palette", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ImageResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Image"
	o = append(o, 0x84, 0xa5, 0x49, 0x6d, 0x61, 0x67, 0x65)
	o = msgp.AppendBytes(o, z.Image)
	// string "DX"
	o = append(o, 0xa2, 0x44, 0x58)
	o = msgp.AppendInt(o, z.DX)
	// string "DY"
	o = append(o, 0xa2, 0x44, 0x59)
	o = msgp.AppendInt(o, z.DY)
	// string "Palette"
	o = append(o, 0xa7, 0x50, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Palette)))
	for za0001 := range z.Palette {
		o, err = z.Palette[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Palette", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ImageResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Image":
			z.Image, bts, err = msgp.ReadBytesBytes(bts, z.Image)
			if err != nil {
				err = msgp.WrapError(err, "Image")
				return
			}
		case "DX":
			z.DX, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DX")
				return
			}
		case "DY":
			z.DY, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DY")
				return
			}
		case "Palette":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Palette")
				return
			}
			if cap(z.Palette) >= int(zb0002) {
				z.Palette = (z.Palette)[:zb0002]
			} else {
				z.Palette = make([]RGBA, zb0002)
			}
			for za0001 := range z.Palette {
				bts, err = z.Palette[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Palette", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ImageResponse) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len(z.Image) + 3 + msgp.IntSize + 3 + msgp.IntSize + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Palette {
		s += z.Palette[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RGBA) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "R":
			z.R, err = dc.ReadUint8()
			if err != nil {
				err = msgp.WrapError(err, "R")
				return
			}
		case "G":
			z.G, err = dc.ReadUint8()
			if err != nil {
				err = msgp.WrapError(err, "G")
				return
			}
		case "B":
			z.B, err = dc.ReadUint8()
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
		case "A":
			z.A, err = dc.ReadUint8()
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RGBA) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "R"
	err = en.Append(0x84, 0xa1, 0x52)
	if err != nil {
		return
	}
	err = en.WriteUint8(z.R)
	if err != nil {
		err = msgp.WrapError(err, "R")
		return
	}
	// write "G"
	err = en.Append(0xa1, 0x47)
	if err != nil {
		return
	}
	err = en.WriteUint8(z.G)
	if err != nil {
		err = msgp.WrapError(err, "G")
		return
	}
	// write "B"
	err = en.Append(0xa1, 0x42)
	if err != nil {
		return
	}
	err = en.WriteUint8(z.B)
	if err != nil {
		err = msgp.WrapError(err, "B")
		return
	}
	// write "A"
	err = en.Append(0xa1, 0x41)
	if err != nil {
		return
	}
	err = en.WriteUint8(z.A)
	if err != nil {
		err = msgp.WrapError(err, "A")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RGBA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "R"
	o = append(o, 0x84, 0xa1, 0x52)
	o = msgp.AppendUint8(o, z.R)
	// string "G"
	o = append(o, 0xa1, 0x47)
	o = msgp.AppendUint8(o, z.G)
	// string "B"
	o = append(o, 0xa1, 0x42)
	o = msgp.AppendUint8(o, z.B)
	// string "A"
	o = append(o, 0xa1, 0x41)
	o = msgp.AppendUint8(o, z.A)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RGBA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "R":
			z.R, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "R")
				return
			}
		case "G":
			z.G, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "G")
				return
			}
		case "B":
			z.B, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
		case "A":
			z.A, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RGBA) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint8Size + 2 + msgp.Uint8Size + 2 + msgp.Uint8Size + 2 + msgp.Uint8Size
	return
}
