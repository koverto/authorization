# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][], and this project adheres to
[Semantic Versioning][].

## Unreleased

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
