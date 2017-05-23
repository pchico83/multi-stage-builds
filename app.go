package main

import log "github.com/Sirupsen/logrus"

func main() {
   var logger log.FieldLogger = log.StandardLogger()
   logger.Info("Hola Docker Meetup!")  
}
