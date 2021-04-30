[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kindlyops/terp)

# eBPF experimental diagnostic interpreter

## terp - eBPF pry bar

An interpreter with diagnostic functions for learning about and validating
eBPF IR.

## installation for homebrew (MacOS/Linux)

    brew install kindlyops/tap/terp

once installed, you can upgrade to a newer version using this command:

    brew upgrade kindlyops/tap/terp

## installation for scoop (Windows Powershell)

To enable the bucket for your scoop installation

    scoop bucket add kindlyops https://github.com/kindlyops/kindlyops-scoop

To install deleterious

    scoop install terp

once installed, you can upgrade to a newer version using this command:

    scoop status
    scoop update terp

## installation from source

    go get github.com/kindlyops/terp
    terp help

## Developer instructions

Want to help add features or fix bugs? Awesome! terp is built using bazel.

    `brew install bazelisk`
    grab the source code from github
    `bazel run :terp-darwin` to compile and run the current version on macOS

### Testing release process

To run goreleaser locally to test changes to the release process configuration:

    goreleaser release --snapshot --skip-publish --rm-dist
