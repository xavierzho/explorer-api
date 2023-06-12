/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package utils

import (
	"math/big"
	"strconv"
	"time"
)

// M is a type shorthand for param input
type M map[string]interface{}

// BigInt is a wrapper over big.Int to implement only unmarshalText
// for json decoding.
type BigInt big.Int

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (b *BigInt) UnmarshalText(text []byte) (err error) {
	var bigInt = new(big.Int)
	err = bigInt.UnmarshalText(text)
	if err != nil {
		return
	}

	*b = BigInt(*bigInt)
	return nil
}

// MarshalText implements the encoding.TextMarshaler
func (b *BigInt) MarshalText() (text []byte, err error) {
	return []byte(b.Int().String()), nil
}

// Int returns b's *big.Int form
func (b *BigInt) Int() *big.Int {
	return (*big.Int)(b)
}

// Time is a wrapper over big.Int to implement only unmarshalText
// for json decoding.
type Time time.Time

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *Time) UnmarshalText(text []byte) (err error) {
	input, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return
	}

	var timestamp = time.Unix(input, 0)
	*t = Time(timestamp)

	return nil
}

// Time returns t's time.Time form
func (t *Time) Time() time.Time {
	return time.Time(*t)
}

// MarshalText implements the encoding.TextMarshaler
func (t *Time) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(t.Time().Unix(), 10)), nil
}
