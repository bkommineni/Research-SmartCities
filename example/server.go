package main

import (
   "fmt"
   "strconv"
   "net"
   "time"
   "log"
   //"io/ioutil"
   "os"
   "encoding/gob"
)

type SensorData struct {
        Temperature float32
        Humidity float32
}

var batch_counter int = 0

func handleConnection(conn net.Conn) {
   //fmt.Println("Connection Accepted!!")
   //var readings []SensorData
   var RESERVOIR_SIZE = 4000
   reservoir := make([]SensorData, RESERVOIR_SIZE)
   //var reservoir [RESERVOIR_SIZE]SensorData//reservoir
   dec := gob.NewDecoder(conn)

   dec.Decode(&reservoir)
   batch_counter = batch_counter + 1
   //fmt.Println("Batch No: ",batch_counter,"No of entries: " , len(reservoir))
   current_time := time.Now().Local()
    new_file, err := os.Create(current_time.Format("20060102150405") +".txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer new_file.Close()
   file, err := os.OpenFile(current_time.Format("20060102150405") +".txt", os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()

   for _, reading := range reservoir { 
     //err := ioutil.WriteFile(current_time.Format("20060102150405") +".txt", []byte("Temperature : " + strconv.FormatFloat(float64(reading.Temperature), 'f', -1, 32) + " Humidity : " + strconv.FormatFloat(float64(reading.Humidity),'f',-1,32)),0600)
	
	len,err := file.WriteString("Temperature : " + strconv.FormatFloat(float64(reading.Temperature),     'f', -1, 32) + " Humidity : " + strconv.FormatFloat(float64(reading.Humidity),'f',-1,32) + "\n")
	_ = len
	if err != nil {
            panic(err)
     }
     //fmt.Println("Temperature ",reading.Temperature,", Humidty ",reading.Humidity)
   }

   conn.Close()
}

func main() {
   fmt.Println("Starting Server.....")
   ln, err := net.Listen("tcp", ":8085")
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

