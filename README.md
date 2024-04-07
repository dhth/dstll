# layitout

✨ Overview
---

`layitout` gives you a high level overview of various constructs in your code
files.

<img src="./layitout.png" alt="layitout" />

Languages supported:

- go
- scala2
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
    - [x] Query methods
- scala
    - [x] Query Objects
- python
    - [ ] Query classes
- [ ] JS
- [ ] TS

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
