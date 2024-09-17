# dstll rust code

Running `dstll` in the [ripgrep][1] repo gives the following output:

```
$ dstll $(git ls-files '**.rs' | head -n 3 ) -p

-> build.rs

fn main()

fn set_windows_exe_options()

fn set_git_revision_hash()

................................................................................

-> crates/cli/src/escape.rs

pub fn escape(bytes: &[u8]) -> String

pub fn escape_os(string: &OsStr) -> String

pub fn unescape(s: &str) -> Vec<u8>

pub fn unescape_os(string: &OsStr) -> Vec<u8>

fn b(bytes: &'static [u8]) -> Vec<u8>

fn empty()

fn backslash()

fn nul()

fn nl()

fn tab()

fn carriage()

fn nothing_simple()

fn nothing_hex0()

fn nothing_hex1()

fn nothing_hex2()

fn invalid_utf8()

................................................................................

-> crates/cli/src/decompress.rs

pub struct DecompressionMatcherBuilder {
    /// The commands for each matching glob.
    commands: Vec<DecompressionCommand>,
    /// Whether to include the default matching rules.
    defaults: bool,
}

struct DecompressionCommand {
    /// The glob that matches this command.
    glob: String,
    /// The command or binary name.
    bin: PathBuf,
    /// The arguments to invoke with the command.
    args: Vec<OsString>,
}

pub struct DecompressionMatcher {
    /// The set of globs to match. Each glob has a corresponding entry in
    /// `commands`. When a glob matches, the corresponding command should be
    /// used to perform out-of-process decompression.
    globs: GlobSet,
    /// The commands for each matching glob.
    commands: Vec<DecompressionCommand>,
}

pub struct DecompressionReaderBuilder {
    matcher: DecompressionMatcher,
    command_builder: CommandReaderBuilder,
}

pub struct DecompressionReader {
    rdr: Result<CommandReader, File>,
}

fn default() -> DecompressionMatcherBuilder

pub fn new() -> DecompressionMatcherBuilder

pub fn build(&self) -> Result<DecompressionMatcher, CommandError>

pub fn defaults(&mut self, yes: bool) -> &mut DecompressionMatcherBuilder

pub fn associate<P, I, A>(
        &mut self,
        glob: &str,
        program: P,
        args: I,
    ) -> &mut DecompressionMatcherBuilder

pub fn try_associate<P, I, A>(
        &mut self,
        glob: &str,
        program: P,
        args: I,
    ) -> Result<&mut DecompressionMatcherBuilder, CommandError>

fn default() -> DecompressionMatcher

pub fn new() -> DecompressionMatcher

pub fn command<P: AsRef<Path>>(&self, path: P) -> Option<Command>

pub fn has_command<P: AsRef<Path>>(&self, path: P) -> bool

pub fn new() -> DecompressionReaderBuilder

pub fn build<P: AsRef<Path>>(
        &self,
        path: P,
    ) -> Result<DecompressionReader, CommandError>

pub fn matcher(
        &mut self,
        matcher: DecompressionMatcher,
    ) -> &mut DecompressionReaderBuilder

pub fn get_matcher(&self) -> &DecompressionMatcher

pub fn async_stderr(
        &mut self,
        yes: bool,
    ) -> &mut DecompressionReaderBuilder

pub fn new<P: AsRef<Path>>(
        path: P,
    ) -> Result<DecompressionReader, CommandError>

fn new_passthru(path: &Path) -> Result<DecompressionReader, CommandError>

pub fn close(&mut self) -> io::Result<()>

fn read(&mut self, buf: &mut [u8]) -> io::Result<usize>

pub fn resolve_binary<P: AsRef<Path>>(
    prog: P,
) -> Result<PathBuf, CommandError>

fn try_resolve_binary<P: AsRef<Path>>(
    prog: P,
) -> Result<PathBuf, CommandError>

fn is_exe(path: &Path) -> bool

fn default_decompression_commands() -> Vec<DecompressionCommand>

fn add(glob: &str, args: &[&str], cmds: &mut Vec<DecompressionCommand>)
```

[1]: https://github.com/BurntSushi/ripgrep
