# bazel-go-hello

This example project demonstrates a very simple, barebones golang "project"
meant to be built from scratch, with nothing more initially installed than Bazel
itself.

The idea here is that given a properly setup repo (i.e. with properly written
BUILD and WORKSPACE files), one need only:

1. Install bazel
2. `git clone this://repo`
3. `cd this_repo`
4. `blaze build :all`

... and the result should be that bazel automatically downloads and installs
(locally, hermetically) an appropriate copy of go (currently 1.8.1), and then
builds the project in its entirety using that copy of go, in an entirely
reproducible way.
