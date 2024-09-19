# QRCode Info Reader

The `qrinfo` command outputs the contents of a QR code along with its encoding parameters.

## Install

Direct downloads are available through the [releases page](https://github.com/rs/qrinfo/releases/latest).

Using [homebrew](http://brew.sh/) on macOS (Go not required):

```
brew install rs/tap/qrinfo
```

From source:

```
go install github.com/rs/jaggr@latest
```

## Usage

The `qrinfo` command supports multiple methods for providing the QR code image data:

### Command Line Argument

Pass the path to the image as the first and only argument:

```
qrinfo /path/to/image
```

### Standard Input

Send the image data via standard input:

```
qrinfo < /path/to/image

some-command | qrinfo
```

### Clipboard

If `qrinfo` is executed without arguments and no piped input, it will attempt to fetch image data from your clipboard, if available.
