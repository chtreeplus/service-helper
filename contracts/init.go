package contracts

import (
	"repository.ch3plus.com/utility/service-helper/bootstrap"
)

type Contract struct {
	micro bootstrap.MicroModule
	redis bootstrap.RedisDB
}

var contract = new(Contract)
