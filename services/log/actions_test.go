package log

import (
	"fmt"
	"testing"
)

func TestMarshalIntoMap(t *testing.T) {
	var param = Params{
		FromBlock:       "0x1",
		Topic0:          "0x2",
		Topic1:          "0x3",
		Topic01Operator: OperatorOr,
		Topic12Operator: OperatorAnd,
	}
	fmt.Println(param.MarshalIntoMap())
}
