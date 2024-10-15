TCP vs UDP Crash Course
# TCP (Transmission Control Protocol)

## Overview
- **DB use TCP all the time**

## TCP Pros and Cons

### Pros
* **Acknowledgment**
    * TCP ensures that data sent from the sender is acknowledged by the receiver. After the receiver successfully receives a data packet, it sends an acknowledgment (ACK) back to the sender.
* **Guaranteed delivery**
    * Data sent from one device will reach the destination device without errors and in the correct order. This is done through mechanisms like acknowledgment, error-checking, and retransmission of lost or corrupted packets.
* **Connection-based**
    * TCP is connection-oriented, meaning a unique connection is established between the sender and receiver before any data is transmitted.
* **Congestion control**
    * If there are many requests, TCP waits and sends data only when the network can handle it.
* **Ordered packets**
    * TCP adds headers to packets and labels them to ensure they are received in order.

### Cons
* **Larger packets**
* **More bandwidth (bps)**
    * Bandwidth is the maximum amount of data that can be transmitted over a network connection in a given amount of time.
* **Slower than UDP**
* **Stateful**
    * If the connection drops, clients can’t resume their work. Both server and clients carry information about their connection.
* **Server memory (DOS)**
    * Need to allocate memory for each connection, which limits the number of TCP connections a server can handle.

## TCP Code Example

```go
package main

import (
    "bufio"
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    // Send a message to the client
    fmt.Fprintln(conn, "Hello.")

    // Read data from the client
    reader := bufio.NewReader(conn)
    for {
        message, _ := reader.ReadString('\n')
        fmt.Println("Received:", message)
    }
}

func main() {
    // Start listening on port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080...")

    for {
        // Accept connection
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn) // Handle each connection concurrently
    }
}
```

## Use Cases
* Texting app
    * Requires ordered delivery

# UDP (User Datagram Protocol)

## Pros and Cons

### Pros
* **Smaller packets**
* **Less bandwidth**
* **Faster than TCP**
* **Stateless**

### Cons
* **No Acknowledgment**
* **No Guaranteed delivery**
    * Only performs checksum
* **Connectionless**
    * Client and server don’t know each other
* **No Congestion control**
* **No Ordered packets**
* **Security**
    * Lack of identification

## UDP Code Example

```go
package main

import (
    "fmt"
    "net"
)

const port = 8081

func main() {
    addr := net.UDPAddr{
        Port: port,
        // IP:   net.ParseIP("0.0.0.0"),
    }

    conn, err := net.ListenUDP("udp4", &addr)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    defer conn.Close()
    fmt.Println("Server is running on port", port)

    buffer := make([]byte, 1024)

    for {
        n, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error: ", err)
            continue
        }

        fmt.Printf("server got: %s from %s\n", string(buffer[:n]), remoteAddr)
    }
}
```

## Use Cases
* Online game
* Video streaming
    * Applications that can tolerate some data loss

[Video Explanation](https://www.youtube.com/watch?v=qqRYkcta6IE&list=PLQnljOFTspQUNnO4p00ua_C5mKTfldiYT&index=5)