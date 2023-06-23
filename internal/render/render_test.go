package render

import (
	"bookings-udemy/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {

	var td models.TemplateData
	r,err:=getSession()
	if err!=nil{
		t.Error(err)
	}

	session.Put(r.Context(),"flash","123")

	result:=AddDefaultData(&td,r)

	if result.Flash!="123"{
		t.Error("flash value of 123 not found in the session")
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplate="./../..templates"
	tc,err:=CreateTemplateCache()
	if err!=nil{
		t.Error(err)
	}
	app.TemplateCache=tc

	r,err:=getSession()
	if err!=nil{
		t.Error(err)
	}

	var ww myWriter

	err=RenderTemplate(&ww,r,"home.page.tmpl",&models.TemplateData{})
	if err!=nil{
		t.Error("Error writing template to browser")
	}

	err=RenderTemplate(&ww,r,"non-existent.page.tmpl",&models.TemplateData{})
	if err==nil{
		t.Error("Rendered template that does not exist")
	}

}

func getSession()(*http.Request,error){

	r,err:=http.NewRequest("GET", "/some-url", nil)

	if err!=nil{
		return nil,err
	}

	ctx:=r.Context()
	ctx, _=session.Load(ctx,r.Header.Get("X-Sessio"))
	r=r.WithContext(ctx)

	return r,nil

}

func TestNewTemplates(t *testing.T){
	NewTemplates(app)
}
func TestCreateTemplateCache(t *testing.T){
	pathToTemplate="./../../templates"
	_,err:=CreateTemplateCache()
	if err!=nil{
		t.Error(err)
	}

}
