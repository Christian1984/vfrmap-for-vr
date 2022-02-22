# FSKneeboard Ingame Panel

This is the ingame panel to actually display the VFR Map inside MSFS, particularly in VR.

The project was forked from [bymaximus/msfs2020-toolbar-window-template](https://github.com/bymaximus/msfs2020-toolbar-window-template). Please consider supporting Maximus for his ongoing efforts!

# How to Build

To build or rebuild the SPB after changing the UI panel template definition run `build.bat` or manually run

`SDK\Tools\bin\fspackagetool.exe christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.xml -nomirroring`

This will generate the SPB at `christian1984-ingamepanel-fskneeboard\Build\Packages\christian1984\Build`. Copy the SPB to `christian1984-ingamepanel-fskneeboard\InGamePanels`.

To deploy the panel to MSFS, copy the package to the MSFS community folder BUT DO NOT COPY the `christian1984-ingamepanel-fskneeboard\Build` directory.