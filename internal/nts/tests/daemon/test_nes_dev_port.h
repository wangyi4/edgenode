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

#ifndef TEST_NES_DEV_PORT_H_
#define TEST_NES_DEV_PORT_H_

#include <CUnit/CUnit.h>

int init_suite_nes_dev_port(void);
int cleanup_suite_nes_dev_port(void);

void test_nes_dev_port_new_device(void);
void test_get_port_rings(void);
void test_scatter_port(void);
void test_scatter_eth_lbp(void);
void test_ctor_eth_port(void);
void test_add_ring_to_ntsqueue(void);
void test_dtor_port(void);

void add_nes_dev_port_suite_to_registry(void);

#endif /* TEST_NES_DEV_PORT_H_ */
