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

var Faketempfile1 = tempfile
var Faketempfile2 = tempfile
var Fakecopy1 = io.Copy
var Fakecopy2 = io.Copy

// Mode defines the shapes used when transforming images.
type Mode int

// Modes supported by the primitive package.
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

// Transform will take the provided image and apply a primitive
// transformation to it, then return a reader to the resulting image.
func Transform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}

	in, err := Faketempfile1("in_", ext)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary input file")
	}
	defer os.Remove(in.Name())
	out, err := Faketempfile2("in_", ext)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary output file")
	}
	defer os.Remove(out.Name())

	// Read image into in file
	_, err = Fakecopy1(in, image)
	if err != nil {
		return nil, errors.New("primitive: failed to copy image into temp input file")
	}

	// Run primitive w/ -i in.Name() -o out.Name()
	stdCombo, err := primitive(in.Name(), out.Name(), numShapes, args...)
	if err != nil {
		return nil, fmt.Errorf("primitive: failed to run the primitive command. stdcombo=%s", stdCombo)
	}

	// read out into a reader, return reader, delete out
	b := bytes.NewBuffer(nil)
	_, err = Fakecopy2(b, out)
	if err != nil {
		return nil, errors.New("primitive: Failed to copy output file into byte buffer")
	}
	return b, nil
}

func primitive(inputFile, outputFile string, numShapes int, args ...string) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d", inputFile, outputFile, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...) //this is where actual fogleman/primitive package is accessed
	b, err := cmd.CombinedOutput()            //baeause it is a command line program
	return string(b), err
}

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
