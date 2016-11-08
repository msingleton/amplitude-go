# Go Amplitude [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/msingleton/amplitude-go) [![Build Status](https://travis-ci.org/msingleton/amplitude-go.svg?branch=master)](https://travis-ci.org/msingleton/amplitude-go)

## Summary

A Go library for [Amplitude](amplitude.com)

## Installation

```sh
go get github.com/msingleton/amplitude-go
```

## Usage

```go
import "github.com/msingleton/amplitude-go"

client := amplitude.New("amplitude-api-key")
client.Event(amplitude.Event{
	EventType:	"joined",
})

```
