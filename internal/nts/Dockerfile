# Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM centos:7.6.1810

ENV http_proxy "http://child-prc.intel.com:913"
ENV https_proxy "http://child-prc.intel.com:913"
ENV ftp_proxy "http://child-prc.intel.com:913"
ENV no_proxy "127.0.0.1,localhost,.intel.com"

RUN echo "proxy=http://child-prc.intel.com:913" >> /etc/yum.conf


RUN yum upgrade ca-certificates && \
    yum install -y epel-release && \
    yum install -y numactl-devel libhugetlbfs-utils iproute python python-pip

RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple configparser
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple docker

#unset proxy
ENV http_proxy=
ENV https_proxy=
ENV ftp_proxy=
ENV no_proxy=

ADD ./nts/nes-daemon /root/
ADD ./nts/kni_docker_daemon.py /root/
ADD ./nts/ovs_docker_daemon.py /root/
ADD ./nts/entrypoint.sh /root/
ADD ./nts/libnes_api_shared.so /root/

ENTRYPOINT ["/root/entrypoint.sh"]
