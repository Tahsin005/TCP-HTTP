## TCP

```
Transmission Control Protocol (TCP) is a primary communication protocol of the internet, though that is changing with HTTP3 (which is not built on TCP) gaining adoption.

TCP is great because it allows ordered data to be safely sent across the internet. For example, let's say we want to send the message "i am live":

text 	binary
i 	    01101001
a 	    01100001
m 	    01101101
l 	    01101100
i 	    01101001
v 	    01110110
e 	    01100101

When data is sent over a network, it is sent in packets. Each message is split into packets, the packets are sent, they arrive (potentially) out of order, and they are reassembled on the other side. And without a protocol like TCP, you can't guarantee that the order is correct...

You might end up with "i am evil" instead of "i am live"! TCP solves this problem.
```

## TCP vs UDP

```
User Datagram Protocol (UDP) is often compared to TCP, as they are both transport layer protocols. Here are the high-level differences between the two:
	            TCP 	UDP
Connection 	    Yes 	No
Handshake 	    Yes 	No
In Order 	    Yes 	No
Blazingly Fast 	No 	    Yes

TCP establishes a connection between sender and receiver with a handshake, and ensures that all the data is sent in order. UDP yeets the data to the receiver and hopes they can make sense of it.

    nc -v localhost 42069

        nc → Netcat, a tool for reading/writing to network connections.
        -v → Verbose mode, gives more connection details.
        localhost → Target host (your own machine).
        42069 → Target port number.

        This command tries to open a TCP connection from your terminal to port 42069 on your local machine, and prints connection details.
```

## UDP Sender

```
What's true of UDP?
- It doesn't require a handshake
```

## Files Vs Network

```
Files and network connections behave very similarly - that's why we started by simply reading and writing to files, then updated our code to be a bit more abstract (the getLinesChannel function) so that it can handle both. From the perspective of your code, files and network connections are both just streams of bytes that you can read from and write to.

All of a sudden, Go's io.Reader (and the very similar io.ReadCloser) and io.Writer interfaces make a lot more sense. They're designed to work with any type of stream, whether it's a file, a network connection, or something else entirely.

Pull vs. Push

When you read from a file, you're in control of the reading process. You decide:

    When to read
    How much to read
    When to stop reading.

You pull data from the file.

When you read from a network connection, the data is pushed to you by the remote server. You don't have control over when the data arrives, how much arrives, or when it stops arriving. Your code has to be ready to receive it when it comes.

The io.Reader and io.Writer interfaces in Go are useful for reading and writing data to and from ____
- Both file and network

Which pushes data to the application that's reading from it, resulting in the reader needing to be prepared to read data?
- Network connections
```

## TCP to HTTP

```
HTTP/1.1 is a text based protocol that works over TCP.

HTTP works because plain text is binary. Because HTTP uses TCP, if the HTTP request or response is too big to fit into a single TCP packet it can be broken up into many packets and reconstructed in the correct order on the other side. TCP guarantees that the data is in order and complete.

At the heart of HTTP is the HTTP-message: the format that the text in an HTTP request or response must use. From RFC 9112 Section 2.1:

start-line CRLF
*( field-line CRLF )
CRLF
[ message-body ]

CRLF (written in plain text as \r\n) is a carriage return followed by a line feed. It's a Microsoft Windows (and HTTP) style newline character.

    I call \r\n "Registered Nurse"... it helps me remember the order of the characters



Let's break down each part:
| Part                 | Example                        | Description                                                       |
|-----------------------|--------------------------------|-------------------------------------------------------------------|
| start-line CRLF       | `POST /users/tahsin HTTP/1.1` | The request (for a request) or status (for a response) line       |
| *( field-line CRLF )  | `Host: google.com`            | Zero or more lines of HTTP headers. These are key-value pairs.    |
| CRLF                  |                                | A blank line that separates the headers from the body.            |
| [ message-body ]      | `{"name": "TheHTTPTahsin"}`   | The body of the message. This is optional.                        |


Both HTTP requests and responses follow this same format, though the contents of each section will differ!
```