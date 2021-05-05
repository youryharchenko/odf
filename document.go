package odf

// Abstraction of the ODF document
type Document struct {

	// private
	container *Container
	//xmlparts  map[string]interface{}
	//body
}

func NewDocument(docType DocumentType) (doc *Document) {
	doc = &Document{
		container: NewContainer(docType),
		//xmlparts:  map[string]interface{}{},
	}
	return
}

func (doc *Document) AddFile(filePath string) *Document {
	return doc
}
