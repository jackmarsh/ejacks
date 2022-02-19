package(default_visibility = ["PUBLIC"])

go_module(
    name = "termbox",
    module = "github.com/gdamore/tcell",
    version = "v1",
    deps = [
    #  ":text",
      ":runewidth",
      ":colorful",
    ],
)

# go_module(
#     name = "text",
#     module = "golang.org/x/text",
#     version = "v0.3.7",
#     deps = [
#     ],
# )

go_module(
    name = "colorful",
    module = "github.com/lucasb-eyer/go-colorful",
    version = "v1.2.0",
    deps = [],
)

go_module(
    name = "runewidth",
    module = "github.com/mattn/go-runewidth",
    version = "v0.0.13",
    deps = [":uniseg"],
)

go_module(
    name = "uniseg",
    module = "github.com/rivo/uniseg",
    version = "v0.1.0",
    deps = [],
)
