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
  // Testing
  // UserCfgPaths[3] := ExpandConstant('C:\Users\chris\Documents\temp\2024-LocalCache\UserCfg.opt');
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

// Get community folder directory
function GetCommunityFolderDir(Value: string): string;
begin
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
  SimConnectDll: String;
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
    SimConnectDll := ExpandConstant('{app}\SimConnect.dll');

    if MsgBox('Do you want to REMOVE all FSKneeboard data (including your license and SimConnect.dll)? (Hint: Answer "NO!" if you are just updating to a new version of FSKneeboard!)',
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
      DeleteFile(SimConnectDll);

      if DirExists(AppFolder) then
        DelTree(AppFolder, True, False, False);
    end;
  end;
end;

procedure OnLinkClick(Sender: TObject);
var
  ErrorCode: Integer;
begin
  ShellExec('open', 'https://github.com/Christian1984/vfrmap-for-vr/blob/master/README.md#installing-the-simconnect-sdk', '', '', SW_SHOWNORMAL, ewNoWait, ErrorCode);
end;

procedure OnSimConnectBrowseButtonClick(Sender: TObject);
var
  FileName: String;
begin
  FileName := SimConnectFileEdit.Text;
  if GetOpenFileName('Select SimConnect.dll', FileName, ExtractFilePath(FileName), 'SimConnect.dll|*.*', 'dll') then
  begin
    SimConnectFileEdit.Text := FileName;
  end;
end;

procedure CreateSimConnectWizardPage(AfterID: Integer);
var
  LinkFont: TFont;
  Linklabel: TNewStaticText;
  SelectFileLabel1: TNewStaticText;
  SelectFileLabel2: TNewStaticText;
begin
  SimConnectWizardPage := CreateOutputMsgPage(
    AfterID,
    'SimConnect.dll Installation',
    'FSKneeboard requires SimConnect.dll to interface with Microsoft Flight Simulator.',
    'The installer could not find an existing SimConnect.dll in the application directory. ' +
    'This file is required for FSKneeboard to communicate with the simulator.'#13#10#13#10 +
    'Please install the Microsoft Flight Simulator SDK to get the required file. You can find the SDK download and installation instructions here:'
  );

  // Create a clickable link
  Linklabel := TNewStaticText.Create(SimConnectWizardPage);
  Linklabel.Caption := 'MSFS SDK Installation Guide';
  Linklabel.Parent := SimConnectWizardPage.Surface;
  Linklabel.Top := 75;
  Linklabel.OnClick := @OnLinkClick;
  Linklabel.Cursor := crHand;

  // Style the link to look like a hyperlink
  LinkFont := TFont.Create;
  LinkFont.Assign(Linklabel.Font);
  LinkFont.Color := clBlue;
  LinkFont.Style := [fsUnderline];
  Linklabel.Font := LinkFont;

  // Add file picker
  SelectFileLabel1 := TNewStaticText.Create(SimConnectWizardPage);
  SelectFileLabel1.Caption := 'After installing the SDK (or if you already have a copy of SimConnect.dll),';
  SelectFileLabel1.Parent := SimConnectWizardPage.Surface;
  SelectFileLabel1.Top := Linklabel.Top + Linklabel.Height + 25;
  SelectFileLabel1.Width := SimConnectWizardPage.SurfaceWidth;


  SelectFileLabel2 := TNewStaticText.Create(SimConnectWizardPage);
  SelectFileLabel2.Caption := 'please point the installer to the location of the SimConnect.dll file:';
  SelectFileLabel2.Parent := SimConnectWizardPage.Surface;
  SelectFileLabel2.Top := SelectFileLabel1.Top + SelectFileLabel1.Height;
  SelectFileLabel2.Width := SimConnectWizardPage.SurfaceWidth;

  SimConnectFileEdit := TEdit.Create(SimConnectWizardPage);
  SimConnectFileEdit.Parent := SimConnectWizardPage.Surface;
  SimConnectFileEdit.Top := SelectFileLabel2.Top + SelectFileLabel2.Height + 5;
  SimConnectFileEdit.Width := SimConnectWizardPage.SurfaceWidth - 80;
  SimConnectFileEdit.Text := 'C:\MSFS 2024 SDK\SimConnect SDK\lib\SimConnect.dll';

  SimConnectFileButton := TButton.Create(SimConnectWizardPage);
  SimConnectFileButton.Parent := SimConnectWizardPage.Surface;
  SimConnectFileButton.Top := SimConnectFileEdit.Top - 2;
  SimConnectFileButton.Left := SimConnectFileEdit.Left + SimConnectFileEdit.Width + 5;
  SimConnectFileButton.Width := 75;
  SimConnectFileButton.Caption := 'Browse...';
  SimConnectFileButton.OnClick := @OnSimConnectBrowseButtonClick;
end;

function GetSimConnectPath(Value: string): string;
begin
    Result := SimConnectPath;
end;