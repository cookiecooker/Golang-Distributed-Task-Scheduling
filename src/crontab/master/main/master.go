package main

import (
	"runtime"
	"fmt"
)

var (
	confFile string
)

//parse command args
func initArgs() {
	flag.StringVar(&confFile, "config", "./master.json")
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //get max processes
}

func main() {

	var (
		err error
	)

	initArgs()

	initEnv() // init env

	//init config

	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	// start api http service
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)

}