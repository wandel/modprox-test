package main

import (
	"log"
	"github.com/pkg/errors"
)

func main(){
	base := errors.New("base error")
	err := errors.Wrap(base, "wrapper error")
	log.Println(err)
}