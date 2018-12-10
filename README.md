# Golang Webview Echo

a skeleton app to build a desktop application using webview and echo in Golang.

- Back-End :
    - Echo : [_Go_] web framework
    - go.rice : [_Go_] embed data in binary
    
- Front-End :
    - AngularJS or `¯\_(ツ)_/¯` [_JS_]
    - webview

### Requirements

* `go` :

    [site officiel Go](https://golang.org/dl/)

* `go.rice` :

    ```bash
    $ go get github.com/GeertJohan/go.rice
    $ go get github.com/GeertJohan/go.rice/rice
    ```

* `dep` :

    [github officiel de Dep](https://github.com/golang/dep)

### Issues

- scrolling is really slow and buggy
- current version (10/12/2018) does not use WkWebview
- had to add Info.plist to avoid blurry content