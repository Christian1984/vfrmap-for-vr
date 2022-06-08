@ECHO OFF

ECHO building panel webpack...

IF %1.==dev. GOTO webpack-build-dev

:webpack-build
CALL npm run build-panel
GOTO build

:webpack-build-dev
CALL npm run build-panel-dev

:build
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO compile panel...
"%MSFS_SDK%\Tools\bin\fspackagetool.exe" "christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.xml" -nomirroring
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO copy spb file...
copy /Y "christian1984-ingamepanel-fskneeboard\Build\Packages\christian1984-ingamepanel-fskneeboard\Build\christian1984-ingamepanel-fskneeboard.spb" "christian1984-ingamepanel-fskneeboard\InGamePanels"
IF %ERRORLEVEL% NEQ 0 GOTO err

GOTO fin

:err
ECHO Build of Ingame-Panel failed! Aborting...
EXIT 1

:fin
ECHO Build of Ingame-Panel finished!