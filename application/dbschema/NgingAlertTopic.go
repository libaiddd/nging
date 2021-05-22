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

type Slice_NgingAlertTopic []*NgingAlertTopic

func (s Slice_NgingAlertTopic) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingAlertTopic) RangeRaw(fn func(m *NgingAlertTopic) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingAlertTopic) GroupBy(keyField string) map[string][]*NgingAlertTopic {
	r := map[string][]*NgingAlertTopic{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingAlertTopic{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingAlertTopic) KeyBy(keyField string) map[string]*NgingAlertTopic {
	r := map[string]*NgingAlertTopic{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingAlertTopic) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingAlertTopic) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingAlertTopic) FromList(data interface{}) Slice_NgingAlertTopic {
	values, ok := data.([]*NgingAlertTopic)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingAlertTopic{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingAlertTopic 报警收信专题关联
type NgingAlertTopic struct {
	base    factory.Base
	objects []*NgingAlertTopic

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Topic       string `db:"topic" bson:"topic" comment:"通知专题" json:"topic" xml:"topic"`
	RecipientId uint   `db:"recipient_id" bson:"recipient_id" comment:"收信账号" json:"recipient_id" xml:"recipient_id"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingAlertTopic) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingAlertTopic) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingAlertTopic) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingAlertTopic) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingAlertTopic) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingAlertTopic) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingAlertTopic) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingAlertTopic) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingAlertTopic) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingAlertTopic) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingAlertTopic) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingAlertTopic) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingAlertTopic) Objects() []*NgingAlertTopic {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingAlertTopic) XObjects() Slice_NgingAlertTopic {
	return Slice_NgingAlertTopic(a.Objects())
}

func (a *NgingAlertTopic) NewObjects() factory.Ranger {
	return &Slice_NgingAlertTopic{}
}

func (a *NgingAlertTopic) InitObjects() *[]*NgingAlertTopic {
	a.objects = []*NgingAlertTopic{}
	return &a.objects
}

func (a *NgingAlertTopic) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingAlertTopic) Short_() string {
	return "nging_alert_topic"
}

func (a *NgingAlertTopic) Struct_() string {
	return "NgingAlertTopic"
}

func (a *NgingAlertTopic) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingAlertTopic) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingAlertTopic) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingAlertTopic) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingAlertTopic:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertTopic(*v))
		case []*NgingAlertTopic:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertTopic(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingAlertTopic) GroupBy(keyField string, inputRows ...[]*NgingAlertTopic) map[string][]*NgingAlertTopic {
	var rows Slice_NgingAlertTopic
	if len(inputRows) > 0 {
		rows = Slice_NgingAlertTopic(inputRows[0])
	} else {
		rows = Slice_NgingAlertTopic(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingAlertTopic) KeyBy(keyField string, inputRows ...[]*NgingAlertTopic) map[string]*NgingAlertTopic {
	var rows Slice_NgingAlertTopic
	if len(inputRows) > 0 {
		rows = Slice_NgingAlertTopic(inputRows[0])
	} else {
		rows = Slice_NgingAlertTopic(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingAlertTopic) AsKV(keyField string, valueField string, inputRows ...[]*NgingAlertTopic) param.Store {
	var rows Slice_NgingAlertTopic
	if len(inputRows) > 0 {
		rows = Slice_NgingAlertTopic(inputRows[0])
	} else {
		rows = Slice_NgingAlertTopic(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingAlertTopic) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingAlertTopic:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertTopic(*v))
		case []*NgingAlertTopic:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertTopic(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingAlertTopic) Add() (pk interface{}, err error) {
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

func (a *NgingAlertTopic) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingAlertTopic) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingAlertTopic) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

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

func (a *NgingAlertTopic) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
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

func (a *NgingAlertTopic) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingAlertTopic) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingAlertTopic) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingAlertTopic) Reset() *NgingAlertTopic {
	a.Id = 0
	a.Topic = ``
	a.RecipientId = 0
	a.Disabled = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingAlertTopic) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Topic"] = a.Topic
		r["RecipientId"] = a.RecipientId
		r["Disabled"] = a.Disabled
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Topic":
			r["Topic"] = a.Topic
		case "RecipientId":
			r["RecipientId"] = a.RecipientId
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingAlertTopic) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "topic":
			a.Topic = param.AsString(value)
		case "recipient_id":
			a.RecipientId = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingAlertTopic) Set(key interface{}, value ...interface{}) {
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
		case "Topic":
			a.Topic = param.AsString(vv)
		case "RecipientId":
			a.RecipientId = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingAlertTopic) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["topic"] = a.Topic
		r["recipient_id"] = a.RecipientId
		r["disabled"] = a.Disabled
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "topic":
			r["topic"] = a.Topic
		case "recipient_id":
			r["recipient_id"] = a.RecipientId
		case "disabled":
			r["disabled"] = a.Disabled
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingAlertTopic) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingAlertTopic) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
