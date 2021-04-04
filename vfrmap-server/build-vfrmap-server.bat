go generate -v .\vfrmap\
go generate -v .\vfrmap\html\leafletjs
go generate -v .\vfrmap\html\premium
go generate -v .\simconnect\

@echo off
date /t>date.txt
set /p datestr=<date.txt
del date.txt

git describe --tags>versionstr.txt
set /p versionstr=<versionstr.txt
del versionstr.txt
@echo on

go build -o vfrmap-for-vr.exe -ldflags "-s -w -X main.buildVersion=%versionstr% -X main.buildTime=%datestr%" -v .\vfrmap\