
# TODOs for FSKneeboard v1.13.0

## Installer Support for MSFS 2020 & 2024
1. [x] Research application IDs (Steam 2024 AppID 2537590) and Windows Store PFN (2020 confirmed, 2024 candidate added) plus strategy to derive Community folders via `UserCfg.opt` (see `research.md`)
2. [x] Implement installer discovery logic:
	- [x] Enumerate UserCfg.opt files for both versions (search known paths + fallback search)
	- [x] Parse InstalledPackagesPath from each UserCfg.opt and build Community folder candidates
	- [x] Deduplicate and classify candidates by version/distribution
	- [x] If one candidate: present confirmation step
	- [x] If multiple candidates: present selection UI (checkboxes) - panel installed to one selected folder, server installed once
	- [x] If none: fall back to existing hardcoded detection logic, then prompt for manual path if still none
3. [x] Update Inno Setup scripts with PowerShell / custom code to run the above discovery before file deployment

## Autostart Feature Extension
4. [x] Add new config booleans: `SteamFs2024`, `WinstoreFs2024` to `globals.go`
5. [x] Update persistence logic in `dbserversettingsmanager.go` for new booleans
6. [x] Add UI controls in `settingspanel.go` for MSFS 2024 options (Steam/Windows Store checkboxes)
7. [x] Register GUI callbacks in `main.go` and update `boolcallback.go` as needed
8. [x] Update `msfsautostart.go` to handle new launch combinations (2024 Steam AppID 2537590, 2024 Windows Store PFN)
9. [x] Update UI help text and layout for the new options

## Bing Maps Layer Replacement
10. [x] Research and select a free-to-use satellite map API (MapTiler recommended; see `research.md`)
11. [x] Implement new MapTiler satellite layer integration (API key handling, caching, attribution)
13. [x] Remove Bing Maps layer completely (code + UI + documentation)
13. [x] Update UI and documentation for map changes (add MapTiler guidance)

## Post-Implementation
14. [ ] Check for dead code!
15. [ ] Bump version numbers for server and panel to v1.13.0
16. [ ] Update README and changelog for v1.13.0
17. [ ] Announce MSFS 2024 support on the website
