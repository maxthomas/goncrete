// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "goncrete"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void submitSort(UUID sessionId,  unitIds)")
  fmt.Fprintln(os.Stderr, "  ServiceInfo about()")
  fmt.Fprintln(os.Stderr, "  bool alive()")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := goncrete.NewActiveLearnerClientServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "submitSort":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SubmitSort requires 2 args")
      flag.Usage()
    }
    arg49 := flag.Arg(1)
    mbTrans50 := thrift.NewTMemoryBufferLen(len(arg49))
    defer mbTrans50.Close()
    _, err51 := mbTrans50.WriteString(arg49)
    if err51 != nil {
      Usage()
      return
    }
    factory52 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt53 := factory52.GetProtocol(mbTrans50)
    argvalue0 := goncrete.NewUUID()
    err54 := argvalue0.Read(jsProt53)
    if err54 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg55 := flag.Arg(2)
    mbTrans56 := thrift.NewTMemoryBufferLen(len(arg55))
    defer mbTrans56.Close()
    _, err57 := mbTrans56.WriteString(arg55)
    if err57 != nil { 
      Usage()
      return
    }
    factory58 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt59 := factory58.GetProtocol(mbTrans56)
    containerStruct1 := goncrete.NewActiveLearnerClientServiceSubmitSortArgs()
    err60 := containerStruct1.ReadField2(jsProt59)
    if err60 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.UnitIds
    value1 := argvalue1
    fmt.Print(client.SubmitSort(value0, value1))
    fmt.Print("\n")
    break
  case "about":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "About requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.About())
    fmt.Print("\n")
    break
  case "alive":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Alive requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Alive())
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
