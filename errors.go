package gokit

import "fmt"

var (
	ErrESCreateIndex     = fmt.Errorf("es: create elasticsearch index error")
	ErrConsulKeyNotExist = fmt.Errorf("consul: key not exist")
)
