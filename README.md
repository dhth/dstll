# layitout

✨ Overview
---

`layitout` gives you a high level overview of various constructs in your code
files.

<img src="./layitout.png" alt="layitout" />

Languages supported:

- go
- scala
- python
- more to come

⚡️ Usage
---

`layitout` accepts a list of file paths from `stdin`.

```bash
git ls-files | layitout
# or
find . -name '*.go' | layitout
# or
fd . --extension=scala | head -n 4 | layitout
# or
ls -1 | layitout
# or
cat <<EOF | layitout
file1.py
dir/file2.py
EOF
```

TODO
---

- go
    - [ ] Query methods
- scala
    - [ ] Query Objects
- python
    - [ ] Query classes
- [ ] JS
- [ ] TS

Examples
---

Running `layitout` in the [go][1] repo gives the following output:

```
$ git ls-files src/strings/**/*.go | grep -v '_test.go' | head -n 5 | layitout


👉 src/strings/builder.go

func noescape(p unsafe.Pointer) unsafe.Pointer

................................................................................

👉 src/strings/clone.go

func Clone(s string) string

................................................................................

👉 src/strings/compare.go

func Compare(a, b string) int

................................................................................

👉 src/strings/reader.go

func NewReader(s string) *Reader

................................................................................

👉 src/strings/replace.go

func NewReplacer(oldnew ...string) *Replacer

func makeGenericReplacer(oldnew []string) *genericReplacer

func getStringWriter(w io.Writer) io.StringWriter

func makeSingleStringReplacer(pattern string, value string) *singleStringReplacer

................................................................................
```

[1]: https://github.com/golang/go
