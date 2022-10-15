# go-logger
My Custom Logger. Uploaded here because I'm lazy to copypasta everywhere. If you wanna use it, please keep in mind that I wrote this with **NO PERFORMANCE ISSUE** in mind.

## Usage
```go
package main

import (
  "errors"

  "github.com/Meonako/go-logger"
)

func init() {
  // You can configure here. You can do it in main too but this way it looks cleaner.
  logger.Settings.LogToFile = false // Default = true
  // Or one liner. Last three is optional
  logger.Settings.Set(true, "go-log", "-", "-", "info", "warn", "error") // "-" means leave it default.
  
  // Thru this function, I provide an easy way to format date.
  logger.Settings.Set(true, "log", "dd_mm_yyyy hh-mm-ss", "dd/mm/yyyy hh:mm:ss") // "dd_mm_yyyy hh-mm-ss" will be replace with "02_01_2006 15-04-05"
  
  // it replace "dd" with "02", "yyyy" with "2006". so it will work no matter the position or symbol. 
  // but keep in mind tho, file name can't include "/", or something that os don't support ( i.e. ":" on windows )
  logger.Settings.Set(true, "-", "yyyy-mm-dd mm!ss!hh", "yyyy/mm/dd mm:ss:hh") // This will work too!
  
  // Or if you prefer readability
  logger.NewSettings(logger.Config{
    LogToFile: true,
    LogFolder: "my-log",
    LogFileName: "mm-dd-yyyy hh_mm_ss",
    DateFormat: "mm-dd-yyyy hh:mm:ss:,
    InfoPrefix: "-", // "-" still means leave it default. If not set || empty string, It will be empty string
    WarnPrefix: "-",
    ErrorPrefix: "-",
  })
  
  logger.Init() // This will required when LogToFile = true
}

func main() {
  logger.Info("ayo")                                              // 15/10/2022 20:37 [ INFO ]: ayo
  logger.Infof("custom %v", "format")                             // 15/10/2022 20:37 [ INFO ]: custom format
  
  // Warn is basically Info but I will show you why later.
  logger.Warn("Test warn")                                        // 15/10/2022 20:37 [ WARN ]: Test warn
  logger.Warnf("test custom %v format", "warn")                   // 15/10/2022 20:37 [ WARN ]: test custom warn format
  
  // Error is not just log, but it will stop your program too. (that's the one and only reason why warn exists)
  // When call, it will log your message then wait for user to press any key, then exit with code 1
  logger.Error("Remove this if you clone it")                     // 15/10/2022 20:37 [ ERROR ]: Remove this if you clone it
                                                                  // Press any key to exit...
                                                            
  logger.Errorf("Remove this %v", "too.")                         // 15/10/2022 20:37 [ ERROR ]: Remove this too.
                                                                  // Press any key to exit...
                                                            
  // Now we will take a look at one of the (actually only me) most used function. Spicy.
  var err error = nil
  logger.WarnIf(err)                                              // This will do nothing because err == nil. But if we do this
  logger.WarnIf(errors.New("test err"))                           // 15/10/2022 20:38 [ WARN ]: test err  <-- from error.Error()
  
  err = errors.New("dont use this")
  
  // But what if you don't wanna log error.Error()? Here's how.
  logger.WarnIf(err, "use this")                                  // 15/10/2022 20:38 [ WARN ]: use this <-- from second argument. 
  // Do note tho, if you pass anything after "use this", it will still only use "use this". For example,
  logger.WarnIf(err, "arg 1", "arg 2")                            // 15/10/2022 20:38 [ WARN ]: arg 1
  
  err = errors.New("unexpected err")
  
  // We have format log too!
  logger.WarnIff(err, "Unknown Error: %v")                        // 15/10/2022 20:38 [ WARN ]: Unknown Error: unexpected err  <-- from error.Error()
  // But.. what if you have more variables? Here's how.
  logger.WarnIff(err, "%v here. Level %v : %v", "error", 1, err)  // 15/10/2022 20:38 [ WARN ]: error here. Level 1 : unexpected err
  
  // Now we have this for error level too. Working the same way as WARN entirely.
  logger.ErrorIf(err)                                             // 15/10/2022 20:38 [ WARN ]: unexpected err
                                                                  // Press any key to exit...
}
```

## Default value
```go
var Settings = Config{
  LogToFile:   true,
  LogFolder:   "log",
  LogFileName: "02_01_2006 15-04-05.log",

  DateFormat:  "02/01/2006 15:04:05",

  InfoPrefix:  "[ INFO ]:",
  WarnPrefix:  "[ WARN ]:",
  ErrorPrefix: "[ ERROR ]:",
}
```
## License even though no one will use it
MIT License

## Q&A
### Why default file name is "dd_mm_yyyy hh-mm-ss"?
* Thai date format
### Why wrote this?
* I mean.. (below) is so pain. it's become boilerplate. I JUST CANT.
```go
if err != nil { 
  log.Fatal("err")
}
```
* Second reason is, most of the time, I log to file too. And log to file kinda pain to do it everytime.
