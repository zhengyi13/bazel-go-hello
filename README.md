# bazel-go-hello

This example project demonstrates a very simple, barebones golang "project"
meant to be built from scratch, with nothing more initially installed than Bazel
itself.

The idea here is that given a properly setup repo (i.e. with properly written
BUILD and WORKSPACE files), one need only:

1. [Install bazel](https://docs.bazel.build/versions/master/install.html)
2. `git clone git@github.com:zhengyi13/bazel-go-hello`
3. `cd bazel-go-hello`
4. `blaze run :bazel-go-hello`

... and the result should be that bazel automatically downloads and installs
(locally, hermetically) an appropriate copy of go (currently 1.9.2), and then
builds the project in its entirety using that copy of go, in an entirely
reproducible way.

By extension, this should be possible for arbitrarily complex projects.

Gazelle support is also included, so that as project complexity grows, BUILD
files can be easily kept up to date.
