package genrepository

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
)

type gConfig struct {
	PkgName string
	LogName string
	Pkgs    []string
}

type Generator struct {
	buf           map[string]*bytes.Buffer
	inputFile     string
	config        gConfig
	structConfigs []structConfig
}

func (g *Generator) SetImportPkg(pkg []string) *Generator {
	g.config.Pkgs = pkg
	return g
}

func (g *Generator) ParseAst(p *Parser, structs []string) *Generator {
	for _, v := range structs {
		g.buf[v] = new(bytes.Buffer)
	}
	g.structConfigs = p.Parse()
	g.config.PkgName = p.pkg.Name
	return g
}

func (g *Generator) Generate() *Generator {
	fmt.Printf("g.config.Pkgs  %#v \n", g.config.Pkgs)
	fmt.Printf("g.config.PkgName  %#v \n", g.config.PkgName)

	if len(g.structConfigs) < 1 {
		panic(errors.New("模型未设置"))
	}

	for i := 0; i < len(g.structConfigs); i++ {
		g.structConfigs[i].config = g.config
	}

	for _, v := range g.structConfigs {
		if _, ok := g.buf[defaultNamer(v.StructName)]; !ok {
			continue
		}

		if err := modelTemplate.Execute(g.buf[defaultNamer(v.StructName)], v); err != nil {
			panic(err)
		}
	}

	return g
}

func (g *Generator) Format() *Generator {
	for k, _ := range g.buf {
		formatedOutput, err := format.Source(g.buf[k].Bytes())
		if err != nil {
			panic(err)
		}
		g.buf[k] = bytes.NewBuffer(formatedOutput)
	}
	return g
}

func (g *Generator) Flush() error {
	for k, _ := range g.buf {
		filename := g.inputFile + "/gen_" + strings.ToLower(k) + ".go"
		if err := ioutil.WriteFile(filename, g.buf[k].Bytes(), 0777); err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}

func NewGenerator(input string) *Generator {
	return &Generator{
		buf:       map[string]*bytes.Buffer{},
		inputFile: input,
	}
}
