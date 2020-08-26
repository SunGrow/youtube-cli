package userset

import (
	"os"
	"strings"
)

// Input Application State
type InputState struct {
	User 		string
	Output 		string
	ConfigFile	string
}

func GetConfigDir() string {
	configDir, isDir := os.LookupEnv("XDG_CONFIG_HOME");
	if isDir {
		configDir = configDir + "/youtube-cli";
	} else {
		configDir = configDir + "/.config";
	}
	_, err := os.Stat(configDir);
	if err != nil {
		os.Mkdir(configDir, os.ModePerm);
	}
	return configDir;
}

func GetDir(file string) (dir string) {
	var tmpArr []string	= strings.Split(dir, "/");
	tmpArr = tmpArr[:len(tmpArr)-1];
	dir = strings.Join(tmpArr, "/");
	return;
}

func CheckDir(dir string) (existed int) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		existed |= 0x1;
	}
	return;
}

func GetWD() string {
	str, err := os.Getwd()
	if err != nil {
		print(err);
	}
	return str;
}
