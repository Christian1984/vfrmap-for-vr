@ECHO off

ECHO building panel webpack...

IF %1.==dev. GOTO webpack-build-dev

:webpack-build
call npm run build-panel
GOTO build

:webpack-build-dev
call npm run build-panel-dev

:build
ECHO compile panel...
"%MSFS_SDK%\Tools\bin\fspackagetool.exe" "christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.xml" -nomirroring

ECHO copy spb file...
copy /Y "christian1984-ingamepanel-fskneeboard\Build\Packages\christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.spb" "christian1984-ingamepanel-fskneeboard\InGamePanels"