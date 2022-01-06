package user_agent

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Factors struct {
	MicroMessengerVersions []string `json:"micro_messenger_versions"`
	OsVersions             []string `json:"os_versions"`
	AppleWebKitVersions    []string `json:"apple_web_kit_versions"`
}

func (f *Factors) Dump(filePath string) error {
	marshaled, err := json.Marshal(f)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, marshaled, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (f *Factors) Load(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, f)
	if err != nil {
		return err
	}

	return nil
}

func (f *Factors) Generate(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	microMessengerVersionsUnique := make(map[string]*struct{})
	osVersionsUnique := make(map[string]*struct{})
	appleWebKitUnique := make(map[string]*struct{})

	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		//UserAgents[i] = line

		ua := New(line)
		version := ua.MicroMessengerVersion()
		if version != "" {
			microMessengerVersionsUnique[version] = nil
		}
		version = ua.OS()
		if version != "" {
			osVersionsUnique[version] = nil
		}
		engine, version := ua.Engine()
		if version != "" && engine == "AppleWebKit" {
			appleWebKitUnique[version] = nil
		}
	}

	f.MicroMessengerVersions = make([]string, len(microMessengerVersionsUnique))
	i := 0
	for ver := range microMessengerVersionsUnique {
		f.MicroMessengerVersions[i] = ver
		i += 1
	}

	f.OsVersions = make([]string, len(osVersionsUnique))
	i = 0
	for ver := range osVersionsUnique {
		f.OsVersions[i] = ver
		i += 1
	}

	f.AppleWebKitVersions = make([]string, len(appleWebKitUnique))
	i = 0
	for ver := range appleWebKitUnique {
		f.AppleWebKitVersions[i] = ver
		i += 1
	}

	return nil
}
