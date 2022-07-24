package main

import "github.com/nxtcoder17/go/pkg/logging"


func main(){
	logger := logging.NewOrDie(&logging.Options{
		Name: "sadsfa",
		Dev: true,
	})
	logger.Infof("asdfa")
}

