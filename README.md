# go-logger
My Custom Logger. Uploaded here because I'm lazy to copypasta everywhere. If you wanna use it, please keep in mind that I wrote this with **NO PERFORMANCE ISSUE** in mind.
## Usage
go```
package main

import (
  "errors"

  "github.com/Meonako/go-logger"
)

func init() {
  // You can configure here. You can do it in main too but this way it looks cleaner.
  logger.Settings.LogToFile = false // Default = true
  // Or one liner
  logger.Settings.Set(true, "go-log", "-", "-", "-", "-") // "-" means leave it default.
  logger.Init() // This will required when LogToFile = true
}

func main() {
  logger.Info("ayo")                                        // 15/10/2022 20:37 [ INFO ]: ayo
  logger.Infof("custom %v", "format")                       // 15/10/2022 20:37 [ INFO ]: custom format
  
  // Warn is basically Info but I will show you why later.
  logger.Warn("Test warn")                                  // 15/10/2022 20:37 [ WARN ]: Test warn
  logger.Warnf("test custom %v format", "warn")             // 15/10/2022 20:37 [ WARN ]: test custom warn format
  
  // Error is not just log, but it will stop your program too. (that's the one and only reason why warn exists)
  // When call, it will log your message then wait for user to press any key, then exit with code 1
  logger.Error("Remove this if you clone it")               // 15/10/2022 20:37 [ ERROR ]: Remove this if you clone it
                                                            // Press any key to exit...
                                                            
  logger.Errorf("Remove this %v", "too.")                   // 15/10/2022 20:37 [ ERROR ]: Remove this too.
                                                            // Press any key to exit...
                                                            
  // Now we will take a look at one of the (actually only me) most used function. Spicy.
  var err error = nil
  logger.WarnIf(err)                                       // This will do nothing because err == nil. But if we do this
  logger.WarnIf(errors.New("test err"))                    // 15/10/2022 20:38 [ WARN ]: test err.  <-- from error.Error()
  // But what if you don't wanna log error.Error()? Here's how.
  logger.WarnIf(errors.New("dont use this"), "eeeeeee")    // 15/10/2022 20:38 [ WARN ]: eeeeeee. <-- from second argument. Do note tho, if you pass anything after "eeeeeee" it will not work!
}
```
