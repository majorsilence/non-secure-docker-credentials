package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/shibukawa/configdir"
)

func main() {
	credentials.Serve(Nonsecuredockercredentials{})
}

// Wincred handles secrets using the Windows credential service.
type Nonsecuredockercredentials struct{}

type cred struct {
	CredsLabel string
	ServerUrl  string
	Secret     string
	Username   string
}

type credList struct {
	Creds []cred
}

// Add adds new credentials to the windows credentials manager.
func (h Nonsecuredockercredentials) Add(creds *credentials.Credentials) error {

	configDirs := configdir.New("majorsilence", "nonsecuredockercredentials")

	folder := configDirs.QueryFolders(configdir.Global)[0]

	configDirs.LocalPath = configDirs.QueryFolders(configdir.Global)[0].Path
	var credsList []credentials.Credentials

	if folder.Exists("settings.json") == false {
		folder.Create("settings.json")
	}

	if folder != nil {
		data, err := folder.ReadFile("settings.json")
		if err != nil {
			fmt.Println("File reading error", err)
			return err
		}

		json.Unmarshal(data, &credsList)
	}

	var credValue credentials.Credentials = *creds
	credValue.Secret = b64.StdEncoding.EncodeToString([]byte(credValue.Secret))
	credValue.Username = b64.StdEncoding.EncodeToString([]byte(credValue.Username))
	credsList = append(credsList, credValue)
	data, _ := json.Marshal(&credsList)
	folder.WriteFile("settings.json", data)

	return nil
}

// Delete removes credentials from the windows credentials manager.
func (h Nonsecuredockercredentials) Delete(serverURL string) error {
	configDirs := configdir.New("majorsilence", "nonsecuredockercredentials")
	var creds []credentials.Credentials

	var credsNew []credentials.Credentials

	folder := configDirs.QueryFolderContainsFile("settings.json")
	if folder != nil {
		data, _ := folder.ReadFile("settings.json")
		json.Unmarshal(data, &creds)
	} else {
		return errors.New("settings.json not found")
	}

	for i := range creds {

		if strings.Compare(creds[i].ServerURL, serverURL) != 0 {
			credsNew = append(credsNew, creds[i])
		}
	}

	data, _ := json.Marshal(&credsNew)
	folder.WriteFile("settings.json", data)

	return nil
}

// Get retrieves credentials from the windows credentials manager.
func (h Nonsecuredockercredentials) Get(serverURL string) (string, string, error) {

	configDirs := configdir.New("majorsilence", "nonsecuredockercredentials")
	var creds []credentials.Credentials

	folder := configDirs.QueryFolderContainsFile("settings.json")
	if folder != nil {
		data, _ := folder.ReadFile("settings.json")
		json.Unmarshal(data, &creds)
	} else {
		return "", "", errors.New("settings.json not found")
	}

	for i := range creds {
		if strings.Compare(creds[i].ServerURL, serverURL) == 0 {
			secret, _ := b64.StdEncoding.DecodeString(creds[i].Secret)
			username, _ := b64.StdEncoding.DecodeString(creds[i].Username)
			return string(username), string(secret), nil
		}

	}

	return "", "", credentials.NewErrCredentialsNotFound()
}

// List returns the stored URLs and corresponding usernames for a given credentials label.
func (h Nonsecuredockercredentials) List() (map[string]string, error) {

	configDirs := configdir.New("majorsilence", "nonsecuredockercredentials")
	var creds []credentials.Credentials
	folder := configDirs.QueryFolderContainsFile("settings.json")

	if folder != nil {

		data, err := folder.ReadFile("settings.json")
		if err != nil {
			fmt.Println("File reading error", err)
			return nil, err
		}
		fmt.Printf(string(data))
		json.Unmarshal(data, &creds)

	} else {
		return nil, errors.New("settings.json not found")
	}

	resp := make(map[string]string)
	for i := range creds {
		fmt.Printf(creds[i].ServerURL)
		username, _ := b64.StdEncoding.DecodeString(creds[i].Username)
		resp[creds[i].ServerURL] = string(username)
	}

	return resp, nil
}
