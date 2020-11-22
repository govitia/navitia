# Changelog

### Note
This repository uses `semver` management, you can read documentation here : https://semver.org/lang/

Tags are always releases to the following form : `vX.Y.Z` to match go module directives.

## 0.3.0
### Added
- Vehicle journey type and route
- Go mod file
- `headsign` filter for jourey and vehicle_journey requests
### Fixed
- Code style issues
- Gitignore missing files
- Some types were omitted but necessary

## v0.2
- **Pretty-printing !** via the `pretty` subpackage
- Paging support
- Bugfix where the response body was never closed
- Limited the size of responses
- `Coverage` has been renamed to `Regions`
- `Regions` (ex-`Coverage`), `RegionByPos` and `RegionByID` have a new parameter needed: `RegionRequest`
- No more `Session.Use`
- Un-export `RemoteErrorsDescriptions`
- PlacesResults support `sort.Interface`
- `PlacesResults` has a new method, `Count`
- No more `String` methods: use pretty !
- New `JourneyResults.Count` to count the number of journeys in the results
- And others, see `git log`
- Exported EmbeddedTypes
- Overhauled testing subsystem

