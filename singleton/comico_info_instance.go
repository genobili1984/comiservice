package singleton

import (
	"sync"
)

type comicoInfoIntance struct {
}

var instance *comicoInfoIntance
var once sync.Once

func GetComicoInfoInstance() *comicoInfoIntance {
	once.Do(func() {
		instance = &comicoInfoIntance{}
	})
	return instance
}

func (ins *comicoInfoIntance) initData() {

}

func (ins *comicoInfoIntance) GetComicoInfo(comicoId int) interface{} {
	return nil
}
