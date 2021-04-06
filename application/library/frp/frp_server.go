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

package frp

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/admpub/confl"
	"github.com/webx-top/com"

	"github.com/admpub/nging/application/dbschema"

	_ "github.com/admpub/frp/assets/frpc/statik"
	_ "github.com/admpub/frp/assets/frps/statik"
	"github.com/admpub/frp/pkg/config"
	frpLog "github.com/admpub/frp/pkg/util/log"
	"github.com/admpub/frp/pkg/util/util"
	"github.com/admpub/frp/server"
)

func SetServerConfigFromDB(conf *dbschema.NgingFrpServer) *config.ServerCommonConf {
	c := config.GetDefaultServerConf()
	c.BindAddr = conf.Addr
	c.BindPort = int(conf.Port)
	c.BindUDPPort = int(conf.UdpPort)
	c.KCPBindPort = int(conf.KcpPort)
	c.ProxyBindAddr = conf.ProxyAddr
	c.VhostHTTPPort = int(conf.VhostHttpPort)
	c.VhostHTTPTimeout = int64(conf.VhostHttpTimeout)
	if c.VhostHTTPTimeout < 1 {
		c.VhostHTTPTimeout = 60
	}
	c.VhostHTTPSPort = int(conf.VhostHttpsPort)

	c.DashboardAddr = conf.DashboardAddr
	c.DashboardPort = int(conf.DashboardPort)
	c.DashboardUser = conf.DashboardUser
	c.DashboardPwd = conf.DashboardPwd
	if conf.LogWay == `console` {
		conf.LogFile = `console`
	}
	c.LogFile = conf.LogFile
	c.LogWay = conf.LogWay
	c.LogLevel = conf.LogLevel
	c.LogMaxDays = int64(conf.LogMaxDays)
	c.Token = conf.Token
	c.SubDomainHost = conf.SubdomainHost
	c.MaxPortsPerClient = int64(conf.MaxPortsPerClient)
	c.TCPMux = conf.TcpMux == `Y`

	// e.g. 1000-2000,2001,2002,3000-4000
	ports, _ := util.ParseRangeNumbers(conf.AllowPorts)
	for _, port := range ports {
		c.AllowPorts[int(port)] = struct{}{}
	}
	return &c
}

func StartServerByConfigFile(filePath string, pidFile string) error {
	ext := filepath.Ext(filePath)
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	switch strings.ToLower(ext) {
	case `.json`:
		r := &dbschema.NgingFrpServer{}
		err = json.Unmarshal(b, r)
		if err != nil {
			return err
		}
		c := SetServerConfigFromDB(r)
		return StartServer(pidFile, c)
	case `.yaml`:
		r := &dbschema.NgingFrpServer{}
		err = confl.Unmarshal(b, r)
		if err != nil {
			return err
		}
		c := SetServerConfigFromDB(r)
		return StartServer(pidFile, c)
	default:
		content := string(b)
		return StartServerByConfig(content, pidFile)
	}
}

func StartServerByConfig(configContent string, pidFile string) error {
	cfg, err := config.UnmarshalServerConfFromIni(configContent)
	if err != nil {
		return err
	}
	return StartServer(pidFile, &cfg)
}

func StartServer(pidFile string, c *config.ServerCommonConf) error {
	once.Do(onceInit)
	err := c.Validate()
	if err != nil {
		return err
	}
	frpLog.InitLog(c.LogWay,
		c.LogFile,
		c.LogLevel,
		c.LogMaxDays, true)
	if len(pidFile) > 0 {
		err := com.WritePidFile(pidFile)
		if err != nil {
			frpLog.Error(err.Error())
			return err
		}
	}
	svr, err := server.NewService(*c)
	if err != nil {
		return err
	}
	frpLog.Info("Start frps success")
	svr.Run()
	return err
}
