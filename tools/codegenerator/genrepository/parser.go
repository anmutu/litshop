package genrepository

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

type fieldConfig struct {
	FieldName  string
	ColumnName string
	FieldType  string
}

type structConfig struct {
	Pkgs         []ImportPkg
	PkgName      string
	LogName      string
	StructName   string
	OnlyFields   []fieldConfig
	OptionFields []fieldConfig
}

type Parser struct {
	dir         string
	pkg         *build.Package
	parsedFiles []*ast.File
}

func (p *Parser) Parse() []structConfig {
	var data []structConfig

	p.getPackage()
	p.parseGoFiles()
	for _, f := range p.parsedFiles {
		data = append(data, p.parseTypes(f)...)
	}

	return data
}

func (p *Parser) getPackage() {
	pkg, err := build.Default.ImportDir(p.dir, build.ImportComment)

	if err != nil {
		log.Fatalf("cannot process directory %s: %s", p.dir, err)
	}

	p.pkg = pkg
}

func (p *Parser) parseGoFiles() {
	var parsedFiles []*ast.File
	fs := token.NewFileSet()
	for _, file := range p.pkg.GoFiles {
		file = p.dir + "/" + file

		parsedFile, err := parser.ParseFile(fs, file, nil, 0)
		if err != nil {
			log.Fatalf("parsing package: %s: %s\n", file, err)
		}
		parsedFiles = append(parsedFiles, parsedFile)
	}

	p.parsedFiles = parsedFiles
}

func (p *Parser) parseTypes(file *ast.File) []structConfig {
	var data []structConfig

	ast.Inspect(file, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}

		for _, spec := range decl.Specs {
			var sc structConfig
			typeSpec, _ok := spec.(*ast.TypeSpec)
			if !_ok {
				continue
			}

			var structType *ast.StructType
			if structType, ok = typeSpec.Type.(*ast.StructType); !ok {
				continue
			}

			sc.StructName = typeSpec.Name.Name
			for _, v := range structType.Fields.List {
				var (
					onlyField   fieldConfig
					optionField fieldConfig
				)

				// 跳过自定义的类型
				if _, _ok := v.Type.(*ast.SelectorExpr); _ok {
					continue
				}

				if t, _ok := v.Type.(*ast.Ident); _ok {
					if t.Name == "Model" {
						onlyField.FieldName = "ID"
						onlyField.FieldType = "uint64"
						onlyField.ColumnName = defaultNamer("ID")
						sc.OnlyFields = append(sc.OnlyFields, onlyField)

						f1 := fieldConfig{}
						f1.FieldName = "CreatedAt"
						f1.FieldType = "time.Time"
						f1.ColumnName = defaultNamer("CreatedAt")
						sc.OptionFields = append(sc.OptionFields, f1)

						f2 := fieldConfig{}
						f2.FieldName = "UpdatedAt"
						f2.FieldType = "time.Time"
						f2.ColumnName = defaultNamer("UpdatedAt")
						sc.OptionFields = append(sc.OptionFields, f2)
						continue
					}
				}

				if v.Tag != nil {
					if strings.Contains(v.Tag.Value, "gorm") && strings.Contains(v.Tag.Value, "unique") ||
						strings.Contains(v.Tag.Value, "primary") {

						if t, _ok := v.Type.(*ast.Ident); _ok {
							onlyField.FieldType = t.String()
						}

						if len(v.Names) > 0 {
							onlyField.FieldName = v.Names[0].String()
							onlyField.ColumnName = defaultNamer(onlyField.FieldName)
						}

						sc.OnlyFields = append(sc.OnlyFields, onlyField)
						continue
					}
				}

				if t, _ok := v.Type.(*ast.Ident); _ok {
					optionField.FieldType = t.Name
				}

				if len(v.Names) > 0 {
					optionField.FieldName = v.Names[0].String()
					optionField.ColumnName = defaultNamer(optionField.FieldName)
				}

				sc.OptionFields = append(sc.OptionFields, optionField)
			}

			data = append(data, sc)
		}

		return true
	})

	return data
}

func NewParser(dir string) *Parser {
	return &Parser{
		dir: dir,
	}
}
