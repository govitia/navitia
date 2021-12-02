# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

# [Unreleased]

## [2.0.0] - 2021-12-01
### Added
- api types support for V2
- go.mod file for V2 (supporting go 1.17)

# [Released]

## [0.4.1] - 2020-12-02
### Changed
- Changelog move to a separate file from the README.md
### Added
- CHANGELOG.md

## [0.4.0] - 2020-11-24
### Fixed
- Coding style on many files
- go.mod file

## [0.3.0] & [0.3.1] - 2020-11-10
### Added
- Vehicle journey type and route
- Go mod file
- `headsign` filter for jourey and vehicle_journey requests
### Fixed
- Code style issues
- Gitignore missing files
- Some types were omitted but necessary

## [0.2.0] - 2017-04-27
### Added
- **Pretty-printing !** via the `pretty` subpackage
- Paging support
- Limited the size of responses
- PlacesResults support `sort.Interface`
- Un-export `RemoteErrorsDescriptions`
- `PlacesResults` has a new method, `Count`
- New `JourneyResults.Count` to count the number of journeys in the results
- And others, see `git log`
- Exported EmbeddedTypes
- Overhauled testing subsystem
### Changed
- `Coverage` has been renamed to `Regions`
- `Regions` (ex-`Coverage`), `RegionByPos` and `RegionByID` have a new parameter needed: `RegionRequest`
### Removed
- No more `Session.Use`
- No more `String` methods: use pretty !
### Fixed
- Bugfix where the response body was never closed