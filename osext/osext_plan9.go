//+build !go1.8

package osext

import (
	"os"
	"strconv"
	"syscall"
)

func executable() (string, error) {
	f, err := os.Open("/proc/" + strconv.Itoa(os.Getpid()) + "/text")
	if err != nil {
		return "", err
	}
	defer f.Close()
	return syscall.Fd2path(int(f.Fd()))
}
