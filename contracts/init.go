package contracts

import (
	"github.com/Decem-Technology/service-helper/bootstrap"
)

type Contract struct {
	micro bootstrap.MicroModule
	redis bootstrap.RedisDB
}

var contract = new(Contract)
