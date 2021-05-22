// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingCollectorExport []*NgingCollectorExport

func (s Slice_NgingCollectorExport) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorExport) RangeRaw(fn func(m *NgingCollectorExport) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorExport) GroupBy(keyField string) map[string][]*NgingCollectorExport {
	r := map[string][]*NgingCollectorExport{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCollectorExport{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCollectorExport) KeyBy(keyField string) map[string]*NgingCollectorExport {
	r := map[string]*NgingCollectorExport{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCollectorExport) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCollectorExport) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCollectorExport) FromList(data interface{}) Slice_NgingCollectorExport {
	values, ok := data.([]*NgingCollectorExport)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCollectorExport{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingCollectorExport 导出规则
type NgingCollectorExport struct {
	base    factory.Base
	objects []*NgingCollectorExport

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	PageRoot    uint   `db:"page_root" bson:"page_root" comment:"根页面ID" json:"page_root" xml:"page_root"`
	PageId      uint   `db:"page_id" bson:"page_id" comment:"页面ID" json:"page_id" xml:"page_id"`
	GroupId     uint   `db:"group_id" bson:"group_id" comment:"组ID" json:"group_id" xml:"group_id"`
	Mapping     string `db:"mapping" bson:"mapping" comment:"字段映射" json:"mapping" xml:"mapping"`
	Dest        string `db:"dest" bson:"dest" comment:"目标" json:"dest" xml:"dest"`
	DestType    string `db:"dest_type" bson:"dest_type" comment:"目标类型" json:"dest_type" xml:"dest_type"`
	Name        string `db:"name" bson:"name" comment:"方案名" json:"name" xml:"name"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Exported    uint   `db:"exported" bson:"exported" comment:"最近导出时间" json:"exported" xml:"exported"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
}

// - base function

func (a *NgingCollectorExport) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingCollectorExport) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCollectorExport) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCollectorExport) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCollectorExport) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCollectorExport) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCollectorExport) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCollectorExport) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCollectorExport) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingCollectorExport) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCollectorExport) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingCollectorExport) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingCollectorExport) Objects() []*NgingCollectorExport {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCollectorExport) XObjects() Slice_NgingCollectorExport {
	return Slice_NgingCollectorExport(a.Objects())
}

func (a *NgingCollectorExport) NewObjects() factory.Ranger {
	return &Slice_NgingCollectorExport{}
}

func (a *NgingCollectorExport) InitObjects() *[]*NgingCollectorExport {
	a.objects = []*NgingCollectorExport{}
	return &a.objects
}

func (a *NgingCollectorExport) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCollectorExport) Short_() string {
	return "nging_collector_export"
}

func (a *NgingCollectorExport) Struct_() string {
	return "NgingCollectorExport"
}

func (a *NgingCollectorExport) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCollectorExport) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCollectorExport) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingCollectorExport) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingCollectorExport:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExport(*v))
		case []*NgingCollectorExport:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExport(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorExport) GroupBy(keyField string, inputRows ...[]*NgingCollectorExport) map[string][]*NgingCollectorExport {
	var rows Slice_NgingCollectorExport
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorExport(inputRows[0])
	} else {
		rows = Slice_NgingCollectorExport(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCollectorExport) KeyBy(keyField string, inputRows ...[]*NgingCollectorExport) map[string]*NgingCollectorExport {
	var rows Slice_NgingCollectorExport
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorExport(inputRows[0])
	} else {
		rows = Slice_NgingCollectorExport(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCollectorExport) AsKV(keyField string, valueField string, inputRows ...[]*NgingCollectorExport) param.Store {
	var rows Slice_NgingCollectorExport
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorExport(inputRows[0])
	} else {
		rows = Slice_NgingCollectorExport(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCollectorExport) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingCollectorExport:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExport(*v))
		case []*NgingCollectorExport:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorExport(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorExport) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.DestType) == 0 {
		a.DestType = "dbAccountID"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingCollectorExport) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.DestType) == 0 {
		a.DestType = "dbAccountID"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingCollectorExport) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCollectorExport) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["dest_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["dest_type"] = "dbAccountID"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingCollectorExport) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.DestType) == 0 {
			a.DestType = "dbAccountID"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.DestType) == 0 {
			a.DestType = "dbAccountID"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingCollectorExport) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingCollectorExport) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCollectorExport) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCollectorExport) Reset() *NgingCollectorExport {
	a.Id = 0
	a.PageRoot = 0
	a.PageId = 0
	a.GroupId = 0
	a.Mapping = ``
	a.Dest = ``
	a.DestType = ``
	a.Name = ``
	a.Description = ``
	a.Created = 0
	a.Exported = 0
	a.Disabled = ``
	return a
}

func (a *NgingCollectorExport) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["PageRoot"] = a.PageRoot
		r["PageId"] = a.PageId
		r["GroupId"] = a.GroupId
		r["Mapping"] = a.Mapping
		r["Dest"] = a.Dest
		r["DestType"] = a.DestType
		r["Name"] = a.Name
		r["Description"] = a.Description
		r["Created"] = a.Created
		r["Exported"] = a.Exported
		r["Disabled"] = a.Disabled
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "PageRoot":
			r["PageRoot"] = a.PageRoot
		case "PageId":
			r["PageId"] = a.PageId
		case "GroupId":
			r["GroupId"] = a.GroupId
		case "Mapping":
			r["Mapping"] = a.Mapping
		case "Dest":
			r["Dest"] = a.Dest
		case "DestType":
			r["DestType"] = a.DestType
		case "Name":
			r["Name"] = a.Name
		case "Description":
			r["Description"] = a.Description
		case "Created":
			r["Created"] = a.Created
		case "Exported":
			r["Exported"] = a.Exported
		case "Disabled":
			r["Disabled"] = a.Disabled
		}
	}
	return r
}

func (a *NgingCollectorExport) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "page_root":
			a.PageRoot = param.AsUint(value)
		case "page_id":
			a.PageId = param.AsUint(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "mapping":
			a.Mapping = param.AsString(value)
		case "dest":
			a.Dest = param.AsString(value)
		case "dest_type":
			a.DestType = param.AsString(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "exported":
			a.Exported = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		}
	}
}

func (a *NgingCollectorExport) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "PageRoot":
			a.PageRoot = param.AsUint(vv)
		case "PageId":
			a.PageId = param.AsUint(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Mapping":
			a.Mapping = param.AsString(vv)
		case "Dest":
			a.Dest = param.AsString(vv)
		case "DestType":
			a.DestType = param.AsString(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Exported":
			a.Exported = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		}
	}
}

func (a *NgingCollectorExport) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["page_root"] = a.PageRoot
		r["page_id"] = a.PageId
		r["group_id"] = a.GroupId
		r["mapping"] = a.Mapping
		r["dest"] = a.Dest
		r["dest_type"] = a.DestType
		r["name"] = a.Name
		r["description"] = a.Description
		r["created"] = a.Created
		r["exported"] = a.Exported
		r["disabled"] = a.Disabled
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "page_root":
			r["page_root"] = a.PageRoot
		case "page_id":
			r["page_id"] = a.PageId
		case "group_id":
			r["group_id"] = a.GroupId
		case "mapping":
			r["mapping"] = a.Mapping
		case "dest":
			r["dest"] = a.Dest
		case "dest_type":
			r["dest_type"] = a.DestType
		case "name":
			r["name"] = a.Name
		case "description":
			r["description"] = a.Description
		case "created":
			r["created"] = a.Created
		case "exported":
			r["exported"] = a.Exported
		case "disabled":
			r["disabled"] = a.Disabled
		}
	}
	return r
}

func (a *NgingCollectorExport) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCollectorExport) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
