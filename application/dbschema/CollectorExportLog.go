// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

// CollectorExportLog 导出日志
type CollectorExportLog struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CollectorExportLog
	namer   func(string) string
	connID  int
	
	Id       	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	PageId   	uint    	`db:"page_id" bson:"page_id" comment:"页面规则ID" json:"page_id" xml:"page_id"`
	ExportId 	uint    	`db:"export_id" bson:"export_id" comment:"导出方案ID" json:"export_id" xml:"export_id"`
	Result   	string  	`db:"result" bson:"result" comment:"结果" json:"result" xml:"result"`
	Status   	string  	`db:"status" bson:"status" comment:"状态" json:"status" xml:"status"`
	Created  	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

func (this *CollectorExportLog) Trans() *factory.Transaction {
	return this.trans
}

func (this *CollectorExportLog) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CollectorExportLog) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CollectorExportLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CollectorExportLog) Objects() []*CollectorExportLog {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CollectorExportLog) NewObjects() *[]*CollectorExportLog {
	this.objects = []*CollectorExportLog{}
	return &this.objects
}

func (this *CollectorExportLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CollectorExportLog) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CollectorExportLog) Short_() string {
	return "collector_export_log"
}

func (this *CollectorExportLog) Struct_() string {
	return "CollectorExportLog"
}

func (this *CollectorExportLog) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *CollectorExportLog) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CollectorExportLog) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CollectorExportLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CollectorExportLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorExportLog) GroupBy(keyField string, inputRows ...[]*CollectorExportLog) map[string][]*CollectorExportLog {
	var rows []*CollectorExportLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*CollectorExportLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*CollectorExportLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *CollectorExportLog) KeyBy(keyField string, inputRows ...[]*CollectorExportLog) map[string]*CollectorExportLog {
	var rows []*CollectorExportLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*CollectorExportLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *CollectorExportLog) AsKV(keyField string, valueField string, inputRows ...[]*CollectorExportLog) map[string]interface{} {
	var rows []*CollectorExportLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *CollectorExportLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorExportLog) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Status) == 0 { this.Status = "idle" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *CollectorExportLog) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.Status) == 0 { this.Status = "idle" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CollectorExportLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CollectorExportLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CollectorExportLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["status"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["status"] = "idle" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CollectorExportLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.Status) == 0 { this.Status = "idle" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Status) == 0 { this.Status = "idle" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *CollectorExportLog) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CollectorExportLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CollectorExportLog) Reset() *CollectorExportLog {
	this.Id = 0
	this.PageId = 0
	this.ExportId = 0
	this.Result = ``
	this.Status = ``
	this.Created = 0
	return this
}

func (this *CollectorExportLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["PageId"] = this.PageId
	r["ExportId"] = this.ExportId
	r["Result"] = this.Result
	r["Status"] = this.Status
	r["Created"] = this.Created
	return r
}

func (this *CollectorExportLog) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["page_id"] = this.PageId
	r["export_id"] = this.ExportId
	r["result"] = this.Result
	r["status"] = this.Status
	r["created"] = this.Created
	return r
}

func (this *CollectorExportLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *CollectorExportLog) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

