package genrepository

import (
	"fmt"
	"text/template"
)

func parseModelTemplateOrPanic(m string) *template.Template {
	tpl, err := template.New("model_template").Parse(m)
	if err != nil {
		panic(err)
	}
	return tpl
}

var commonTemplate = parseModelTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}
type FieldData struct {
		Value interface{} %sjson:"value" form:"value"%s
		Symbol string %sjson:"symbol" form:"symbol"%s
	
}
`, "`", "`", "`", "`"))

var modelTemplate = parseModelTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}

{{$LogName := .LogName}}

{{range .Pkgs}}
{{ if .Pkg }}
import "{{.Pkg}}"
{{end}}
{{end}}

	// New{{.StructName}} new
	func New{{.StructName}}() *{{.StructName}} {
		return (new({{.StructName}})).init()
	}

	// init
	func (m *{{.StructName}}) init() *{{.StructName}} {
		m.InitDB()
		return m
	}

	// Add add one record
	func (m *{{.StructName}}) Save()(err error) {
		return 
	}

	// Delete delete record
	func (m *{{.StructName}}) Delete()(err error) {
		return
	}
	
	// Updates update record
	func (m *{{.StructName}}) Updates(updateOpt map[string]interface{})(err error) {
		if err = m.DB.Model(&{{.StructName}}{}).Where("id = ?",m.ID).Updates(updateOpt).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			return
		}
		return 
	}

	// Get{{.StructName}}All get all record
	func (m *{{.StructName}}) Get{{.StructName}}All()(ret []*{{.StructName}},err error){
		if err = m.DB.Find(&ret).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			return
		}
		return
	}
	
	// Get{{.StructName}}Count get count
	func (m *{{.StructName}}) Get{{.StructName}}Count()(ret int64){
		m.DB.Model(&{{.StructName}}{}).Count(&ret)
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
	func (m *{{.StructName}}) Get{{$StructName}}List(q *Query{{$StructName}}Form)(ret []*{{$StructName}},err error){
		db := m.DB
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
		func (m *{{$StructName}}) SetQueryBy{{.FieldName}}({{.ColumnName}} {{.FieldType}})*{{$StructName}} {
			m.{{.FieldName}} = {{.ColumnName}}
			return m
		}
		// GetBy{{.FieldName}} get one record by {{.FieldName}}
		func (m *{{$StructName}})GetBy{{.FieldName}}()(err error){
			if err = m.DB.First(m,"{{.ColumnName}} = ?",m.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				return
			}
			return
		}
		// DeleteBy{{.FieldName}} delete record by {{.FieldName}}
		func (m *{{$StructName}}) DeleteBy{{.FieldName}}()(err error) {
			if err= m.DB.Delete(m,"{{.ColumnName}} = ?",m.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				return
				}
			return
		}
	{{end}}
`, "`", "`", "`", "`", "`", "`", "`", "`"))
