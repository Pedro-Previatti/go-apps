package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string      // define map to pass strings to templates
	IntMap    map[string]int         // define map to pass ints to templates
	FloatMap  map[string]float32     // define map to pass floats to templates
	DataMap   map[string]interface{} // define map to pass data to templates
	CSRFToken string                 // define security token to pass to templates
	Flash     string                 // define flash to pass to templates
	Warning   string                 // define warning to pass to templates
	Error     string                 // define error to pass to templates
}
