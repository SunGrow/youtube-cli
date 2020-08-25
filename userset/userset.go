package userset

import (
	"os"
)

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

func CheckUser(user string, configDir string) int {
	var existed int = 0;
	_, err := os.Stat(configDir);
	if err != nil {
		existed |= 0x1;
		os.Mkdir(configDir, os.ModePerm);
	}
	_, err = os.Stat(configDir + "/" + user);
	if err != nil {
		existed |= 0x2;
		os.Mkdir(configDir + "/" + user, os.ModePerm);
	}
	return existed;
}
