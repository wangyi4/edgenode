#! /bin/sh

set -x

REPO=edgecontroller-vm.sh.intel.com:5000
VER=19.06

# tag images
docker tag appliance:1.0 $REPO/edgenode-appliance:$VER
docker tag edgednssvr:1.0 $REPO/edgenode-edgednssvr:$VER
docker tag nts:1.0 $REPO/edgenode-nts:$VER

# upload images
docker push $REPO/edgenode-appliance:$VER
docker push $REPO/edgenode-edgednssvr:$VER
docker push $REPO/edgenode-nts:$VER
