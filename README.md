F5 Kubernetes BIG-IP Controller
===============================

The F5 BIG-IP Controller for [Kubernetes](http://kubernetes.io/) makes F5 BIG-IP
[Local Traffic Manager](<https://f5.com/products/big-ip/local-traffic-manager-ltm)
services available to applications running in Kubernetes.

Documentation
-------------

For instruction on how to use this component, see the
[F5 Kubernetes BIG-IP Controller docs](http://clouddocs.f5.com/products/connectors/k8s-bigip-ctlr/latest/).

For guides on this and other solutions for Kubernetes, see the
[F5 Kubernetes Solution Guides](http://clouddocs.f5.com/containers/latest/kubernetes).


Running
-------

The official docker image is `f5networks/k8s-bigip-ctlr`.

Usually, the controller is deployed in Kubernetes. However, the controller can be run locally for development testing.

```shell
docker run f5networks/k8s-bigip-ctlr /app/bin/k8s-bigip-ctlr <args>
```


Building
--------

The official images are built using docker, but the adventurous can use standard go build tools. 

### Official Build

Prerequisites:
- Docker

```shell
git clone https://github.com/f5networks/k8s-bigip-ctlr.git
cd  k8s-bigip-ctlr

# Use docker to build the release artifacts, into a local "_docker_workspace" directory, then put into docker iamges
make prod
```


### Alternate, unofficial build

A normal go and godep toolchain can be used as well

Prerequisites:
- go 1.7
- $GOPATH pointing at a valid go workspace
- godep (Only needed to modify vendor's packages)
- python
- virtualenv

```shell
mkdir -p $GOPATH/src/github.com/F5Networks
cd $GOPATH/src/github.com/F5Networks
git clone https://github.com/f5networks/k8s-bigip-ctlr.git
cd k8s-bigip-ctlr

# Build all packages, and run unit tests
make all test
```

To make changes to vendor dependencies, see [Devel](DEVEL.md)
