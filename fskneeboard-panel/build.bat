@ECHO off

ECHO building panel webpack...
npm run build-panel

ECHO compile panel...
"%MSFS_SDK%\Tools\bin\fspackagetool.exe" "christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.xml" -nomirroring

ECHO copy spb file...
copy /Y "christian1984-ingamepanel-fskneeboard\Build\Packages\christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.spb" "christian1984-ingamepanel-fskneeboard\InGamePanels"