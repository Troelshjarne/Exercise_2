# Exercise_2
## Timekeeper
The Timekeeper is the oldest current member of the system, as determined by Lamport Timestamps or by being the only member of the system.  
The Timekeeper maintains the Lamport Timestamp of the system. When a node intends to send a request to the rest of the system, they first ask the Timekeeper for the current Lamport Timestamp. Upon being requested a Lamport Timestamp, the Timekeeper responds to the request with the current Lamport Timestamp, then increments the timestamp and broadcasts it to all other nodes (for the sake of maintaining the current Timestamp, should the Timekeeper drop out of the system).  
Any of the following actions result in an increase of the Lamport Timestamp:
 - A client joins the system.
 - A client requests access to the critical section.
 - A client relinquishes access to the critical section.
The system assumes all members act in good faith; that no one would lie about Timestamps.
