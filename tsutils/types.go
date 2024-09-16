package tsutils

const (
	nodeTypeIdentifier = "identifier"
	nodeTypeModifiers  = "modifiers"
	nodeTypeParameters = "parameters"
)

type Result struct {
	FPath   string
	Results []string
	Err     error
}
type FileType uint

const (
	FTNone FileType = iota
	FTGo
	FTPython
	FTRust
	FTScala
)
