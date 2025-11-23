For v1.13.0 of FSKneeboard, I want to implement the following changes:

1. The FSKneeboard Installer should support both Flight Simulator 2020 and Flight Simulator FS 2024. It should scan for existing installations of both application and then react to the result:

  - The installer found only ONE installation of either MSFS 2020 OR MSFS 2024: Proceed with the default directories for the installation that was detected (still show the existing confirmation step)
  - The installer found installations of each MSFS version: Allow the user to install FSKneeboard for MSFS 2020, MSFS 2024 or both.
  - The installer could not detect an installation automatically: Prompt user for MSFS version and Community Folder location.

  AI Agent plan:

  - research application IDs and default install directories for MSFS 2024, use the existing MSFS 2020 data as reference (Windows Store and Steam)
  - line out an implementation plan for adding the above mentioned script to the inno setup scripts.

2. The autostart feature should be extended so that it can start MSFS 2020 and MSFS 2024.

  AI Agent plan:

  - research application IDs and default install directories for MSFS 2024, use the existing MSFS 2020 data as reference (Windows Store and Steam)
  - add another dropdown "Flight Simulator" above the existing dropdown "Flight Simulator Version" with options: "MSFS 2020" and "MSFS 2024"
  - extend the existing code to process the input of these two fields and deduce the CLI command that needs to be run accordingly for "MSFS 2024 Steam", "MSFS 2024 Windows Store" and "MSFS 2024 Steam", "MSFS 2020 Windows Store"
  - line out an implementation plan.

3. The Bing Maps layer must be deactivated (as the service was shut down by MS) or, better yet, be replaced by a free to use map API that has satellite images. It's okay if the user has to sign up and bring his own api key, as was implemented with the existing Bing Maps layer.

Post-Implementation:

- Bump the version number for both the server and panel components to v1.13.0.
- Update the README.
- Update the changelog accordingly.

For all of the above, do not start implementation right away. Start by creating an implementation plan in `artifacts/v1-13-0/implementation-plan.md` and create actionable TODOs in `artifacts/v1-13-0/todo.md`. Organize the task list by features, and add task numbers to the items in the TODO list so that I can point the implementation agent to each task (e.g. "Implement Task 11"). Use `[ ]` for each task the agent can then check.

When doing your research, create an outline of your results in `artifacts/v1-13-0/research/[task-id].md`.

# Note to self:

- Update Website to highlight MSFS 2024 support.