package contracts

import (
	"github.com/saylom99/service-helper/bootstrap"
)

type Contract struct {
	micro bootstrap.MicroModule
	redis bootstrap.RedisDB
}

var contract = new(Contract)
