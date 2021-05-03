package odf

import "path"

// Container class, ODF file management
type Container struct {

	// private

}

// Return a Container instance based on template argument.
func NewContainer(target interface{}) (container *Container) {
	if target == nil {
		return new(Container)
	}

	switch target := target.(type) {
	case *Container:
		container = target
	case Container:
		container = &target
	case string:
		// Return a Container instance based on template argument
		tmplName, ok := ODF_TEMPLATES[target]
		if ok {
			tmplPath := path.Join(ODF_TEMPLATES_DIR, tmplName)

		} else {
			container = &Container{}
			err := container.open(target)
			if err != nil {
				return nil
			}
		}

	}

	return
}

func (container *Container) open(filePath string) error {

}
