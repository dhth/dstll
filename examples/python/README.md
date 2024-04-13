# dstll python code

Running `dstll` in the [flask][1] repo gives the following output:

```
$ git ls-files src/flask/**/*.py | grep -v '__init__.py' | head -n 3 | dstll -plain=true

-> src/flask/app.py

def _make_timedelta(value: timedelta | int | None) -> timedelta | None

def __init__(
        self,
        import_name: str,
        static_url_path: str | None = None,
        static_folder: str | os.PathLike[str] | None = "static",
        static_host: str | None = None,
        host_matching: bool = False,
        subdomain_matching: bool = False,
        template_folder: str | os.PathLike[str] | None = "templates",
        instance_path: str | None = None,
        instance_relative_config: bool = False,
        root_path: str | None = None,
    ) -> timedelta | None

def get_send_file_max_age(self, filename: str | None) -> int | None

def send_static_file(self, filename: str) -> Response

def open_resource(self, resource: str, mode: str = "rb") -> t.IO[t.AnyStr]

def open_instance_resource(self, resource: str, mode: str = "rb") -> t.IO[t.AnyStr]

def create_jinja_environment(self) -> Environment

def create_url_adapter(self, request: Request | None) -> MapAdapter | None

def raise_routing_exception(self, request: Request) -> t.NoReturn

def update_template_context(self, context: dict[str, t.Any]) -> None

def make_shell_context(self) -> dict[str, t.Any]

def run(
        self,
        host: str | None = None,
        port: int | None = None,
        debug: bool | None = None,
        load_dotenv: bool = True,
        **options: t.Any,
    ) -> None

def test_client(self, use_cookies: bool = True, **kwargs: t.Any) -> FlaskClient

def test_cli_runner(self, **kwargs: t.Any) -> FlaskCliRunner

def handle_http_exception(
        self, e: HTTPException
    ) -> HTTPException | ft.ResponseReturnValue

def handle_user_exception(
        self, e: Exception
    ) -> HTTPException | ft.ResponseReturnValue

def handle_exception(self, e: Exception) -> Response

def log_exception(
        self,
        exc_info: (tuple[type, BaseException, TracebackType] | tuple[None, None, None]),
    ) -> None

def dispatch_request(self) -> ft.ResponseReturnValue

def full_dispatch_request(self) -> Response

def finalize_request(
        self,
        rv: ft.ResponseReturnValue | HTTPException,
        from_error_handler: bool = False,
    ) -> Response

def make_default_options_response(self) -> Response

def ensure_sync(self, func: t.Callable[..., t.Any]) -> t.Callable[..., t.Any]

def async_to_sync(
        self, func: t.Callable[..., t.Coroutine[t.Any, t.Any, t.Any]]
    ) -> t.Callable[..., t.Any]

def url_for(
        self,
        /,
        endpoint: str,
        *,
        _anchor: str | None = None,
        _method: str | None = None,
        _scheme: str | None = None,
        _external: bool | None = None,
        **values: t.Any,
    ) -> str

def make_response(self, rv: ft.ResponseReturnValue) -> Response

def preprocess_request(self) -> ft.ResponseReturnValue | None

def process_response(self, response: Response) -> Response

def do_teardown_request(
        self,
        exc: BaseException | None = _sentinel,  # type: ignore[assignment]
    ) -> None

def do_teardown_appcontext(
        self,
        exc: BaseException | None = _sentinel,  # type: ignore[assignment]
    ) -> None

def app_context(self) -> AppContext

def request_context(self, environ: WSGIEnvironment) -> RequestContext

def test_request_context(self, *args: t.Any, **kwargs: t.Any) -> RequestContext

def wsgi_app(
        self, environ: WSGIEnvironment, start_response: StartResponse
    ) -> cabc.Iterable[bytes]

def __call__(
        self, environ: WSGIEnvironment, start_response: StartResponse
    ) -> cabc.Iterable[bytes]

................................................................................

-> src/flask/blueprints.py

def get_send_file_max_age(self, filename: str | None) -> int | None

def send_static_file(self, filename: str) -> Response

def open_resource(self, resource: str, mode: str = "rb") -> t.IO[t.AnyStr]

................................................................................
```

[1]: https://github.com/pallets/flask
