![Tripletex Golang Client][banner]

![Build & Deploy][build-badge]
[![Total alerts][lgtm-badge]][lgtm-alerts]
[![Maintainability][codeclimate-badge]][codeclimate]

## About

The Tripletex Client for Golang simplifies integrations against [Tripletex 2.0 API][tripletex-docs]. The client is generated using [swagger-go][swagger-go].

# Documentation

Apart from this README, you can find details and examples of using the SDK in
the following places:

- [SDK Documentation][sdk-doc]
- [API Documentation][tripletex-docs]

You can find an example of how this library is used at [cobraz/jira-to-tripletex][jira-to-tripletex].

## Contributing

We love contributions! 🙏 Bug reports and pull requests are welcome on [GitHub][github].

## Known issues

Tripletex defines consumers with `application/json; charset=utf-8` a few places. This is not according to spesifications of ordinary Swagger files, and not something [swagger-go][swagger-go] supports straight out-of-the-box. Therefore, you'll have to define this as a separate `Producer`. You can follow this [issue here](https://github.com/bjerkio/tripletex-go/issues/1)

## Missing features?

Due to Tripletex not having implemented `operationId` in their Swagger files, we are not automatically updating this library based on changes in their Swagger file. This is a temporary situation, and as soon as their Swagger file includes this we will update automatically. This will mean that this library will automatically publish new versions of the client at automatically when Tripletex releases new functionality to their API.

[banner]: ./.github/header.png
[build-badge]: https://github.com/bjerkio/tripletex-go/workflows/build/badge.svg
[lgtm-badge]: https://img.shields.io/lgtm/alerts/g/bjerkio/tripletex-go.svg?logo=lgtm&logoWidth=18
[lgtm-alerts]: https://lgtm.com/projects/g/bjerkio/tripletex-go/alerts/
[codeclimate-badge]: https://api.codeclimate.com/v1/badges/807d8fb5210aa406713d/maintainability
[codeclimate]: https://codeclimate.com/github/bjerkio/tripletex-go/maintainability
[tripletex-docs]: https://tripletex.no/v2-docs/#/
[swagger-go]: https://github.com/go-swagger/go-swagger
[jira-to-tripletex]: https://github.com/cobraz/jira-to-tripletex
[github]: https://github.com/bjerkio/tripletex-go
[sdk-doc]: https://pkg.go.dev/mod/github.com/bjerkio/tripletex-go
