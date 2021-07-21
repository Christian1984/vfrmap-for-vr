; Script generated by the Inno Setup Script Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

#ifndef ApplicationVersion
#define ApplicationVersion "0.0.1"
#endif

#define MyAppName "FSKneeboard PRO"
#define MyAppPublisher "Christian-Alexander Hoffmann"
#define MyAppURL "https://www.fskneeboard.com/"
#define MyAppExeName "fskneeboard.exe"

[Setup]
; NOTE: The value of AppId uniquely identifies this application. Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{85F960AA-B1FF-4C01-BEF0-74D87689AE8E}
AppName={#MyAppName}
AppVersion={#ApplicationVersion}
AppVerName={#MyAppName} v{#ApplicationVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={autopf}\{#MyAppName}
DisableProgramGroupPage=yes
LicenseFile=..\LICENSE.txt
; Remove the following line to run in administrative install mode (install for all users.)
PrivilegesRequired=lowest
OutputDir=..\dist
OutputBaseFilename=Install-FSKneeboard-PRO-v{#ApplicationVersion}
SetupIconFile=..\fskneeboard-server\vfrmap\winres\fskneeboard.ico
Compression=lzma
SolidCompression=yes
WizardStyle=modern

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "..\dist\pro\fskneeboard-server\{#MyAppExeName}"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\dist\pro\CHANGELOG.md"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\dist\pro\README.pdf"; DestDir: "{app}"; Flags: ignoreversion isreadme
Source: "..\dist\pro\fskneeboard-server\copy-your-license-file-here.txt"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-server\fskneeboard-autostart-steam.bat"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-server\fskneeboard-autostart-windows-store.bat"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-server\charts\copy-your-charts-here.txt"; DestDir: "{app}\charts"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-server\charts\traffic-pattern.png"; DestDir: "{app}\charts"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-server\autosave\autosaves-will-go-here.txt"; DestDir: "{app}\autosave"; Flags: ignoreversion

Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\layout.json"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\manifest.json"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\InGamePanels\christian1984-ingamepanel-fskneeboard.spb"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard\InGamePanels"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\Textures\Menu\toolbar\ICON_TOOLBAR_CHRISTIAN1984_INGAMEPANEL_VFRMAPFORVR.svg"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard\html_ui\Textures\Menu\toolbar"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\FSKneeboardPanel.css"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\FSKneeboardPanel.html"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel"; Flags: ignoreversion
Source: "..\dist\pro\fskneeboard-panel\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel\FSKneeboardPanel.js"; DestDir: "{code:GetCommunityFolderDir}\christian1984-ingamepanel-fskneeboard\html_ui\InGamePanels\FSKneeboardPanel"; Flags: ignoreversion

Source: "{code:GetLicenseFile}"; DestDir: "{app}"; DestName: "fskneeboard.lic"; Flags: external
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Icons]
Name: "{autoprograms}\FSKneeboard\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; IconFilename: "{app}\{#MyAppExeName}"
Name: "{autoprograms}\FSKneeboard\{#MyAppName} + MSFS (Windows Store)"; Filename: "{app}\fskneeboard-autostart-windows-store.bat"; IconFilename: "{app}\{#MyAppExeName}"
Name: "{autoprograms}\FSKneeboard\{#MyAppName} + MSFS (Steam)"; Filename: "{app}\fskneeboard-autostart-steam.bat"; IconFilename: "{app}\{#MyAppExeName}"
Name: "{autoprograms}\FSKneeboard\Docs - Readme"; Filename: "{app}\README.pdf"
Name: "{autoprograms}\FSKneeboard\Docs - Changelog"; Filename: "{app}\CHANGELOG.md"
Name: "{autoprograms}\FSKneeboard\Charts-Folder"; Filename: "{app}\charts"
Name: "{autoprograms}\FSKneeboard\Autosave-Folder"; Filename: "{app}\autosave"
Name: "{autoprograms}\FSKneeboard\Web - Join us on Discord"; Filename: "https://discord.fskneeboard.com"
; Name: "{autoprograms}\FSKneeboard\Web - Upgrade to PRO"; Filename: "https://fskneeboard.com/buy-now"
Name: "{autodesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent
Filename: "https://discord.fskneeboard.com"; Flags: nowait shellexec runasoriginaluser postinstall; Description: "Join us on Discord!"

[UninstallDelete]
;This works if it is installed in custom location
Type: files; Name: "{app}\latestcheck.json"; 
Type: files; Name: "{app}\SimConnect.dll"; 

[Code]
var
  CommunityFolderDirWizardPage: TInputDirWizardPage;
  CommunityFolderDir: String;
  LicenseFileWizardPage: TInputFileWizardPage;
  LicenseFile: String;

function GetCommunityFolderDir(Value: string): string;
begin
    Result := CommunityFolderDir;
end;

function GetLicenseFile(Value: string): string;
begin
    Result := LicenseFile;
end;

procedure InitializeWizard;
var
  AfterID: Integer;
  communityFolder: String;
  winstoreCommunityFolder: String;
  steamCommunityFolder: String;
  communityFolderDirWizardDescription: String;

begin
  AfterID := wpSelectDir;

  winstoreCommunityFolder := ExpandConstant('{localappdata}\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\LocalCache\Packages\Community');
  steamCommunityFolder := ExpandConstant('{localappdata}\Packages\Microsoft.FlightDashboard_8wekyb3d8bbwe\LocalCache\Packages\Community');
  
  communityFolderDirWizardDescription := 'WARNING: Your Flight Simulator Community Folder could NOT be auto-detected! Please set the path to your community folder manually:'#13#10#13#10
    '- WINDOWS STORE USERS: If you have purchased MSFS through the Windows Store, you will typically find it under'#13#10#13#10 + 'C:\Users\[username]\AppData\Local\Packages\Microsoft.FlightSimulator_8wekyb3d8bbwe\ '#13#10 + 'LocalCache\Packages\Community'#13#10#13#10
    + '- STEAM USERS: If you have purchased MSFS through Steam, the default path for your Community Folder would typically be'#13#10#13#10 + 'C:\Users\[username]\AppData\Local\Packages\Microsoft.FlightDashboard_8wekyb3d8bbwe\ '#13#10 + 'LocalCache\Packages\Community';

  if DirExists(winstoreCommunityFolder) then begin
    communityFolder := winstoreCommunityFolder;
    communityFolderDirWizardDescription := 'SUCCESS! Automatically detected your Flight Simulator Community Folder.';
    Log('winstoreCommunityFolder found!');
  end else if DirExists(steamCommunityFolder) then begin
    communityFolder := steamCommunityFolder;
    communityFolderDirWizardDescription := 'SUCCESS! Automatically detected your Flight Simulator Community Folder.';
    Log('steamCommunityFolder found!');
  end else begin
    communityFolder := 'C:\'
  end;

  communityFolderDirWizardDescription := communityFolderDirWizardDescription + ''#13#10#13#10 + 'PLEASE NOTE: If the FSKneeboard-Ingame-Panel does NOT appear inside the game, then double-check that you have this directory right!';

  CommunityFolderDirWizardPage := CreateInputDirPage(
        AfterID,
        'Select Community Folder Location',
        'Please tell us where your Flight Simulator Community Folder is located!',
        communityFolderDirWizardDescription,
        False, '');
  CommunityFolderDirWizardPage.Add('&Microsoft Flight Simulator Community Folder:');
  CommunityFolderDirWizardPage.Values[0] := winstoreCommunityFolder;
  AfterID := CommunityFolderDirWizardPage.ID;

  LicenseFileWizardPage := CreateInputFilePage(
      AfterID,
      'Select License File',
      'Please tell us where we can find your license file...'#13#10
      + '(Typically that would be something like C:\Users\[username]\Downloads)', '');
  LicenseFileWizardPage.Add('&License File:', 'FSKneeboard License Files|*.lic', '.lic');
  LicenseFileWizardPage.Values[0] := ExpandConstant('{%USERPROFILE}\Downloads\fskneeboard.lic');
  AfterID := LicenseFileWizardPage.ID;
end;

function NextButtonClick(CurrPageID: Integer): Boolean;
begin
  if CurrPageID = CommunityFolderDirWizardPage.ID then begin
    CommunityFolderDir := CommunityFolderDirWizardPage.Values[0];
    Log('CommunityFolderDir is: ' + CommunityFolderDir);
  end else if CurrPageID = LicenseFileWizardPage.ID then begin
    LicenseFile := LicenseFileWizardPage.Values[0];
    Log('LicenseFile is: ' + LicenseFile);

    if not FileExists(LicenseFile) then begin
        MsgBox('The License File ' + LicenseFile + ' does not exists. Please select your license file!'#13#10#13#10 + 'If you do not have a license file yet, please purchase a copy of FSKneeboard at https://fskneeboard.com/buy-now', mbError, MB_OK);
        Result := False
        Exit;
    end;
  end;
  Result := True;
end;

function UpdateReadyMemo(Space, NewLine, MemoUserInfoInfo, MemoDirInfo, MemoTypeInfo,
  MemoComponentsInfo, MemoGroupInfo, MemoTasksInfo: String): String;
var
  S: String;
begin
  S := 'FSKneeboard PRO is Ready for Installation!' + NewLine;
  S := S + NewLine;
  S := S + 'FSKneeboard Server Component will be installed to:' + NewLine;
  S := S + Space + ExpandConstant('{app}') + NewLine;
  S := S + NewLine;
  S := S + 'FSKneeboard Ingame Panel will be installed to:' + NewLine;
  S := S + Space + CommunityFolderDir + '\christian1984-ingamepanel-fskneeboard' + NewLine;
  S := S + NewLine;
  S := S + 'FSKneeboard License File will be copied from:' + NewLine;
  S := S + Space + LicenseFile + NewLine;
  S := S + 'to:' + NewLine;
  S := S + Space + ExpandConstant('{app}\fskneeboard.lic') + NewLine;
  S := S + NewLine;
  S := S + 'Join us on Discord at https://discord.fskneeboard.com';
  Result := S;
end;