package ttyinput

import (
	"fmt"
	"os"
	"youtube-cli/pagebuild"
	"youtube-cli/userset"
)

// Functions return the number of args they read beside themselves
func GetArgFunctionMap()(map[string]func(userset.InputState)) {
	return map[string]func(userset.InputState) {
		"feed"		:genUserFeedPage,
		"sublist"	:addUserSubList,
	}
}

func genUserFeedPage(state userset.InputState) {
	userset.CheckDir(userset.GetDir(state.ConfigFile));
	channels := pagebuild.ParceCSVSubListInput(state.ConfigFile);
	pagebuild.BuildHTMLFeedPage(channels, state.Output);
	fmt.Printf("%s feed page generated\n", state.User);
}

func GetSubListReqParam() (fFormat int, opType int) {
	fmt.Print(
		"Do you want to:\n",
		"1) Overwrite a sublist\n",
		"2) Append to a sublist?\n",
	);
	fmt.Scan(&opType);

	fmt.Print(
		"You want to do it with:\n",
		"1) CSV file\n",
		"2) XML file\n",
		"3) Plaintext links\n",
	);
	fmt.Scan(&fFormat);
	return;
}

func addUserSubList(state userset.InputState) {
	userset.CheckDir(userset.GetDir(state.ConfigFile));
	fFormat, opType := GetSubListReqParam();

	var file *os.File;
	var err error;

	switch opType {
	case 1: 
		file, err = os.Create(state.ConfigFile);
		break;
	case 2:
		file, err = os.OpenFile(state.ConfigFile, os.O_RDWR|os.O_APPEND, os.ModePerm);
		break;
	default:
		return;
	}
	if err != nil {
		fmt.Print(err);
	}
	defer file.Close();
	genUserSubList(file, fFormat);
}


func genUserSubList(file *os.File, fileFormat int) {
	var path string;
	var channelInfos [][]string;
	switch fileFormat {
	case 1:
		fmt.Print("Write a PATH to your CSV\n");
		fmt.Scan(&path);
		channelInfos = pagebuild.ParceCSVSubListInput(path);
		break;
	case 2: 
		fmt.Print("Write a PATH to your subscription_manager XML\n");
		fmt.Scan(&path);
		channelInfos = pagebuild.ParceXMLSubListInput(path);
		break;
	case 3:
		channelInfos = pagebuild.GetTTYSubListInput(channelInfos);
		break;
	default: 
		break;
	}
	pagebuild.BuildCSVPage(channelInfos, file);
}

