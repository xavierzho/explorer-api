package iface

import "github.com/xavierzho/explorer-api/utils"

type Client interface {
	Call(module Module, action string, param utils.M, outcome any) error
}
