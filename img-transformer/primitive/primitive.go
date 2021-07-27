package primitive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Mode int

const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedrect
	ModeBeziers
	ModeRotatedellipse
	ModePolygon
)

func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

func Transform(image io.Reader, extension string, numShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}

	in, err := tmpfile("_in", extension)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary input file")
	}
	defer os.Remove(in.Name())

	out, err := tmpfile("_out", extension)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary output file")
	}
	defer os.Remove(out.Name())

	_, err = io.Copy(in, image)
	if err != nil {
		return nil, errors.New("primitive: failed to copy image to temp input file")
	}

	_, err = primitive(in.Name(), out.Name(), numShapes, args...)
	if err != nil {
		return nil, errors.New("primitive: failed to run primitive command")
	}

	b := bytes.NewBuffer(nil)

	_, err = io.Copy(b, out)
	if err != nil {
		return nil, errors.New("primitive: failed to copy out to buffer")
	}

	return b, nil
}

func primitive(input, output string, numShapes int, args ...string) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d", input, output, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func tmpfile(name, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", name)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary input file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", name, ext))
}
