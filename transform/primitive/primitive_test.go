package primitive

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var Mymux = mux.NewRouter()

func myWithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}
func myCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	return 0, errors.New("fake error")
}
func mytempfile(prefix, ext string) (*os.File, error) {
	return nil, errors.New("test generates error")
}

// func TestModify1(t *testing.T) {

// 	request, _ := http.NewRequest("GET", "/modify/fakeowl.png", nil)
// 	Mymux.HandleFunc("/modify/{someParameter}", controller.Modify)
// 	response := httptest.NewRecorder()
// 	Mymux.ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "StatusOk found")
// }

//wrong directory
func TestTempfileErr(t *testing.T) {
	directory := "/home/gslab/src/gophcsvrm/pre.go"
	tempfile(directory, "")
}

//negative test case
func TestTransform1(t *testing.T) {
	var mode1 Mode
	mode1 = 7
	testFil, _ := os.Open("/home/gslab/goworkspace/src/gophercises/transform/img/fakemonalisa.png")
	Transform(testFil, "png", 2, myWithMode(mode1))
}

//positive test with no model
func TestTransform(t *testing.T) {
	var mode1 Mode
	testFil, _ := os.Open("/home/gslab/goworkspace/src/gophercises/transform/img/fakemonalisa.png")
	Fakecopy2 = myCopy
	Transform(testFil, "png", 2, myWithMode(mode1))
}

//wrong-> mode, image, extension, num of shapes, mode options
func TestTransformErr(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("")
	Transform(r, "", 0, myWithMode(mode1))
}

// //mock io copy
func TestTransform3(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("/home/gslab/goworkspace/src/gophercises/transform/img")
	Fakecopy1 = myCopy
	Transform(r, "", 0, myWithMode(mode1))
}

func TestTransform4(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("/home/gslab/goworkspace/src/gophercises/transform/img")
	Faketempfile2 = mytempfile
	Transform(r, "", 0, myWithMode(mode1))
}

func TestTransform2(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("/home/gslab/goworkspace/src/gophercises/transform/img")
	Faketempfile1 = mytempfile
	Transform(r, "", 0, myWithMode(mode1))
}

func TestWithmode(t *testing.T) {
	var mode1 Mode
	WithMode(mode1)
}
