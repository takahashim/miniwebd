# miniwebd: mini web server

Zero configuration web server for static contents.

[![GitHub release](http://img.shields.io/github/release/takahashim/miniwebd.svg)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)][license]
[![CircleCI](https://circleci.com/gh/takahashim/miniwebd.svg?style=svg)][circleci]

[release]: https://github.com/takahashim/miniwebd/releases
[license]: https://github.com/takahashim/miniwebd/blob/master/LICENSE
[circleci]: https://circleci.com/gh/takahashim/miniwebd

## Usage

1. Rename content directory as `html` or `htdocs` or `content`
2. Copy the file `miniwebd` (or `miniwebd.exe` in Windows) in the same directory of content
3. Execute `iniwebd`

### Note

* The value of `DocumentRoot` and `Port` is pre-defined.
    * `DocumentRoot`: `html` or `htdocs` or `content`
    * `Port`: 22222
* Dot-files and directories cannot be accessed.
