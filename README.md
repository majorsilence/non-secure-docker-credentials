# non-secure-docker-credentials
Play around with custom docker credential helpers


See https://github.com/docker/docker-credential-helpers for upstream supported credential helpers.  

This project is a torn apart copy of wincred to learn how the docker credential helpers work and teach myself some basics of go.


# Warning
All credentials are saved in plain text.  You should not use this if that is a problem.



# Build


```powershell
go build -o docker-credential-nonsecuredockercredentials.exe
```

# Usage

Find your dockers config.json.

Modify credStore section to be nonsecuredockercredentials

```json
	"credsStore": "nonsecuredockercredentials",
```
