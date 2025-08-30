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