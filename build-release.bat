
@echo off
IF %1.==. GOTO err

echo BUILD README.pdf MANUALLY!!!
echo ============================
echo Preparing...
del /s /q dist\*.* >nul 2>&1
rmdir /s /q dist\ >nul 2>&1

echo Build Ingame Panel...
cd fskneeboard-panel\

rem call build.bat

echo Build FSKneeboard FREE...
cd ..\fskneeboard-server\

call build-fskneeboard-server-FREE.bat

echo Packaging FSKneeboard FREE...
robocopy .\ ..\dist\free\fskneeboard-server fskneeboard.exe fskneeboard-autostart-steam.bat fskneeboard-autostart-windows-store.bat /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\fskneeboard-panel\christian1984-ingamepanel-fskneeboard ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard /s /e /NFL /NDL /NJH /NJS /nc /ns /np
del /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\*.* >nul 2>&1
rmdir /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\ >nul 2>&1
del /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\index.html >nul 2>&1

echo Build FSKneeboard PRO...

call build-fskneeboard-server-PRO.bat

echo Packaging FSKneeboard PRO...
robocopy .\ ..\dist\pro\fskneeboard-server fskneeboard.exe fskneeboard-autostart-steam.bat fskneeboard-autostart-windows-store.bat /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\fskneeboard-panel\christian1984-ingamepanel-fskneeboard ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard /s /e /NFL /NDL /NJH /NJS /nc /ns /np
del /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\*.* >nul 2>&1
rmdir /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\ >nul 2>&1
del /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\index.html >nul 2>&1

robocopy .\charts\ ..\dist\pro\fskneeboard-server\charts traffic-pattern.png /NFL /NDL /NJH /NJS /nc /ns /np
mkdir ..\dist\pro\fskneeboard-server\autosave\

echo Creating hints...
echo Convert your charts to .png-files and copy them here! > ..\dist\pro\fskneeboard-server\charts\copy-your-charts-here.txt
echo Your autosaves will go here! Run 'fskneeboard --autosave 5' to create autosaves every 5 minutes... > ..\dist\pro\fskneeboard-server\autosave\autosaves-will-go-here.txt
echo Place your fskneeboard.lic license file right here inside this folder! > ..\dist\pro\fskneeboard-server\copy-your-license-file-here.txt

echo Copying README.pdf...
robocopy ..\ ..\dist\free\ README.pdf /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\ ..\dist\pro\ README.pdf /NFL /NDL /NJH /NJS /nc /ns /np

echo Copying CHANGELOG.md...
robocopy ..\ ..\dist\free\ CHANGELOG.md /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\ ..\dist\pro\ CHANGELOG.md /NFL /NDL /NJH /NJS /nc /ns /np

echo Zipping...
cd ..\dist
powershell -Command "Compress-Archive .\free\* .\fskneeboard-free-v%1.zip"
powershell -Command "Compress-Archive .\pro\* .\fskneeboard-pro-v%1.zip"

echo Build Installers...
cd ..\setup
"%programfiles(x86)%\Inno Setup 6\ISCC.exe" /Q[p] "fskneeboard-free.iss" /DApplicationVersion=%1
"%programfiles(x86)%\Inno Setup 6\ISCC.exe" /Q[p] "fskneeboard-pro.iss" /DApplicationVersion=%1

cd ..

echo BUILD FINISHED!
goto fin

:err
echo Please provide a version number for the build!

:fin

