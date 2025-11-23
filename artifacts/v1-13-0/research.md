# Research Results for v1.13.0

## Task 1: MSFS 2024 Application IDs & Default Install / Data Directories

Goal: Enable installer + autostart logic to reliably detect MSFS 2020 and MSFS 2024 (Steam + Microsoft Store) and derive Community folder paths.

### 1.1 Known (Confirmed) MSFS 2020 Data
Steam App ID (2020): 1250410
Microsoft Store Package Family Name (PFN 2020): Microsoft.FlightSimulator_8wekyb3d8bbwe

Executable locations (launch targets):
- Steam (default): C:\Program Files (x86)\Steam\steamapps\common\MicrosoftFlightSimulator\FlightSimulator.exe
- MS Store: Executable resides in a WindowsApps folder with versioned subfolder; direct launching should use the shell protocol or app execution alias rather than hardcoding path. A reliable start method uses: "shell:AppsFolder\\Microsoft.FlightSimulator_8wekyb3d8bbwe!App" or the existing current approach you use.

User data & configuration:
- MS Store UserCfg.opt: %LOCALAPPDATA%\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\LocalCache\UserCfg.opt
- Steam UserCfg.opt: %APPDATA%\Microsoft Flight Simulator\UserCfg.opt

Community / Official base path derivation:
- Parse last (or matching) line in UserCfg.opt: InstalledPackagesPath "<ABSOLUTE_PATH>". Community folder = <ABSOLUTE_PATH>\Community.

### 1.2 Confirmed & Pending Identifiers
Confirmed (Steam) App ID MSFS 2024: 2537590 (from Steam store URL: https://store.steampowered.com/app/2537590/Microsoft_Flight_Simulator_2024/)

Pending Confirmation (Microsoft Store PFN): Microsoft.Limitless_8wekyb3d8bbwe

PowerShell verification command (run on target machine):
   Get-AppxPackage -Name *FlightSimulator* | Select Name, PackageFullName, InstallLocation

If two packages appear, distinguish 2020 vs 2024 by Name or InstallLocation metadata (2024 should contain "2024"). Capture the exact PFN string returned and substitute everywhere below.

### 1.3 Proposed Strategy for MSFS 2024 Detection
We must not guess static paths. Instead, perform multi-step discovery for both versions:

1. Detect Microsoft Store packages:
   - Run (Inno Setup script via PowerShell / registry): Enumerate AppxPackages filtering DisplayName or PFN containing "FlightSimulator".
   - Expected PFNs (2024 likely): Microsoft.FlightSimulator2024_8wekyb3d8bbwe or Microsoft.FlightSimulator.XYZ_8wekyb3d8bbwe. (Exact PFN to be confirmed.)
   - If two distinct FlightSimulator PFNs found, treat them as 2020 and 2024. Use version numbers or package display name to distinguish (DisplayName should contain year).

2. Detect Steam installations:
   - Steam library folders: Read Steam config VDF file: C:\Program Files (x86)\Steam\steamapps\libraryfolders.vdf
   - For each library: Check for app manifests:
     - appmanifest_1250410.acf (MSFS 2020)
     - appmanifest_<APPID_2024>.acf (MSFS 2024). (APPID for 2024 to be confirmed; placeholder APPID_2024.)
   - If manifest present, read "installdir" field to build path: <library>/steamapps/common/<installdir>/FlightSimulator.exe.

3. Distinguish versions:
   - If both a 2020 PFN and a 2024 PFN (or both Steam app manifests) found, prompt user to select installation targets (2020, 2024, or both).
   - If only one variant found, proceed with confirmation step.
   - If none found, ask user to specify MSFS version and Community folder path, then validate by checking existence of UserCfg.opt and FlightSimulator.exe.

### 1.4 Community Folder Path Resolution for 2024
Use same approach as 2020: locate UserCfg.opt for each install.
Heuristic search order (per version):
1. MSFS 2024 Known location (Windows Store): %LOCALAPPDATA%\Packages\Microsoft.Limitless_8wekyb3d8bbwe\LocalCache\UserCfg.opt
2. MSFS 2024 Known location (Steam): %APPDATA%\Microsoft Flight Simulator 2024\UserCfg.opt
3. MSFS 2020 Known location (Windows Store): %LOCALAPPDATA%\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\LocalCache\UserCfg.opt
4. MSFS 2020 Known location (Steam 1): %LOCALAPPDATA%\Packages\Microsoft.FlightDashboard_8wekyb3d8bbwe\LocalCache\UserCfg.opt
6. MSFS 2020 Known location (Steam 2): %APPDATA%\Microsoft Flight Simulator\UserCfg.opt
7. Fallback: Search fixed drive roots for file pattern UserCfg.opt limited to depth 4 with candidate names.
8. Parse InstalledPackagesPath.

Validation: Confirm Community folder contains at least one of: "Community" directory exists and maybe "Official" directory sibling.

### 1.5 Registry / Package Queries (Implementation Notes)
- Microsoft Store listing via PowerShell: Get-AppxPackage -Name *FlightSimulator* | Select Name, InstallLocation.
- Steam app IDs via registry (unreliable for path): HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\Steam App <APPID>.
- Use Inno Setup [Code] section with ShellExec or custom PowerShell script execution for complex discovery if needed.

### 1.6 Autostart Launch Commands (Preliminary for 2024)
- Steam: steam://run/<APPID>
- Microsoft Store: explorer.exe shell:AppsFolder\\<PFN>!App
(Confirm actual PFN & AppID for 2024 before finalizing.)

### 1.7 Outstanding Unknowns to Confirm
- Final PFN for MSFS 2024
- Final Steam App ID for MSFS 2024
Once confirmed, substitute placeholders PFN_2024 and APPID_2024 everywhere.

### 1.8 Risks & Mitigations
- Risk: Path or naming changes -> Mitigation: Use dynamic discovery (enumerate packages & Steam manifests).
- Risk: User custom install path -> Mitigation: Always parse UserCfg.opt InstalledPackagesPath.
- Risk: Multiple library folders -> Mitigation: Iterate all libraryfolders.vdf entries.
- Risk: Performance of drive search if metadata missing -> Mitigation: Limit depth & file pattern, prompt user sooner.

### 1.9 Data Structure Suggestion (Installer Script)
Define a record per detected install:
Record fields: Version (2020|2024), Distribution (Steam|Store), RootPath, UserCfgPath, PackagesPath, CommunityPath, LaunchCommand, Status (Detected|UserProvided|Invalid).

## Task 8: Free Satellite Map API Options & Recommendation

Goal: Replace Bing Maps satellite imagery with a sustainable alternative requiring user API key (if necessary).

### 8.1 Candidate Services Overview
1. MapTiler Satellite
   - API Key: Required (Free tier: limited monthly tiles, generous for hobby use)
   - URL template: https://api.maptiler.com/tiles/satellite/{z}/{x}/{y}.jpg?key=KEY
   - Attribution: "© MapTiler © OpenStreetMap contributors" + data sources.
   - Pros: Simple integration (raster tiles), good global coverage.
   - Cons: Rate limits; must enforce key usage.
2. Mapbox Satellite
   - API Key (token) required; Pricing after free tier; Terms restrict exposing raw tiles.
   - More complex licensing; avoid for broad distribution without user-managed account.
3. HERE Maps
   - Requires app_id/app_code or API key; Freemium tier with limits; Terms may restrict caching.
4. Esri World Imagery (ArcGIS)
   - Requires API key for production; License restrictions + usage tracking.
5. Google Maps
   - Not allowed for direct tile use in Leaflet without JS API; pricing & legal friction.
6. OpenAerialMap / NASA / Sentinel
   - Open data but not a direct plug-in global, consistent tiled high-res source for easy drop-in.

### 8.2 Selection Rationale
Recommend MapTiler Satellite as primary replacement:
- Clear API.
- Acceptable free tier for typical usage (users can upgrade individually).
- Straightforward attribution.
- Works with existing Leaflet stack (similar to Bing integration pattern).

### 8.3 Integration Plan
1. Configuration:
   - New setting: MAPTILER_API_KEY (user-provided; store alongside existing API keys).
   - UI: Replace Bing Maps section with "Satellite (MapTiler)" instructions; if key absent, show disabled state.
2. Layer Definition (Leaflet):
   - tileLayer("https://api.maptiler.com/tiles/satellite/{z}/{x}/{y}.jpg?key=" + key, { maxZoom: 20, attribution: '© MapTiler © OpenStreetMap contributors' })
3. Graceful Fallback:
   - If key missing: Do not register layer OR register with message overlay prompting to add key.
4. Caching Strategy:
   - Reuse existing maptilecache mechanism (if generic) else create per-provider subdirectory (maptilecache/maptiler/...). Respect provider ToS (avoid long-term offline replay if disallowed).
5. Attribution & License:
   - Ensure attribution string appears on map when layer active.
6. API Key Validation:
   - Optional: On first use, test one known tile (z=1,x=0,y=0). If HTTP 4xx, mark key invalid.

### 8.4 Data Structure
Extend existing map layer registry with fields:
- id: "satellite-maptiler"
- provider: "MapTiler"
- requiresKey: true
- active: boolean
- urlTemplate: template
- attribution: string
- maxZoom

### 8.5 Deactivation of Bing Maps
Steps:
1. Remove Bing layer registration & UI controls.
2. Migrate any Bing-specific settings -> Archive section in config (so upgrade does not crash).
3. If user previously supplied Bing key, show notice about removal & alternative.

### 8.6 Edge Cases
- Invalid API key -> Show status indicator + disable layer.
- Rate limit exceeded -> Detect HTTP 429 responses; surface warning.
- Offline mode -> Display cached tiles only; if none, fall back to default non-satellite layer.

### 8.7 Testing Plan
- Unit: Key parsing & layer registration conditional logic.
- Integration: Launch map with and without key; verify attribution; tile fetch success.
- Performance: Confirm tile caching reduces subsequent load times.

### 8.8 Future Enhancements
- Allow multiple satellite providers selectable by user.
- Add dynamic quality selection (lower-res tiles for VR performance).

## 9. Next Actions Triggered by This Research
- Replace placeholders PFN_2024 and APPID_2024 once confirmed.
- Implement detection functions (Task 2/3 follow-up).
- Implement MapTiler integration (Task 9 onward).

## 10. Items Requiring Confirmation
Please provide (or approve acquisition method for) actual MSFS 2024 Steam App ID and Microsoft Store PFN to finalize detection & autostart commands.
