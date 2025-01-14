# Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

export PATH=$PATH:/usr/local/share/openvswitch/scripts
#remove bridges if there are any
ovs-vsctl list-br | while read line ; do ovs-vsctl del-br $line ; done

# Adding bridge
ovs-vsctl add-br $1 -- set bridge $1 datapath_type=netdev
 
mkdir /tmp/openvswitch
chown qemu:qemu /tmp/openvswitch

exit 0


