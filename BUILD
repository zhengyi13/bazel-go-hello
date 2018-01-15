load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/example/project",
)

go_library(
    name = "go_default_library",
    srcs = ["hello.go"],
    importpath = "github.com/example/project",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "project",
    embed = [":go_default_library"],
    importpath = "github.com/example/project",
    visibility = ["//visibility:public"],
)
