package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Gabriel2233/gophercises/img-transformer/primitive"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body>
        <form action="/upload" enctype="multipart/form-data" method="post">
            <input type="file" name="image" />
            <button type="submit">Upload Image</button>
        </form>
        </body></html>`
		fmt.Fprint(w, html)
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		ext := filepath.Ext(header.Filename)[1:]
		onDisk, err := tmpfile("", ext)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer onDisk.Close()

		_, err = io.Copy(onDisk, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/modify/"+filepath.Base(onDisk.Name()), http.StatusFound)

	})

	mux.HandleFunc("/modify/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("./img/" + filepath.Base(r.URL.Path))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()

		ext := filepath.Ext(f.Name())[1:]

		modeStr := r.FormValue("mode")
		if modeStr == "" {
			renderModeChoices(w, r, f, ext)
			return
		}

		mode, err := strconv.Atoi(modeStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		numShapesStr := r.FormValue("n")
		if numShapesStr == "" {
			renderNumShapesChoices(w, r, f, ext, primitive.Mode(mode))
			return
		}

		numShapes, err := strconv.Atoi(numShapesStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_ = numShapes

		http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
	})

	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img/", fs))

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func renderNumShapesChoices(w http.ResponseWriter, r *http.Request, f io.ReadSeeker, ext string, mode primitive.Mode) {
	opts := []genOpts{
		{N: 5, M: mode},
		{N: 10, M: mode},
		{N: 15, M: mode},
		{N: 20, M: mode},
	}

	images, err := genImages(f, ext, opts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `<html><body>
		{{range .}}
        <a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
            <img style="width: 20%;" src="/img/{{.Name}}" />
        </a>
		{{end}}
		</html></body>`
	tpl := template.Must(template.New("").Parse(html))

	type dataStruct struct {
		Name      string
		Mode      primitive.Mode
		NumShapes int
	}

	var data []dataStruct
	for i, img := range images {
		data = append(data, dataStruct{
			Name:      filepath.Base(img),
			Mode:      opts[i].M,
			NumShapes: opts[i].N,
		})
	}
	tpl.Execute(w, data)
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, f io.ReadSeeker, ext string) {
	opts := []genOpts{
		{N: 10, M: primitive.ModeCircle},
		{N: 10, M: primitive.ModeRect},
		{N: 10, M: primitive.ModePolygon},
		{N: 10, M: primitive.ModeCombo},
	}

	images, err := genImages(f, ext, opts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `<html><body>
		{{range .}}
        <a href="/modify/{{.Name}}?mode={{.Mode}}">
            <img style="width: 20%;" src="/img/{{.Name}}" />
        </a>
		{{end}}
		</html></body>`
	tpl := template.Must(template.New("").Parse(html))

	type dataStruct struct {
		Name string
		Mode primitive.Mode
	}

	var data []dataStruct
	for i, img := range images {
		data = append(data, dataStruct{
			Name: filepath.Base(img),
			Mode: opts[i].M,
		})
	}
	tpl.Execute(w, data)
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	for _, opt := range opts {
		rs.Seek(0, 0)
		img, err := genImage(rs, ext, opt.N, opt.M)
		if err != nil {
			return nil, err
		}
		ret = append(ret, img)
	}

	return ret, nil
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	result, err := primitive.Transform(r, ext, numShapes, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}

	outFile, err := tmpfile("", ext)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, result)

	return outFile.Name(), nil
}

func tmpfile(name, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", name)
	if err != nil {
		return nil, errors.New("main: failed to create temporary input file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
