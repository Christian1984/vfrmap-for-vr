@ECHO off
ECHO This script tries to repair your FSKneeboard installation by cleaning up the database.
ECHO:
ECHO IMPORTANT: Please close FSKneeboard and Flight Simulator before you proceed!
ECHO:
ECHO Do you wish to proceed? This will reset your FSKneeboard Database and clear your Maptile Cache! (The charts folder will NOT be affected)
SET /P FSKNEEBOARD_CLEANUP="(Y)es / (N)o > "

IF "%FSKNEEBOARD_CLEANUP%" == "Y" GOTO repair
IF "%FSKNEEBOARD_CLEANUP%" == "y" GOTO repair

:abort
ECHO:
ECHO Repair script aborted!
GOTO end

:repair
del /s /q .\fskneeboard.db >nul 2>&1
del /s /q .\fskneeboard.db.lock >nul 2>&1
del /s /q .\maptilecache\*.* >nul 2>&1
rmdir /s /q .\maptilecache >nul 2>&1

ECHO:
ECHO Repair script finished! Please restart both FSKneeboard and Flight Simulator!

:end
PAUSE