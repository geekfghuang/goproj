package main

import (
	"fmt"
	"os/exec"
)

func main() {
	result, err := exec.Command("/root/geekfghuang/spark-2.2.1-bin-hadoop2.7/bin/spark-submit",
		"/root/geekfghuang/spark-2.2.1-bin-hadoop2.7/python/my_script.py").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))

	result, err = exec.Command("hadoop",
		"fs", "-cat", "/user/root/input_wordcount/file2").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))

	/* my_script.py：
	from pyspark import SparkConf, SparkContext

	conf = SparkConf().setMaster("local").setAppName("My App")
	sc = SparkContext(conf = conf)
	lines = sc.textFile("/root/geekfghuang/hbase/README.txt")
	print lines.count()
	*/

	/* 程序输出：
	34

	new file
	hadoop file
	hadoop new world
	hadoop free home
	hadoop free school

	*/
}