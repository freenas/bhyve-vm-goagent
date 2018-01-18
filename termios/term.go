package termios

import (
	"log"
	"syscall"
)

const DEBUG int = 0

func Read(vconsole string) []byte {
	var MaxRead int = 1024
	var fd, numread int
	var err error
	var guestInfo []byte

	fd, err = syscall.Open(vconsole, syscall.O_RDONLY, 0666)
	if err != nil {
		panic(err.Error())
	}

	SetTerm(&fd)

	if DEBUG == 1 {
		log.Println("===> Open device READ: ", vconsole)
	}

	buffer := make([]byte, MaxRead)
	numread, err = syscall.Read(fd, buffer)
	if err != nil {
		panic(err.Error())
	}

	if numread < MaxRead {
		MaxRead = numread
	}
	guestInfo = append(buffer[:MaxRead])
	syscall.Close(fd)

	if DEBUG == 1 {
		log.Println("===> READ COMMAND: ", string(guestInfo))
	}

	return guestInfo
}

func Write(vconsole string, guestInfo []byte) {
	var fd int
	var err error

	fd, err = syscall.Open(vconsole, syscall.O_WRONLY, 0666)
	if err != nil {
		panic(err.Error())
	}

	SetTerm(&fd)
	if DEBUG == 1 {
		log.Println("===> Open device WRITE: ", vconsole)
	}
	syscall.Write(fd, guestInfo)
	if DEBUG == 1 {
		log.Println("===> Writting: ", string(guestInfo))
	}
	syscall.Close(fd)
}
