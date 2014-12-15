package exec

import (
	"os/exec"
	"time"
)

func ExecTimeout(d time.Duration, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Start(); err != nil {
		return err
	}

	if d <= 0 {
		return cmd.Wait() //command execute
	}

	done := make(chan error) //create new chan
	go func() {
		done <- cmd.Wait() //异步执行
	}()

	select {
	case <-time.After(d):
		cmd.Process.Kill() //超时即杀死进程
		return <-done

	case err := <-done:
		return err
	}

}
