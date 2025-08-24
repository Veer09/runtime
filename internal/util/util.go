package util

import (
	"os"
	"math"

	"golang.org/x/sys/unix"
)

func NewSocketPair(name string) (*os.File, *os.File, error) {
	fd, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}
	return os.NewFile(uintptr(fd[0]), name + "-parent"), os.NewFile(uintptr(fd[1]), name + "-child"), nil
}

func CloseExecFrom(minFd int) error {
	err := unix.CloseRange(uint(minFd), math.MaxInt32, unix.CLOSE_RANGE_CLOEXEC)
	if err != nil {
		return err
	}
	return nil
}