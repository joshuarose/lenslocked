package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir defaults to views/layouts/ but can be changed
	LayoutDir = "views/layouts/"
	//TemplateExt holds the extension to be used for templates, defaults to .gohtml
	TemplateExt = ".gohtml"
)

// NewView is an exported function for creating View structs
func NewView(layout string, files ...string) *View {
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
