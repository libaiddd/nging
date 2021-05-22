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

type Slice_NgingUserRole []*NgingUserRole

func (s Slice_NgingUserRole) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingUserRole) RangeRaw(fn func(m *NgingUserRole) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingUserRole) GroupBy(keyField string) map[string][]*NgingUserRole {
	r := map[string][]*NgingUserRole{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingUserRole{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingUserRole) KeyBy(keyField string) map[string]*NgingUserRole {
	r := map[string]*NgingUserRole{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingUserRole) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingUserRole) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingUserRole) FromList(data interface{}) Slice_NgingUserRole {
	values, ok := data.([]*NgingUserRole)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingUserRole{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingUserRole 用户角色
type NgingUserRole struct {
	base    factory.Base
	objects []*NgingUserRole

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name        string `db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created     uint   `db:"created" bson:"created" comment:"添加时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	ParentId    uint   `db:"parent_id" bson:"parent_id" comment:"父级ID" json:"parent_id" xml:"parent_id"`
	PermCmd     string `db:"perm_cmd" bson:"perm_cmd" comment:"指令集权限(多个用“,”分隔开)" json:"perm_cmd" xml:"perm_cmd"`
	PermAction  string `db:"perm_action" bson:"perm_action" comment:"行为权限(多个用“,”分隔开)" json:"perm_action" xml:"perm_action"`
}

// - base function

func (a *NgingUserRole) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingUserRole) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingUserRole) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingUserRole) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingUserRole) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingUserRole) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingUserRole) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingUserRole) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingUserRole) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingUserRole) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingUserRole) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingUserRole) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingUserRole) Objects() []*NgingUserRole {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingUserRole) XObjects() Slice_NgingUserRole {
	return Slice_NgingUserRole(a.Objects())
}

func (a *NgingUserRole) NewObjects() factory.Ranger {
	return &Slice_NgingUserRole{}
}

func (a *NgingUserRole) InitObjects() *[]*NgingUserRole {
	a.objects = []*NgingUserRole{}
	return &a.objects
}

func (a *NgingUserRole) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingUserRole) Short_() string {
	return "nging_user_role"
}

func (a *NgingUserRole) Struct_() string {
	return "NgingUserRole"
}

func (a *NgingUserRole) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingUserRole) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingUserRole) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingUserRole) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingUserRole:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRole(*v))
		case []*NgingUserRole:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRole(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingUserRole) GroupBy(keyField string, inputRows ...[]*NgingUserRole) map[string][]*NgingUserRole {
	var rows Slice_NgingUserRole
	if len(inputRows) > 0 {
		rows = Slice_NgingUserRole(inputRows[0])
	} else {
		rows = Slice_NgingUserRole(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingUserRole) KeyBy(keyField string, inputRows ...[]*NgingUserRole) map[string]*NgingUserRole {
	var rows Slice_NgingUserRole
	if len(inputRows) > 0 {
		rows = Slice_NgingUserRole(inputRows[0])
	} else {
		rows = Slice_NgingUserRole(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingUserRole) AsKV(keyField string, valueField string, inputRows ...[]*NgingUserRole) param.Store {
	var rows Slice_NgingUserRole
	if len(inputRows) > 0 {
		rows = Slice_NgingUserRole(inputRows[0])
	} else {
		rows = Slice_NgingUserRole(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingUserRole) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingUserRole:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRole(*v))
		case []*NgingUserRole:
			err = DBI.FireReaded(a, queryParam, Slice_NgingUserRole(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingUserRole) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
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

func (a *NgingUserRole) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
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

func (a *NgingUserRole) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingUserRole) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

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

func (a *NgingUserRole) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
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

func (a *NgingUserRole) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingUserRole) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingUserRole) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingUserRole) Reset() *NgingUserRole {
	a.Id = 0
	a.Name = ``
	a.Description = ``
	a.Created = 0
	a.Updated = 0
	a.Disabled = ``
	a.ParentId = 0
	a.PermCmd = ``
	a.PermAction = ``
	return a
}

func (a *NgingUserRole) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Name"] = a.Name
		r["Description"] = a.Description
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		r["Disabled"] = a.Disabled
		r["ParentId"] = a.ParentId
		r["PermCmd"] = a.PermCmd
		r["PermAction"] = a.PermAction
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Name":
			r["Name"] = a.Name
		case "Description":
			r["Description"] = a.Description
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "ParentId":
			r["ParentId"] = a.ParentId
		case "PermCmd":
			r["PermCmd"] = a.PermCmd
		case "PermAction":
			r["PermAction"] = a.PermAction
		}
	}
	return r
}

func (a *NgingUserRole) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "parent_id":
			a.ParentId = param.AsUint(value)
		case "perm_cmd":
			a.PermCmd = param.AsString(value)
		case "perm_action":
			a.PermAction = param.AsString(value)
		}
	}
}

func (a *NgingUserRole) Set(key interface{}, value ...interface{}) {
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
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "ParentId":
			a.ParentId = param.AsUint(vv)
		case "PermCmd":
			a.PermCmd = param.AsString(vv)
		case "PermAction":
			a.PermAction = param.AsString(vv)
		}
	}
}

func (a *NgingUserRole) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["name"] = a.Name
		r["description"] = a.Description
		r["created"] = a.Created
		r["updated"] = a.Updated
		r["disabled"] = a.Disabled
		r["parent_id"] = a.ParentId
		r["perm_cmd"] = a.PermCmd
		r["perm_action"] = a.PermAction
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "name":
			r["name"] = a.Name
		case "description":
			r["description"] = a.Description
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		case "disabled":
			r["disabled"] = a.Disabled
		case "parent_id":
			r["parent_id"] = a.ParentId
		case "perm_cmd":
			r["perm_cmd"] = a.PermCmd
		case "perm_action":
			r["perm_action"] = a.PermAction
		}
	}
	return r
}

func (a *NgingUserRole) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingUserRole) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
