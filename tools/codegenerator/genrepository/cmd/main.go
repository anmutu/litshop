package main

import (
	"flag"
	"github.com/google/martian/log"
	"litshop/tools/codegenerator/genrepository"
	"os"
	"strings"
)

type conf struct {
	dir     string
	pkg     []genrepository.ImportPkg
	structs []string
}

var config = conf{}

func main() {
	parseFlags()

	parser := genrepository.NewParser(config.dir)
	generator := genrepository.NewGenerator(config.dir).SetImportPkg(config.pkg)
	err := generator.ParseAst(parser, config.structs).Generate().Format().Flush()
	if err != nil {
		log.Errorf("flush err : %#v \n", err)
	}
}

func parseFlags() {
	var paramDir, paramStructs, paramPkgs string
	flag.StringVar(&paramStructs, "structs", "", "模型Struct，以逗号分割多个")
	flag.StringVar(&paramDir, "dir", "", "模型文件夹")
	flag.StringVar(&paramPkgs, "pkgs", "", "导入的依赖包")
	flag.Parse()

	if paramStructs == "" || paramDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	config.dir = paramDir
	config.structs = strings.Split(paramStructs, ",")
	if len(config.structs) > 0 {
		for _, p := range strings.Split(paramPkgs, ",") {
			config.pkg = append(config.pkg, genrepository.ImportPkg{
				Pkg: p,
			})
		}
	}
}
