
# TODOs for FSKneeboard v1.13.0

## Installer Support for MSFS 2020 & 2024
1. [x] Research application IDs (Steam 2024 AppID 2537590) and Windows Store PFN (2020 confirmed, 2024 candidate added) plus strategy to derive Community folders via `UserCfg.opt` (see `research.md`)
2. [ ] Implement installer discovery logic:
	- [ ] Enumerate UserCfg.opt files for both versions (search known paths + fallback search)
	- [ ] Parse InstalledPackagesPath from each UserCfg.opt and build Community folder candidates
	- [ ] Deduplicate and classify candidates by version/distribution
	- [ ] If one candidate: present confirmation step
	- [ ] If multiple candidates: present selection UI (checkboxes) for 2020 / 2024
	- [ ] If none: prompt user for manual Community folder path & version selection
3. [ ] Update Inno Setup scripts with PowerShell / custom code to run the above discovery before file deployment

## Autostart Feature Extension
4. [ ] Add a new "Flight Simulator" dropdown above the existing "Flight Simulator Version" dropdown (options: MSFS 2020, MSFS 2024)
5. [ ] Extend code to process both dropdowns and deduce the correct CLI command for all combinations (MSFS 2020/2024, Steam/Windows Store)
6. [ ] Update UI and help text for the new autostart options

## Bing Maps Layer Replacement
7. [ ] Deactivate Bing Maps layer (service shutdown)
8. [x] Research and select a free-to-use satellite map API (MapTiler recommended; see `research.md`)
9. [ ] Implement new MapTiler satellite layer integration (API key handling, caching, attribution)
10. [ ] Update UI and documentation for map changes (remove Bing instructions, add MapTiler guidance)

## Post-Implementation
11. [ ] Check for dead code!
12. [ ] Bump version numbers for server and panel to v1.13.0
13. [ ] Update README and changelog for v1.13.0
14. [ ] Announce MSFS 2024 support on the website
