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