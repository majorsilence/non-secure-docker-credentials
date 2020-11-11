$ErrorActionPreference = "Stop"
$CURRENTPATH=$pwd.Path

go build -o docker-credential-nonsecuredockercredentials.exe

Copy-Item "$CURRENTPATH\docker-credential-nonsecuredockercredentials.exe" "$CURRENTPATH\docker-credential-nonsecuredockercredentials\tools\docker-credential-nonsecuredockercredentials.exe" 

cd "$CURRENTPATH\docker-credential-nonsecuredockercredentials"
choco pack

cd "$CURRENTPATH"