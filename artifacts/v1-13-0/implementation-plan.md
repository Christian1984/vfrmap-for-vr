# Implementation Plan for FSKneeboard v1.13.0

## 1. Installer Support for MSFS 2020 and MSFS 2024
- Research application IDs and default install directories for MSFS 2024 (Windows Store and Steam). (See `research.md` Task 1 findings.)
- Compare with existing MSFS 2020 detection logic.
- Update installer logic to:
  - Detect community folders (scan for `UserCfg.opt` and derive from there; do this for the existing hardcoded logic as well)
  - If only one community folder is found, proceed with this and display the existing confirmation screen.
  - If multiple are found, allow user to select one (or multiple) for installation of the panel component.
  - If none are found, prompt user for Community Folder location.
- Update Inno Setup scripts accordingly.

## 2. Autostart Feature Extension
- Research application IDs and install directories for MSFS 2024 (Windows Store and Steam).
- Add a new dropdown in the UI: "Flight Simulator" (options: MSFS 2020, MSFS 2024).
- Place this above the existing "Flight Simulator Version" dropdown.
- Update logic to process both dropdowns and deduce the correct CLI command for launching the selected simulator/version/store.
- Update documentation and help text as needed.

## 3. Bing Maps Layer Replacement
- Deactivate Bing Maps layer (service shutdown).
- Research and select a free-to-use map API with satellite imagery (e.g., MapTiler, OpenStreetMap, etc.). (See `research.md` Task 8 findings.)
- Implement support for the new map API, allowing users to provide their own API key if required.
- Update UI and documentation to reflect the change.

## 4. Post-Implementation Tasks
- Bump version numbers for server and panel to v1.13.0.
- Update README and changelog.
- Announce MSFS 2024 support on the website.
