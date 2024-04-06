package tsutils

type Result struct {
	FPath   string
	Results []string
	Err     error
}
type FileType uint

const (
	FTNone FileType = iota
	FTScala
	FTGo
	FTPython
)
