package input

import (
	"fmt"
	"strings"

	"github.com/diegojromerolopez/congolway/pkg/base"
)

// GolReader : tasked with reading a Game of Life from files
type GolReader struct {
	readGol base.GolInterface
}

// NewGolReader : returns a new pointer to GolReader
func NewGolReader(g base.GolInterface) *GolReader {
	return &GolReader{g}
}

// ReadFile : read a file from a path
func (gr *GolReader) ReadFile(filename string, gconf *base.GolConf) (base.GolInterface, error) {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex < 0 {
		return nil, fmt.Errorf("File \"%s\" has no extension. Only .txt and .cells files are allowed", filename)
	}
	fileExtension := filename[lastDotIndex:]

	if fileExtension == ".txt" {
		return gr.ReadCongolwayFile(filename)
	} else if fileExtension == ".cells" {
		return gr.ReadCellsFile(filename, gconf)
	} else if fileExtension == ".life" {
		return gr.ReadLifeFile(filename, gconf)
	} else if fileExtension == ".gif" {
		return gr.ReadGifFile(filename, gconf)
	}
	return nil, fmt.Errorf(
		"File extension \"%s\" not recognized. "+
			"Only .txt, .cells and .gif are allowed", fileExtension)
}
