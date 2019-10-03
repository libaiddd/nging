/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package file

import (
	"fmt"
	"strings"

	"github.com/admpub/nging/application/dbschema"
	"github.com/admpub/nging/application/library/fileupdater"
	"github.com/admpub/nging/application/model/base"
	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"
)

func NewEmbedded(ctx echo.Context, fileMdls ...*File) *Embedded {
	var fileM *File
	if len(fileMdls) > 0 {
		fileM = fileMdls[0]
	} else {
		fileM = NewFile(ctx)
	}
	return &Embedded{
		FileEmbedded: &dbschema.FileEmbedded{},
		base:         base.New(ctx),
		File:         fileM,
	}
}

type Embedded struct {
	*dbschema.FileEmbedded
	base    *base.Base
	File    *File
	updater *fileupdater.FileUpdater
}

func (f *Embedded) Updater() *fileupdater.FileUpdater {
	if f.updater == nil {
		f.updater = fileupdater.New(f.base.Context, f)
	}
	return f.updater
}

// DeleteByTableID 删除嵌入文件
func (f *Embedded) DeleteByTableID(project string, table string, tableID uint64) error {
	_, err := f.ListByOffset(nil, nil, 0, -1, db.And(
		db.Cond{`table_id`: tableID},
		db.Cond{`table_name`: table},
		db.Cond{`project`: project},
	))
	if err != nil {
		return err
	}
	var ids []uint64
	for _, row := range f.Objects() {
		err = f.File.SetField(nil, `used_times`, db.Raw(`used_times-1`), db.And(
			db.Cond{`used_times`: db.Gt(0)},
			db.Cond{`id`: db.In(strings.Split(row.FileIds, `,`))},
		))
		if err != nil {
			return err
		}
		ids = append(ids, row.Id)
	}
	if len(ids) > 0 {
		err = f.Delete(nil, db.Cond{`id`: db.In(ids)})
	}
	return err
}

func (f *Embedded) UpdateByFileID(project string, table string, field string, tableID uint64, fileID uint64) error {
	err := f.File.UpdateUnrelation(project, table, field, tableID, fileID)
	if err != nil {
		return err
	}
	err = f.File.Incr(fileID)
	if err != nil {
		return err
	}
	m := &dbschema.FileEmbedded{}
	err = m.Get(nil, db.And(
		db.Cond{`table_id`: tableID},
		db.Cond{`table_name`: table},
		db.Cond{`field_name`: field},
	))
	if err != nil {
		if err != db.ErrNoMoreRows {
			return err
		}
		m.Reset()
		m.FieldName = field
		m.TableName = table
		m.Project = project
		m.TableId = tableID
		m.FileIds = fmt.Sprint(fileID)
		_, err = m.Add()
	}
	return err
}

func (f *Embedded) UpdateEmbedded(embedded bool, project string, table string, field string, tableID uint64, fileIds ...interface{}) error {

	err := f.File.UpdateUnrelation(project, table, field, tableID, fileIds...)
	if err != nil {
		return err
	}

	m := &dbschema.FileEmbedded{}
	err = m.Get(nil, db.And(
		db.Cond{`table_id`: tableID},
		db.Cond{`table_name`: table},
		db.Cond{`field_name`: field},
	))
	if err != nil {
		if err != db.ErrNoMoreRows {
			return err
		}
		if len(fileIds) < 1 {
			return nil
		}
		// 不存在时，添加
		m.Reset()
		m.FieldName = field
		m.TableName = table
		m.Project = project
		m.TableId = tableID
		if embedded {
			m.Embedded = `Y`
		} else {
			m.Embedded = `N`
		}
		m.FileIds = ""
		err = f.File.Incr(fileIds...)
		if err != nil {
			return err
		}
		for _, v := range fileIds {
			m.FileIds += fmt.Sprintf("%v,", v)
		}
		m.FileIds = strings.TrimSuffix(m.FileIds, ",")
		_, err = m.Add()
		return err
	}
	if len(fileIds) < 1 { // 删除关联记录
		err = f.Delete(nil, `id`, m.Id)
		if err != nil {
			return err
		}
	}
	var fidsString string
	fidList := make([]string, len(fileIds))
	for k, v := range fileIds {
		s := fmt.Sprint(v)
		fidList[k] = s
		fidsString += s + `,`
	}
	fidsString = strings.TrimSuffix(fidsString, ",")
	if m.FileIds == fidsString {
		return nil
	}
	ids := strings.Split(m.FileIds, ",")
	var (
		delIds []interface{}
		newIds []interface{}
	)
	//已删除引用
	for _, v := range ids {
		if !com.InSlice(v, fidList) {
			delIds = append(delIds, v)
		}
	}
	//新增引用
	for _, v := range fidList {
		if !com.InSlice(v, ids) {
			newIds = append(newIds, v)
		}
	}
	if len(delIds) > 0 {
		err := f.File.Decr(delIds...)
		if err != nil {
			return err
		}
		err = f.File.SetFields(nil, echo.H{
			`table_id`:   0,
			`table_name`: ``,
			`field_name`: ``,
		}, db.Cond{`used_times`: 0})
		if err != nil {
			return err
		}
	}
	if len(newIds) > 0 {
		err := f.File.Incr(newIds...)
		if err != nil {
			return err
		}
	}
	m.FileIds = fidsString
	err = f.SetField(nil, `file_ids`, m.FileIds, db.Cond{`id`: m.Id})
	return err
}

// RelationEmbeddedFiles 关联嵌入的文件
// @param project 项目名称
// @param table 表名称
// @param field 被嵌入的字段名
// @param tableID 表中行主键ID
// @param v 内容
// @return
// @author AdamShen <swh@admpub.com>
func (f *Embedded) RelationEmbeddedFiles(project string, table string, field string, tableID uint64, v string) error {
	var (
		files []interface{}
		fids  []interface{} //旧文件ID
	)
	EmbeddedRes(v, func(file string, fid int64) {
		var exists bool
		if fid > 0 {
			exists = com.InSliceIface(fid, fids)
		} else {
			exists = com.InSliceIface(file, files)
		}
		if exists {
			return
		}
		files = append(files, file)
		if fid > 0 {
			fids = append(fids, fid)
		}
	})
	if len(fids) < 1 && len(files) > 0 {
		fids = f.File.GetIDByViewURLs(files)
	}
	err := f.UpdateEmbedded(true, project, table, field, tableID, fids...)
	return err
}

func (f *Embedded) RelationFiles(project string, table string, field string, tableID uint64, v string, seperator ...string) error {
	var (
		files []interface{}
		fids  []interface{} //旧文件ID
	)
	RelatedRes(v, func(file string, fid int64) {
		var exists bool
		if fid > 0 {
			exists = com.InSliceIface(fid, fids)
		} else {
			exists = com.InSliceIface(file, files)
		}
		if exists {
			return
		}
		files = append(files, file)
		if fid > 0 {
			fids = append(fids, fid)
		}
	})
	if len(fids) < 1 && len(files) > 0 {
		fids = f.File.GetIDByViewURLs(files)
	}
	err := f.UpdateEmbedded(false, project, table, field, tableID, fids...)
	return err
}