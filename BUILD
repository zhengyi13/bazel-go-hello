load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/zhengyi13/bazel-go-hello
gazelle(name = "gazelle")

go_binary(
    name = "hello",
    embed = [":bazel-go-hello_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "bazel-go-hello_lib",
    srcs = ["hello.go"],
    importpath = "github.com/zhengyi13/bazel-go-hello",
    visibility = ["//visibility:private"],
    deps = ["@com_github_golang_glog//:glog"],
)
