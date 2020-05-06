# GoPlag

## Dependency

 - Go 1.13+
 - Java 1.8+
 - ANTLR v4.8+

## Install

### Install dependencies

```shell
# apt install -y curl default-jdk
# curl -O https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
# tar xvf go1.14.2.linux-amd64.tar.gz
# chown -R root:root ./go
# mv go /usr/local
```

### Setting Go paths

Filename: `~/.profile`

```shell
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

```shell
# source ~/.profile
```

> For Debian 10+

### Download ANTLR

```shell
# cd /usr/local/lib
# curl -O https://www.antlr.org/download/antlr-4.8-complete.jar
```

### Create execute file for ANTLR

Filename `~/.profile`

```shell
...
...
export CLASSPATH=".:/usr/local/lib/antlr-4.8-complete.jar:$CLASSPATH"
```

Filename: `/usr/bin/antlr4`

```shell
#!/bin/bash
java -jar /usr/local/lib/antlr-4.8-complete.jar
```

### Install ANTLR runtime for Go

```shell
# go get -u github.com/antlr/antlr4/runtime/Go/antlr
```

### Clone GoPlag and build

```shell
# cd $GOPATH
# git clone {git repo of goplag}
# cd goplag
# go build main.go
# mv main /usr/bin/goplag
```

## Usgae

```shell
# goplag [-base {base file}] -a {file1} -b {file2}
```

## Reference

 - [Winnowing: Local Algorithms for Document Fingerprinting](http://theory.stanford.edu/~aiken/publications/papers/sigmod03.pdf)

## License

[MIT](https://opensource.org/licenses/MIT)
