package pagebuild

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"
)


type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Title string `xml:"title"`
	Link Link `xml:"link"`
	Author Author `xml:"author"`
	Published string `xml:"published"`
	Media Media `xml:"group"`
}

type Author struct {
	XMLName xml.Name `xml:"author"`
	Name string `xml:"name"`
	Uri string `xml:"uri"`
}

type Link struct {
	XMLName xml.Name `xml:"link"`
	Rel string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Media struct {
	XMLName xml.Name `xml:"group"`
	Title string `xml:"title"`
	Thumbnail Thumbnail `xml:"thumbnail"`
	Description string `xml:"description"`
	Community Community `xml:"community"`
	Content Content `xml:"content"`
}

type Thumbnail struct {
	XMLName xml.Name `xml:"thumbnail"`
	Url string `xml:"url,attr"`
	Width string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}

type Community struct {
	XMLName xml.Name `xml:"community"`
	StarRating StarRating `xml:"starRating"`
	Statistics Statistics `xml:"statistics"`
}

type Content struct {
	XMLName xml.Name `xml:"content"`
	Url string `xml:"url"`
}

type StarRating struct {
	XMLName xml.Name `xml:"starRating"`
	Count string `xml:"count,attr"`
	Average string `xml:"average,attr"`
}

type Statistics struct {
	XMLName xml.Name `xml:"statistics"`
	Views string `xml:"views,attr"`
}


func ParceXMLChannelFeed(channelAddr string) (channelInfo []map[string]string){
	// Get channel feed
	resp, err := http.Get(channelAddr);
	if err != nil {
    	fmt.Println(err);
	}
	defer resp.Body.Close();
	
	// Read body from response
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
    	fmt.Println(err);
	}
	var channel Feed;
	xml.Unmarshal(body, &channel);
	for i := 0; i < len(channel.Entries); i++ {
		channelInfo = append(channelInfo, map[string]string{
			"author"		:channel.Entries[i].Author.Name,
			"authorLink"	:channelAddr,
			"title"			:channel.Entries[i].Title,
			"link"			:channel.Entries[i].Link.Href,
			"published"		:channel.Entries[i].Published,
			"views"			:channel.Entries[i].Media.Community.Statistics.Views,
			"rate"			:channel.Entries[i].Media.Community.StarRating.Average,
			"rateCount"		:channel.Entries[i].Media.Community.StarRating.Count,
			"thumbnail"		:channel.Entries[i].Media.Thumbnail.Url,
			"thumbnailW"	:channel.Entries[i].Media.Thumbnail.Width,
			"thumbnailH"	:channel.Entries[i].Media.Thumbnail.Height,
			"descr"			:channel.Entries[i].Media.Description,
			"video"			:channel.Entries[i].Media.Content.Url,
		});
	}
	
	return;
}


type OPML struct {
	XMLName xml.Name `xml:"opml"`
	Version string `xml:"version,attr"`
	Body Body `xml:"body"`
}

type Body struct {
	XMLName xml.Name `xml:"body"`
	HeadOutline HeadOutline `xml:"outline"`
}

type HeadOutline struct {
	XMLName xml.Name  `xml:"outline"`
	Outline []Outline `xml:"outline"`
}

type Outline struct {
	XMLName xml.Name `xml:"outline"`
	Text string `xml:"text,attr"`
	ChannelTitle string `xml:"title,attr"`
	LinkType string `xml:"rss,attr"`
	ChannelLink string `xml:"xmlUrl,attr"`
}

func ParceXMLSubListInput(inputDir string) (channelInfo [][]string) {
    file, err := os.Open(inputDir);
    if err != nil {
    	fmt.Println(err)
    }
    defer file.Close()
    byteValue, _ := ioutil.ReadAll(file)   // Unmarshal takes a []byte and fills the rss struct with the values found in the xmlFile
	var page OPML;
	xml.Unmarshal(byteValue, &page);
	for i := 0; i < len(page.Body.HeadOutline.Outline); i++ {
		channelInfo = append(channelInfo, []string{page.Body.HeadOutline.Outline[i].ChannelTitle, page.Body.HeadOutline.Outline[i].ChannelLink});
	}
	return;
}

func ParceCSVSubListInput(inputDir string) (channels [][]string) {
    file, err := os.Open(inputDir);
    if err != nil {
    	fmt.Println(err)
    }
    defer file.Close()
	reader := csv.NewReader(file);
	str, err := reader.ReadAll()
	return str;
}

func GetTTYSubListInput(channelInfos [][]string) (chInf [][]string) {
	chInf = channelInfos;
	var i int = 1;
	for i == 1 {
		var tmpStrTitle string;
		var tmpStrLink string;
		fmt.Print("Enter a channel name:\n");
		fmt.Scan(&tmpStrTitle);
		fmt.Print("Enter a channel URL:\n");
		fmt.Scan(&tmpStrLink);
		chInf = append(chInf, []string{tmpStrTitle, tmpStrLink});
		fmt.Print("Do you wish to continue?\n1) Yes\n2) No\n");
		fmt.Scan(&i);
	}
	return;
}

func BuildCSVPage(channelInfos [][]string, outFile *os.File) {
	writer := csv.NewWriter(outFile);
	err := writer.WriteAll(channelInfos);
	if err != nil {
		fmt.Println(err);
	}
}


func BuildHTMLFeedPage(channels [][]string, out string) {
	file, err := os.Create(out);
	if err != nil {
		fmt.Println(err);
		return;
	}
	defer file.Close();

	var videos []map[string]string;
	for _, channel := range channels {
		for _, video := range ParceXMLChannelFeed(channel[1]){
			videos = append(videos, video);
		}
	}
	sort.Slice(videos, func(i int, j int)bool {
		return videos[i]["published"] > videos[j]["published"];
	});
	FillHTMLFeedPage(file, videos);
}

func FillHTMLFeedPage(file *os.File, videos []map[string]string) {
	hrs, min, sec := time.Now().Local().Clock();
	day, month, year := time.Now().Date();
	var datetime string;
	datetime = fmt.Sprintf("%d:%d:%d %d-%s-%d", hrs, min, sec, year, month.String(), day);

	file.WriteString("<html>\n");
	file.WriteString("<head>\n");
	file.WriteString("<title>YouTube Subscription Feed (" + datetime + ")</title>\n");
	file.WriteString("</head>\n");
	file.WriteString("<body>\n");
	file.WriteString("<h1>YouTube Subscription Feed (" + datetime + ")</h1>\n")
	for i, video := range videos {
		file.WriteString("<h2><img src="+video["thumbnail"]+" witdth="+video["thumbnailW"]+" height="+video["thumbnailH"]+" alt=img></img></h2>\n");
		file.WriteString("<h2><a href="+video["link"]+">"+video["title"]+"</a></h2>\n");
		file.WriteString("<p><a href="+video["authorLink"]+"> Author: "+video["author"]+"</a></p>\n")
		file.WriteString("<table style=\"width:30%\">\n");
		file.WriteString("<tr>\n");
		file.WriteString("<td><p>Views: "+video["views"]+"</p></td>\n");
		file.WriteString("<td><p>Date: "+video["published"]+"</p></td>\n");
		file.WriteString("</tr>\n");
		file.WriteString("</table>");
		if i > 100 { break; }
	}
	file.WriteString("</body>\n");
	file.WriteString("</html>\n");
}
