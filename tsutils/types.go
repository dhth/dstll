package tsutils

const (
	nodeTypeIdentifier     = "identifier"
	nodeTypeModifiers      = "modifiers"
	nodeTypeTypeParameters = "type_parameters"
	nodeTypeParameters     = "parameters"
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
