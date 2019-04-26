package main

import (
	"os/exec" 
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	cmd = exec.Command("pwd")

	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println(output)

}