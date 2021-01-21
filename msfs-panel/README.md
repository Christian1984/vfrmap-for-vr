# msfs-panel

This is the ingame panel to actually display the VFR Map inside your Sim, particularly in VR.

The project was forked from [bymaximus/msfs2020-toolbar-window-template](https://github.com/bymaximus/msfs2020-toolbar-window-template). Please consider supporting Maximus for his ongoing efforts!

# How to build

To build the SPB if you have changed UI panel template definition run `build.bat` or manually

`SDK\Tools\bin\fspackagetool.exe christian1984-ingamepanel-vfrmapforvr\Build\christian1984-ingamepanel-vfrmapforvr.xml -nomirroring`

It will generate the SPB at `christian1984\Build\Packages\christian1984\Build` copy the SPB to `christian1984\InGamePanels`.

Copy the package to community folder BUT DO NOT COPY the `christian1984\Build` directory.