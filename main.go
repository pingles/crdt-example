package main

import (
  "flag"
  "log"
  "github.com/hashicorp/memberlist"
  crdt "github.com/pingles/crdt-go"
  "os"
)

func main() {
  var join string
  var port int
  var name string
  var http string
  
  hostname, _ := os.Hostname()
  
  flag.StringVar(&name, "name", hostname, "name of node, must be unique")
  flag.StringVar(&join, "join", "", "address of other node to join")
  flag.StringVar(&http, "http", ":8080", "http binding")
  flag.IntVar(&port, "listen", 7946, "listen port")
  flag.Parse()

  counter := crdt.NewCounter(name)
  gossiper := NewMemberlistGossiper(name, counter)

  config := memberlist.DefaultLocalConfig()
  config.Delegate = gossiper
  config.BindPort = port
  config.Name = name
  list, err := memberlist.Create(config)

  if err != nil {
    log.Fatal(err)
  }
  
  if join != "" {
    _, err = list.Join([]string{join})
    if err != nil {
      log.Println("error joining:", err)
    }
  }
  
  log.Printf("listening %s:%d", config.BindAddr, config.BindPort)
  
  for _, member := range list.Members() {
    log.Printf("member %s: %s:%d\n", member.Name, member.Addr, member.Port)
  }
  
  ServeCounter(counter, http)
}