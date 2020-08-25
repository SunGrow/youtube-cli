package ttyinput

import (
	"fmt"
	"os"
	"youtube-cli/pagebuild"
	"youtube-cli/userset"
)


func genUserFeedPage(user string, configDir string) {
	userset.CheckUser(user, configDir);
	channels := pagebuild.ParceCSVSubListInput(configDir+"/"+user+"/"+user+"_sub_list.csv");

	pagebuild.BuildHTMLFeedPage(channels, configDir+"/"+user);
	fmt.Printf("User Feed page for %s generated\n", user);
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
		var i int = 1;
		for i == 1 {
			var tmpStrTitle string;
			var tmpStrLink string;
			fmt.Print("Enter a channel name:\n");
			fmt.Scan(&tmpStrTitle);
			fmt.Print("Enter a channel URL:\n");
			fmt.Scan(&tmpStrLink);
			channelInfos = append(channelInfos, []string{tmpStrTitle, tmpStrLink});
			fmt.Print("Do you wish to continue?\n1) Yes\n2) No\n");
			fmt.Scan(&i);
		}
	default: 
		break;
	}
	pagebuild.BuildCSVPage(channelInfos, file);
}

func addUserSubList(user string, configDir string) {
	userset.CheckUser(user, configDir);
	fmt.Print(
		"Do you want to:\n",
		"1) Overwrite a sublist\n",
		"2) Append to a sublist?\n",
	);
	var operationType int;
	fmt.Scan(&operationType);

	fmt.Print(
		"You want to do it with:\n",
		"1) CSV file\n",
		"2) XML file\n",
		"3) Plaintext links\n",
	);
	var fileFormat int;
	fmt.Scan(&fileFormat);

	var file *os.File;
	var err error;
	switch operationType {
	case 1: 
		file, err = os.Create(configDir+"/"+user+"/"+user+"_sub_list.csv");
		break;
	case 2:
		file, err = os.OpenFile(configDir+"/"+user+"/"+user+"_sub_list.csv", os.O_RDWR|os.O_APPEND,os.ModePerm);
		break;
	default:
		return;
	}
	if err != nil {
		fmt.Print(err);
	}
	defer file.Close();
	genUserSubList(file, fileFormat);
}

// Functions return the number of args they read beside themselves
func GetArgFunctionMap()(map[string]func(string, string)) {
	return map[string]func(string, string) {
		"feed"		:genUserFeedPage,
		"sublist"	:addUserSubList,
	}
}
