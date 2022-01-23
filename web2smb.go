package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/hirochachacha/go-smb2"
)

type Smb struct {
	Namefs string
	Datafs string
}

func SmbGo(Datafs string) {
	conn, err := net.Dial("tcp", "localhost:445")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "Login",
			Password: "Password",
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer s.Logoff()

	fs, err := s.Mount("smbgo")
	if err != nil {
		panic(err)
	}
	defer fs.Umount()

	f, err := fs.OpenFile("DirectumDataTest.doc", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write([]byte(Datafs))
	if err != nil {
		panic(err)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
