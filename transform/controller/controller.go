package controller

import (
	"errors"
	"fmt"
	"gophercises/transform/primitive"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var Faketempfile1 = tempfile
var Faketempfile2 = tempfile
var Fakecopy = io.Copy
var Faketransform = primitive.Transform
var FakegenImages = genImages

//RenderNumShapeChoices ...
func RenderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) {
	opts := []genOpts{
		{N: 1, M: mode}, //these number of shapes are connected with
		{N: 1, M: mode}, //<a href="/modify/521721869.png?mode=1&n=2">
		{N: 1, M: mode}, //this '&n=2' parameter
		{N: 1, M: mode}, //we are not increasing these because of need of quick processing
	}
	imgs, err := genImages(rs, ext, opts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name      string
		Mode      primitive.Mode
		NumShapes int
	}
	var data []dataStruct
	for i, img := range imgs {
		data = append(data, dataStruct{
			Name:      filepath.Base(img),
			Mode:      opts[i].M,
			NumShapes: opts[i].N,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "no", 404) //insted of panic(err)
	}
}

//RenderModeChoices ...
func RenderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	opts := []genOpts{
		{N: 1, M: primitive.ModeTriangle},
		{N: 1, M: primitive.ModeCircle},
		{N: 1, M: primitive.ModePolygon},
		{N: 1, M: primitive.ModeCombo},
	}
	imgs, err := FakegenImages(rs, ext, opts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name string
		Mode primitive.Mode
	}
	var data []dataStruct
	for i, img := range imgs {
		data = append(data, dataStruct{
			Name: filepath.Base(img),
			Mode: opts[i].M,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "no", 404) //insted of panic(err)
	}
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err := genImage(rs, ext, opt.N, opt.M)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
	}
	return ret, nil
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	out, err := Faketransform(r, ext, numShapes, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}
	outFile, err := Faketempfile1("", ext)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, out)
	return outFile.Name(), nil
}

//Upload ...
func Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image") //this "image" is form-data key
	if err != nil {                          //we must use same name "image"
		http.Error(w, err.Error(), http.StatusBadRequest) //in request form data
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	onDisk, err := Faketempfile2("", ext)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	defer onDisk.Close()
	_, err = Fakecopy(onDisk, file)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/modify/"+filepath.Base(onDisk.Name()), http.StatusFound)
}

//Homepage -------------------------
func Homepage(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="image">
			<button type="submit">Upload Image</button>
		</form>
		</body></html>`
	fmt.Fprint(w, html)
}

//Modify ...
func Modify(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	ext := filepath.Ext(f.Name())[1:]
	modeStr := r.FormValue("mode")
	if modeStr == "" {
		RenderModeChoices(w, r, f, ext)
		return
	}
	mode, err := strconv.Atoi(modeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nStr := r.FormValue("n")
	if nStr == "" {
		RenderNumShapeChoices(w, r, f, ext, primitive.Mode(mode))
		return
	}
	numShapes, err := strconv.Atoi(nStr) //numshapes is used through its
	if err != nil {                      //struct which is Datastruct
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = numShapes //this numshapes is used nowhere
	http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
}

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil {
		return nil, errors.New("main: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
