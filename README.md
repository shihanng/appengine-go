[![CircleCI](https://circleci.com/gh/shihanng/appengine-go/tree/master.svg?style=svg&circle-token=fe6d23ae8712e43e2d346bfeaa44b84d6a00b8cf)](https://circleci.com/gh/shihanng/appengine-go/tree/master)

# appengine-go

## Project Structure

- `GOPATH` starts from `.`.
- The "main"/"driver" of the app is in `aeloader`.
- The other packages of the app live in `applib`.
- External packages are vendored with [`dep`](https://github.com/golang/dep).

```
src
└── app
    ├── aeloader
    ├── applib
    │   ├── guestbook
    │   └── sign
    └── vendor
```

## Deployment

```bash
appcfg.py -A [YOUR_PROJECT_ID] -V [YOUR_VERSION_ID] update src/app/aeloader
```

## Local Development

```bash
dev_appserver.py src/app/aeloader
```

## Development Environment

We can use [direnv](https://direnv.net/) to help us setup
an environment that kind of "replace" `go` with `goapp`
from the [App Engine SDK for Go](https://cloud.google.com/appengine/docs/standard/go/download#appengine_sdk).

In our `.envrc` we can have something like the following:

```bash
export PATH=`pwd`/bin:$HOME/bin/fakego:$HOME/bin/go_appengine:$PATH
export GOROOT=$HOME/bin/go_appengine/goroot-1.8
export GOPATH=`pwd`
```

In `$HOME/bin/fakego` we have a shell script name `go`:

```sh
#!/bin/sh
# Remove gohack from path and run goapp.
PATH=${PATH#$HOME/bin/go_appengine} goapp $@
```

just to tell our editor that we are using `goapp` version of `go`.
