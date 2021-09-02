package crud

import "strings"

// PoDefContent .
func PoDefContent() string {
	return `
//Package po generated by 'freedom new-po'
package po
{{.Import}}
{{.Content}}

// GetChanges .
func (obj *{{.Name}})GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *{{.Name}}) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

{{range .SetMethods}}
// Set{{.Name}} .
func (obj *{{.ObjectName}}) Set{{.Name}} ({{.Variable}} {{.VariableType}}) {
	obj.{{.Name}} = {{.Variable}} 
	obj.Update("{{.Column}}", {{.Variable}})
}
{{ end }}

{{range .AddMethods}}
// Add{{.Name}} .
func (obj *{{.ObjectName}}) Add{{.Name}} ({{.Variable}} {{.VariableType}}) {
	obj.{{.Name}} += {{.Variable}} 
	obj.Update("{{.Column}}", gorm.Expr("{{.Column}} + ?", {{.Variable}}))
}
{{ end }}
`
}

// FunTemplatePackage .
func FunTemplatePackage() string {
	source := `
	package repository
	import (
		"errors"
		"github.com/8treenet/freedom"
		"gorm.io/gorm"
		"time"
		"{{.PackagePath}}"
		"fmt"
		"strings"
	)
	
	// GORMRepository .
	type GORMRepository interface {
		db() *gorm.DB
		Worker() freedom.Worker
	}

	type saveObject interface {
		TableName() string
		Location() map[string]interface{}
		GetChanges() map[string]interface{}
	}	

	// Builder .
	type Builder interface {
		Execute(db *gorm.DB, object interface{}) error
	}

	// Pager .
	type Pager struct {
		pageSize  int
		page      int
		totalPage int
		fields    []string
		orders    []string
	}

	// NewDescPager .
	func NewDescPager(column string, columns ...string) *Pager {
		return newDefaultPager("desc", column, columns...)
	}

	// NewAscPager .
	func NewAscPager(column string, columns ...string) *Pager {
		return newDefaultPager("asc", column, columns...)
	}

	// NewDescOrder .
	func newDefaultPager(sort, field string, args ...string) *Pager {
		fields := []string{field}
		fields = append(fields, args...)
		orders := []string{}
		for index := 0; index < len(fields); index++ {
			orders = append(orders, sort)
		}
		return &Pager{
			fields: fields,
			orders: orders,
		}
	}

	// Order .
	func (p *Pager) Order() interface{} {
		if len(p.fields) == 0 {
			return nil
		}
		args := []string{}
		for index := 0; index < len(p.fields); index++ {
			args = append(args, fmt.Sprintf("$$wave%s$$wave %s", p.fields[index], p.orders[index]))
		}

		return strings.Join(args, ",")
	}

	// TotalPage .
	func (p *Pager) TotalPage() int {
		return p.totalPage
	}

	// SetPage .
	func (p *Pager) SetPage(page, pageSize int) *Pager {
		p.page = page
		p.pageSize = pageSize
		return p
	}
	
	// Execute .
	func (p *Pager) Execute(db *gorm.DB, object interface{}) (e error) {
		if p.page != 0 && p.pageSize != 0 {
			var count64 int64
			e = db.Model(object).Count(&count64).Error
			count := int(count64)
			if e != nil {
				return
			}
			if count != 0 {
				//Calculate the length of the pagination
				if count%p.pageSize == 0 {
					p.totalPage = count / p.pageSize
				} else {
					p.totalPage = count / p.pageSize + 1
				}
			}
			db = db.Offset((p.page - 1) * p.pageSize).Limit(p.pageSize)
		}
		
		orderValue := p.Order()
		if orderValue != nil {
			db = db.Order(orderValue)
		}
		
		resultDB := db.Find(object)
		if resultDB.Error != nil {
			return resultDB.Error
		}
		return
	}

	
	func ormErrorLog(repo GORMRepository, model, method string, e error, expression ...interface{}) {
		if e == nil || e == gorm.ErrRecordNotFound {
			return
		}
		repo.Worker().Logger().Errorf("error: %v, model: %s, method: %s", e, model, method)
	}
`
	return strings.ReplaceAll(source, "$$wave", "`")
}

// FunTemplate .
func FunTemplate() string {
	return `
	// find{{.Name}} .
	func find{{.Name}}(repo GORMRepository, result *po.{{.Name}}, builders ...Builder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}", e, result)
		}()
		db := repo.db()
		if len(builders) == 0 {
			e = db.Where(result).Last(result).Error
			return
		}
		e = builders[0].Execute(db.Limit(1), result)
		return
	}
	
	// find{{.Name}}ListByPrimarys .
	func find{{.Name}}ListByPrimarys(repo GORMRepository, primarys ...interface{}) (results []po.{{.Name}}, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ListByPrimarys", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}sByPrimarys", e, primarys)
		}()

		e = repo.db().Find(&results, primarys).Error
		return
	}
	
	// find{{.Name}}ByWhere .
	func find{{.Name}}ByWhere(repo GORMRepository, query string, args []interface{}, builders ...Builder) (result po.{{.Name}}, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ByWhere", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}ByWhere", e, query, args)
		}()
		db := repo.db()
		if query != "" {
			db = db.Where(query, args...)
		}
		if len(builders) == 0 {
			e = db.Last(&result).Error
			return
		}
	
		e = builders[0].Execute(db.Limit(1), &result)
		return
	}
	
	// find{{.Name}}ByMap .
	func find{{.Name}}ByMap(repo GORMRepository, query map[string]interface{}, builders ...Builder) (result po.{{.Name}}, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ByMap", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}ByMap", e, query)
		}()

		db := repo.db().Where(query)
		if len(builders) == 0 {
			e = db.Last(&result).Error
			return
		}
	
		e = builders[0].Execute(db.Limit(1), &result)
		return
	}
	
	// find{{.Name}}List .
	func find{{.Name}}List(repo GORMRepository, query po.{{.Name}}, builders ...Builder) (results []po.{{.Name}}, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}List", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}s", e, query)
		}()
		db := repo.db().Where(query)
	
		if len(builders) == 0 {
			e = db.Find(&results).Error
			return
		}
		e = builders[0].Execute(db, &results)
		return
	}
	
	// find{{.Name}}ListByWhere .
	func find{{.Name}}ListByWhere(repo GORMRepository, query string, args []interface{}, builders ...Builder) (results []po.{{.Name}}, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ListByWhere", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}sByWhere", e, query, args)
		}()
		db := repo.db()
		if query != "" {
			db = db.Where(query, args...)
		}
	
		if len(builders) == 0 {
			e = db.Find(&results).Error
			return
		}
		e = builders[0].Execute(db, &results)
		return
	}
	
	// find{{.Name}}ListByMap .
	func find{{.Name}}ListByMap(repo GORMRepository, query map[string]interface{}, builders ...Builder) (results []po.{{.Name}}, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ListByMap", e, now)
			ormErrorLog(repo, "{{.Name}}", "find{{.Name}}sByMap", e, query)
		}()

		db := repo.db().Where(query)
	
		if len(builders) == 0 {
			e = db.Find(&results).Error
			return
		}
		e = builders[0].Execute(db, &results)
		return
	}
	
	// create{{.Name}} .
	func create{{.Name}}(repo GORMRepository, object *po.{{.Name}}) (rowsAffected int64, e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "create{{.Name}}", e, now)
			ormErrorLog(repo, "{{.Name}}", "create{{.Name}}", e, *object)
		}()

		db := repo.db().Create(object)
		rowsAffected = db.RowsAffected
		e = db.Error
		return
	}

	// save{{.Name}} .
	func save{{.Name}}(repo GORMRepository, object saveObject) (rowsAffected int64, e error) {
		if len(object.Location()) == 0 {
			return -1, errors.New("location cannot be empty")
		}
		updateValues := object.GetChanges()
		if len(updateValues) == 0 {
			return -1, nil
		}
		
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "save{{.Name}}", e, now)
			ormErrorLog(repo, "{{.Name}}", "save{{.Name}}", e, object)
		}()

		db := repo.db().Table(object.TableName()).Where(object.Location()).Updates(updateValues)
		e = db.Error
		rowsAffected = db.RowsAffected
		return
	}
`
}
