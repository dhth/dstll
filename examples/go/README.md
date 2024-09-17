# dstll go code

Running `dstll` in the [go][1] repo gives the following output:

```
$ dstll $(git ls-files src/io/**/*.go | grep -v '_test.go' | head -n 3) -p

-> src/io/fs/glob.go

type GlobFS interface {
	FS

	// Glob returns the names of all files matching pattern,
	// providing an implementation of the top-level
	// Glob function.
	Glob(pattern string) ([]string, error)
}

func Glob(fsys FS, pattern string) (matches []string, err error)

func globWithLimit(fsys FS, pattern string, depth int) (matches []string, err error)

func cleanGlobPath(path string) string

func glob(fs FS, dir, pattern string, matches []string) (m []string, e error)

func hasMeta(path string) bool

................................................................................

-> src/io/fs/format.go

func FormatFileInfo(info FileInfo) string

func FormatDirEntry(dir DirEntry) string

................................................................................

-> src/io/fs/fs.go

type FS interface {
	// Open opens the named file.
	//
	// When Open returns an error, it should be of type *PathError
	// with the Op field set to "open", the Path field set to name,
	// and the Err field describing the problem.
	//
	// Open should reject attempts to open names that do not satisfy
	// ValidPath(name), returning a *PathError with Err set to
	// ErrInvalid or ErrNotExist.
	Open(name string) (File, error)
}

type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

type DirEntry interface {
	// Name returns the name of the file (or subdirectory) described by the entry.
	// This name is only the final element of the path (the base name), not the entire path.
	// For example, Name would return "hello.go" not "home/gopher/hello.go".
	Name() string

	// IsDir reports whether the entry describes a directory.
	IsDir() bool

	// Type returns the type bits for the entry.
	// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
	Type() FileMode

	// Info returns the FileInfo for the file or subdirectory described by the entry.
	// The returned FileInfo may be from the time of the original directory read
	// or from the time of the call to Info. If the file has been removed or renamed
	// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
	// If the entry denotes a symbolic link, Info reports the information about the link itself,
	// not the link's target.
	Info() (FileInfo, error)
}

type ReadDirFile interface {
	File

	// ReadDir reads the contents of the directory and returns
	// a slice of up to n DirEntry values in directory order.
	// Subsequent calls on the same file will yield further DirEntry values.
	//
	// If n > 0, ReadDir returns at most n DirEntry structures.
	// In this case, if ReadDir returns an empty slice, it will return
	// a non-nil error explaining why.
	// At the end of a directory, the error is io.EOF.
	// (ReadDir must return io.EOF itself, not an error wrapping io.EOF.)
	//
	// If n <= 0, ReadDir returns all the DirEntry values from the directory
	// in a single slice. In this case, if ReadDir succeeds (reads all the way
	// to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory,
	// ReadDir returns the DirEntry list read until that point and a non-nil error.
	ReadDir(n int) ([]DirEntry, error)
}

type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() any           // underlying data source (can return nil)
}

type FileMode uint32

type PathError struct {
	Op   string
	Path string
	Err  error
}

func ValidPath(name string) bool

func errInvalid() error

func errPermission() error

func errExist() error

func errNotExist() error

func errClosed() error

func (m FileMode) String() string

func (m FileMode) IsDir() bool

func (m FileMode) IsRegular() bool

func (m FileMode) Perm() FileMode

func (m FileMode) Type() FileMode

func (e *PathError) Error() string

func (e *PathError) Unwrap() error

func (e *PathError) Timeout() bool
```

[1]: https://https://github.com/golang/go.com/pallets/flask
