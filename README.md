# miniwebd: mini web server

Zero configuration web server for static contents.

[![GitHub release](http://img.shields.io/github/release/takahashim/miniwebd.svg)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)][license]

[release]: https://github.com/takahashim/miniwebd/releases
[license]: https://github.com/takahashim/miniwebd/blob/master/LICENSE

## Usage

1. copy content directory as "html" or "htdocs" or "content"
2. copy miniwebd (or miniwebd.exe in Windows) in the same directory
3. execute miniwebd

## Note

* The value of `DocumentRoot` and `Port` is pre-defined.
    * `DocumentRoot`: "html" or "htdocs" or "content"
    * `Port`: 22222
* Dot-files and directories cannot be accessed.
