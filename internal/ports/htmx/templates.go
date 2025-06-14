package htmx

import (
	"embed"
	"html/template"
)

// Embed htmx templates from filesystem at compile time.
//
//go:embed views/*
var views embed.FS

// readTemplates reads templates from embedded filesystem.
func readTemplates() (*template.Template, error) {
	tmpl := template.New("")

	// func are functions which can be used from templates.
	fncs := template.FuncMap{
		"dict": mapFromValues,
	}
	tmpl = tmpl.Funcs(fncs)

	// read templates from embedded filesystem
	return tmpl.ParseFS(views, "views/*")
}

// mapFromValues creates map from values
// where each odd value becomes a map key
// and each even value becomes a map value.
//
// mapFromValues("Name", "alice", "Age", 20)
// creates
//
//	map[string]any{
//		"Name": "alice",
//		"Age": 20,
//	}
func mapFromValues(values ...any) map[string]any {
	dict := make(map[string]any)
	for i := 0; i < len(values); i += 2 {
		key := values[i].(string)
		dict[key] = values[i+1]
	}
	return dict
}
