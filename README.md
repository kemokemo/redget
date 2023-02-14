# redget

A simple Redmine API client tool to get a issue's description.

## Install

### Homebrew

```sh
brew install kemokemo/tap/redget
```

### Scoop

First, add my scoop-bucket.

```sh
scoop bucket add kemokemo-bucket https://github.com/kemokemo/scoop-bucket.git
```

Next, install this app by running the following.

```sh
scoop install redget
```

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/redget/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH`.

## Preparation

Please set the two `env` below:

- `REDMINE_URL`: the root URL of the Redmine
- `REDMINE_KEY`: the API access key of the above Redmine

## Usage

```sh
$ redget -n 12345
# you can get the description of the issue #12345 here.
```

