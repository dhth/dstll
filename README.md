# layitout

✨ Overview
---

`layitout` gives you a high level overview of various constructs in your code
files.

<p align="center">
  <img src="https://tools.dhruvs.space/images/layitout/layitout.gif" alt="Usage" />
</p>

Languages supported:

- go
- scala2
- python
- more to come

Motivation
---

Sometimes, you want to quickly understand how a project is organized. It could
be a new repo you're working on or a specific part of a project you're
unfamiliar with. When given a list of files you're curious about, `layitout`
shows you a list of signatures representing different constructs found in those
files, such as classes, functions, objects, etc.

💾 Installation
---

**go**:

```sh
go install github.com/dhth/layitout@latest
```

⚡️ Usage
---

`layitout` can be run in both CLI and TUI mode. The former is the default.

### CLI Mode

In CLI mode, `layitout` accepts a list of file paths from `stdin`.

```bash
git ls-files | layitout
# or
git ls-files | layitout -plain=true
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

### TUI Mode

In TUI mode, `layitout` allows you to manually query for constructs with the
help of a file browser.

```bash
layitout -mode=tui
```

TODO
---

- go
    - [x] Query methods
- scala
    - [x] Query Objects
- python
    - [ ] Query classes
- [ ] JS
- [ ] TS

Screenshots
---

<p align="center">
  <img src="https://tools.dhruvs.space/images/layitout/layitout-1.png" alt="Usage" />
</p>

<p align="center">
  <img src="https://tools.dhruvs.space/images/layitout/layitout-2.png" alt="Usage" />
</p>

<p align="center">
  <img src="https://tools.dhruvs.space/images/layitout/layitout-3.png" alt="Usage" />
</p>


Examples
---

Running `layitout` in the [scala][1] repo gives the following output:

```
$ git ls-files src/compiler/scala/tools/tasty | head -n 3 | layitout

👉 src/compiler/scala/tools/tasty/AttributeUnpickler.scala

object AttributeUnpickler

def attributes(reader: TastyReader): Attributes

................................................................................

👉 src/compiler/scala/tools/tasty/Attributes.scala

object Attributes

private class ConcreteAttributes(val isJava: Boolean) extends Attributes

................................................................................

👉 src/compiler/scala/tools/tasty/ErasedTypeRef.scala

object ErasedTypeRef

class ErasedTypeRef(qualifiedName: TypeName, arrayDims: Int)

def signature: String

def encode: ErasedTypeRef

def apply(tname: TastyName): ErasedTypeRef

def name(qual: TastyName, tname: SimpleName, isModule: Boolean)

def specialised(qual: TastyName, terminal: String, isModule: Boolean, arrayDims: Int = 0): ErasedTypeRef

................................................................................
```

More examples can be found [here](./examples).

[1]: https://github.com/scala/scala
