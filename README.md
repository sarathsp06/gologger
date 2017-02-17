![Logo of the project](http://www.davesgames.net/papercraft/png/gallery-logs-01.png)

# gologger
![Build Status](https://api.travis-ci.org/sarathsp06/gologger.svg?branch=master)
[![Join the chat at https://gitter.im/gologger/Lobby](https://badges.gitter.im/gologger/Lobby.svg)](https://gitter.im/gologger/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Go Report Card](https://goreportcard.com/badge/github.com/sarathsp06/gologger)](https://goreportcard.com/report/github.com/sarathsp06/gologger)
> Logger library for the gophers

An openionated logging library for golang.Gologger writes logs as json in the following format

```json
{
    "time":"2016-11-10T16:11:46.591304719+05:30",
    "process_name":"sample_logger",
    "host_name":"sarath.local",
    "process_id":29422,
    "level":"ERROR",
    "file_name":"/home/sarath/go/src/github.com/sarathsp06/gologger/sample/main.go",
    "line_num":13,
    "log_msg":"error happened"
}
```  

That much info will be how following will be written
```golang
logger.Error("error happened")
```

## Installing / Getting started

Just like any other go library

```shell
go get github.com/sarathsp06/gologger
```
But if one needs to get a particular version then use 

```shell
go get gopkg.in/sarathsp06/gologger.vx #x is the version Number
``` 


## Developing

Here's a brief intro about what a developer must do in order to start developing
the project further:

1. Get it using go get or clone
2. Make changes and make a pull reqeuest with Updated README if featue addition 


## Features

What's all the bells and whistles this project can perform?
* It can log - Obviously
* Easy buffered logging 
* You can set write to logger for redirecting logs
* log message will contain much deeper details like line number of error , process id,host_name  etc


## How to use
This is how one may use the library

import the package 
```golang
import logger "github.com/sarathsp06/gologger"
``` 
Initialize the  package 
```golang
if err := logger.InitLogger("INFO", ".", "sample_logger"); err != nil {
		panic(err.Error())
}
```
Log as you wish with formated and normal message

```golang
logger.Error("error happened")
logger.Debug("Debug message",":Sasa","sasa")
logger.Info("error happened")
logger.Warning("error happened")
logger.Errorf("error happened %s", "Yo")
logger.Debugf("debug message : %s", "YoYo")
```

Explore more ....


Here is a sample code

```golang
package main

import (
	"os"
	logger "github.com/sarathsp06/gologger"
)

func main() {
	if err := logger.InitLogger("INFO", ".", "sample_logger"); err != nil {
		panic(err.Error())
	}
	logger.SetLogWriter(os.Stdout)
	logger.Error("error happened")
	logger.Debug("Debug message")
	logger.Info("error happened")
	logger.Warning("error happened")
	logger.Errorf("error happened %s", "Yo")
	logger.Debugf("debug message : %s", "YoYo")
}
```

## Licensing


"The code in this project is licensed under MIT license."
