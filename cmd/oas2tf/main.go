package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mantisec/terraform-provider-umbrella/internal/openapi"
	"github.com/mantisec/terraform-provider-umbrella/internal/translate"
)

func main() {
	var (
		spec    = flag.String("spec", "", "OpenAPI spec path")
		outDir  = flag.String("out", "generated", "output Go dir")
		docsDir = flag.String("docs", "docs", "docs root dir")
	)
	flag.Parse()
	if *spec == "" {
		fmt.Println("-spec is required")
		os.Exit(1)
	}

	resources, err := openapi.ParseSpec(*spec)
	if err != nil {
		panic(err)
	}

	for _, rd := range resources {
		files, err := translate.GenerateAll(rd, *outDir, *docsDir)
		if err != nil {
			panic(err)
		}
		for path, data := range files {
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				panic(err)
			}
			if err := os.WriteFile(path, data, 0o644); err != nil {
				panic(err)
			}
			fmt.Println("generated", path)
		}
	}
}
