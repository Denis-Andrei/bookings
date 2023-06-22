package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/Denis-Andrei/bookings/pkg/config"
	"github.com/Denis-Andrei/bookings/pkg/models"
)


var app *config.Appconfig

//NewTemplate sets the config for the template package
func NewTemplates(a *config.Appconfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}
	
	
	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// this is an arbitrary choise and is not needed 
	// this will hold bytes
	//	****Execute that buffer directly and then write it out.***
	// And the only reason I'm doing this is for finer grained error checking, because once I have declared
	buffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buffer, td)

	if err !=nil {
		log.Println(err)
	}

	//render the template

	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	//another way to create a Map without make keyword
	myCache := map[string]*template.Template{}

	//get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range thorugh all files
	for _, page := range pages {
		name := filepath.Base(page) // returns the last element of a path which will be the file name e.g home.page.tmpl
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil

}
