
# gimme

An easy, fast and portable way to retrieve files you copy around repeatedly.

**gimme** works by traversing a list of directories looking for the file you want. As soon as the file is found, it is copied to your current working directory. In short, it is given to you.

1. [installation](#installation)
1. [setting up][setup]
1. [usage](#usage)
1. [todo](#todo)
1. [license](#license)


<a name="installation"></a>
## installation:

Using Go tools:

```bash
⇒ go get github.com/GochoMugo/gimme
```


<a name="setup"></a>
## setting up:

On *nix machines, you could add the following line to `~/.bashrc` (or equivalent file):

```sh
export GIMME_PATH="path/to/some/dir:another/path"
```

On Windows machines, you need to use the Control panel.


<a name="usage"></a>
## usage:

Getting some files fast:

```bash
⇒ gimme LICENSE-MIT .jshintrc
```

Showing help information:

```bash
⇒ gimme
```


<a name="todo"></a>
## todo:

* [ ] allow templates, compiled while giving user
* [ ] allow path prefixes, which in turn allows files with same name
* [ ] more information for Windows users
* [ ] default gimme path, thus not amust to set environment variable
* [ ] add version information


<a name="license"></a>
## license:

**The MIT License (MIT)**

Copyright (c) 2015 GochoMugo <mugo@forfuture.co.ke>

