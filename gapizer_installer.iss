[Setup]
AppName=GAPIzer
AppVersion=1.0.0
DefaultDirName={commonpf}\GAPIzer
DefaultGroupName=GAPIzer
OutputDir=Output
OutputBaseFilename=GAPIzer-Installer
Compression=zip
SolidCompression=yes
AppPublisher=SHM SOFTWARE SOLUTIONS
AppPublisherURL=SHM SOFTWARE SOLUTIONS
AppSupportURL=https://github.com/rshmdev/gapizer
AppUpdatesURL=https://github.com/rshmdev/gapizer

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"
Name: "portuguese"; MessagesFile: "compiler:Languages\BrazilianPortuguese.isl"

[Files]
Source: "bin\gapizer.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\GAPIzer CLI"; Filename: "{app}\gapizer.exe"

[Run]
Filename: "{app}\gapizer.exe"; Description: "Rodar GAPIzer após a instalação"; Flags: shellexec skipifsilent

[Registry]
Root: HKLM; Subkey: "System\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{app}"; Flags: preservestringtype
