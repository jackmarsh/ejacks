package(default_visibility = ["PUBLIC"])

go_module(
    name = "tcell",
    module = "github.com/gdamore/tcell/v2",
    version = "v2.5.4",
    install = ["..."],
    deps = [
        ":x_text",
        ":x_term",
        ":x_sys",
        ":runewidth",
        ":colorful",
        ":encoding",
    ],
)

go_module(
    name = "x_text",
    install = ["..."],
    module = "golang.org/x/text",
    strip = [
        "cmd",
        "message/pipeline",
    ],
    version = "v0.3.6",
)

go_module(
    name = "x_sys",
    install = ["..."],
    module = "golang.org/x/sys",
    version = "v0.0.0-20210415045647-66c3f260301c",
)

go_module(
    name = "x_term",
    module = "golang.org/x/term",
    version = "v0.0.0-20210615171337-6886f2dfbf5b",
    deps = [":x_sys"],
)

go_module(
    name = "colorful",
    module = "github.com/lucasb-eyer/go-colorful",
    version = "v1.2.0",
    deps = [],
)

go_module(
    name = "runewidth",
    module = "github.com/mattn/go-runewidth",
    version = "v0.0.14",
    deps = [":uniseg"],
)

go_module(
    name = "uniseg",
    module = "github.com/rivo/uniseg",
    version = "v0.1.0",
    deps = [],
)

go_module(
    name = "encoding",
    module = "github.com/gdamore/encoding",
    version = "v1.0.0",
    deps = [":x_text"],
)
