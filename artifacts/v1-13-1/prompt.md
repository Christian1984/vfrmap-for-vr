For v1.13.1 of FSKneeboard, I want to implement the following changes:

1. for licensing compliance, remove the embedded simconnect.dll:
   - cleanup the entire bindata-stuff from `simconnect.go`, and stop the "automatic deployment" of `SimConnect.dll`.
   - check for the existence of SimConnect.dll and show an error via the UI if `SimConnect.dll` dows not exist (pointing the user to the readme where the download of SimConnect.dll will be explained.)

2. add a page to the installer that prompts the user to install the MSFS Developer Tools and/or point it to the location of the SimConnect.dll
   - prompt the user to download the SDK from https://sdk.flightsimulator.com/msfs2024/files/installers/1.5.7/MSFS2024_SDK_Core_Installer_1.5.7.zip
   - ... and install it to the default directory
   - add a filechooser that defaults to C:\MSFS 2024 SDK\SimConnect SDK\lib\SimConnect.dll
   - during installation, copy the SimConnect.dll to the installation directory of FSKneeboard server.

3. update the uninstaller as well:
   - uninstall SimConnect.dll only if the user choose to remove everything (just append it to the existing prompt and extend the handling accordingly).

4. extend the prerequisites section of the readme to have a section that explains downloading and installing the MSFS SDK and how to copy the `SimConnect.dll` from the SDK install directory to the (default) FSKneeboard install directory (next to fskneeboard.exe).

Post-Implementation:

- Bump the version number for both the server and panel components to v1.13.1.
- Update the README.
- Update the changelog accordingly.

For all of the above, do not start implementation right away. Start by creating an implementation plan in `artifacts/v1-13-1/implementation-plan.md` and create actionable TODOs in `artifacts/v1-13-1/todo.md`. Organize the task list by features, and add task numbers to the items in the TODO list so that I can point the implementation agent to each task (e.g. "Implement Task 11"). Use `[ ]` for each task the agent can then check.

When doing your research, create an outline of your results in `artifacts/v1-13-1/research/[task-id].md`.