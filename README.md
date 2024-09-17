# dstll

✨ Overview
---

`dstll` *(short for "distill")* gives you a high level overview of various
constructs in your code files.

<p align="center">
  <img src="https://tools.dhruvs.space/images/dstll/dstll-1.png" alt="Usage" />
</p>

Languages supported:

- go
- python
- rust
- scala 2
- more to come

Motivation
---

Sometimes, you want to quickly understand how a project is organized. It could
be a new repo you're working on or a specific part of a project you're
unfamiliar with. When given a list of files you're curious about, `dstll`
shows you a list of signatures representing different constructs found in those
files, such as classes, functions, objects, etc.

💾 Installation
---

**go**:

```sh
go install github.com/dhth/dstll@latest
```

⚡️ Usage
---

```bash
# print findings to stdout
dstll [PATH ...]

# write findings to a directory
dstll write [PATH ...] -o /var/tmp/findings

# serve findings via a web server
dstll serve [PATH ...] -o /var/tmp/findings

# open TUI
dstll tui
```

<p align="center">
  <img src="https://tools.dhruvs.space/images/dstll/dstll-2.png" alt="Usage" />
</p>

<p align="center">
  <img src="https://tools.dhruvs.space/images/dstll/dstll-3.png" alt="Usage" />
</p>

<p align="center">
  <img src="https://tools.dhruvs.space/images/dstll/dstll-4.png" alt="Usage" />
</p>

🛠️ Configuration
---

Create a configuration file that looks like the following. By default,
`dstll` will look for this file at `~/.config/dstll/dstll.yml`.

```toml
view-file-command = ["your", "command"]
# for example, ["bat", "--style", "plain", "--paging", "always"]
# will run 'bat --style plain --paging always <file-path>'
```

Examples
---

Running `dstll` in the [scala][1] repo gives the following output:

```
$ dstll $(git ls-files src/compiler/scala/tools/tasty | head -n 3)

-> src/compiler/scala/tools/tasty/ErasedTypeRef.scala

object ErasedTypeRef

class ErasedTypeRef(qualifiedName: TypeName, arrayDims: Int)

def apply(tname: TastyName): ErasedTypeRef

def name(qual: TastyName, tname: SimpleName, isModule: Boolean)

def specialised(qual: TastyName, terminal: String, isModule: Boolean, arrayDims: Int = 0): ErasedTypeRef

................................................................................

-> src/compiler/scala/tools/tasty/Attributes.scala

object Attributes

private class ConcreteAttributes(val isJava: Boolean) extends Attributes

................................................................................

-> src/compiler/scala/tools/tasty/AttributeUnpickler.scala

object AttributeUnpickler

def attributes(reader: TastyReader): Attributes
```

More examples can be found [here](./examples).

[1]: https://github.com/scala/scala
