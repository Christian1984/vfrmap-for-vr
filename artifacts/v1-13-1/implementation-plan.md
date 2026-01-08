# FSKneeboard v1.13.1 Implementation Plan

This document outlines the plan for implementing the changes for FSKneeboard v1.13.1.

## 1. Remove Embedded `SimConnect.dll`

The current implementation embeds `SimConnect.dll`. This will be removed and replaced with a check and user guidance.

- **Analysis:** The `simconnect.go` file uses `go-bindata` to embed the DLL. This will be removed.
- **Server-Side Changes:**
    - The server will check for `SimConnect.dll` in its root directory upon startup.
    - If the DLL is not found, a blocking popup will notify the user to first download and properly place SimConnect.dll.
    - The popup should contain a link to the appropriate readme section (e.g. https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md#downloading-simconnect)

## 2. Installer Updates

The Inno Setup installer scripts (`fskneeboard-free.iss` and `fskneeboard-pro.iss`) will be updated to handle the `SimConnect.dll` dependency.

- **New Installer Page:** A new wizard page will be added to the installer.
- **User Prompt:** This page will inform the user about the need for `SimConnect.dll` and prompt them to download and install the MSFS SDK from the official source: `https://sdk.flightsimulator.com/msfs2024/files/installers/1.5.7/MSFS2024_SDK_Core_Installer_1.5.7.zip`.
- **File Picker:** A file input field will allow the user to point to the location of `SimConnect.dll`. This field will default to `C:\MSFS 2024 SDK\SimConnect SDK\lib\SimConnect.dll`.
- **File Copy:** The installer will copy the user-selected `SimConnect.dll` into the FSKneeboard installation directory.

## 3. Uninstaller Updates

The uninstaller will be modified to handle the removal of `SimConnect.dll`.

- **Conditional Deletion:** The uninstaller already asks the user if they want to remove all settings and data. This prompt will be extended to include `SimConnect.dll`.
- **Implementation:** The Inno Setup script will be modified to delete `SimConnect.dll` from the installation directory only if the user confirms the complete removal.

## 4. Documentation Updates

The `README.md` file will be updated to reflect these changes.

- **Prerequisites Section:** A new section will be added under "Prerequisites".
- **Instructions:** This section will provide clear, step-by-step instructions for:
    1. Downloading and installing the MSFS SDK.
    2. Locating the `SimConnect.dll` within the SDK installation folder.
    3. Manually copying `SimConnect.dll` to the FSKneeboard server installation directory if they skipped the step during installation.

## 5. Post-Implementation Tasks

After the core features are implemented, the following tasks will be completed.

- **Versioning:**
    - The version in `fskneeboard-server/go.mod` will be updated.
    - The version in `fskneeboard-panel/package.json` will be updated to `1.13.1`.
- **Changelog:** `CHANGELOG.md` will be updated with the changes for version 1.13.1.
- **README:** The `README.md` will be reviewed and updated to ensure it is consistent with the new functionality.