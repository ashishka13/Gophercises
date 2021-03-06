package controller

import (
	"errors"
	"fmt"
	"gophercises/transform/primitive"
	"html/template"
	"io"
	"io/ioutil"
	"log"
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

// WithMode is an option for the Transform function that will define the
// mode you want to use. By default, ModeTriangle will be used.
func WithMode(mode primitive.Mode) func() []string {
	return func() []string {
		fmt.Println("this is withmode ", []string{"-m", fmt.Sprintf("%d", mode)})
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

//RenderNumShapeChoices ____________________
func RenderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) {
	opts := []genOpts{
		{N: 1, M: mode}, //these number of shapes are connected with
		{N: 1, M: mode}, //<a href="/modify/521721869.png?mode=1&n=2">
		{N: 1, M: mode}, //this '&n=2' parameter
		{N: 1, M: mode}, //we are not increasing these because of need of quick processing
	}
	imgs, err := genImages(rs, ext, opts...)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//"/modify/320985324.png?mode=8&n=1"
	//this will show urls for all the 4 images

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

//RenderModeChoices _________________________
func RenderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	opts := []genOpts{
		{N: 1, M: primitive.ModeTriangle},
		{N: 1, M: primitive.ModeCircle},
		{N: 1, M: primitive.ModePolygon},
		{N: 1, M: primitive.ModeCombo},
	}
	imgs, err := FakegenImages(rs, ext, opts...) //... means take in all the options inside the "opts" slice
	if err != nil {
		log.Println(err)
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
	err = tpl.Execute(w, data) //this will execute the HTML code snippet on webpage
	if err != nil {
		http.Error(w, "no", 404) //insted of panic(err)
	}
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string //... is argument to varadic function
	for _, opt := range opts {
		rs.Seek(0, 0) //start making changes from 0,0 location from inside the file
		f, err := genImage(rs, ext, opt.N, opt.M)
		fmt.Println("this is gen image f ", f)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ret = append(ret, f) //appenf "f" filenames to the array
	}
	return ret, nil
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	fmt.Println("this is mode struct ", mode)
	out, err := Faketransform(r, ext, numShapes, WithMode(mode))
	// fmt.Println("outs name ", out)
	if err != nil {
		log.Println(err)
		return "", err
	}
	outFile, err := Faketempfile1("", ext) //generate temporary file for local storage
	// fmt.Println("outfile name immediately ", outFile.Name())
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, out) //copy generated image's io.reader into the empty temp file
	fmt.Println("outfile name ", outFile.Name())
	return outFile.Name(), nil
}

//Homepage ----------------------
func Homepage(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="image">
			<button type="submit">Upload Image</button>
		</form>
		</body></html>`
	fmt.Fprint(w, html)
}

//Upload _______________________
func Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image") //this "image" is form-data key and also written in "input tag name"
	if err != nil {                          //we must use same name "image"
		http.Error(w, err.Error(), http.StatusBadRequest) //in request form data
		return
	} //file uploaded
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:] //file extension extracted
	onDisk, err := Faketempfile2("", ext)    //create an empty file with same extension as uploaded image
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	defer onDisk.Close()
	_, err = Fakecopy(onDisk, file) //copy the contents of uploaded file into the temp file and keep the original file intact
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/modify/"+filepath.Base(onDisk.Name()), http.StatusFound) //now image processing phase
	//filepath.base will return the filename.png and will add it to /modify/ url
}

//Modify _____________________
func Modify(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./img/" + filepath.Base(r.URL.Path)) //open the image in the URL
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	ext := filepath.Ext(f.Name())[1:] //again extract extension for modified output files
	modeStr := r.FormValue("mode")    //extract mode from http name tag
	fmt.Println("mode str is:", modeStr)
	if modeStr == "" {
		RenderModeChoices(w, r, f, ext)
		return
	}
	mode, err := strconv.Atoi(modeStr) //convert "mode" string into integer
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nStr := r.FormValue("n") //extract number of shapes to be generated from name tag
	fmt.Println("n str is: ", nStr)
	if nStr == "" {
		RenderNumShapeChoices(w, r, f, ext, primitive.Mode(mode))
		return
	}
	numShapes, err := strconv.Atoi(nStr) //convert "numberOfShapes" string into integer
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = numShapes //this numshapes is used nowhere after here
	http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
}

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil {
		log.Println(err)
		return nil, errors.New("main: failed to create temporary file")
	}
	fmt.Println("tempfile name: ", in.Name())
	defer os.Remove(in.Name()) //delete the temp file after new image generation. these filrs are without extension
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
