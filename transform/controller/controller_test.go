package controller

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var Mymux = mux.NewRouter()

func Mytempfile(prefix, ext string) (*os.File, error) {
	return nil, errors.New("test generated temp error")
}

func Mycopy(dst io.Writer, src io.Reader) (written int64, err error) {
	return 0, errors.New("testing generated copy error")
}

func Mytransform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	return nil, errors.New("test generated primitive.Transform error")
}

func MygenImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	return nil, errors.New("testing generated genImage error")
}

func TestHomepage(t *testing.T) {
	Mymux.HandleFunc("/", Homepage)                //test this url
	request, _ := http.NewRequest("GET", "/", nil) //this is type of request to this url
	response := httptest.NewRecorder()             //this is function used to record our request
	Mymux.ServeHTTP(response, request)             //our request is served to mux router
}

//negative
func TestTempfileErr(t *testing.T) {

	tempfile("/fakeDirectoryName/", "")
}

//modify test cases

//plain url entered
func TestModifyURL(t *testing.T) {

	Mymux.HandleFunc("/modify/", Modify)                  //test this url
	request, _ := http.NewRequest("GET", "/modify/", nil) //this is type of request to this url
	response := httptest.NewRecorder()                    //this is function used to record our request
	Mymux.ServeHTTP(response, request)                    //our request is served to mux router
}

//provide correct file
func TestModify1(t *testing.T) {

	request, _ := http.NewRequest("GET", "/modify/fakeowl.png", nil)
	Mymux.HandleFunc("/modify/{someParameter}", Modify) //we provide {some} because
	//{} is used to specify variable in url. also we can add it to our main url.
	//but it is optional. because here the parameters are sent through test cases
	//and not from main
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "StatusOk found")
}

//provide correct file with mode but not number of shapes
func TestModify3(t *testing.T) {

	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=3", nil)
	//	(if nStr == "") will hit here
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
}

//send string in number of modes
func TestModify5(t *testing.T) {

	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=a&n=df", nil)
	//numShapes, err := strconv.Atoi(nStr) string sent insted on int
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
}

//send string in number of shapes
func TestModify4(t *testing.T) {

	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=2&n=df", nil)
	//numShapes, err := strconv.Atoi(nStr) string sent insted on int
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
}

//provide correct filem, mode and number of shapes
func TestModify2(t *testing.T) {

	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=2&n=10", nil)

	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)

	assert.Equal(t, 302, response.Code, "Status Found")
}

//mock fakecopy2
func TestModify6(t *testing.T) {
	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=2&n=10", nil)
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)

	assert.Equal(t, 302, response.Code, "Status Found")
}

//modify test complete

////////////////////////////////////////////////////////////////////////////////////////

//upload test

func TestUploadURL(t *testing.T) {
	Mymux.HandleFunc("/upload", Upload)                   //just like we type url in address bar
	request, _ := http.NewRequest("POST", "/upload", nil) //write request for Mux router with method type
	response := httptest.NewRecorder()                    //initialize response recorder for our request
	Mymux.ServeHTTP(response, request)                    //just like we press enter key on keyboard
}

//positive testing
func TestUpload1(t *testing.T) {
	//postman gives us all the necessary header parameters and file upload procedure as below
	Mymux.HandleFunc("/upload", Upload)
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"image\"; filename=\"fakeowl.png\"\r\nContent-Type: image/png\r\n\r\n\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	request, _ := http.NewRequest("POST", "/upload", payload)

	request.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")

	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
}

//mock io.copy for error
func TestUpload3(t *testing.T) {
	Mymux.HandleFunc("/upload", Upload)
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"image\"; filename=\"fakeowl.png\"\r\nContent-Type: image/png\r\n\r\n\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	request, _ := http.NewRequest("POST", "/upload", payload)

	request.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")

	response := httptest.NewRecorder()
	Fakecopy = Mycopy
	Mymux.ServeHTTP(response, request)
}

//mock tempfile1 for error
func TestUpload2(t *testing.T) {
	//postman gives us all the necessary header parameters and file upload procedure as below
	Mymux.HandleFunc("/upload", Upload)
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"image\"; filename=\"fakeowl.png\"\r\nContent-Type: image/png\r\n\r\n\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	request, _ := http.NewRequest("POST", "/upload", payload)

	request.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")

	response := httptest.NewRecorder()
	Faketempfile2 = Mytempfile
	Mymux.ServeHTTP(response, request)
}

//mock tempfile1
func TestMockTempfile1(t *testing.T) {
	Faketempfile1 = Mytempfile
	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=3", nil)
	//	(if nStr == "") will hit here
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
}

//mock transform
func TestMockTransform(t *testing.T) {
	Faketransform = Mytransform
	request, _ := http.NewRequest("GET", "/modify/fakeowl.png?mode=3", nil)
	//	(if nStr == "") will hit here
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
}
func TestMockGenImages(t *testing.T) {
	FakegenImages = MygenImages

	request, _ := http.NewRequest("GET", "/modify/fakeowl.png", nil)
	Mymux.HandleFunc("/modify/{someParameter}", Modify)
	response := httptest.NewRecorder()
	Mymux.ServeHTTP(response, request)
	if response.Code == 200 {
		t.Error("expecting error not getting error")
	}
}
