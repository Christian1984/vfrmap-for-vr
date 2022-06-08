
@ECHO OFF
IF %1.==. GOTO err_version_missing

ECHO Preparing...
del /s /q dist\*.* >nul 2>&1
rmdir /s /q dist\ >nul 2>&1

ECHO Build Readme...
CALL npm run build-readme
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO Build Ingame Panel...
cd fskneeboard-panel\

CALL build.bat
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO Build FSKneeboard FREE...
cd ..\fskneeboard-server\

CALL build-fskneeboard-server-FREE.bat gui
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO Packaging FSKneeboard FREE...
robocopy .\ ..\dist\free\fskneeboard-server fskneeboard.exe repair-fskneeboard.bat /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\fskneeboard-panel\christian1984-ingamepanel-fskneeboard ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard /s /e /NFL /NDL /NJH /NJS /nc /ns /np
del /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\*.* >nul 2>&1
rmdir /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\ >nul 2>&1
del /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\index.html >nul 2>&1

ECHO Build FSKneeboard PRO...

CALL build-fskneeboard-server-PRO.bat gui
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO Packaging FSKneeboard PRO...
robocopy .\ ..\dist\pro\fskneeboard-server fskneeboard.exe repair-fskneeboard.bat /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\fskneeboard-panel\christian1984-ingamepanel-fskneeboard ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard /s /e /NFL /NDL /NJH /NJS /nc /ns /np
del /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\*.* >nul 2>&1
rmdir /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\ >nul 2>&1
del /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\index.html >nul 2>&1

robocopy .\charts\ ..\dist\pro\fskneeboard-server\charts traffic-pattern.png /NFL /NDL /NJH /NJS /nc /ns /np
robocopy .\charts\approach ..\dist\pro\fskneeboard-server\charts\approach MDW.png /NFL /NDL /NJH /NJS /nc /ns /np
robocopy .\charts\weather ..\dist\pro\fskneeboard-server\charts\weather weather_forecast_chart.png /NFL /NDL /NJH /NJS /nc /ns /np
mkdir ..\dist\pro\fskneeboard-server\autosave\

ECHO Creating hints...
ECHO Convert your charts to .png-files and copy them here! > ..\dist\pro\fskneeboard-server\charts\copy-your-charts-here.txt
ECHO Your autosaves will go here! Run 'fskneeboard --autosave 5' to create autosaves every 5 minutes... > ..\dist\pro\fskneeboard-server\autosave\autosaves-will-go-here.txt
ECHO Place your fskneeboard.lic license file right here inside this folder! > ..\dist\pro\fskneeboard-server\copy-your-license-file-here.txt

ECHO Copying README.pdf...
robocopy ..\ ..\dist\free\ README.pdf /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\ ..\dist\pro\ README.pdf /NFL /NDL /NJH /NJS /nc /ns /np

ECHO Copying CHANGELOG.md...
robocopy ..\ ..\dist\free\ CHANGELOG.md /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\ ..\dist\pro\ CHANGELOG.md /NFL /NDL /NJH /NJS /nc /ns /np

ECHO Zipping...
cd ..\dist
powershell -Command "Compress-Archive .\free\* .\fskneeboard-free-v%1.zip"
IF %ERRORLEVEL% NEQ 0 GOTO err

powershell -Command "Compress-Archive .\pro\* .\fskneeboard-pro-v%1.zip"
IF %ERRORLEVEL% NEQ 0 GOTO err

ECHO Build Installers...
cd ..\setup
"%programfiles(x86)%\Inno Setup 6\ISCC.exe" /Q[p] "fskneeboard-free.iss" /DApplicationVersion=%1
IF %ERRORLEVEL% NEQ 0 GOTO err

"%programfiles(x86)%\Inno Setup 6\ISCC.exe" /Q[p] "fskneeboard-pro.iss" /DApplicationVersion=%1
IF %ERRORLEVEL% NEQ 0 GOTO err

cd ..

GOTO fin

:err_version_missing
ECHO Please provide a version number for the build (e.g. 1.8.0)! Aborting...
EXIT 1

:err
ECHO Build failed: Something went wrong! Aborted!
EXIT 1

:fin
ECHO BUILD FINISHED!