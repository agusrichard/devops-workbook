# System Design Primer

</br>

## List of Contents:

### 1. [Performance vs scalability](#content-1)
### 2. [Latency vs throughput](#content-2)
### 3. [Availability vs consistency](#content-3)
### 4. [Domain Name System](#content-4)
### 5. [Content delivery network](#content-5)
### 6. [Load balancer](#content-6)
### 7. [Reverse proxy (web server)](#content-7)
### 8. [Application layer](#content-8)

</br>

---

## Contents:
## [Performance vs scalability](https://github.com/donnemartin/system-design-primer#performance-vs-scalability) <span id="content-1"></span>
- A service is scalable if it results in increased performance in a manner proportional to resources added.
- If you have a performance problem, your system is slow for a single user.
- If you have a scalability problem, your system is fast for a single user but slow under heavy load.
- Increasing performance in general means serving more units of work, but it can also be to handle larger units of work, such as when datasets grow.
- Introducing redundancy is an important first line of defense against failures. An always-on service is said to be scalable if adding resources to facilitate redundancy does not result in a loss of performance.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Latency vs throughput](https://github.com/donnemartin/system-design-primer#latency-vs-throughput) <span id="content-2"></span>
- Latency is the time to perform some action or to produce some result.
- Throughput is the number of such actions or results per unit of time.
- Generally, you should aim for maximal throughput with acceptable latency.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Availability vs consistency](https://github.com/donnemartin/system-design-primer#availability-vs-consistency) <span id="content-1"></span>
- CAP theorem:
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/bgLMI2u.png)
- In a distributed computer system, you can only support two of the following guarantees:
  - Consistency - Every read receives the most recent write or an error
  - Availability - Every request receives a response, without guarantee that it contains the most recent version of the information
  - Partition Tolerance - The system continues to operate despite arbitrary partitioning due to network failures
- Networks aren't reliable, so you'll need to support partition tolerance. You'll need to make a software tradeoff between consistency and availability.
- CP - consistency and partition tolerance: Waiting for a response from the partitioned node might result in a timeout error. CP is a good choice if your business needs require atomic reads and writes.
- AP - availability and partition tolerance: Responses return the most readily available version of the data available on any node, which might not be the latest. Writes might take some time to propagate when the partition is resolved. 
- AP is a good choice if the business needs allow for eventual consistency or when the system needs to continue working despite external errors.

### Consistency patterns
- Weak consistency
  - After a write, reads may or may not see it. A best effort approach is taken.
  - This approach is seen in systems such as memcached. Weak consistency works well in real time use cases such as VoIP, video chat, and realtime multiplayer games. For example, if you are on a phone call and lose reception for a few seconds, when you regain connection you do not hear what was spoken during connection loss.
- Eventual consistency
  - After a write, reads will eventually see it (typically within milliseconds). Data is replicated asynchronously.
  - This approach is seen in systems such as DNS and email. Eventual consistency works well in highly available systems.
- Strong consistency
  - After a write, reads will see it. Data is replicated synchronously.
  - This approach is seen in file systems and RDBMSes. Strong consistency works well in systems that need transactions.

### Availability patterns
- There are two complementary patterns to support high availability: fail-over and replication.
- Fail-over
  - Active-passive
    - With active-passive fail-over, heartbeats are sent between the active and the passive server on standby
    - If the heartbeat is interrupted, the passive server takes over the active's IP address and resumes service.
    - The length of downtime is determined by whether the passive server is already running in 'hot' standby or whether it needs to start up from 'cold' standby. Only the active server handles traffic.
    - Active-passive failover can also be referred to as master-slave failover.
  - Active-active
    - In active-active, both servers are managing traffic, spreading the load between them.
    - Active-active failover can also be referred to as master-master failover.
- Disadvantage(s): failover
  - Fail-over adds more hardware and additional complexity.
  - There is a potential for loss of data if the active system fails before any newly written data can be replicated to the passive.

### Availability in numbers
- Availability is often quantified by uptime (or downtime) as a percentage of time the service is available. Availability is generally measured in number of 9s--a service with 99.99% availability is described as having four 9s.
- 99.9% availability - three 9s -> Downtime per year: 8h 45min 57s
- 99.99% availability - four 9s -> Downtime per year: 52min 35.7s
- Availability in parallel vs in sequence:
  - In sequence:
    - Availability (Total) = Availability (Foo) * Availability (Bar)
  - In parallel
    - Availability (Total) = 1 - (1 - Availability (Foo)) * (1 - Availability (Bar))

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Domain Name System](https://github.com/donnemartin/system-design-primer#domain-name-system) <span id="content-4"></span>
- A Domain Name System (DNS) translates a domain name such as www.example.com to an IP address.
- DNS is hierarchical, with a few authoritative servers at the top level. Your router or ISP provides information about which DNS server(s) to contact when doing a lookup.
- Lower level DNS servers cache mappings, which could become stale due to DNS propagation delays.
- DNS results can also be cached by your browser or OS for a certain period of time, determined by the time to live (TTL).
- NS record (name server) - Specifies the DNS servers for your domain/subdomain.
- MX record (mail exchange) - Specifies the mail servers for accepting messages.
- A record (address) - Points a name to an IP address.
- CNAME (canonical) - Points a name to another name or CNAME (example.com to www.example.com) or to an A record.
- Some DNS services can route traffic through various methods:
  - Weighted round robin
    - Prevent traffic from going to servers under maintenance
    - Balance between varying cluster sizes
    - A/B testing
  - Latency-based
  - Geolocation-based
- Disadvantage(s): DNS
  - Accessing a DNS server introduces a slight delay, although mitigated by caching described above.
  - DNS server management could be complex and is generally managed by governments, ISPs, and large companies.
  - DNS services have recently come under DDoS attack, preventing users from accessing websites such as Twitter without knowing Twitter's IP address(es).

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Content Delivery Network](https://github.com/donnemartin/system-design-primer#content-delivery-network) <span id="content-5"></span>
- A content delivery network (CDN) is a globally distributed network of proxy servers, serving content from locations closer to the user.
- Generally, static files such as HTML/CSS/JS, photos, and videos are served from CDN, although some CDNs such as Amazon's CloudFront support dynamic content. The site's DNS resolution will tell clients which server to contact.
- Serving content from CDNs can significantly improve performance in two ways:
  - Users receive content from data centers close to them
  - Your servers do not have to serve requests that the CDN fulfills
- Push CDNs
  - Push CDNs receive new content whenever changes occur on your server. You take full responsibility for providing content, uploading directly to the CDN and rewriting URLs to point to the CDN. You can configure when content expires and when it is updated. Content is uploaded only when it is new or changed, minimizing traffic, but maximizing storage.
  - Sites with a small amount of traffic or sites with content that isn't often updated work well with push CDNs. Content is placed on the CDNs once, instead of being re-pulled at regular intervals.
- Pull CDNs
  - Pull CDNs grab new content from your server when the first user requests the content. You leave the content on your server and rewrite URLs to point to the CDN. This results in a slower request until the content is cached on the CDN.
  - A time-to-live (TTL) determines how long content is cached. Pull CDNs minimize storage space on the CDN, but can create redundant traffic if files expire and are pulled before they have actually changed.
  - Sites with heavy traffic work well with pull CDNs, as traffic is spread out more evenly with only recently-requested content remaining on the CDN.
- Disadvantage(s): CDN
  - CDN costs could be significant depending on traffic, although this should be weighed with additional costs you would incur not using a CDN.
  - Content might be stale if it is updated before the TTL expires it.
  - CDNs require changing URLs for static content to point to the CDN.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Load balancer](https://github.com/donnemartin/system-design-primer#load-balancer) <span id="content-6"></span>
- Image
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/h81n9iK.png)
- Load balancers distribute incoming client requests to computing resources such as application servers and databases.
- In each case, the load balancer returns the response from the computing resource to the appropriate client. Load balancers are effective at:
  - Preventing requests from going to unhealthy servers
  - Preventing overloading resources
  - Helping to eliminate a single point of failure
- To protect against failures, it's common to set up multiple load balancers, either in active-passive or active-active mode.
- Load balancers can route traffic based on various metrics, including:
  - Random
  - Least loaded
  - Session/cookies
  - Round robin or weighted round robin
  - Layer 4
  = Layer 7
- Layer 4 load balancing
  - Layer 4 load balancers look at info at the transport layer to decide how to distribute requests. Generally, this involves the source, destination IP addresses, and ports in the header, but not the contents of the packet. Layer 4 load balancers forward network packets to and from the upstream server, performing Network Address Translation (NAT).
- Layer 7 load balancing
  - Layer 7 load balancers look at the application layer to decide how to distribute requests. This can involve contents of the header, message, and cookies. Layer 7 load balancers terminate network traffic, reads the message, makes a load-balancing decision, then opens a connection to the selected server. For example, a layer 7 load balancer can direct video traffic to servers that host videos while directing more sensitive user billing traffic to security-hardened servers.
- At the cost of flexibility, layer 4 load balancing requires less time and computing resources than Layer 7, although the performance impact can be minimal on modern commodity hardware.
- Horizontal scaling
  - Load balancers can also help with horizontal scaling, improving performance and availability. Scaling out using commodity machines is more cost efficient and results in higher availability than scaling up a single server on more expensive hardware, called Vertical Scaling.
- Disadvantage(s): horizontal scaling
  - Scaling horizontally introduces complexity and involves cloning servers
    - Servers should be stateless: they should not contain any user-related data like sessions or profile pictures
    - Sessions can be stored in a centralized data store such as a database (SQL, NoSQL) or a persistent cache (Redis, Memcached)
  - Downstream servers such as caches and databases need to handle more simultaneous connections as upstream servers scale out
- Disadvantage(s): load balancer:
  - The load balancer can become a performance bottleneck if it does not have enough resources or if it is not configured properly.
  - Introducing a load balancer to help eliminate a single point of failure results in increased complexity.
  - A single load balancer is a single point of failure, configuring multiple load balancers further increases complexity.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Reverse proxy (web server)](https://github.com/donnemartin/system-design-primer#reverse-proxy-web-server) <span id="content-7"></span>
- A reverse proxy is a web server that centralizes internal services and provides unified interfaces to the public.
- Requests from clients are forwarded to a server that can fulfill it before the reverse proxy returns the server's response to the client.
- Benefits of using reverse proxy:
  - Increased security - Hide information about backend servers, blacklist IPs, limit number of connections per client
  - Increased scalability and flexibility - Clients only see the reverse proxy's IP, allowing you to scale servers or change their configuration
  - Compression - Compress server responses
  - Caching - Return the response for cached requests
  - Static content - Serve static content directly
- Load balancer vs reverse proxy
  - Deploying a load balancer is useful when you have multiple servers. Often, load balancers route traffic to a set of servers serving the same function.
  - Reverse proxies can be useful even with just one web server or application server, opening up the benefits described in the previous section.
  - Solutions such as NGINX and HAProxy can support both layer 7 reverse proxying and load balancing.
- Disadvantage(s): reverse proxy
  - Introducing a reverse proxy results in increased complexity.
  - A single reverse proxy is a single point of failure, configuring multiple reverse proxies (ie a failover) further increases complexity.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Application layer](https://github.com/donnemartin/system-design-primer#application-layer) <span id="content-8"></span>
- Separating out the web layer from the application layer (also known as platform layer) allows you to scale and configure both layers independently.
- Adding a new API results in adding application servers without necessarily adding additional web servers.
- The single responsibility principle advocates for small and autonomous services that work together. Small teams with small services can plan more aggressively for rapid growth.
- Microservices
  - Related to this discussion are microservices, which can be described as a suite of independently deployable, small, modular services.
  - Each service runs a unique process and communicates through a well-defined, lightweight mechanism to serve a business goal.
  - Pinterest, for example, could have the following microservices: user profile, follower, feed, search, photo upload, etc.
- Service Discovery
  - Systems such as Consul, Etcd, and Zookeeper can help services find each other by keeping track of registered names, addresses, and ports.
  - Health checks help verify service integrity and are often done using an HTTP endpoint. Both Consul and Etcd have a built in key-value store that can be useful for storing config values and other shared data.
- Disadvantage(s): application layer
  - Adding an application layer with loosely coupled services requires a different approach from an architectural, operations, and process viewpoint (vs a monolithic system).
  - Microservices can add complexity in terms of deployments and operations.


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://github.com/donnemartin/system-design-primer#next-steps