@echo off
echo copy premium modules...
del /s /q _vendor\premium\*.* >nul 2>&1
rmdir /s /q _vendor\premium\ >nul 2>&1
robocopy _vendor\premium_src _vendor\premium /MIR /XD .git /s /e /NFL /NDL /NJH /NJS /nc /ns /np

echo generate bindata...
go generate -v .\vfrmap\
go generate -v .\vfrmap\html\fontawesome
go generate -v .\vfrmap\html\leafletjs
go generate -v .\vfrmap\html\maps
go generate -v .\vfrmap\html\premium
go generate -v .\simconnect\

date /t>date.txt
set /p datestr=<date.txt
del date.txt

git describe --tags>versionstr.txt
set /p versionstr=<versionstr.txt
del versionstr.txt

echo build project...
go build -o fskneeboard.exe -ldflags "-s -w -X main.buildVersion=%versionstr% -X main.buildTime=%datestr% -X main.pro=true" -v .\vfrmap\

echo cleanup...
rem del /s /q _vendor\premium\*.* >nul 2>&1
rem rmdir /s /q _vendor\premium\ >nul 2>&1