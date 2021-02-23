package main

import (
	"fmt"
	"log"
	"os"
)

//FooReader defines and ioReader to read from stdin
type FooReader struct{}

func (fooReader *FooReader)