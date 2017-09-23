# Version 0.4.7

[Documentation](README.md)

## Changelog since 0.4.6

### Codebase

- Added default context encoding

## Changelog since 0.4.5

### Codebase

- Added ability to set app name in make commands [#5](https://github.com/k8s-community/k8sapp/pull/5), [@rumyantseva](https://github.com/rumyantseva))

## Changelog since 0.4.4

### Configuration

- Added k8s-community infrastructure configuration

## Changelog since 0.4.3

### Configuration

- Used separate infrastructure and kube-context variables [#4](https://github.com/k8s-community/k8sapp/pull/4), [@rumyantseva](https://github.com/rumyantseva))

## Changelog since 0.4.2

### Codebase

- Fixed certs directory inside a container

## Changelog since 0.4.1

### Codebase

- Fixed of creating certificate in non-existing directory

## Changelog since 0.4.0

### Codebase

- Implemented http.ResponseWriter interface in BitRoute Control

## Changelog since 0.3.5

### Codebase

- Added info handler which contains detailed statistics with HTTP codes
- Added system signals handling (reload, shutdown, maintenance)
- Used system signal handling and info handler

## Changelog since 0.3.4

### Configuration

- Fixed image registry tag in respect that every account represent as namespace

## Changelog since 0.3.3

### Configuration

- CPU/Memory limits was decreased down to 10m/30Mi

## Changelog since 0.3.2

### Configuration

- Changed configuration of Makefile/replicas related with namespace for `k8s-community` accounts

## Changelog since 0.3.1

### Documentation

- Changes related with account `k8s-community`: [#1](https://github.com/k8s-community/k8sapp/pull/1), [@takama](https://github.com/takama))

### Configuration

- Changed configuration of Makefile related with `k8s-community` account and registry [#1](https://github.com/k8s-community/k8sapp/pull/1), [@takama](https://github.com/takama))

## Changelog since 0.3.0

### Configuration

- Added Helm charts ([#29](https://github.com/takama/k8sapp/pull/29), [@takama](https://github.com/takama))
- Added deployment functionality in commands of Makefile  ([#29](https://github.com/takama/k8sapp/pull/29), [@takama](https://github.com/takama))

## Changelog since 0.2.2

### Documentation

- Added usage description of the loggers: [xlog](https://github.com/rs/xlog), [logrus](https://github.com/sirupsen/logrus)

### Codebase

- Added routers interface ([#22](https://github.com/takama/k8sapp/pull/22), [@takama](https://github.com/takama))
- Implemented Bit-Route interface  ([#23](https://github.com/takama/k8sapp/pull/23), [@takama](https://github.com/takama))
- Implemented httprouter interface ([#24](https://github.com/takama/k8sapp/pull/24), [@takama](https://github.com/takama))
- Added environment configuration ([#26](https://github.com/takama/k8sapp/pull/26), [@takama](https://github.com/takama))
- Refactoring of the packages relations ([#25](https://github.com/takama/k8sapp/pull/25), [@takama](https://github.com/takama))
- Added health/ready handlers ([#27](https://github.com/takama/k8sapp/pull/27), [@takama](https://github.com/takama))


## Changelog since 0.2.1

### Tests

- Added Travis CI [travis-ci.org](https://travis-ci.org/takama/k8sapp)
- Added code coverage data/bot [codecov.io](https://codecov.io/gh/takama/k8sapp)

## Changelog since 0.2.0

### Tests

- Simplify standard logger tests

## Changelog since 0.1.0

### Configuration

- Added Makefile ([#15](https://github.com/takama/k8sapp/pull/15), [@takama](https://github.com/takama))
- Added Dockerfile ([#15](https://github.com/takama/k8sapp/pull/15), [@takama](https://github.com/takama))
- Added dep package manager ([#15](https://github.com/takama/k8sapp/pull/15), [@takama](https://github.com/takama))
- Added SSL certificates ([#15](https://github.com/takama/k8sapp/pull/15), [@takama](https://github.com/takama))

### Codebase

- Added base application ([#13](https://github.com/takama/k8sapp/pull/13), [@takama](https://github.com/takama))
- Added Logger interface ([#14](https://github.com/takama/k8sapp/pull/14), [@takama](https://github.com/takama))
- Added XLog Logger ([#14](https://github.com/takama/k8sapp/pull/14), [@takama](https://github.com/takama))
- Added logrus Logger ([#14](https://github.com/takama/k8sapp/pull/14), [@takama](https://github.com/takama))

## Changelog since 0.0.0

### Documentation

- Added description of the application main criteria
- Added [Roadmap](https://github.com/takama/k8sapp/wiki/Roadmap)
