package main

import (
	"bytes"
	"io"
	"path"
	"path/filepath"

	"github.com/GeertJohan/go.rice"
	"github.com/flosch/pongo2"
)

type Pongo2Loader struct {
	box *rice.Box
}

func NewPongo2TemplatesLoader() (*Pongo2Loader, error) {
	fs := &Pongo2Loader{}

	p2l, err := rice.FindBox("templates")
	if err != nil {
		return nil, err
	}

	fs.box = p2l
	return fs, nil
}

func (fs *Pongo2Loader) Get(path string) (io.Reader, error) {
	myBytes, err := fs.box.Bytes(path)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(myBytes), nil
}

func (fs *Pongo2Loader) Abs(base, name string) string {
	me := path.Join(filepath.Dir(base), name)
	return me
}

func populateTemplatesMap(tSet *pongo2.TemplateSet, tMap map[string]*pongo2.Template) error {

	templates := [...]string{
		"index.html",
		"404.html",
		"401.html",
		"oops.html",

		"display/audio.html",
		"display/image.html",
		"display/video.html",
		"display/pdf.html",
		"display/file.html",
	}

	for _, tName := range templates {
		tpl, err := tSet.FromFile(tName)
		if err != nil {
			return err
		}

		tMap[tName] = tpl
	}

	return nil
}