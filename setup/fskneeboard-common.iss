procedure CurUninstallStepChanged(CurUninstallStep: TUninstallStep);
var
  AppFolder: String;
  MaptileFolder: String;
  ChartsFolder: String;
  AutosaveFolder: String;
  LogsFolder: String;
  PdfImporterFolder: String;
  DbFile: String;
  DbLockFile: String;
  LicenseFile: String;
begin
  if CurUninstallStep = usPostUninstall then
  begin
    AppFolder := ExpandConstant('{app}');
    MaptileFolder := ExpandConstant('{app}\maptilecache');
    ChartsFolder := ExpandConstant('{app}\charts');
    AutosaveFolder := ExpandConstant('{app}\autosave');
    LogsFolder := ExpandConstant('{app}\logs');
    PdfImporterFolder := ExpandConstant('{app}\pdf-importer');
    DbFile := ExpandConstant('{app}\fskneeboard.db');
    DbLockFile := ExpandConstant('{app}\fskneeboard.db.lock');
    LicenseFile := ExpandConstant('{app}\fskneeboard.lic');

    if MsgBox('Do you want to REMOVE all FSKneeboard data (including your license)? (Hint: Answer "NO!" if you are just updating to a new version of FSKneeboard!)',
      mbConfirmation, MB_YESNO) = IDYES
    then begin
      if DirExists(MaptileFolder) then
        DelTree(MaptileFolder, True, True, True);
      if DirExists(ChartsFolder) then
        DelTree(ChartsFolder, True, True, True);
      if DirExists(AutosaveFolder) then
        DelTree(AutosaveFolder, True, True, True);
      if DirExists(LogsFolder) then
        DelTree(LogsFolder, True, True, True);
      if DirExists(PdfImporterFolder) then
        DelTree(PdfImporterFolder, True, True, True);

      DeleteFile(DbFile);
      DeleteFile(DbLockFile);
      DeleteFile(LicenseFile);

      if DirExists(AppFolder) then
        DelTree(AppFolder, True, False, False);
    end;
  end;
end;