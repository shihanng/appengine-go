[![CircleCI](https://circleci.com/gh/shihanng/appengine-go/tree/master.svg?style=svg&circle-token=fe6d23ae8712e43e2d346bfeaa44b84d6a00b8cf)](https://circleci.com/gh/shihanng/appengine-go/tree/master)

# appengine-go

## Deployment

```bash
appcfg.py -A [YOUR_PROJECT_ID] -V [YOUR_VERSION_ID] update aeloader
```

## Local Development

```bash
dev_appserver.py aeloader
```

## Development Environment

We can use [direnv](https://direnv.net/) to help us setup
an environment that kind of "replace" `go` with `goapp`
from the [App Engine SDK for Go](https://cloud.google.com/appengine/docs/standard/go/download#appengine_sdk).

In our `.envrc` we can write the following:

```bash
export PATH=<path-to-app-engine-sdk-for-go>/go_appengine
export PATH=<path-to-alias>/go:$PATH

export GOROOT=<path-to-app-engine-sdk-for-go>/go_appengine/goroot-1.8
```

Here is the content of `go`

```sh
#!/bin/sh
# Remove gohack from path and run goapp.
PATH=${PATH#$HOME/bin/gae_go} goapp $@
```

Basically you can put `go` wherever you like as long as the
proper path is provided to `.envrc` as the above.
With this trick, tests within the editor should run without problem.
