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

type ImportPkg struct {
	Pkg string
}

type gConfig struct {
	PkgName string
	LogName string
	Pkgs    []ImportPkg
}

type Generator struct {
	buf           map[string]*bytes.Buffer
	inputFile     string
	config        gConfig
	structConfigs []structConfig
}

func (g *Generator) SetImportPkg(pkg []ImportPkg) *Generator {
	g.config.Pkgs = pkg
	return g
}

func (g *Generator) ParseAst(p *Parser, structs []string) *Generator {
	for _, v := range structs {
		g.buf[defaultNamer(v)] = new(bytes.Buffer)
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
		g.structConfigs[i].PkgName = g.config.PkgName
		g.structConfigs[i].LogName = g.config.LogName
		g.structConfigs[i].Pkgs = g.config.Pkgs
	}

	g.buf["common"] = &bytes.Buffer{}
	if err := commonTemplate.Execute(g.buf["common"], g.config); err != nil {
		panic(err)
	}

	for _, v := range g.structConfigs {
		_, ok := g.buf[defaultNamer(v.StructName)]
		//fmt.Printf("ok %#v \n", ok)
		if !ok {
			continue
		}

		if err := modelTemplate.Execute(g.buf[defaultNamer(v.StructName)], v); err != nil {
			panic(err)
		}

		fmt.Printf("v %#v %#v \n", v, g.buf[defaultNamer(v.StructName)].String())
	}

	return g
}

func (g *Generator) Format() *Generator {
	for k, _ := range g.buf {
		formatedOutput, err := format.Source(g.buf[k].Bytes())
		if err != nil {
			fmt.Printf("err k %#v %s \n", k, string(g.buf[k].Bytes()))
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
