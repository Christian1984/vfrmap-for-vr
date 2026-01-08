
# FSKneeboard v1.13.1 Todo List

## Feature: SimConnect Removal

- [ ] **Task 1:** Remove `go-bindata` usage and embedded `SimConnect.dll` from `fskneeboard-server/freemium_src/gosrc/simconnect/simconnect.go`.
- [ ] **Task 2:** Implement a check for `SimConnect.dll` in the server's root directory on startup.
- [ ] **Task 3:** If `SimConnect.dll` is missing, stop further startup and send a status to the UI, displaying a popup.
- [ ] **Task 4:** In the Popup, display a notification if `SimConnect.dll` is missing, with a link to the README.

## Feature: Installer Updates

- [ ] **Task 5:** Add a new wizard page to `setup/fskneeboard-free.iss` and `setup/fskneeboard-pro.iss` to prompt for `SimConnect.dll`.
- [ ] **Task 6:** Add a link to the MSFS SDK download on the new installer page.
- [ ] **Task 7:** Add a file picker to the new installer page, defaulting to `C:\MSFS 2024 SDK\SimConnect SDK\lib\SimConnect.dll`.
- [ ] **Task 8:** Implement the logic to copy the selected `SimConnect.dll` to the installation directory.

## Feature: Uninstaller Updates

- [ ] **Task 9:** Modify the uninstaller prompt in `setup/fskneeboard-free.iss` and `setup/fskneeboard-pro.iss` to include an option for removing `SimConnect.dll`.
- [ ] **Task 10:** Implement the logic to delete `SimConnect.dll` based on user choice during uninstallation.

## Feature: Documentation Updates

- [ ] **Task 11:** Update the "Prerequisites" section of `README.md` with instructions for downloading the MSFS SDK and copying `SimConnect.dll`.

## Feature: Post-Implementation Tasks

- [ ] **Task 12:** Bump the server version to `v1.13.1` in `fskneeboard-server/go.mod` (if applicable) and other relevant files.
- [ ] **Task 13:** Bump the panel version to `1.13.1` in `fskneeboard-panel/package.json`.
- [ ] **Task 14:** Update `CHANGELOG.md` with the changes for v1.13.1.
- [ ] **Task 15:** Review and finalize `README.md`.
