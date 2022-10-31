package models

import "Bookme/internal/forms"

//TemplateData holds data sent from handlers to templates
type TemplateData struct{
	//map from string to string
	StringMap map[string]string 
	IntMap map[string]int 
	FloatMap map[string]float32
	Data 	map[string]interface{}
	CSRFToken string 
	Flash string 
	Warning string 
	Form *forms.Form
	Error string 
	IsAuthenticated int 
	

}