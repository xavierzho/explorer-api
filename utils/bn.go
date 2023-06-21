package utils

import "math/big"

// BN is a wrapper over big.Int to implement only unmarshalText
// for json decoding.
type BN big.Int

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (b *BN) UnmarshalText(text []byte) (err error) {
	var bigInt = new(big.Int)
	err = bigInt.UnmarshalText(text)
	if err != nil {
		return
	}

	*b = BN(*bigInt)
	return nil
}

// MarshalText implements the encoding.TextMarshaler
func (b *BN) MarshalText() (text []byte, err error) {
	return []byte(b.Int().String()), nil
}

// Int returns b's *big.Int form
func (b *BN) Int() *big.Int {
	return (*big.Int)(b)
}
