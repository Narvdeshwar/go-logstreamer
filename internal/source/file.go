package source

type FileSource struct {
	Path string
}

func NewFileSource(path string) *FileSource {
	return &FileSource{Path: path}
}
