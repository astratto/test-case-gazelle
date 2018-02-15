# Test repo for [bazel-gazelle's issue 117](https://github.com/bazelbuild/bazel-gazelle/issues/117)

Test repository to try and reproduce the issue described in https://github.com/bazelbuild/bazel-gazelle/issues/117

## Steps to reproduce

1. `dep ensure` to generate the `Gopkg.lock` file.

2. `bazel run :gazelle` generates a `testapp/BUILD.bazel` file correct, but that points to vendor

```
$ cat testapp/BUILD.bazel
...
    deps = ["//vendor/github.com/BurntSushi/toml:go_default_library"],
)
```

3. `bazel run testapp` works correctly

```
$ bazel run testapp
...
2018/02/15 18:32:03 Count:  100
```

4. `rm -rf vendor && bazel run :gazelle` generates a lowercase import

```
$ cat testapp/BUILD.bazel
...
    deps = ["@com_github_burntsushi_toml//:go_default_library"],
)
```

5. `bazel run testapp` now fails

```
$ bazel run testapp
(18:32:45) ERROR: /Users/stefanotortarolo/Dev/go/src/github.com/astratto/test-case-gazelle/testapp/BUILD.bazel:3:1: no such package '@com_github_burntsushi_toml//': The repository could not be resolved and referenced by '//testapp:go_default_library'
...
```
