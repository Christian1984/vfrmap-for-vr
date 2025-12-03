// Common variables used by installer discovery functions
// Note: These must be declared in the main installer files as well:
//   DetectedCommunityFolders: TArrayOfString;
//   CommunityFolderVersions: TArrayOfString; 
//   CommunityFolderDir: String;

// Parse UserCfg.opt file to extract InstalledPackagesPath

function ParseUserCfgOpt(FilePath: String): String;
var
  Lines: TArrayOfString;
  i: Integer;
  Line: String;
  StartPos: Integer;
begin
  Result := '';
  if FileExists(FilePath) then
  begin
    LoadStringsFromFile(FilePath, Lines);
    for i := 0 to GetArrayLength(Lines) - 1 do
    begin
      Line := Trim(Lines[i]);
      if Pos('InstalledPackagesPath', Line) > 0 then
      begin
        StartPos := Pos('"', Line);
        if StartPos > 0 then
        begin
          Delete(Line, 1, StartPos);
          StartPos := Pos('"', Line);
          if StartPos > 0 then
          begin
            Delete(Line, StartPos, Length(Line) - StartPos + 1);
            Result := Line;
            Exit;
          end;
        end;
      end;
    end;
  end;
end;

// Discover community folders by searching for UserCfg.opt files
procedure DiscoverCommunityFolders();
var
  UserCfgPaths: TArrayOfString;
  i, Count: Integer;
  PackagesPath, CommunityPath: String;
  Version: String;
begin
  SetArrayLength(DetectedCommunityFolders, 10); // Reserve space
  SetArrayLength(CommunityFolderVersions, 10);
  Count := 0;

  // Known paths for MSFS 2020
  SetArrayLength(UserCfgPaths, 6);
  UserCfgPaths[0] := ExpandConstant('{localappdata}\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\LocalCache\UserCfg.opt');
  UserCfgPaths[1] := ExpandConstant('{localappdata}\Packages\Microsoft.FlightDashboard_8wekyb3d8bbwe\LocalCache\UserCfg.opt');
  UserCfgPaths[2] := ExpandConstant('{userappdata}\Microsoft Flight Simulator\UserCfg.opt');
  
  // Known paths for MSFS 2024
  UserCfgPaths[3] := ExpandConstant('{localappdata}\Packages\Microsoft.Limitless_8wekyb3d8bbwe\LocalCache\UserCfg.opt');
  UserCfgPaths[4] := ExpandConstant('{userappdata}\Microsoft Flight Simulator 2024\UserCfg.opt');
  UserCfgPaths[5] := ExpandConstant('{localappdata}\Packages\Microsoft.FlightDashboard2024_8wekyb3d8bbwe\LocalCache\UserCfg.opt'); // Potential Steam 2024 path

  for i := 0 to GetArrayLength(UserCfgPaths) - 1 do
  begin
    if FileExists(UserCfgPaths[i]) then
    begin
      PackagesPath := ParseUserCfgOpt(UserCfgPaths[i]);
      if PackagesPath <> '' then
      begin
        CommunityPath := PackagesPath + '\Community';
        if DirExists(CommunityPath) then
        begin
          // Determine version based on path
          if (Pos('Microsoft.FlightSimulator_8wekyb3d8bbwe', UserCfgPaths[i]) > 0) or 
             (Pos('Microsoft.FlightDashboard_8wekyb3d8bbwe', UserCfgPaths[i]) > 0) or
             (Pos('Microsoft Flight Simulator\UserCfg.opt', UserCfgPaths[i]) > 0) then
            Version := 'MSFS 2020'
          else if (Pos('Microsoft.Limitless_8wekyb3d8bbwe', UserCfgPaths[i]) > 0) or
                  (Pos('Microsoft Flight Simulator 2024\UserCfg.opt', UserCfgPaths[i]) > 0) or
                  (Pos('Microsoft.FlightDashboard2024_8wekyb3d8bbwe', UserCfgPaths[i]) > 0) then
            Version := 'MSFS 2024'
          else
            Version := 'Unknown';

          // Check for duplicates
          if Count = 0 then
          begin
            DetectedCommunityFolders[Count] := CommunityPath;
            CommunityFolderVersions[Count] := Version;
            Count := Count + 1;
          end else begin
            // Check if this path is already added
            if CommunityPath <> DetectedCommunityFolders[Count-1] then
            begin
              DetectedCommunityFolders[Count] := CommunityPath;
              CommunityFolderVersions[Count] := Version;
              Count := Count + 1;
            end;
          end;
        end;
      end;
    end;
  end;

  // Resize arrays to actual count
  SetArrayLength(DetectedCommunityFolders, Count);
  SetArrayLength(CommunityFolderVersions, Count);
end;

// Install panel to additional community folders
procedure InstallPanelToAdditionalFolders();
var
  i: Integer;
  SourceDir: String;
  TargetDir: String;
  FolderPaths: TArrayOfString;
begin
  if GetArrayLength(DetectedCommunityFolders) <= 1 then
    Exit; // No additional folders to process

  // Parse selected folders from semicolon-separated string
  FolderPaths := StrSplit(CommunityFolderDir, ';');
  
  SourceDir := DetectedCommunityFolders[0] + '\christian1984-ingamepanel-fskneeboard';
  
  // Copy to each additional selected folder (skip first one as it's already installed)
  for i := 1 to GetArrayLength(FolderPaths) - 1 do
  begin
    TargetDir := FolderPaths[i] + '\christian1984-ingamepanel-fskneeboard';
    Log('Copying panel from ' + SourceDir + ' to ' + TargetDir);
    
    if DirExists(SourceDir) then
    begin
      // Create target directory
      ForceDirectories(TargetDir);
      
      // Copy all files recursively
      if not DirCopy(SourceDir, TargetDir, True) then
        Log('Warning: Failed to copy panel to ' + TargetDir);
    end;
  end;
end;

// Get community folder directory (handles multiple folders)
function GetCommunityFolderDir(Value: string): string;
begin
    // For multiple folders, return the first one for the main installation
    // Additional copies will be handled by CurStepChanged
    if GetArrayLength(DetectedCommunityFolders) > 1 then
      Result := DetectedCommunityFolders[0]
    else
      Result := CommunityFolderDir;
end;

// Helper function to split strings
function StrSplit(Text: String; Separator: String): TArrayOfString;
var
  i, p: Integer;
  Dest: TArrayOfString;
begin
  i := 0;
  repeat
    SetArrayLength(Dest, i + 1);
    p := Pos(Separator, Text);
    if p > 0 then
    begin
      Dest[i] := Copy(Text, 1, p - 1);
      Text := Copy(Text, p + Length(Separator), Length(Text));
      i := i + 1;
    end
    else
    begin
      Dest[i] := Text;
      Text := '';
    end;
  until Length(Text) = 0;
  Result := Dest;
end;

// Helper function to copy directory recursively
function DirCopy(SourcePath, DestPath: String; Overwrite: Boolean): Boolean;
var
  FindRec: TFindRec;
  SourceFilePath, DestFilePath: String;
begin
  Result := True;
  
  if FindFirst(SourcePath + '\*', FindRec) then
  begin
    try
      repeat
        if (FindRec.Name <> '.') and (FindRec.Name <> '..') then
        begin
          SourceFilePath := SourcePath + '\' + FindRec.Name;
          DestFilePath := DestPath + '\' + FindRec.Name;
          
          if FindRec.Attributes and FILE_ATTRIBUTE_DIRECTORY <> 0 then
          begin
            // Directory
            if not DirExists(DestFilePath) then
              ForceDirectories(DestFilePath);
            if not DirCopy(SourceFilePath, DestFilePath, Overwrite) then
              Result := False;
          end
          else
          begin
            // File
            if not FileCopy(SourceFilePath, DestFilePath, Overwrite) then
              Result := False;
          end;
        end;
      until not FindNext(FindRec);
    finally
      FindClose(FindRec);
    end;
  end;
end;

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