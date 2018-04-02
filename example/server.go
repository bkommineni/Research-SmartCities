package main

import (
   "fmt"
   "net"
   "encoding/gob"
)

type SensorData struct {
        Temperature float32
        Humidity float32
}

var batch_counter int = 0

func handleConnection(conn net.Conn) {
   fmt.Println("Connection Accepted!!")
   var readings []SensorData
   dec := gob.NewDecoder(conn)

   dec.Decode(&readings)
   batch_counter = batch_counter + 1
   fmt.Println("Batch No: ",batch_counter,"No of entries: " , len(readings))
   for _, reading := range readings {
     fmt.Println("Temperature ",reading.Temperature,", Humidty ",reading.Humidity)
   }

   conn.Close()
}

func main() {
   fmt.Println("Starting Server.....")
   ln, err := net.Listen("tcp", ":8081")
   if err != nil {
      panic(err)
   }
   fmt.Println("Ready to accept connections!!")
   for {
      conn, err := ln.Accept()
      if err != nil {
         continue
      }
      go handleConnection(conn)
   }
}

