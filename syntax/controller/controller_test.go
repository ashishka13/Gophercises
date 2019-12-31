package controller

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
)

func myCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	return 0, errors.New("testing generated copy error")
}

//testing hello api only
func TestHello(t *testing.T) {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Hello) //test this

	request, err := http.NewRequest("GET", "/", nil) //this is type of request to this url
	response := httptest.NewRecorder()               //this is function used to record our request
	mux.ServeHTTP(response, request)                 //our request is served to mux router

	if err != nil {
		t.Error("hello err")
	}
}

//basic testing of source code handler using /debug/ only and not the path
func TestSourceCodeHandler(t *testing.T) {
	mux := mux.NewRouter()
	mux.HandleFunc("/debug/", SourceCodeHandler) //main code with highliting
	request, err := http.NewRequest("GET", "/debug/", nil)
	response := httptest.NewRecorder()
	mux.ServeHTTP(response, request)

	if err != nil {
		t.Error("debug error")
	}
}

//some dummy url sent to debug. "path" and "line" are also provided
func TestSourceCodeHandler1(t *testing.T) {
	mux := mux.NewRouter()
	mux.HandleFunc("/debug/", SourceCodeHandler) //main code with highliting
	request, err := http.NewRequest("GET", "/debug/?line=71&path=%2Fhome%2Fgslab%2Fgoworkspace%2Fsrc%2Fgophercises%2Fsyntax%2Fmain.go", nil)
	response := httptest.NewRecorder()
	mux.ServeHTTP(response, request)
	if err != nil {
		t.Error("response fault")
	}
}

//this involves panic line
func TestPanicFun(t *testing.T) {
	mux := mux.NewRouter()
	mux.HandleFunc("/panic/", PanicDemo)

	request, _ := http.NewRequest("GET", "/panic/", nil)
	response := httptest.NewRecorder()
	handler := DevMw(mux)                //the stack created by panic is handled and highlited in DevMw
	handler.ServeHTTP(response, request) //so devmw is called using a handler of DevMw type, and mux is sent as argumrnt

	assert.Equal(t, 500, response.Code)
	// if err != nil {
	// 	t.Error("panic err")
	// }
}

func TestMockIOcopy(t *testing.T) {
	mux := mux.NewRouter()
	Fakecopy1 = myCopy
	mux.HandleFunc("/debug/", SourceCodeHandler) //main code with highliting
	request, err := http.NewRequest("GET", "/debug/?line=71&path=%2Fhome%2Fgslab%2Fgoworkspace%2Fsrc%2Fgophercises%2Fsyntax%2Fmain.go", nil)
	response := httptest.NewRecorder()
	mux.ServeHTTP(response, request)
	if err != nil {
		t.Error("response fault")
	}
}
