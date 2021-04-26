
@echo off
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
robocopy .\ ..\dist\free\fskneeboard-server fskneeboard.exe /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\fskneeboard-panel\christian1984-ingamepanel-fskneeboard ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard /s /e /NFL /NDL /NJH /NJS /nc /ns /np
del /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\*.* >nul 2>&1
rmdir /s /q ..\dist\free\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\ >nul 2>&1

echo Build FSKneeboard PRO...

call build-fskneeboard-server-PRO.bat

echo Packaging FSKneeboard PRO...
robocopy .\ ..\dist\pro\fskneeboard-server fskneeboard.exe /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\fskneeboard-panel\christian1984-ingamepanel-fskneeboard ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard /s /e /NFL /NDL /NJH /NJS /nc /ns /np
del /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\*.* >nul 2>&1
rmdir /s /q ..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\Build\ >nul 2>&1

robocopy .\charts\ ..\dist\pro\fskneeboard-server\charts traffic-pattern.png /NFL /NDL /NJH /NJS /nc /ns /np
echo Convert your charts to .png-files and copy them here! > ..\dist\pro\fskneeboard-server\charts\copy-your-charts-here.txt
echo Place your fskneeboard.lic license file right here inside this folder! > ..\dist\pro\fskneeboard-server\copy-your-license-file-here.txt

echo Copy README.pdf
robocopy ..\ ..\dist\free\ README.pdf /NFL /NDL /NJH /NJS /nc /ns /np
robocopy ..\ ..\dist\pro\ README.pdf /NFL /NDL /NJH /NJS /nc /ns /np


echo BUILD FINISHED!

