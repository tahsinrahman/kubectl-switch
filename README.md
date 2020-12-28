### Kubectl Switch
A simple kubectl plugin to switch between clusters.
Supports fuzzy searching for cluster names.

#### Installing
If you have `go` installed on your local machine
```
$ make build
```

Otherwise, you can build it inside docker container.
You only need to have `make` and `docker`.
By default it builds binary for macos.
```
$ make build-docker
```

But you can configure it by specifying your desired os and architecture.
https://golang.org/doc/install/source#environment
```bash
$ make build-docker OS=linux ARCH=amd64
```

#### Usage
```bash
$ kubectl switch

# or specify the cluster name
$ kubectl switch <cluster-name>
```