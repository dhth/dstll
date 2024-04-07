# layitout with go

Running `layitout` in the [go][1] repo gives the following output:

```
$ git ls-files src/io/**/*.go | grep -v '_test.go' | head -n 3 | layitout


👉 src/io/fs/format.go

func FormatFileInfo(info FileInfo) string

func FormatDirEntry(dir DirEntry) string

................................................................................

👉 src/io/fs/fs.go

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

................................................................................

👉 src/io/fs/glob.go

func Glob(matches []string, err error)

func globWithLimit(matches []string, err error)

func cleanGlobPath(path string) string

func glob(m []string, e error) string

func hasMeta(path string) bool

................................................................................
```

[1]: https://https://github.com/golang/go.com/pallets/flask
