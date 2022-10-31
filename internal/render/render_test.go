package render 

import (
	"Bookme/internal/models"
	"net/http"
	"testing"
)

//tedt the func AddeDefaulTData 
func TestAddDefaultData(t *testing.T){
	var td models.TemplateData 
	//we need request we use getsession
	r, err := getSession()
	if err != nil{
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	//cal the function ---
	result := AddDefaultData(&td, r)

	if result.Flash != "123"{
		t.Error("flash value of 123 not found in session")
	}
}


func TestRenderTemplate(t *testing.T){
	//path to the templates
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil{
		t.Error(err)
	}
	app.TemplateCache = tc 
	r, err := getSession()
	if err != nil{
		t.Error(err)
	}
	var ww myWriter 

	err = Template(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err == nil{
		t.Error("rendered template that does not exist")
	}
}

func getSession() (*http.Request, error){
	//get request with some-url and a body nil 
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil{
		return nil, err 
	}
	//get the context from the request we made
	ctx := r.Context()
	//attach the session data to the context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	//put my contact back to my request
	r = r.WithContext(ctx)
	return r, nil 
}


func TestNewTemplae(t *testing.T){
	NewRenderer(app)
}

func TestCreateTemplateCache(t *testing.T){
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil{
		t.Error(err)
	}
}