package main

import (
	"github.com/opslevel/opslevel-go/v2023"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(opslevel.StructTagAnalyzer)
}
