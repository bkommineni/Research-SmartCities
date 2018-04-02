package main

import (
        "fmt"
        "log"
        "encoding/gob"
        "net"
        "github.com/d2r2/go-dht"
        "time"
)

type SensorData struct {
        Temperature float32
        Humidity float32
}

var batch_counter int = 0

 func main() {
   //in seconds give inerval to send data
   var interval_to_send_data time.Duration = 5
   readings := make([]SensorData,0)
   ticker := time.NewTicker(time.Second * interval_to_send_data)

   go func() {
      for t := range ticker.C {
        _ = t

        fmt.Println("Client")
        fmt.Println("Start Client")
        conn, err := net.Dial("tcp", "localhost:8081")
        if err != nil {
           log.Fatal("Connection error", err)
        }
        batch_counter = batch_counter + 1
        fmt.Println("Batch No: ",batch_counter,"No of entries: " , len(readings))
        for _, reading := range readings {
          fmt.Println("Temperature ",reading.Temperature,", Humidty ",reading.Humidity);
        }
        encoder := gob.NewEncoder(conn)
        encoder.Encode(readings)
        conn.Close()
        fmt.Println("Done sending to server");

        readings = []SensorData{}
    }
   }()

   for {
     sensorType := dht.DHT11

     temperature, humidity,err :=
             dht.ReadDHTxx(sensorType, 27,false)
     if err == nil {
       sensorReading := SensorData{Temperature:temperature,Humidity:humidity}
       readings = append(readings,sensorReading)
     }
   }

   //uncomment when trying to send data only for specific amount of time
   //without sending continuously
   /*time.Sleep(time.Millisecond * 5000)
   ticker.Stop()
   fmt.Println("Ticker stopped")*/
}

