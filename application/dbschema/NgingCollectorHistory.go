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

type Slice_NgingCollectorHistory []*NgingCollectorHistory

func (s Slice_NgingCollectorHistory) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorHistory) RangeRaw(fn func(m *NgingCollectorHistory) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorHistory) GroupBy(keyField string) map[string][]*NgingCollectorHistory {
	r := map[string][]*NgingCollectorHistory{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCollectorHistory{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCollectorHistory) KeyBy(keyField string) map[string]*NgingCollectorHistory {
	r := map[string]*NgingCollectorHistory{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCollectorHistory) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCollectorHistory) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCollectorHistory) FromList(data interface{}) Slice_NgingCollectorHistory {
	values, ok := data.([]*NgingCollectorHistory)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCollectorHistory{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingCollectorHistory 采集历史
type NgingCollectorHistory struct {
	base    factory.Base
	objects []*NgingCollectorHistory

	Id           uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	ParentId     uint64 `db:"parent_id" bson:"parent_id" comment:"父ID" json:"parent_id" xml:"parent_id"`
	PageId       uint   `db:"page_id" bson:"page_id" comment:"页面ID" json:"page_id" xml:"page_id"`
	PageParentId uint   `db:"page_parent_id" bson:"page_parent_id" comment:"父页面ID" json:"page_parent_id" xml:"page_parent_id"`
	PageRootId   uint   `db:"page_root_id" bson:"page_root_id" comment:"入口页面ID" json:"page_root_id" xml:"page_root_id"`
	HasChild     string `db:"has_child" bson:"has_child" comment:"是否有子级" json:"has_child" xml:"has_child"`
	Url          string `db:"url" bson:"url" comment:"页面网址" json:"url" xml:"url"`
	UrlMd5       string `db:"url_md5" bson:"url_md5" comment:"页面网址MD5" json:"url_md5" xml:"url_md5"`
	Title        string `db:"title" bson:"title" comment:"页面标题" json:"title" xml:"title"`
	Content      string `db:"content" bson:"content" comment:"页面内容MD5" json:"content" xml:"content"`
	RuleMd5      string `db:"rule_md5" bson:"rule_md5" comment:"规则标识MD5" json:"rule_md5" xml:"rule_md5"`
	Data         string `db:"data" bson:"data" comment:"采集到的数据" json:"data" xml:"data"`
	Created      uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Exported     uint   `db:"exported" bson:"exported" comment:"最近导出时间" json:"exported" xml:"exported"`
}

// - base function

func (a *NgingCollectorHistory) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingCollectorHistory) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCollectorHistory) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCollectorHistory) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCollectorHistory) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCollectorHistory) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCollectorHistory) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCollectorHistory) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCollectorHistory) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingCollectorHistory) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCollectorHistory) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingCollectorHistory) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingCollectorHistory) Objects() []*NgingCollectorHistory {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCollectorHistory) XObjects() Slice_NgingCollectorHistory {
	return Slice_NgingCollectorHistory(a.Objects())
}

func (a *NgingCollectorHistory) NewObjects() factory.Ranger {
	return &Slice_NgingCollectorHistory{}
}

func (a *NgingCollectorHistory) InitObjects() *[]*NgingCollectorHistory {
	a.objects = []*NgingCollectorHistory{}
	return &a.objects
}

func (a *NgingCollectorHistory) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCollectorHistory) Short_() string {
	return "nging_collector_history"
}

func (a *NgingCollectorHistory) Struct_() string {
	return "NgingCollectorHistory"
}

func (a *NgingCollectorHistory) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCollectorHistory) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCollectorHistory) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingCollectorHistory) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCollectorHistory:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorHistory(*v))
		case []*NgingCollectorHistory:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorHistory(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorHistory) GroupBy(keyField string, inputRows ...[]*NgingCollectorHistory) map[string][]*NgingCollectorHistory {
	var rows Slice_NgingCollectorHistory
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorHistory(inputRows[0])
	} else {
		rows = Slice_NgingCollectorHistory(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCollectorHistory) KeyBy(keyField string, inputRows ...[]*NgingCollectorHistory) map[string]*NgingCollectorHistory {
	var rows Slice_NgingCollectorHistory
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorHistory(inputRows[0])
	} else {
		rows = Slice_NgingCollectorHistory(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCollectorHistory) AsKV(keyField string, valueField string, inputRows ...[]*NgingCollectorHistory) param.Store {
	var rows Slice_NgingCollectorHistory
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorHistory(inputRows[0])
	} else {
		rows = Slice_NgingCollectorHistory(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCollectorHistory) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCollectorHistory:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorHistory(*v))
		case []*NgingCollectorHistory:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorHistory(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorHistory) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.HasChild) == 0 {
		a.HasChild = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingCollectorHistory) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.HasChild) == 0 {
		a.HasChild = "N"
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

func (a *NgingCollectorHistory) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCollectorHistory) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["has_child"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["has_child"] = "N"
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

func (a *NgingCollectorHistory) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.HasChild) == 0 {
			a.HasChild = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.HasChild) == 0 {
			a.HasChild = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
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

func (a *NgingCollectorHistory) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingCollectorHistory) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCollectorHistory) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCollectorHistory) Reset() *NgingCollectorHistory {
	a.Id = 0
	a.ParentId = 0
	a.PageId = 0
	a.PageParentId = 0
	a.PageRootId = 0
	a.HasChild = ``
	a.Url = ``
	a.UrlMd5 = ``
	a.Title = ``
	a.Content = ``
	a.RuleMd5 = ``
	a.Data = ``
	a.Created = 0
	a.Exported = 0
	return a
}

func (a *NgingCollectorHistory) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["ParentId"] = a.ParentId
		r["PageId"] = a.PageId
		r["PageParentId"] = a.PageParentId
		r["PageRootId"] = a.PageRootId
		r["HasChild"] = a.HasChild
		r["Url"] = a.Url
		r["UrlMd5"] = a.UrlMd5
		r["Title"] = a.Title
		r["Content"] = a.Content
		r["RuleMd5"] = a.RuleMd5
		r["Data"] = a.Data
		r["Created"] = a.Created
		r["Exported"] = a.Exported
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "ParentId":
			r["ParentId"] = a.ParentId
		case "PageId":
			r["PageId"] = a.PageId
		case "PageParentId":
			r["PageParentId"] = a.PageParentId
		case "PageRootId":
			r["PageRootId"] = a.PageRootId
		case "HasChild":
			r["HasChild"] = a.HasChild
		case "Url":
			r["Url"] = a.Url
		case "UrlMd5":
			r["UrlMd5"] = a.UrlMd5
		case "Title":
			r["Title"] = a.Title
		case "Content":
			r["Content"] = a.Content
		case "RuleMd5":
			r["RuleMd5"] = a.RuleMd5
		case "Data":
			r["Data"] = a.Data
		case "Created":
			r["Created"] = a.Created
		case "Exported":
			r["Exported"] = a.Exported
		}
	}
	return r
}

func (a *NgingCollectorHistory) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "parent_id":
			a.ParentId = param.AsUint64(value)
		case "page_id":
			a.PageId = param.AsUint(value)
		case "page_parent_id":
			a.PageParentId = param.AsUint(value)
		case "page_root_id":
			a.PageRootId = param.AsUint(value)
		case "has_child":
			a.HasChild = param.AsString(value)
		case "url":
			a.Url = param.AsString(value)
		case "url_md5":
			a.UrlMd5 = param.AsString(value)
		case "title":
			a.Title = param.AsString(value)
		case "content":
			a.Content = param.AsString(value)
		case "rule_md5":
			a.RuleMd5 = param.AsString(value)
		case "data":
			a.Data = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "exported":
			a.Exported = param.AsUint(value)
		}
	}
}

func (a *NgingCollectorHistory) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint64(vv)
		case "ParentId":
			a.ParentId = param.AsUint64(vv)
		case "PageId":
			a.PageId = param.AsUint(vv)
		case "PageParentId":
			a.PageParentId = param.AsUint(vv)
		case "PageRootId":
			a.PageRootId = param.AsUint(vv)
		case "HasChild":
			a.HasChild = param.AsString(vv)
		case "Url":
			a.Url = param.AsString(vv)
		case "UrlMd5":
			a.UrlMd5 = param.AsString(vv)
		case "Title":
			a.Title = param.AsString(vv)
		case "Content":
			a.Content = param.AsString(vv)
		case "RuleMd5":
			a.RuleMd5 = param.AsString(vv)
		case "Data":
			a.Data = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Exported":
			a.Exported = param.AsUint(vv)
		}
	}
}

func (a *NgingCollectorHistory) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["parent_id"] = a.ParentId
		r["page_id"] = a.PageId
		r["page_parent_id"] = a.PageParentId
		r["page_root_id"] = a.PageRootId
		r["has_child"] = a.HasChild
		r["url"] = a.Url
		r["url_md5"] = a.UrlMd5
		r["title"] = a.Title
		r["content"] = a.Content
		r["rule_md5"] = a.RuleMd5
		r["data"] = a.Data
		r["created"] = a.Created
		r["exported"] = a.Exported
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "parent_id":
			r["parent_id"] = a.ParentId
		case "page_id":
			r["page_id"] = a.PageId
		case "page_parent_id":
			r["page_parent_id"] = a.PageParentId
		case "page_root_id":
			r["page_root_id"] = a.PageRootId
		case "has_child":
			r["has_child"] = a.HasChild
		case "url":
			r["url"] = a.Url
		case "url_md5":
			r["url_md5"] = a.UrlMd5
		case "title":
			r["title"] = a.Title
		case "content":
			r["content"] = a.Content
		case "rule_md5":
			r["rule_md5"] = a.RuleMd5
		case "data":
			r["data"] = a.Data
		case "created":
			r["created"] = a.Created
		case "exported":
			r["exported"] = a.Exported
		}
	}
	return r
}

func (a *NgingCollectorHistory) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCollectorHistory) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
