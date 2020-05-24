package base

// GolInterface : minimal Gol interface.
type GolInterface interface {
	// Sort of an initializer
	Init(name, description, rules, gridType, rowsLimitation, colsLimitation string, rows, cols, generation, neighborhoodType int)
	// Dummy-property methods
	Name() string
	Description() string
	// Row and cols related methods
	Rows() int
	Cols() int
	LimitRows() bool
	SetLimitRows(limitRows bool)
	LimitCols() bool
	SetLimitCols(limitCols bool)
	// Cloning
	Clone() GolInterface
	// Indexing methods
	Get(i int, j int) int
	Set(i int, j int, value int)
	SetAll(value int)
	// Rules methods
	Rules() string
	SetRules(rules string)
	// Debug methods
	DbgStdout()
	// Generation methods
	Generation() int
	SetGeneration(generation int)
	// Equality methods
	GridEquals(g GolInterface, mode string) bool
	Equals(g GolInterface) bool
	EqualsError(g GolInterface) error
	// NeighborhoodType-related methods
	NeighborhoodTypeString() string
	SetNeighborhoodType(neighborhoodType int)
	SetNeighborhoodTypeString(neighborhoodType string)
	// Concurrency-related methods
	GetProcesses() int
	SetProcesses(processes int)
	// Compute the next generation for a game of life instance
	NextGeneration() GolInterface
}
