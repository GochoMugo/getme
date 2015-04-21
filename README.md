
# getme

[![Build Status](https://travis-ci.org/GochoMugo/getme.svg)](https://travis-ci.org/GochoMugo/getme)

An easy, fast and portable way to retrieve files somewhere.

**getme** works by traversing a list of directories looking for the file you want. As soon as the file is found, it is copied to your current working directory. It can also retrieve the file from a remote source, such as Github. In short, it is given to you.

1. [features](#features)
1. [installation](#installation)
1. [usage](#usage)
1. [configuration](#configuration)
1. [utilities](#utils)
1. [license](#license)


<a name=”features”></a>
## features:

* [ ] compiles templates with custom values from you
* [ ] allows path prefixes, which in turn allows files with same name
* [ ] no configuration
* [ ] versioned
* [ ] programmatic adding of predefined paths e.g. Dropbox folder
* [ ] aware of hosted git repos e.g. Github
* [ ] works [on Windows](#windows) too. :sweat:


<a name="installation"></a>
## installation:

Using Go tools:

```bash
⇒ go get github.com/GochoMugo/getme
```


<a name="usage"></a>
## usage:

### local (getting files from your disk):

```bash
⇒ getme local LICENSE-MIT .jshintrc
```

This will make `getme` look for the files *LICENSE-MIT* and *.jshintrc* in the configured template directories. See [adding template directories](#dir)

### remote:

#### getting files from Github:

```bash
⇒ getme remote --github GochoMugo/getme master lib/getme.go
```

*GochoMugo/getme* is a **Github shorthand**, *master* is the repo's **branch**, *lib/getme.go* is the **path to the target file** from the root of the repo.

### help information:

```bash
⇒ getme help
```


<a name="configuration"></a>
## configuration:

`getme` can be used as is, without any configurations. It will automatically look for files in directory **.getme** in your Home directory.

<a name="dir"></a>
### add template directories:

You may want it to look for files in other directories. This is possible using the `$GETME_PATH` variable.

On *nix machines, you could add the following line to `~/.bashrc` (or equivalent file):

```sh
export GETME_PATH="path/to/some/dir:another/path"
```

On Windows machines, you need to use the Control panel.


<a name="utils"></a>
## utilities:

To increase productivity (every boss likes that), I've added some utilities to the [utils](https://github.com/GochoMugo/getme/tree/master/utils) directory that can be quite useful. These utilities may not be as portable as the `getme` program itself but we shall try our best.


<a name="license"></a>
## license:

**The MIT License (MIT)**

Copyright (c) 2015 GochoMugo <mugo@forfuture.co.ke>

