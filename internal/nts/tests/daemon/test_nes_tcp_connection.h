/*******************************************************************************
* Copyright 2019 Intel Corporation. All rights reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*******************************************************************************/

#ifndef TEST_NES_TCP_CONNECTION_H_
#define TEST_NES_TCP_CONNECTION_H_

#include <CUnit/CUnit.h>
#include <sys/socket.h>
#include <sys/un.h>
#include "ctrl/nes_tcp_connection.h"

int init_suite_nes_tcp_connection(void);
int cleanup_suite_nes_tcp_connection(void);

void add_nes_tcp_connection_suite_to_registry(void);

#endif /* TEST_NES_TCP_CONNECTION_H_ */
