# non-secure-docker-credentials
Play around with custom docker credential helpers


See https://github.com/docker/docker-credential-helpers for upstream supported credential helpers.  

This project is a torn apart copy of wincred to learn how the docker credential helpers work and teach myself some basics of go.


# Warning
All credentials are saved in plain text (base 64).  You should not use this if that is a problem.



# Build


```powershell
go build -o docker-credential-nonsecuredockercredentials.exe
```

Or run the build script to also include generating a choco package.

```powershell
.\build.ps1
```



# Usage

Find your dockers config.json.   On windows __C:\Users\[Your User Name]\.docker\config.json__.

Modify credsStore section to be nonsecuredockercredentials

```json
	"credsStore": "nonsecuredockercredentials",
```

Add the executable to the system path.

```powershell
[Environment]::SetEnvironmentVariable("the path to the folder with docker-credential-nonsecuredockercredentials.exe", $env:Path, [System.EnvironmentVariableTarget]::Machine)
```
