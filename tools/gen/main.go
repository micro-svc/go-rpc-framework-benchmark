package main

import (
	"flag"
	"io/ioutil"
	"strings"
	"time"

	"github.com/flosch/pongo2/v4"
	"github.com/gxxgle/go-utils/json"
	ulog "github.com/gxxgle/go-utils/log"
)

var (
	resultFilename = "../go-rpc-framework-benchmark.result"
	docsTmplFile   = "./gen/docs.tmpl"
	readmeTmplFile = "./gen/readme.tmpl"
	colors         = []string{"#264653", "#2a9d8f", "#e9c46a", "#f4a261", "#e76f51"}
	gen            = flag.String("gen", "docs", "gen result type. docs, readme")
)

type Result struct {
	Name      string
	Package   string
	Transport string
	Codec     string
	TPS       int64
	Color     string
	Detail    string
}

func (r Result) OK() bool {
	return len(r.Name) > 0 && len(r.Package) > 0 && r.TPS > 0
}

func main() {
	flag.Parse()
	ulog.ColorConsole()
	data, err := ioutil.ReadFile(resultFilename)
	ulog.FatalIfError(err)

	var (
		current Result
		results = []Result{}
	)

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			if current.OK() {
				results = append(results, current)
				current = Result{}
			}

			continue
		}

		// name
		if strings.HasPrefix(line, "###") {
			current.Name = strings.TrimPrefix(line, "###")
			current.Name = strings.TrimSpace(current.Name)
			continue
		}

		if strings.HasPrefix(line, "$ bin/") {
			line = strings.TrimPrefix(line, "$ bin/")
			for _, str := range strings.Split(line, " ") {
				str = strings.TrimSpace(str)
				// package name
				if strings.HasSuffix(str, "-cli") {
					current.Package = strings.TrimSuffix(str, "-cli")
				}
				// transport
				if strings.HasPrefix(str, "--transport=") {
					current.Transport = strings.TrimPrefix(str, "--transport=")
				}
				// codec
				if strings.HasPrefix(str, "--codec=") {
					current.Codec = strings.TrimPrefix(str, "--codec=")
				}
			}
			continue
		}

		// tps
		if strings.HasPrefix(line, "{") {
			current.Detail += line + "\n"
			current.TPS = json.GetFromString(line, "tps").Int()
			continue
		}
	}

	switch *gen {
	case "docs":
		genDocs(results)
	case "readme":
		genReadme(results)
	}
}

func setColor(rsts []Result) {
	idx := 0
	seted := make(map[string]string)
	for i, r := range rsts {
		if clr := seted[r.Package]; len(clr) > 0 {
			rsts[i].Color = clr
			continue
		}

		clr := colors[idx%len(colors)]
		rsts[i].Color = clr
		seted[r.Package] = clr
		idx++
	}
}

func genDocs(rsts []Result) {
	setColor(rsts)
	// sort.Slice(rsts, func(i, j int) bool { return rsts[i].TPS > rsts[j].TPS })
	tmpl, err := pongo2.FromFile(docsTmplFile)
	ulog.FatalIfError(err)
	data, err := tmpl.Execute(pongo2.Context{
		"UpdatedAt": time.Now().Format("2006-01-02"),
		"Results":   rsts,
	})
	ulog.FatalIfError(err)
	println(data)
}

func genReadme(rsts []Result) {
	tmpl, err := pongo2.FromFile(readmeTmplFile)
	ulog.FatalIfError(err)
	data, err := tmpl.Execute(pongo2.Context{
		"UpdatedAt": time.Now().Format("2006-01-02"),
		"Results":   rsts,
	})
	ulog.FatalIfError(err)
	println(data)
}
