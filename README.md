# Exercise_2
## Architecture
This implementation is completely distributed, such that it is completely peer-to-peer and every client is equal.  
[Serf](https://github.com/hashicorp/serf) is used, mainly for two specefic purposes. Firstly to connect all the client to one another, and secondly to maintain the lamport time.  
Any given client can, at any time, decide that it wishes to gain access to the critical section, at which point it sends out a request to every other client, to allow it access.
The client is subsequently placed into a queue by every other client, and only once every other client has answered it's request, will the client be granted access to the critical section.

## Access requests
Asking the other clients for permission is all well and good, but inevitably two clients' access requests will collide, and the system has to somehow decide, which client has priority.  
This is why, with every access request a client makes, it first asks Serf for the current lamport time (which Serf then increments, to avoid collisions) and attaches that time to it's access request message.  

When a client receives an access request message, there are then 4 possible scenarios:
 1. The client is currently idle, and simply responds to the request.
 2. The client is currently itself in the critital section, puts the request on hold, and responds once it has exited the critial section.
 3. The client is itself waiting for access to the critical section and made it's request at an earlier lamport time (own request pre-dates incoming request). Similairly to Scenario 2, the client waits until it has both gained and relinqueshed access to the critical section before answering the incoming request.
 4. The client is itself waiting for access to the critical section and made it's request at a later lamport time (incoming request pre-dates own request). The client then reluctantly grants permission to the incoming request over its own. The client then implicitly waits for the other client to finish, since its request is stuck in scenario 3 for the other client.

## System faults
Firstly, this system relies on honesty and is easily vulnerable to malicious actors. If a client decides to simply lie about what the lamport time was when they made their access request, there is nothing stopping them, and they can completely skip the queue. If one were inclined to shut the system down, they could refuse to answer access requests from other clients, or refuse to relinquish their current access to the critial section.
