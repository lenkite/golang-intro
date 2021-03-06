package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lenkite/golang-course/pkg/ppt"
)

const basicsPath = "tmpl/basics.html"
const addr = ":3000"

var config = ppt.Config{Name: "go-basics", Path: basicsPath, RefExt: ".go", RefDir: "samples", HotReload: true}
var pptBasics *ppt.Presentation

func init() {
	ppt, err := ppt.New(config)
	if err != nil {
		log.Panicf("Initializing ppt %q: %v", basicsPath, err)
	}
	pptBasics = ppt
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveGoBasicsPresentation)
	fmt.Printf("Serving %q on %q...\n", basicsPath, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func serveGoBasicsPresentation(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Creating ppt for %q...\n", basicsPath)
	_, err := pptBasics.WriteTo(w)
	fmt.Printf("Wrote ppt %q to response", pptBasics.Name)
	if err != nil {
		log.Panicf("Writing %q: %v", pptBasics.Name, err)
	}
}

// func loadPresentationTemplate(ppt string) (*template.Template, error) {
// 	tmpl, err := template.ParseFiles(ppt)
// 	if err != nil {
// 		return nil, fmt.Errorf("Loading ppt from: %q: %v", ppt, err)
// 	}
// 	dir := "samples"
// 	if err := addChildTemplates(tmpl, dir, ".go"); err != nil {
// 		return nil, fmt.Errorf("Loading template from: %q: %v", dir, err)
// 	}
// 	fmt.Printf("Templates Loaded%s\n", tmpl.DefinedTemplates())
// 	return tmpl, nil
// }

// func addChildTemplates(parent *template.Template, dir string, ext string) error {
// 	gofiles, err := listFiles(dir, ext)
// 	if err != nil {
// 		return fmt.Errorf("Error listing %q files: %v", ext, err)
// 	}
// 	if len(gofiles) == 0 {
// 		log.Panic("No go files found!")
// 	}
// 	for _, f := range gofiles {
// 		sample, err := ioutil.ReadFile(f)
// 		if err != nil {
// 			return fmt.Errorf("Cannot read go file %q: %v", f, err)
// 		}
// 		n := strings.TrimPrefix(f, "samples/")
// 		t, err := template.New(n).Parse(string(sample))
// 		if err != nil {
// 			return fmt.Errorf("Parsing %q: %v", f, err)
// 		}
// 		parent.AddParseTree(n, t.Tree)
// 	}
// 	return nil
// }

// func listFiles(dir string, ext string) ([]string, error) {
// 	var files []string
// 	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
// 		if filepath.Ext(path) == ext {
// 			files = append(files, path)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("Walking dir %s: %v", dir, err)
// 	}
// 	return files, err
// }
