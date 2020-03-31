# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][], and this project adheres to
[Semantic Versioning][].

## Unreleased

## v1.4.2 - 2020-03-31

### Changed

- Updated to github.com/golang/protobuf v1.3.5
- Updated to github.com/micro/go-micro v2.4.0

## v1.4.1 - 2020-03-30

### Changed

- Updated to github.com/koverto/micro@v2.0.1

## v1.4.0 - 2020-03-05

### Changed

- Updated to github.com/koverto/micro@v1.2.0
- Updated to github.com/koverto/uuid@v1.3.0

## v1.3.0 - 2020-03-02

### Added

- `claims.ContextKeyJTI`
- `claims.ContextKeySUB`
- Added token ID as `Claims.ID`
- Added token expiry as `Claims.ExpiresAt`

### Changed

- Consolidated `TokenRequest` and `TokenResponse` into `Claims`
- Moved token subject to `Claims.Subject` from `Claims.UserID`

## v1.2.0 - 2020-02-22

### Added

- `claims.ContextKeyUID`

## v1.1.0 - 2020-02-22

### Added

- `Authorization.TokenResponse` now includes `userID`

## v1.0.0 - 2020-02-21

### Added

- Protobuf API
- `Authorization.Create`
- `Authorization.Validate`
- `Authorization.Invalidate`

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
