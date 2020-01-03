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

//wrong directory gives error
func TestTempfile(t *testing.T) {
	directory := "/home/gslab/src/gophcsvrm/pre.go"
	_, err := tempfile(directory, "")
	if err == nil {
		t.Error(err)
	}
}

//negative test case gives error
func TestTransform1(t *testing.T) {
	var mode1 Mode
	mode1 = 7
	testFil, _ := os.Open("/home/gslab/goworkspace/src/gophercises/transform/img/fakemonalisa.png")
	_, err := Transform(testFil, "png", 2, myWithMode(mode1))
	if err != nil {
		t.Error(err)
	}
}

//positive test with no model gives error
func TestTransform(t *testing.T) {
	var mode1 Mode
	testFil, _ := os.Open("/home/gslab/goworkspace/src/gophercises/transform/img/fakemonalisa.png")
	Fakecopy2 = myCopy
	_, err := Transform(testFil, "png", 2, myWithMode(mode1))
	if err == nil {
		t.Error(err)
	}
}

//wrong-> mode, image, extension, num of shapes, mode options gives error
func TestTransformErr(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("")
	_, err := Transform(r, "", 0, myWithMode(mode1))
	if err == nil {
		t.Error(err)
	}
}

// //mock io copy
func TestTransform3(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("/home/gslab/goworkspace/src/gophercises/transform/img")
	Fakecopy1 = myCopy
	_, err := Transform(r, "", 0, myWithMode(mode1))
	if err == nil {
		t.Error(err)
	}
}

//temporary file mock
func TestTransform4(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("/home/gslab/goworkspace/src/gophercises/transform/img")
	Faketempfile2 = mytempfile
	_, err := Transform(r, "", 0, myWithMode(mode1))
	if err == nil {
		t.Error(err)
	}
}

//temporary file mock
func TestTransform2(t *testing.T) {
	var mode1 Mode
	r := strings.NewReader("/home/gslab/goworkspace/src/gophercises/transform/img")
	Faketempfile1 = mytempfile
	_, err := Transform(r, "", 0, myWithMode(mode1))
	if err == nil {
		t.Error(err)
	}
}
