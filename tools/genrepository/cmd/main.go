package main

import (
	"flag"
	"litshop/tools/genrepository"
	"os"
	"strings"
)

type conf struct {
	dir     string
	pkg     []string
	structs []string
}

var config = conf{}

func main() {
	parseFlags()

	parser := genrepository.NewParser(config.dir)
	generator := genrepository.NewGenerator(config.dir).SetImportPkg(config.pkg)
	generator.ParseAst(parser, config.structs).Generate().Format().Flush()
}

func parseFlags() {
	var paramDir, paramStructs, paramPkgs string
	flag.StringVar(&paramStructs, "structs", "", "模型Struct，以逗号分割多个")
	flag.StringVar(&paramDir, "dir", "", "模型文件夹")
	flag.StringVar(&paramPkgs, "pkgs", "", "导入的依赖包")
	flag.Parse()

	if paramStructs == "" || paramPkgs == "" || paramDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	config.dir = paramDir
	config.structs = strings.Split(paramStructs, ",")
	config.pkg = strings.Split(paramPkgs, ",")
}
