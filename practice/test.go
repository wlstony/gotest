package main

import (
	"fmt"
	"strconv"
	"syscall"
)

func CheckPID(pid string) (exist bool, err error) {
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return
	}

	err = syscall.Kill(pidInt, 0)
	if err != nil {
		return
	}

	exist = true
	return
}
func alwaysFalse() bool {
	return false
}

func main() {
	c := ConfigOne{}
	fmt.Println(&c)
}

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string  {
	return fmt.Sprintf("%v", c)
}