package main

type Result struct {
	fPath   string
	results []string
	err     error
}

type FileType uint

const (
	FTNone FileType = iota
	FTScala
	FTGo
	FTPython
)
