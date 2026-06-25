package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sam-maton/go-web-starter-components/internal/template"
)

func cleanOutputDir() {
	os.RemoveAll("public")
	os.MkdirAll("public", 0755)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createPages(app template.Application, data templateData) {
	for k := range app.Cache {
		folder := "public/" + strings.TrimSuffix(k, ".html")
		path := strings.TrimSuffix(folder, "index")
		os.MkdirAll(path, 0755)

		f, err := os.Create(path + "/index.html")
		check(err)
		defer f.Close()

		app.Render(k, f, data)
	}
}

func copyStaticFiles() error {

	return filepath.WalkDir("ui/static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		assetPath := "public" + strings.TrimPrefix(path, "ui")
		if d.IsDir() {
			os.MkdirAll(assetPath, 0755)
		}

		if d.Type().IsRegular() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			in, err := os.OpenFile(path, os.O_RDONLY, info.Mode())
			if err != nil {
				return err
			}

			defer in.Close()

			out, err := os.OpenFile(assetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
			if err != nil {
				return err
			}

			defer out.Close()

			if _, err := io.Copy(out, in); err != nil {
				return err
			}

			err = out.Sync()
			if err != nil {
				return err
			}
		}

		return nil
	})
}
