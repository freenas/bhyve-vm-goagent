package termios

import (
	"log"
	"syscall"
)

const DEBUG int = 1

func NewConnection(vconsole string) int {
	fd, err := syscall.Open(vconsole, syscall.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}

	SetTerm(&fd)
	return fd
}

func CloseConnection(fd int) {
	syscall.Close(fd)
}

func Read(fd int) []byte {
	var MaxRead int = 1024
	var numread int
	var err error
	var guestInfo []byte

	buffer := make([]byte, 1024)
	numread, err = syscall.Read(fd, buffer)
	if err != nil {
		panic(err.Error())
	}

	if numread < MaxRead {
		MaxRead = numread
	}
	guestInfo = append(buffer[:MaxRead])

	if DEBUG == 1 {
		log.Println("===> READ COMMAND: ", string(guestInfo[:MaxRead]))
	}

	return guestInfo
}

func Write(fd int, guestInfo []byte) {
	_, err := syscall.Write(fd, guestInfo)
	if err != nil {
		panic(err.Error())
	}
	if DEBUG == 1 {
		log.Println("===> Writting: ", string(guestInfo))
	}
}
