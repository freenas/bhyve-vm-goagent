// +build freebsd netbsd openbsd

package termios

import (
	"syscall"
	"unsafe"
)

func SetTerm(fd *int) {
	termios := syscall.Termios{}
	termios.Cflag |= syscall.CS8 | syscall.CREAD | syscall.CLOCAL | syscall.B115200
	termios.Cflag &^= syscall.CSIZE | syscall.PARENB
	termios.Iflag &^= syscall.BRKINT | syscall.ICRNL | syscall.INPCK | syscall.ISTRIP | syscall.IXON
	termios.Oflag &^= syscall.OPOST
	termios.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.IEXTEN | syscall.ISIG
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0
	termios.Ispeed = syscall.B115200
	termios.Ospeed = syscall.B115200

	syscall.Syscall6(syscall.SYS_IOCTL, uintptr(*fd),
		uintptr(syscall.TIOCSETAW), uintptr(unsafe.Pointer(&termios)),
		0, 0, 0)
}
