package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)
func RunContainerInitProcess(command string, args []string) error {
	log.Infof("command %s", command)
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func NewParentProcess(tty bool, command string) *exec.Cmd  {
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot:                     "",
		Credential:                 nil,
		Ptrace:                     false,
		Setsid:                     false,
		Setpgid:                    false,
		Setctty:                    false,
		Noctty:                     false,
		Ctty:                       0,
		Foreground:                 false,
		Pgid:                       0,
		Pdeathsig:                  0,
		Cloneflags:                 syscall.CLONE_NEWUTS | syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
		Unshareflags:               0,
		UidMappings:                nil,
		GidMappings:                nil,
		GidMappingsEnableSetgroups: false,
		AmbientCaps:                nil,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}