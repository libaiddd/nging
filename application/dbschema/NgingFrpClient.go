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

type Slice_NgingFrpClient []*NgingFrpClient

func (s Slice_NgingFrpClient) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpClient) RangeRaw(fn func(m *NgingFrpClient) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpClient) GroupBy(keyField string) map[string][]*NgingFrpClient {
	r := map[string][]*NgingFrpClient{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFrpClient{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFrpClient) KeyBy(keyField string) map[string]*NgingFrpClient {
	r := map[string]*NgingFrpClient{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFrpClient) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFrpClient) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFrpClient) FromList(data interface{}) Slice_NgingFrpClient {
	values, ok := data.([]*NgingFrpClient)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFrpClient{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFrpClient FRP客户端设置
type NgingFrpClient struct {
	base    factory.Base
	objects []*NgingFrpClient

	Id                uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name              string `db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Disabled          string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	ServerAddr        string `db:"server_addr" bson:"server_addr" comment:"" json:"server_addr" xml:"server_addr"`
	ServerPort        uint   `db:"server_port" bson:"server_port" comment:"" json:"server_port" xml:"server_port"`
	HttpProxy         string `db:"http_proxy" bson:"http_proxy" comment:"" json:"http_proxy" xml:"http_proxy"`
	PoolCount         uint   `db:"pool_count" bson:"pool_count" comment:"" json:"pool_count" xml:"pool_count"`
	TcpMux            string `db:"tcp_mux" bson:"tcp_mux" comment:"必须跟服务端一致" json:"tcp_mux" xml:"tcp_mux"`
	User              string `db:"user" bson:"user" comment:"" json:"user" xml:"user"`
	DnsServer         string `db:"dns_server" bson:"dns_server" comment:"" json:"dns_server" xml:"dns_server"`
	LoginFailExit     string `db:"login_fail_exit" bson:"login_fail_exit" comment:"" json:"login_fail_exit" xml:"login_fail_exit"`
	Protocol          string `db:"protocol" bson:"protocol" comment:"tcp or kcp" json:"protocol" xml:"protocol"`
	HeartbeatInterval int64  `db:"heartbeat_interval" bson:"heartbeat_interval" comment:"" json:"heartbeat_interval" xml:"heartbeat_interval"`
	HeartbeatTimeout  int64  `db:"heartbeat_timeout" bson:"heartbeat_timeout" comment:"" json:"heartbeat_timeout" xml:"heartbeat_timeout"`
	LogFile           string `db:"log_file" bson:"log_file" comment:"" json:"log_file" xml:"log_file"`
	LogWay            string `db:"log_way" bson:"log_way" comment:"console or file" json:"log_way" xml:"log_way"`
	LogLevel          string `db:"log_level" bson:"log_level" comment:"" json:"log_level" xml:"log_level"`
	LogMaxDays        uint   `db:"log_max_days" bson:"log_max_days" comment:"" json:"log_max_days" xml:"log_max_days"`
	Token             string `db:"token" bson:"token" comment:"" json:"token" xml:"token"`
	AdminAddr         string `db:"admin_addr" bson:"admin_addr" comment:"" json:"admin_addr" xml:"admin_addr"`
	AdminPort         uint   `db:"admin_port" bson:"admin_port" comment:"" json:"admin_port" xml:"admin_port"`
	AdminUser         string `db:"admin_user" bson:"admin_user" comment:"" json:"admin_user" xml:"admin_user"`
	AdminPwd          string `db:"admin_pwd" bson:"admin_pwd" comment:"" json:"admin_pwd" xml:"admin_pwd"`
	Start             string `db:"start" bson:"start" comment:"要启动的代理节点名称，留空代表全部，多个用半角逗号隔开" json:"start" xml:"start"`
	Metas             string `db:"metas" bson:"metas" comment:"meta值" json:"metas" xml:"metas"`
	Extra             string `db:"extra" bson:"extra" comment:"" json:"extra" xml:"extra"`
	Uid               uint   `db:"uid" bson:"uid" comment:"" json:"uid" xml:"uid"`
	GroupId           uint   `db:"group_id" bson:"group_id" comment:"" json:"group_id" xml:"group_id"`
	Type              string `db:"type" bson:"type" comment:"" json:"type" xml:"type"`
	Created           uint   `db:"created" bson:"created" comment:"" json:"created" xml:"created"`
	Updated           uint   `db:"updated" bson:"updated" comment:"" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingFrpClient) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFrpClient) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFrpClient) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFrpClient) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFrpClient) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFrpClient) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFrpClient) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFrpClient) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFrpClient) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFrpClient) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFrpClient) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFrpClient) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFrpClient) Objects() []*NgingFrpClient {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFrpClient) XObjects() Slice_NgingFrpClient {
	return Slice_NgingFrpClient(a.Objects())
}

func (a *NgingFrpClient) NewObjects() factory.Ranger {
	return &Slice_NgingFrpClient{}
}

func (a *NgingFrpClient) InitObjects() *[]*NgingFrpClient {
	a.objects = []*NgingFrpClient{}
	return &a.objects
}

func (a *NgingFrpClient) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFrpClient) Short_() string {
	return "nging_frp_client"
}

func (a *NgingFrpClient) Struct_() string {
	return "NgingFrpClient"
}

func (a *NgingFrpClient) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFrpClient) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFrpClient) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingFrpClient) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFrpClient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpClient(*v))
		case []*NgingFrpClient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpClient(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpClient) GroupBy(keyField string, inputRows ...[]*NgingFrpClient) map[string][]*NgingFrpClient {
	var rows Slice_NgingFrpClient
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpClient(inputRows[0])
	} else {
		rows = Slice_NgingFrpClient(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFrpClient) KeyBy(keyField string, inputRows ...[]*NgingFrpClient) map[string]*NgingFrpClient {
	var rows Slice_NgingFrpClient
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpClient(inputRows[0])
	} else {
		rows = Slice_NgingFrpClient(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFrpClient) AsKV(keyField string, valueField string, inputRows ...[]*NgingFrpClient) param.Store {
	var rows Slice_NgingFrpClient
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpClient(inputRows[0])
	} else {
		rows = Slice_NgingFrpClient(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFrpClient) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFrpClient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpClient(*v))
		case []*NgingFrpClient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpClient(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpClient) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ServerAddr) == 0 {
		a.ServerAddr = "0.0.0.0"
	}
	if len(a.TcpMux) == 0 {
		a.TcpMux = "Y"
	}
	if len(a.LoginFailExit) == 0 {
		a.LoginFailExit = "Y"
	}
	if len(a.Protocol) == 0 {
		a.Protocol = "tcp"
	}
	if len(a.LogFile) == 0 {
		a.LogFile = "console"
	}
	if len(a.LogWay) == 0 {
		a.LogWay = "console"
	}
	if len(a.LogLevel) == 0 {
		a.LogLevel = "info"
	}
	if len(a.Type) == 0 {
		a.Type = "web"
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

func (a *NgingFrpClient) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ServerAddr) == 0 {
		a.ServerAddr = "0.0.0.0"
	}
	if len(a.TcpMux) == 0 {
		a.TcpMux = "Y"
	}
	if len(a.LoginFailExit) == 0 {
		a.LoginFailExit = "Y"
	}
	if len(a.Protocol) == 0 {
		a.Protocol = "tcp"
	}
	if len(a.LogFile) == 0 {
		a.LogFile = "console"
	}
	if len(a.LogWay) == 0 {
		a.LogWay = "console"
	}
	if len(a.LogLevel) == 0 {
		a.LogLevel = "info"
	}
	if len(a.Type) == 0 {
		a.Type = "web"
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

func (a *NgingFrpClient) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFrpClient) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["server_addr"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["server_addr"] = "0.0.0.0"
		}
	}
	if val, ok := kvset["tcp_mux"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["tcp_mux"] = "Y"
		}
	}
	if val, ok := kvset["login_fail_exit"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["login_fail_exit"] = "Y"
		}
	}
	if val, ok := kvset["protocol"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["protocol"] = "tcp"
		}
	}
	if val, ok := kvset["log_file"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["log_file"] = "console"
		}
	}
	if val, ok := kvset["log_way"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["log_way"] = "console"
		}
	}
	if val, ok := kvset["log_level"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["log_level"] = "info"
		}
	}
	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "web"
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

func (a *NgingFrpClient) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.ServerAddr) == 0 {
			a.ServerAddr = "0.0.0.0"
		}
		if len(a.TcpMux) == 0 {
			a.TcpMux = "Y"
		}
		if len(a.LoginFailExit) == 0 {
			a.LoginFailExit = "Y"
		}
		if len(a.Protocol) == 0 {
			a.Protocol = "tcp"
		}
		if len(a.LogFile) == 0 {
			a.LogFile = "console"
		}
		if len(a.LogWay) == 0 {
			a.LogWay = "console"
		}
		if len(a.LogLevel) == 0 {
			a.LogLevel = "info"
		}
		if len(a.Type) == 0 {
			a.Type = "web"
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
		if len(a.ServerAddr) == 0 {
			a.ServerAddr = "0.0.0.0"
		}
		if len(a.TcpMux) == 0 {
			a.TcpMux = "Y"
		}
		if len(a.LoginFailExit) == 0 {
			a.LoginFailExit = "Y"
		}
		if len(a.Protocol) == 0 {
			a.Protocol = "tcp"
		}
		if len(a.LogFile) == 0 {
			a.LogFile = "console"
		}
		if len(a.LogWay) == 0 {
			a.LogWay = "console"
		}
		if len(a.LogLevel) == 0 {
			a.LogLevel = "info"
		}
		if len(a.Type) == 0 {
			a.Type = "web"
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

func (a *NgingFrpClient) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFrpClient) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFrpClient) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFrpClient) Reset() *NgingFrpClient {
	a.Id = 0
	a.Name = ``
	a.Disabled = ``
	a.ServerAddr = ``
	a.ServerPort = 0
	a.HttpProxy = ``
	a.PoolCount = 0
	a.TcpMux = ``
	a.User = ``
	a.DnsServer = ``
	a.LoginFailExit = ``
	a.Protocol = ``
	a.HeartbeatInterval = 0
	a.HeartbeatTimeout = 0
	a.LogFile = ``
	a.LogWay = ``
	a.LogLevel = ``
	a.LogMaxDays = 0
	a.Token = ``
	a.AdminAddr = ``
	a.AdminPort = 0
	a.AdminUser = ``
	a.AdminPwd = ``
	a.Start = ``
	a.Metas = ``
	a.Extra = ``
	a.Uid = 0
	a.GroupId = 0
	a.Type = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingFrpClient) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Name"] = a.Name
		r["Disabled"] = a.Disabled
		r["ServerAddr"] = a.ServerAddr
		r["ServerPort"] = a.ServerPort
		r["HttpProxy"] = a.HttpProxy
		r["PoolCount"] = a.PoolCount
		r["TcpMux"] = a.TcpMux
		r["User"] = a.User
		r["DnsServer"] = a.DnsServer
		r["LoginFailExit"] = a.LoginFailExit
		r["Protocol"] = a.Protocol
		r["HeartbeatInterval"] = a.HeartbeatInterval
		r["HeartbeatTimeout"] = a.HeartbeatTimeout
		r["LogFile"] = a.LogFile
		r["LogWay"] = a.LogWay
		r["LogLevel"] = a.LogLevel
		r["LogMaxDays"] = a.LogMaxDays
		r["Token"] = a.Token
		r["AdminAddr"] = a.AdminAddr
		r["AdminPort"] = a.AdminPort
		r["AdminUser"] = a.AdminUser
		r["AdminPwd"] = a.AdminPwd
		r["Start"] = a.Start
		r["Metas"] = a.Metas
		r["Extra"] = a.Extra
		r["Uid"] = a.Uid
		r["GroupId"] = a.GroupId
		r["Type"] = a.Type
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Name":
			r["Name"] = a.Name
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "ServerAddr":
			r["ServerAddr"] = a.ServerAddr
		case "ServerPort":
			r["ServerPort"] = a.ServerPort
		case "HttpProxy":
			r["HttpProxy"] = a.HttpProxy
		case "PoolCount":
			r["PoolCount"] = a.PoolCount
		case "TcpMux":
			r["TcpMux"] = a.TcpMux
		case "User":
			r["User"] = a.User
		case "DnsServer":
			r["DnsServer"] = a.DnsServer
		case "LoginFailExit":
			r["LoginFailExit"] = a.LoginFailExit
		case "Protocol":
			r["Protocol"] = a.Protocol
		case "HeartbeatInterval":
			r["HeartbeatInterval"] = a.HeartbeatInterval
		case "HeartbeatTimeout":
			r["HeartbeatTimeout"] = a.HeartbeatTimeout
		case "LogFile":
			r["LogFile"] = a.LogFile
		case "LogWay":
			r["LogWay"] = a.LogWay
		case "LogLevel":
			r["LogLevel"] = a.LogLevel
		case "LogMaxDays":
			r["LogMaxDays"] = a.LogMaxDays
		case "Token":
			r["Token"] = a.Token
		case "AdminAddr":
			r["AdminAddr"] = a.AdminAddr
		case "AdminPort":
			r["AdminPort"] = a.AdminPort
		case "AdminUser":
			r["AdminUser"] = a.AdminUser
		case "AdminPwd":
			r["AdminPwd"] = a.AdminPwd
		case "Start":
			r["Start"] = a.Start
		case "Metas":
			r["Metas"] = a.Metas
		case "Extra":
			r["Extra"] = a.Extra
		case "Uid":
			r["Uid"] = a.Uid
		case "GroupId":
			r["GroupId"] = a.GroupId
		case "Type":
			r["Type"] = a.Type
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFrpClient) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "server_addr":
			a.ServerAddr = param.AsString(value)
		case "server_port":
			a.ServerPort = param.AsUint(value)
		case "http_proxy":
			a.HttpProxy = param.AsString(value)
		case "pool_count":
			a.PoolCount = param.AsUint(value)
		case "tcp_mux":
			a.TcpMux = param.AsString(value)
		case "user":
			a.User = param.AsString(value)
		case "dns_server":
			a.DnsServer = param.AsString(value)
		case "login_fail_exit":
			a.LoginFailExit = param.AsString(value)
		case "protocol":
			a.Protocol = param.AsString(value)
		case "heartbeat_interval":
			a.HeartbeatInterval = param.AsInt64(value)
		case "heartbeat_timeout":
			a.HeartbeatTimeout = param.AsInt64(value)
		case "log_file":
			a.LogFile = param.AsString(value)
		case "log_way":
			a.LogWay = param.AsString(value)
		case "log_level":
			a.LogLevel = param.AsString(value)
		case "log_max_days":
			a.LogMaxDays = param.AsUint(value)
		case "token":
			a.Token = param.AsString(value)
		case "admin_addr":
			a.AdminAddr = param.AsString(value)
		case "admin_port":
			a.AdminPort = param.AsUint(value)
		case "admin_user":
			a.AdminUser = param.AsString(value)
		case "admin_pwd":
			a.AdminPwd = param.AsString(value)
		case "start":
			a.Start = param.AsString(value)
		case "metas":
			a.Metas = param.AsString(value)
		case "extra":
			a.Extra = param.AsString(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "type":
			a.Type = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingFrpClient) Set(key interface{}, value ...interface{}) {
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
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "ServerAddr":
			a.ServerAddr = param.AsString(vv)
		case "ServerPort":
			a.ServerPort = param.AsUint(vv)
		case "HttpProxy":
			a.HttpProxy = param.AsString(vv)
		case "PoolCount":
			a.PoolCount = param.AsUint(vv)
		case "TcpMux":
			a.TcpMux = param.AsString(vv)
		case "User":
			a.User = param.AsString(vv)
		case "DnsServer":
			a.DnsServer = param.AsString(vv)
		case "LoginFailExit":
			a.LoginFailExit = param.AsString(vv)
		case "Protocol":
			a.Protocol = param.AsString(vv)
		case "HeartbeatInterval":
			a.HeartbeatInterval = param.AsInt64(vv)
		case "HeartbeatTimeout":
			a.HeartbeatTimeout = param.AsInt64(vv)
		case "LogFile":
			a.LogFile = param.AsString(vv)
		case "LogWay":
			a.LogWay = param.AsString(vv)
		case "LogLevel":
			a.LogLevel = param.AsString(vv)
		case "LogMaxDays":
			a.LogMaxDays = param.AsUint(vv)
		case "Token":
			a.Token = param.AsString(vv)
		case "AdminAddr":
			a.AdminAddr = param.AsString(vv)
		case "AdminPort":
			a.AdminPort = param.AsUint(vv)
		case "AdminUser":
			a.AdminUser = param.AsString(vv)
		case "AdminPwd":
			a.AdminPwd = param.AsString(vv)
		case "Start":
			a.Start = param.AsString(vv)
		case "Metas":
			a.Metas = param.AsString(vv)
		case "Extra":
			a.Extra = param.AsString(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingFrpClient) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["name"] = a.Name
		r["disabled"] = a.Disabled
		r["server_addr"] = a.ServerAddr
		r["server_port"] = a.ServerPort
		r["http_proxy"] = a.HttpProxy
		r["pool_count"] = a.PoolCount
		r["tcp_mux"] = a.TcpMux
		r["user"] = a.User
		r["dns_server"] = a.DnsServer
		r["login_fail_exit"] = a.LoginFailExit
		r["protocol"] = a.Protocol
		r["heartbeat_interval"] = a.HeartbeatInterval
		r["heartbeat_timeout"] = a.HeartbeatTimeout
		r["log_file"] = a.LogFile
		r["log_way"] = a.LogWay
		r["log_level"] = a.LogLevel
		r["log_max_days"] = a.LogMaxDays
		r["token"] = a.Token
		r["admin_addr"] = a.AdminAddr
		r["admin_port"] = a.AdminPort
		r["admin_user"] = a.AdminUser
		r["admin_pwd"] = a.AdminPwd
		r["start"] = a.Start
		r["metas"] = a.Metas
		r["extra"] = a.Extra
		r["uid"] = a.Uid
		r["group_id"] = a.GroupId
		r["type"] = a.Type
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "name":
			r["name"] = a.Name
		case "disabled":
			r["disabled"] = a.Disabled
		case "server_addr":
			r["server_addr"] = a.ServerAddr
		case "server_port":
			r["server_port"] = a.ServerPort
		case "http_proxy":
			r["http_proxy"] = a.HttpProxy
		case "pool_count":
			r["pool_count"] = a.PoolCount
		case "tcp_mux":
			r["tcp_mux"] = a.TcpMux
		case "user":
			r["user"] = a.User
		case "dns_server":
			r["dns_server"] = a.DnsServer
		case "login_fail_exit":
			r["login_fail_exit"] = a.LoginFailExit
		case "protocol":
			r["protocol"] = a.Protocol
		case "heartbeat_interval":
			r["heartbeat_interval"] = a.HeartbeatInterval
		case "heartbeat_timeout":
			r["heartbeat_timeout"] = a.HeartbeatTimeout
		case "log_file":
			r["log_file"] = a.LogFile
		case "log_way":
			r["log_way"] = a.LogWay
		case "log_level":
			r["log_level"] = a.LogLevel
		case "log_max_days":
			r["log_max_days"] = a.LogMaxDays
		case "token":
			r["token"] = a.Token
		case "admin_addr":
			r["admin_addr"] = a.AdminAddr
		case "admin_port":
			r["admin_port"] = a.AdminPort
		case "admin_user":
			r["admin_user"] = a.AdminUser
		case "admin_pwd":
			r["admin_pwd"] = a.AdminPwd
		case "start":
			r["start"] = a.Start
		case "metas":
			r["metas"] = a.Metas
		case "extra":
			r["extra"] = a.Extra
		case "uid":
			r["uid"] = a.Uid
		case "group_id":
			r["group_id"] = a.GroupId
		case "type":
			r["type"] = a.Type
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFrpClient) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFrpClient) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
