http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/archive/cc0b6dd26fc79f48482ebf1792009c4b6a0d3242.tar.gz",
    strip_prefix = "rules_go-cc0b6dd26fc79f48482ebf1792009c4b6a0d3242",
    sha256 = "548c975c3281db38dacacd6cb490d2e15081607bf93b07829624b1f0fdaf3303",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/archive/6f36218e9eb6c9e9b4974fe0b20d893cb0b6c53c.tar.gz",
    strip_prefix = "bazel-gazelle-6f36218e9eb6c9e9b4974fe0b20d893cb0b6c53c",
    sha256 = "7b4f3b09c160d4ee8e9d19371a7ad782c3336d9fb3576d4e9f933b0343757c26",
)

# TODO: pin a release once they have one - note that we're using a fork for stability anyway
http_archive(
    name = "com_github_scele_rules_go_dep",
    urls = ["https://github.com/scele/rules_go_dep/archive/1116a26789f5013fe137c56024dd849625c76c2b.tar.gz"],
    strip_prefix = "rules_go_dep-1116a26789f5013fe137c56024dd849625c76c2b",
    sha256 = "0c982ad30defa3e38057693bbba15d4622da16244a571e1e27f53384fdc7d8d2",
)

# Set up the go toolchain
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()
go_register_toolchains("1.9.4")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

# Dependencies are managed via dep's Gopkg.lock and converted via rules_go_dep
load("@com_github_scele_rules_go_dep//dep:dep.bzl", "dep_import")
dep_import(
    name = "godeps",
    prefix = "github.com/astratto/test-case-gazelle",
    gopkg_lock = "//:Gopkg.lock",
)
load("@godeps//:Gopkg.bzl", "go_deps")
go_deps()
