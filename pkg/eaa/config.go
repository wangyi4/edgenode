// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
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

package eaa

import "github.com/open-ness/edgenode/pkg/util"

// CertsInfo describes paths for certs used in configuration
type CertsInfo struct {
	CaRootKeyPath  string `json:"CaRootKeyPath"`
	CaRootPath     string `json:"CaRootPath"`
	ServerCertPath string `json:"ServerCertPath"`
	ServerKeyPath  string `json:"ServerKeyPath"`
}

// Config describes EAA JSON config file
type Config struct {
	TLSEndpoint       string        `json:"TlsEndpoint"`
	OpenEndpoint      string        `json:"OpenEndpoint"`
	InternalEndpoint  string        `json:"InternalEndpoint"`
	HeartbeatInterval util.Duration `json:"HeartbeatInterval"`
	Certs             CertsInfo     `json:"Certs"`
}
