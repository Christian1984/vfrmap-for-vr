@ECHO off
SET noguiflag=
IF %1.==gui. SET noguiflag=-H=windowsgui
IF %1.==dev. GOTO build

ECHO copy premium modules...
del /s /q _vendor\premium\*.* >nul 2>&1
rmdir /s /q _vendor\premium\ >nul 2>&1
npm run build-server-pro
robocopy _vendor\premium_src\gosrc _vendor\premium /MIR /XD .git /s /e /NFL /NDL /NJH /NJS /nc /ns /np

:build
ECHO generate bindata...
go generate -v .\vfrmap\
go generate -v .\vfrmap\server
go generate -v .\vfrmap\html\fontawesome
go generate -v .\vfrmap\html\leafletjs
go generate -v .\vfrmap\html\freemium
go generate -v .\vfrmap\html\premium
go generate -v .\vfrmap\gui\res
go generate -v .\simconnect\

date /t>date.txt
SET /p datestr=<date.txt
del date.txt

git describe --tags>versionstr.txt
SET /p versionstr=<versionstr.txt
del versionstr.txt

ECHO create winres meta data...
cd vfrmap
go-winres make --product-version=git-tag --file-version=git-tag
cd ..

ECHO build project...
go build -o fskneeboard.exe -ldflags "-s -w -X main.buildVersion=%versionstr% -X main.buildTime=%datestr% -X main.pro=true %noguiflag%" -v .\vfrmap\

ECHO cleanup...
REM del /s /q _vendor\premium\*.* >nul 2>&1
REM rmdir /s /q _vendor\premium\ >nul 2>&1