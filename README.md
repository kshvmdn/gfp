## gfp (GitHub follower path)

> Find a path of followers between two GitHub users.

### Contents

- [Demo](#demo)
- [Installation](#installation)
- [Setup](#setup)
- [Usage](#usage)
- [Contribute](#contribute)
- [Credits](#credits)
- [License](#license)

### Demo

_Coming soon._

### Installation

You should have Go [installed](https://golang.org/doc/install#download) and [configured](https://golang.org/doc/install#testing).

Install with `go get`:

```sh
$ go get -u -v github.com/kshvmdn/gfp/...
$ which gfp
$GOPATH/bin/gfp
```

Or, install directly via source:

```sh
$ git clone https://github.com/kshvmdn/gfp.git $GOPATH/src/github.com/kshvmdn/gfp
$ cd $_
$ make install && make
$ ./gfp
```

### Setup

gfp uses the [GitHub API](https://api.github.com), so you'll be required to register an application. 

Head over to [this page](https://github.com/settings/tokens) and [generate a new token](https://github.com/settings/tokens/new). Give it a name (anything works) and hit `Generate token` (no scope required).

Copy the token and export it as an environment variable.

```sh
$ export GITHUB_ACCESS_TOKEN=<access token>
```

### Usage

gfp takes two GitHub usernames: an origin and a target.

The program starts at the _origin_ user and recursively builds a directed graph based on who this user is following. This process is repeated until the _target_ is reached (or the API rate limit is exceeded).

You can generally expect a different path each run (particularly for _popular_ accounts), this is because the traversal order is unpredicatable (we're running everything concurrently, so it's not possible to control order).

gfp is built on top of a concurrent job queue with 6 workers by default. You can change this number with the `-workers` flag.

Log output is redirected to `/dev/null`, use the `-show-log` flag to disable this.

View the usage dialogue with the `-help` flag.

```sh
$ gfp -help
usage: gfp [options] origin target
  -show-log
        show log output
  -version
        print version and exit
  -workers int
        number of workers (default 6)
```

#### Examples

```sh
$ gfp kshvmdn torvalds
kshvmdn -> ... -> torvalds
```

```sh
$ gfp -show-log kshvmdn torvalds
2017/05/28 05:53:09 worker 1: kshvmdn
2017/05/28 05:53:10 worker 4: ...
2017/05/28 05:53:10 worker 5: ...
2017/05/28 05:53:10 worker 2: ...
2017/05/28 05:53:10 worker 6: ...
...
kshvmdn -> ... -> torvalds
```

### Contribute

gfp is completely open source. Feel free to open an [issue](https://github.com/kshvmdn/gfp/issues) or a [pull request](https://github.com/kshvmdn/gfp/pulls).

Prior to submitting work, please ensure your changes comply with [Golint](https://github.com/golang/lint), test this with `make lint`.

#### Todo

- [ ] Find shortest path, might be impossible to find the shortest path (due to rate limits), but it _is_ possible to do something like: run until we find x paths and return the shortest of those x.
- [ ] I think we can do some cool stuff with the resultant graph. Perhaps add a flag to export it once the computation is complete (maybe do something with [awalterschulze/gographviz](https://github.com/awalterschulze/gographviz)?).
- [ ] Better-handle rate limit (perhaps it's possible to pause execution and wait until more requests are available?).
- [ ] Handle the case where we run out of users to traverse and haven't found the target (also, similar: origin user is following 0 people).

### Credits

- Idea inspired by [ryanmcdermott/spotifind](https://github.com/ryanmcdermott/spotifind).
- Job queue based on [this article](http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html).

### License

gfp source code is released under the [MIT license](./LICENSE).
