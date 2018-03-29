package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir defaults to views/layouts/ but can be changed
	LayoutDir = "views/layouts/"
	// TemplateDir is prepending on view names
	TemplateDir = "views/"
	//TemplateExt holds the extension to be used for templates, defaults to .gohtml
	TemplateExt = ".gohtml"
)

// NewView is an exported function for creating View structs
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View is a struct that contains all template info
type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutsFiles returns a slice of strings representing
// the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// addTemplatePath takes in a slice of strings
// representing file paths for templates
// and it prepends TemplateDir directory to each string
//
// Eg the input {"home"} would result in ouput {"views/home"}
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes in a slice of strings
// representing file extensions for templates
// and it appends TemplateExt extension to each string
//
// Eg the input {"home"} would result in ouput {"home.gohtml"}
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
