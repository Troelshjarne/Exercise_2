# Exercise_2
## Timekeeper
The Timekeeper is the oldest current member of the system, as determined by Lamport timestamps or by being the only member of the system.  
The Timekeeper maintains the Lamport Timestamp of the system. When a node intends to send a request to the rest of the system, they first ask the Timekeeper for the current Lamport Timestamp. Upon being requested a Lamport Timestamp, the Timekeeper responds to the request with the current Lamport Timestamp, then increments the timestamp and brodcasts it to all other nodes (for the sake of maintaining the current Timestamp, should the Timekeeper drop out of the system).
