// http://hadoop.apache.org/docs/current/hadoop-project-dist/hadoop-hdfs/WebHDFS.html
// https://github.com/vladimirvivien/gowfs

package main

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"log"
	"io/ioutil"
)

func main() {
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: "localhost:50070", User: "gouser"})
	if err != nil{
		log.Fatal(err)
	}

	checksum, err := fs.GetFileChecksum(gowfs.Path{Name: "/user/root/input_wordcount/file2"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checksum)

	data, err := fs.Open(gowfs.Path{Name:"/user/root/input_wordcount/file2"}, 0, 512, 2048)
	rcvData, _ := ioutil.ReadAll(data)
	fmt.Println(string(rcvData))

	/* 增加配置：
	hdfs-site.xml
	<property>
    	<name>dfs.webhdfs.enabled</name>
    	<value>true</value>
	</property>

	core-site.xml
	<property>
    	<name>hadoop.http.staticuser.user</name>
    	<value>gouser</value>
	</property>
	*/

	/* 程序输出：
	{MD5-of-0MD5-of-512CRC32 000002000000000000000000cd19ed08102fe3791de77156abf23edd00000000 28}
	new file
	hadoop file
	hadoop new world
	hadoop free home
	hadoop free school

	*/
}