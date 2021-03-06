## go-auto-commenter
A Go tool to add comments automatically to all the exported methods.

[![Build Status](https://www.travis-ci.com/diptomondal007/go-auto-commenter.svg?branch=main)](https://www.travis-ci.com/diptomondal007/go-auto-commenter)
[![Coverage Status](https://coveralls.io/repos/github/diptomondal007/go-auto-commenter/badge.svg)](https://coveralls.io/github/diptomondal007/go-auto-commenter)
[![Go Report Card](https://goreportcard.com/badge/github.com/diptomondal007/go-auto-commenter)](https://goreportcard.com/report/github.com/diptomondal007/go-auto-commenter)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

#### Installation
```shell
1. Clone this repo
2. cd to the cloned directory
3. Type ./install in terminal and hit enter
```

#### Uninstallation
```shell
1. Clone this repo
2. cd to the cloned directory
3. Type ./uninstall in terminal and hit enter
```

#### To auto comment all the files in current directory
```shell
autocomment .
```

#### To auto comment all the files of a directory
```shell
autocomment [DIRNAME]
```

#### To auto comment files
```shell
autocomment [FILENAME]...
```

#### Example
##### before ->
```go
func AB() {

}
```
##### after ->
```go
// AB ...
func AB(){
}
```

### License
go-auto-commenter is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/diptomondal007/go-auto-commenter/blob/main/LICENSE)
