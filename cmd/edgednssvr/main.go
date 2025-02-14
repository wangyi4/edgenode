// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	logger "github.com/open-ness/common/log"
	edgedns "github.com/open-ness/edgenode/pkg/edgedns"
	"github.com/open-ness/edgenode/pkg/edgedns/grpc"
	"github.com/open-ness/edgenode/pkg/edgedns/storage"
	"github.com/open-ness/edgenode/pkg/util"
)

var log = logger.DefaultLogger.WithField("main", nil)

func main() {
	logLvl := flag.String("log", "info", "Log level.\nSupported values: "+
		"debug, info, notice, warning, error, critical, alert, emergency")
	syslogAddr := flag.String("syslog", "", "Syslog address")
	v4 := flag.String("4", "", "IPv4 listener address")
	port := flag.Int("port", 5053, "listener UDP port")
	sock := flag.String("sock", "/run/edgedns.sock", "API socket path")
	db := flag.String("db", "/var/lib/edgedns/rrsets.db",
		"Database file path")
	fwdr := flag.String("fwdr", "8.8.8.8", "Forwarder")
	hbInterval := flag.Int("hb", 60, "Heartbeat interval in s")
	flag.Parse()

	lvl, err := logger.ParseLevel(*logLvl)
	if err != nil {
		log.Errf("Failed to parse log level: %s", err.Error())
		os.Exit(1)
	}
	logger.SetLevel(lvl)

	err = logger.ConnectSyslog(*syslogAddr)
	if err != nil {
		log.Errf("Syslog(%s) connection failed: %s", *syslogAddr, err.Error())
		os.Exit(1)
	}

	sockPath := path.Dir(*sock)
	if _, err = os.Stat(sockPath); os.IsNotExist(err) {
		err = os.MkdirAll(sockPath, 0750)
		if err != nil {
			log.Err(err)
			os.Exit(1)
		}
	}

	dbPath := path.Dir(*db)
	if _, err = os.Stat(dbPath); os.IsNotExist(err) {
		err = os.MkdirAll(dbPath, 0750)
		if err != nil {
			log.Err(err)
			os.Exit(1)
		}
	}

	cfg := edgedns.Config{
		Addr4: *v4,
		Port:  *port,
	}

	stg := &storage.BoltDB{
		Filename: *db,
	}

	ctl := &grpc.ControlServer{
		Sock: *sock,
	}

	srv := edgedns.NewResponder(cfg, stg, ctl)
	srv.SetDefaultForwarder(*fwdr)
	err = srv.Start()
	defer srv.Stop()

	if err != nil {
		log.Err(err)
		os.Exit(1)
	}

	// Heartbeat routine
	var interval util.Duration
	interval.Duration = time.Second * time.Duration(*hbInterval)
	util.Heartbeat(context.Background(), interval, func() {
		// TODO: implementation of modules checking
		log.Info("Heartbeat")
	})

	// Receive OS signals and listener errors from Start()
	signal.Notify(srv.Sig, syscall.SIGINT, syscall.SIGTERM)
	sig := <-srv.Sig
	switch sig {
	case syscall.SIGCHLD:
		log.Err("Child lisenter/service unexpectedly died")
	default:
		log.Infof("Signal (%v) received, shutting down", sig)
	}
}
