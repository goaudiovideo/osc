# osc
Go-based Open Sound Control (OSC) Library

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE]


## Overview

[osc][] provides a Go-based [Open Sound Control (OSC)][osc.org] library. In
addition to being able to create and send user specified OSC messages, [osc][]
provides a Go library for specific audio equipment.

### Supported Equipment

[osc][] (partially) supports the following equipment:

- Behringer X32 Digital Mixer

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check
```

To update and view the test coverage report:

```bash
$ make cover
```

## License

[osc][] is released under the MIT license. Please see the
[LICENSE][] file for more information.

[osc]: https://github.com/goaudiovideo/osc
[osc.org]: http://opensoundcontrol.org/
[godoc badge]: https://godoc.org/github.com/goaudiovideo/osc?status.svg
[godoc link]: https://godoc.org/github.com/goaudiovideo/osc
[LICENSE]: https://github.com/goaudiovideo/osc/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/goaudiovideo/osc
[report card]: https://goreportcard.com/report/github.com/goaudiovideo/osc
