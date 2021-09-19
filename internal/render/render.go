package render

import (
	"github.com/Tomonaga007/bookings/internal/config"
	"github.com/Tomonaga007/bookings/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter,r *http.Request, tmpl string, td *models.TemplateData)  {
	var tc map[string]*template.Template
	if app.UseCache{
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")

	}
	//buff := new(bytes.Buffer)

	td = AddDefaultData(td,r)
	t.Execute(w,td)
	//buff.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template,error){

	cashedTemplates := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil{
		return cashedTemplates, err
	}

	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cashedTemplates, err
		}
		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil{
			return cashedTemplates, err
		}
		if len(matches) > 0{
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return cashedTemplates, err
			}
		}
		cashedTemplates[name] = ts
	}
	return cashedTemplates, nil
}

