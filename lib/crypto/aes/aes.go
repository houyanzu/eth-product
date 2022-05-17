package aes

type Decoder interface {
	Decode(key []byte) Encoder
	ToBytes() []byte
}

type Encoder interface {
	Encode(key []byte) Decoder
	ToString() string
}
