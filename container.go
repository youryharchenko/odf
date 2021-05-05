package odf

import (
	"fmt"
	"os"
	"path"
)

// Container class, ODF file management

// Representation of the ODF file.
type Container struct {
	Path string
	// private
	parts     map[string]interface{}
	partsTs   map[string]interface{}
	pathLike  string
	packaging string
}

// Return a Container instance based on template argument.
func NewContainer(docType DocumentType) (container *Container) {
	container = &Container{
		parts:     map[string]interface{}{},
		partsTs:   map[string]interface{}{},
		pathLike:  "",
		packaging: "zip",
	}

	tmplName := docType.Template()
	tmplPath := path.Join(ODF_TEMPLATES_DIR, tmplName)

	err := container.Open(tmplPath)
	if err != nil {
		panic(fmt.Errorf("Template '%s' not found", tmplPath))
	}
	container = container.Clone()
	return
}

// Load the content of an ODF file
func (container *Container) Open(filePath string) error {
	inf, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	container.pathLike = filePath
	container.Path = filePath
	if inf.IsDir() {
		return container.readFolder()
	}
	return container.readZip()
}

// Make a copy of this container with no path.
func (container *Container) Clone() (newContainer *Container) {
	return
}

func (container *Container) readFolder() error {
	return nil
}

func (container *Container) readZip() error {
	return nil
}
