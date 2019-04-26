package main

import (
	"fmt"
	"os/exec"
	"context"
	"time"
)

type result struct {
	output []byte
	err error
}

func main() {

	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
		cmd *exec.Cmd
	)

	resultChan = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"pwd")
		output, err = cmd.CombinedOutput()

		resultChan <- &result {
			err: err,
			output: output
		}
	}()

	time.Sleep(1 * time.Second)

	cancelFunc()

	res = <- resultChan

}