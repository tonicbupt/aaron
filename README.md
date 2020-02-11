# Aaron

![duck](https://raw.githubusercontent.com/tonicbupt/aaron/master/pic/duck.jpg)

Aaron 就是一个朋友的名字啦.

为什么叫这个呢, 因为 Aaron 会帮助你这个选择困难症.

## Install

```
$ go get github.com/tonicbupt/aaron
$ aaron -h
NAME:
   Aaron - Aaron helps you to make decision!

USAGE:
   aaron [global options] command [command options] [arguments...]

VERSION:
   0x336ae74f

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --apikey value  API Key from random.org [$AARON_APIKEY]
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```

## Usage

* 去 [random.org](https://www.random.org/) 注册一个账号, 生成一个 API Key.
* 然后用这个 apikey 来跑

	```
	$ export AARON_APIKEY=[your api key]
	
	$ aaron x y z
	Aaron says y
	
	$ aaron 吃饭 睡觉 打豆豆
	Aaron says 打豆豆
	
	$ aaron --apikey [your api key] 他爱我 他不爱我
	Aaron says 他不爱我
	```
