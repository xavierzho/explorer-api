package logs

import (
	"github.com/Jonescy/explorer-api"
	"reflect"
	"strings"
)

type Action interface {
	explorer.Module
	GetLogs(params Params) ([]Log, error)
}
type Operator string

const (
	OperatorAnd = "and"
	OperatorOr  = "or"
)

type Params struct {
	FromBlock       string   `json:"fromBlock,omitempty"`
	ToBlock         string   `json:"toBlock,omitempty"`
	Address         string   `json:"address"`
	Topic0          string   `json:"topic0"`
	Topic1          string   `json:"topic1,omitempty"`
	Topic2          string   `json:"topic2,omitempty"`
	Topic3          string   `json:"topic3,omitempty"`
	Topic01Operator Operator `json:"topic0_1_opr,omitempty"`
	Topic12Operator Operator `json:"topic1_2_opr,omitempty"`
	Topic23Operator Operator `json:"topic2_3_opr,omitempty"`
	Topic02Operator Operator `json:"topic0_2_opr,omitempty"`
	Topic03Operator Operator `json:"topic0_3_opr,omitempty"`
	Topic13Operator Operator `json:"topic1_3_opr,omitempty"`
	Page            string   `json:"page,omitempty"`
	Offset          string   `json:"offset,omitempty"`
}

func (p *Params) MarshalIntoMap() map[string]string {
	v := reflect.ValueOf(*p)
	t := v.Type()
	m := make(map[string]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		name := field.Name
		val := v.FieldByName(field.Name)
		// don't check if it's omitted
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		tags := strings.Split(tag, ",")
		if tags[0] != "" {
			name = tags[0]
		}
		if strings.Contains(tag, "omitempty") {
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()
			if reflect.DeepEqual(current, zero) {
				continue
			}
		}

		m[name] = v.Field(i).String()
	}
	return m
}

type Log struct {
	TimeStamp        string   `json:"timeStamp"`
	GasUsed          string   `json:"gasUsed"`
	Address          string   `json:"address"`
	LogIndex         string   `json:"logIndex"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	GasPrice         string   `json:"gasPrice"`
}
