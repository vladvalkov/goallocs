package build

import (
	"bytes"
	"io"
	"os/exec"
)

func Build(targetPackage string) (io.Reader, error) {
	_, err := exec.LookPath("go")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "build", `-gcflags=-S`, "-o", "/dev/null", targetPackage)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()
	if err != nil {
		return &errb, err
	}

	return &errb, nil
}
