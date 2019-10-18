package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//通过URL返回页面HTML
func getHttp(url string) (result string, err error) {
	resp, err1 := http.Get(url)

	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println(url, "已读完。")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

//分析笑话列表页面，返回笑话列表URL
func getURLs(html string) (urls [][]string) {

	ret := regexp.MustCompile(`<h1 class="post-title"><a href="(?s:(.*?))">`)
	urls = ret.FindAllStringSubmatch(html, -1)
	return

}

//分析笑话页面，返回笑话标题和内容
func getJoke(html string) (title string, content string) {
	ret1 := regexp.MustCompile(`<h1 class="post-title">(?s:(.*?))</h1>`)
	titles := ret1.FindAllStringSubmatch(html, 1)

	for _, t := range titles {
		title = t[1]
		break
	}
	ret2 := regexp.MustCompile(`<section class="post-content">(?s:(.*?))</section>`)
	contents := ret2.FindAllStringSubmatch(html, 1)
	for _, c := range contents {
		content = strings.Replace(c[1], "<p>", "", -1)
		content = strings.Replace(content, "</p>", "", -1)
		content = strings.Replace(content, "<br>", "", -1)
		content = strings.TrimSpace(content)

		break
	}
	return
}

func spiderPage(url string, page int, spideChan chan int) {
	//获取第page页内容
	result, err := getHttp(url)
	if err != nil {
		fmt.Println(" err is ....", err)
	}
	//返回该页中的所有笑话连接
	urls := getURLs(result)
	cMap := make(map[string]string)
	for i, v := range urls {
		fmt.Println(url, "page:", i)
		//返回每个笑话页的内容
		upage, err := getHttp(v[1])
		if err != nil {
			fmt.Println("--#######err :", err)
		}
		// fmt.Println("upage size:", len(upage))
		//提取笑话页中的笑话
		title, content := getJoke(upage)
		cMap[title] = content

	}

	mapSaveFile(cMap, "page"+strconv.Itoa(page)+".txt")
	spideChan <- page

}

//将爬行结果写入文件
func mapSaveFile(cMap map[string]string, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.create error:", err)
	}

	for k, v := range cMap {

		f.WriteString(k + "\r\n" + v + "\r\n")
		f.WriteString("---------------------------------\r\n")
	}

	defer f.Close()
	fmt.Println(fileName, "页保存完毕。。。。")
}

//爬行主程序
func runSpider(url string, start int, end int) {
	//提示信息
	fmt.Println("url:", url, " start:", start, " end:", end)
	spideChan := make(chan int)
	for i := start; i <= end; i++ {
		fmt.Println("开始爬第", i, "页")
		tURL := url + strconv.Itoa(i)
		go spiderPage(tURL, i, spideChan)
	}
	//	保护多go程并发
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页爬行完毕..........。\n", <-spideChan)
	}

}

func main() {

	var url string
	var start, end int
	fmt.Print("Input URL :")
	fmt.Scan(&url)
	fmt.Print("Input start page:")
	fmt.Scan(&start)
	fmt.Print("Input end page:")
	fmt.Scan(&end)
	runSpider(url, start, end)

}
