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

## 5. Open Questions Requiring Clarification

### 5.1 Installer Scope & Backwards Compatibility
- **Question**: Should the new installer logic **replace** the existing Community folder detection entirely, or should it be **additive** (keeping current MSFS 2020 detection as fallback)?
- **Current state**: The existing installer already has hardcoded paths for MSFS 2020 Windows Store and Steam (`Microsoft.FlightSimulator_8wekyb3d8bbwe` and `Microsoft.FlightDashboard_8wekyb3d8bbwe`)
- **Impact**: This affects whether we refactor the existing logic or add parallel detection
- **Decision**: Use existing detection logic as fallback! If multiple community folders where detected, the user can pick any number of these with checkboxes. The panel should be installed in each of the selected community folders, while the server should only be installed once.

### 5.2 MSFS 2024 Windows Store PFN Verification
- **Question**: The research shows `Microsoft.Limitless_8wekyb3d8bbwe` as the candidate PFN for MSFS 2024, but this needs confirmation.
- **Current state**: We have a PowerShell verification command, but haven't run it on a machine with MSFS 2024 installed
- **Impact**: Without the correct PFN, we can't reliably detect MSFS 2024 Windows Store installations
- **Decision**: We use this for now. My research and communication with MSFS 2024 owners indicates this information is reliable.

### 5.3 Autostart Feature Integration Points
- **Question**: Where exactly in the codebase is the autostart feature implemented? Need to locate the current UI and backend logic.
- **Current state**: We found `msfsautostart.go` with hardcoded Steam commands, but haven't located the UI components
- **Impact**: Need to understand the current architecture before adding the new dropdown
- **Decision**: That is where the logic resides, correct. Config booleans are located in `globals.go` as `SteamFs` (for existing MSFS 2020 Steam launcher), `WinstoreFs` (for existing MSFS 2020 Windows Store launcher) and `MsfsAutostart` (for enabling the feature globally). I'm okay with just adding `SteamFs2024` and `WindtoreFs2024` for now. Persistence of these booleans is implemented in `dbserversettingsmanager.go`. The UI is implemented in `settingspanel.go`. Also register GUI-Callbacks in `main.go` Look for other occurences across the codebase, e.g. `boolcallback.go`.

### 5.4 Bing Maps Removal Strategy
- **Question**: Should we completely remove Bing Maps code/UI, or just disable it with a deprecation notice?
- **Current state**: Found multiple Bing references in README and codebase
- **Impact**: Affects user migration experience and potential rollback scenarios
- **Decision**: Step 1: Remove it. Bing Maps are dead! Step 2: Implement MapTiler.

### 5.5 Version Numbering & Release Strategy
- **Question**: Should this be released as v1.13.0 for both server and panel components simultaneously, or can they be versioned independently?
- **Current state**: Both components seem to share version numbers based on existing patterns
- **Impact**: Affects build scripts and compatibility matrix
- **Decision**: They are tightly coupled. v.1.13.0 is good for both!
