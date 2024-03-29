# bazel-go-hello

This example project demonstrates a very simple, barebones golang "project"
meant to be built from scratch, with nothing more initially installed than Bazel
itself.

The idea here is that given a properly setup repo (i.e. with properly written
BUILD and WORKSPACE files), one need only:

1. [Install bazel](https://docs.bazel.build/versions/master/install.html)
2. `git clone git@github.com:zhengyi13/bazel-go-hello`
3. `cd bazel-go-hello`
4. `blaze run :hello`

... and the result should be that bazel automatically downloads and installs
(locally, hermetically) an appropriate copy of go (currently 1.19.3), and then
builds the project in its entirety using that copy of go, in an entirely
reproducible way.

By extension, this should be possible for arbitrarily complex projects.

Gazelle support is also included, so that as project complexity grows, BUILD
files can be easily kept up to date.

## Updating WORKSPACE

As a project grows, you may add new (external) modules/dependencies. Typically, you'll probably do something like

```shell
go run hello.go
```

... and if you've added a hot new module to your imports, you'll find it's automatically added to your go.mod before being built and run.

... and then you'll discover that `bazel run :hello` is now failing, and complaining about something like your new module not being there.

When this happens, you can:

```shell
bazel run //:gazelle -- update_repos github.com/user/project
```

... and a new repository will be added to your WORKSPACE.
