package genrepository

import (
	"fmt"
	"text/template"
)

func parseModelTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("model_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var modelTemplate = parseModelTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}

{{$LogName := .LogName}}

import (
	{{range .ImportPkgs}}
	"{{.Pkg}}"
	{{end}}
)

type {{.StructName}} struct {
	DB *gorm.DB
}

	// New{{.StructName}} new
	func New{{.StructName}}() *{{.StructName}} {
		return (new({{.StructName}})).init()
	}

	// init
	func (t *{{.StructName}}) init()(err error) {
		t.InitDB()
		return t
	}

	// Add add one record
	func (t *{{.StructName}}) Save()(err error) {
		return 
	}

	// Delete delete record
	func (t *{{.StructName}}) Delete()(err error) {
		return
	}
	
	// Updates update record
	func (t *{{.StructName}}) Updates(m map[string]interface{})(err error) {
		if err = t.DB.Model(&{{.StructName}}{}).Where("id = ?",t.ID).Updates(m).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			return
		}
		return 
	}

	// Get{{.StructName}}All get all record
	func Get{{.StructName}}All()(ret []*{{.StructName}},err error){
		if err = t.DB.Find(&ret).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			return
		}
		return
	}
	
	// Get{{.StructName}}Count get count
	func Get{{.StructName}}Count()(ret int64){
		t.db.Model(&{{.StructName}}{}).Count(&ret)
		return
	}

	{{$StructName := .StructName}}
	type Query{{$StructName}}Form struct{
	{{range .OptionFields}} {{.FieldName}} *FieldData %sjson:"{{.FieldName}}" form:"{{.FieldName}}"%s; {{end}}
		Order []string %sjson:"order" form:"order"%s
		PageNum int %sjson:"pageNum" form:"pageNum"%s
		PageSize int %sjson:"pageSize" form:"pageSize"%s
		}

	//  Get{{$StructName}}List get {{$StructName}} list some field value or some condition
	func Get{{$StructName}}List(q *Query{{$StructName}}Form)(ret []*{{$StructName}},err error){
		db := t.DB
		// order
		if len(q.Order)>0{
			for _,v:=range q.Order {
				db = db.Order(v)
			}
		}
		// pageSize
		if q.PageSize!=0{
			db = db.Limit(q.PageSize)
		}
		// pageNum
		if q.PageNum!=0{
			q.PageNum = (q.PageNum - 1) * q.PageSize
			db = db.Offset(q.PageNum)
		}
	{{range .OptionFields}} 
		// {{.FieldName}}
		if q.{{.FieldName}}!=nil{
			db = db.Where("{{.ColumnName}}" +q.{{.FieldName}}.Symbol +"?",q.{{.FieldName}}.Value)
		}  ; {{end}}
		if err = db.Find(&ret).Error;err!=nil{
			return	
		}
		return 
	}
	{{range .OnlyFields}}
		// QueryBy{{.FieldName}} query cond by {{.FieldName}}
		func (t *{{$StructName}}) SetQueryBy{{.FieldName}}({{.ColumnName}} {{.FieldType}})*{{$StructName}} {
			t.{{.FieldName}} = {{.ColumnName}}
			return  t
		}
		// GetBy{{.FieldName}} get one record by {{.FieldName}}
		func (t *{{$StructName}})GetBy{{.FieldName}}()(err error){
			if err = t.db.First(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				return
			}
			return
		}
		// DeleteBy{{.FieldName}} delete record by {{.FieldName}}
		func (t *{{$StructName}}) DeleteBy{{.FieldName}}()(err error) {
			if err= t.db.Delete(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				return
				}
			return
		}
	{{end}}
`, "`", "`", "`", "`", "`", "`", "`", "`"))
