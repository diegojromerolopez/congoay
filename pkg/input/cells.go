package input

import (
	"bufio"
	"io"
	"os"

	"github.com/diegojromerolopez/congolway/pkg/base"
	"github.com/diegojromerolopez/congolway/pkg/statuses"
)

// ReadCellsFile : create a new Game of life from a cells file
func (gr *GolReader) ReadCellsFile(filename string, rowsLimitation string,
	colsLimitation string, generation int, neighborhoodType int) (base.GolInterface, error) {
	file, fileError := os.Open(filename)
	defer file.Close()

	if fileError != nil {
		return nil, fileError
	}

	reader := bufio.NewReader(file)

	// Name of the GOL pattern
	// TODO: store name of GOL pattern
	_, nameError := readCellFileLine(reader)
	if nameError != nil {
		return nil, nameError
	}
	gridLine := ""
	// Description of the pattern
	for true {
		line, err := readCellFileLine(reader)
		if err != nil {
			return nil, err
		}
		if line[0:1] != "!" {
			gridLine = line
			break
		}
	}

	cellValueCorrespondence := map[string]int{".": statuses.DEAD, "O": statuses.ALIVE}
	rows := 0
	cols := len(gridLine)
	cells := make([]int, 0, cols)
	lastLoop := false
	for true {
		rows++
		for j := 0; j < cols; j++ {
			cells = append(cells, cellValueCorrespondence[gridLine[j:j+1]])
		}
		if lastLoop {
			break
		}
		line, err := readCellFileLine(reader)
		gridLine = line
		if err == io.EOF {
			lastLoop = true
		} else {
			if err != nil {
				return nil, err
			}
		}
	}
	g := gr.readGol
	g.Init(rows, cols, rowsLimitation, colsLimitation, generation, neighborhoodType)
	for rowI := 0; rowI < rows; rowI++ {
		for colI := 0; colI < cols; colI++ {
			g.Set(rowI, colI, cells[rowI*cols+colI])
		}
	}
	return g, nil
}

func readCellFileLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		return line, err
	} else if err != nil {
		return "", err
	}
	return line[:len(line)-1], nil
}
