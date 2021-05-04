package odf

import (
	"fmt"
	"path"
)

// Container class, ODF file management
type Container struct {

	// private

}

// Return a Container instance based on template argument.
func NewContainer(docType DocumentType) (container *Container) {

	tmplName := docType.Template()
	tmplPath := path.Join(ODF_TEMPLATES_DIR, tmplName)
	container = new(Container)
	err := container.open(tmplPath)
	if err != nil {
		panic(fmt.Errorf("Template '%s' not found", tmplPath))
	}
	container = container.clone()
	return
}

func (container *Container) open(filePath string) error {
	return nil
}

func (container *Container) clone() (newContainer *Container) {
	return
}
