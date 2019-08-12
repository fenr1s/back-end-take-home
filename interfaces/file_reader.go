package interfaces

//FileReader contract
type FileReader interface {
	ReadFromFile(path string) (lines [][]string, err error)
}
