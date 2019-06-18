package engine

import (
	hb "github.com/aymerick/raymond"
	hasher "github.com/cnf/structhash"
	"log"
	"path/filepath"
)

const (
	FIELD_REALPATH      = "realpath"
	FIELD_TEMPLATE_NAME = "templateName"
)

type Template struct {
	template *hb.Template
	context  *hb.DataFrame
	realpath string
	hash     string
}

//
// Add template variables to DataFrame
// - realpath ~ full path of template
// - templateName ~ filename of template
//
func templateContext(realpath string) *hb.DataFrame {
	frame := hb.NewDataFrame()
	frame.Set(FIELD_REALPATH, realpath)

	_, filename := filepath.Split(realpath)
	frame.Set(FIELD_TEMPLATE_NAME, filename)

	return frame
}

func (template *Template) Hash() (string, error) {
	if len(template.hash) == 0 {
		hashed, err := hasher.Hash(template.template, 1)
		if err != nil {
			return "", err
		}

		template.hash = hashed
	}

	return template.hash
}

func (template *Template) Reload() {
	templ, err := hb.ParseFile(template.realpath)
	if err != nil {
		log.Printf("error can't parse template file in %#v, caused by %#v", template.realpath, err)
		return
	}

	// Use hash as diff resolutions
	hashed, err := hasher.Hash(templ, 1)
	if err != nil {
		log.Printf("error can't hash template file in %#v, caused by %#v", template.realpath, err)
	}

	if hashed == template.hash {
		log.Printf("newer template are exactly same as old template, doesn't need to reload the template %s", template.realpath)
		return
	} else {
		log.Printf("reloading template %s", template.realpath)
		template.template = templ
	}
}

func NewTemplate(templ *hb.Template, realpath string, hashed string) *Template {
	return Template{templ, templateContext(realpath), realpath, hashed}
}

func ParseTemplate(realpath string) (template *Template, err error) {
	templ, err := hb.ParseFile(realpath)
	if err != nil {
		return nil, err
	}

	hashed, err := hasher.Hash(templ, 1)
	if err != nil {
		log.Printf("error can't hash template file in %#v, caused by %#v", template.realpath, err)
	}

	return NewTemplate(templ, realpath, hashed), err
}
