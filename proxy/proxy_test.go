package proxy

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumberString2Hex(t *testing.T) {
	var s = "88888"
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {

	}
	fmt.Println(i)
	var h = fmt.Sprintf("%d", s)
	fmt.Println(h)
}
