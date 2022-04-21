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
### 9. [Database](#content-9)
### 10. [Cache](#content-10)
### 11. [Asynchronism](#content-11)
### 12. [Communication](#content-12)

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

## [Database](https://github.com/donnemartin/system-design-primer#database) <span id="content-9"></span>

### Relational database management system (RDBMS)
- ACID is a set of properties of relational database transactions.
  - Atomicity - Each transaction is all or nothing
  - Consistency - Any transaction will bring the database from one valid state to another
  - Isolation - Executing transactions concurrently has the same results as if the transactions were executed serially
  - Durability - Once a transaction has been committed, it will remain so
- Master-slave replication
  - The master serves reads and writes, replicating writes to one or more slaves, which serve only reads.
  - If the master goes offline, the system can continue to operate in read-only mode until a slave is promoted to a master or a new master is provisioned.
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/C9ioGtn.png)
  - Disadvantage(s): master-slave replication
    - Additional logic is needed to promote a slave to a master.
- Master-master replication
  - Both masters serve reads and writes and coordinate with each other on writes. If either master goes down, the system can continue to operate with both reads and writes.
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/krAHLGg.png)
  - Disadvantage(s): master-master replication
    - You'll need a load balancer or you'll need to make changes to your application logic to determine where to write.
    - Most master-master systems are either loosely consistent (violating ACID) or have increased write latency due to synchronization.
    - Conflict resolution comes more into play as more write nodes are added and as latency increases.
- Disadvantage(s): replication
  - There is a potential for loss of data if the master fails before any newly written data can be replicated to other nodes.
  - Writes are replayed to the read replicas. If there are a lot of writes, the read replicas can get bogged down with replaying writes and can't do as many reads.
  - The more read slaves, the more you have to replicate, which leads to greater replication lag.
  - On some systems, writing to the master can spawn multiple threads to write in parallel, whereas read replicas only support writing sequentially with a single thread.
  - Replication adds more hardware and additional complexity.
- Federation
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/U3qV33e.png)
  - Federation (or functional partitioning) splits up databases by function. For example, instead of a single, monolithic database, you could have three databases: forums, users, and products, resulting in less read and write traffic to each database and therefore less replication lag.
  - With no single central master serializing writes you can write in parallel, increasing throughput.
  - Disadvantage(s): federation
    - Federation is not effective if your schema requires huge functions or tables.
    - You'll need to update your application logic to determine which database to read and write.
    - Joining data from two databases is more complex with a server link.
  Federation adds more hardware and additional complexity.
- Sharding
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/wU8x5Id.png)
  - Sharding distributes data across different databases such that each database can only manage a subset of the data. Taking a users database as an example, as the number of users increases, more shards are added to the cluster.
  - Similar to the advantages of federation, sharding results in less read and write traffic, less replication, and more cache hits. Index size is also reduced, which generally improves performance with faster queries.
  - If one shard goes down, the other shards are still operational, although you'll want to add some form of replication to avoid data loss. Like federation, there is no single central master serializing writes, allowing you to write in parallel with increased throughput.
  - Common ways to shard a table of users is either through the user's last name initial or the user's geographic location.
  - Disadvantage(s): sharding
    - You'll need to update your application logic to work with shards, which could result in complex SQL queries.
    - Data distribution can become lopsided in a shard. For example, a set of power users on a shard could result in increased load to that shard compared to others.
    - Rebalancing adds additional complexity. A sharding function based on consistent hashing can reduce the amount of transferred data.
    - Joining data from multiple shards is more complex.
    - Sharding adds more hardware and additional complexity
- Denormalization
  - Denormalization attempts to improve read performance at the expense of some write performance.
  - Redundant copies of the data are written in multiple tables to avoid expensive joins.
  - Once data becomes distributed with techniques such as federation and sharding, managing joins across data centers further increases complexity. Denormalization might circumvent the need for such complex joins.
  - In most systems, reads can heavily outnumber writes 100:1 or even 1000:1. A read resulting in a complex database join can be very expensive, spending a significant amount of time on disk operations.
  - Disadvantage(s): denormalization
    - Data is duplicated.
    - Constraints can help redundant copies of information stay in sync, which increases complexity of the database design.
    - A denormalized database under heavy write load might perform worse than its normalized counterpart.
- SQL tuning
  - It's important to benchmark and profile to simulate and uncover bottlenecks.
  - Tighten up the schema
    - MySQL dumps to disk in contiguous blocks for fast access.
    - Use CHAR instead of VARCHAR for fixed-length fields.
    CHAR effectively allows for fast, random access, whereas with VARCHAR, you must find the end of a string before moving onto the next one.
    - Use TEXT for large blocks of text such as blog posts. TEXT also allows for boolean searches. Using a TEXT field results in storing a pointer on disk that is used to locate the text block.
    - Use INT for larger numbers up to 2^32 or 4 billion.
    - Use DECIMAL for currency to avoid floating point representation errors.
    - Avoid storing large BLOBS, store the location of where to get the object instead.
    - VARCHAR(255) is the largest number of characters that can be counted in an 8 bit number, often maximizing the use of a byte in some RDBMS.
    - Set the NOT NULL constraint where applicable to improve search performance.
  - Use good indices
    - Columns that you are querying (SELECT, GROUP BY, ORDER BY, JOIN) could be faster with indices.
    - Indices are usually represented as self-balancing B-tree that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time.
    - Placing an index can keep the data in memory, requiring more space.
    - Writes could also be slower since the index also needs to be updated.
    - When loading large amounts of data, it might be faster to disable indices, load the data, then rebuild the indices.
  - Avoid expensive joins
  - Partition tables
  - Tune the query cache

### NoSQL
- NoSQL is a collection of data items represented in a key-value store, document store, wide column store, or a graph database.
- Data is denormalized, and joins are generally done in the application code. Most NoSQL stores lack true ACID transactions and favor eventual consistency.
- BASE is often used to describe the properties of NoSQL databases. In comparison with the CAP Theorem, BASE chooses availability over consistency.
  - Basically available - the system guarantees availability.
  - Soft state - the state of the system may change over time, even without input.
  - Eventual consistency - the system will become consistent over a period of time, given that the system doesn't receive input during that period.
- Key-value store
  - A key-value store generally allows for O(1) reads and writes and is often backed by memory or SSD.
  - Key-value stores provide high performance and are often used for simple data models or for rapidly-changing data, such as an in-memory cache layer.
- Document store
  - A document store is centered around documents (XML, JSON, binary, etc), where a document stores all information for a given object. Document stores provide APIs or a query language to query based on the internal structure of the document itself. Note, many key-value stores include features for working with a value's metadata, blurring the lines between these two storage types.
  - Based on the underlying implementation, documents are organized by collections, tags, metadata, or directories. Although documents can be organized or grouped together, documents may have fields that are completely different from each other.
  - Some document stores like MongoDB and CouchDB also provide a SQL-like language to perform complex queries. DynamoDB supports both key-values and documents.
- Wide column store
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/n16iOGk.png)
  - A wide column store's basic unit of data is a column (name/value pair). A column can be grouped in column families (analogous to a SQL table). Super column families further group column families. You can access each column independently with a row key, and columns with the same row key form a row. Each value contains a timestamp for versioning and for conflict resolution.
  - Wide column stores offer high availability and high scalability. They are often used for very large data sets.
- Graph database
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/fNcl65g.png)
  - In a graph database, each node is a record and each arc is a relationship between two nodes. Graph databases are optimized to represent complex relationships with many foreign keys or many-to-many relationships.

### SQL or NoSQL
- Image:
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/wXGqG5f.png)
- Reasons for SQL:
  - Structured data
  - Strict schema
  - Relational data
  - Need for complex joins
  - Transactions
  - Clear patterns for scaling
  - More established: developers, community, code, tools, etc
  - Lookups by index are very fast
- Reasons for NoSQL:
  - Semi-structured data
  - Dynamic or flexible schema
  - Non-relational data
  - No need for complex joins
  - Store many TB (or PB) of data
  - Very data intensive workload
  - Very high throughput for IOPS
- Sample data well-suited for NoSQL:
  - Rapid ingest of clickstream and log data
  - Leaderboard or scoring data
  - Temporary data, such as a shopping cart
  - Frequently accessed ('hot') tables
  - Metadata/lookup tables


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Cache](https://github.com/donnemartin/system-design-primer#cache) <span id="content-10"></span>
- Image
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/Q6z24La.png)
- Caching improves page load times and can reduce the load on your servers and databases.
- In this model, the dispatcher will first lookup if the request has been made before and try to find the previous result to return, in order to save the actual execution.
- Popular items can skew the distribution, causing bottlenecks. Putting a cache in front of a database can help absorb uneven loads and spikes in traffic.
- Client caching: Caches can be located on the client side (OS or browser), server side, or in a distinct cache layer.
- CDN caching
- Web server caching
- Reverse proxies and caches such as Varnish can serve static and dynamic content directly. Web servers can also cache requests, returning responses without having to contact application servers.
- Database caching
- Your database usually includes some level of caching in a default configuration, optimized for a generic use case. Tweaking these settings for specific usage patterns can further boost performance.
- Application caching
  - In-memory caches such as Memcached and Redis are key-value stores between your application and your data storage.
  - Since the data is held in RAM, it is much faster than typical databases where data is stored on disk. RAM is more limited than disk, so cache invalidation algorithms such as least recently used (LRU) can help invalidate 'cold' entries and keep 'hot' data in RAM.
  - Redis has the following additional features:
    - Persistence option
    - Built-in data structures such as sorted sets and lists
  - There are multiple levels you can cache that fall into two general categories: database queries and objects:
    - Row level
    - Query-level
    - Fully-formed serializable objects
    - Fully-rendered HTML
  - Generally, you should try to avoid file-based caching, as it makes cloning and auto-scaling more difficult.
  - Caching at the database query level
    - Whenever you query the database, hash the query as a key and store the result to the cache. This approach suffers from expiration issues:
      - Hard to delete a cached result with complex queries
      - If one piece of data changes such as a table cell, you need to delete all cached queries that might include the changed cell
  - Caching at the object level
    - See your data as an object, similar to what you do with your application code. Have your application assemble the dataset from the database into a class instance or a data structure(s)
  - Suggestions of what to cache:
    - User sessions
    - Fully rendered web pages
    - Activity streams
    - User graph data

### When to update the cache
- Since you can only store a limited amount of data in cache, you'll need to determine which cache update strategy works best for your use case.
- Cache-aside
  - ![](https://github.com/donnemartin/system-design-primer/raw/master/images/ONjORqk.png)
  - The application is responsible for reading and writing from storage. The cache does not interact with storage directly. The application does the following:
    - Look for entry in cache, resulting in a cache miss
    - Load entry from the database
    - Add entry to cache
    - Return entry
  - Code:
    ```python
    def get_user(self, user_id):
        user = cache.get("user.{0}", user_id)
        if user is None:
            user = db.query("SELECT * FROM users WHERE user_id = {0}", user_id)
            if user is not None:
                key = "user.{0}".format(user_id)
                cache.set(key, json.dumps(user))
        return user
    ```
  - Memcached is generally used in this manner.
  - Subsequent reads of data added to cache are fast. Cache-aside is also referred to as lazy loading. Only requested data is cached, which avoids filling up the cache with data that isn't requested.
  - Disadvantage(s): cache-aside
    - Each cache miss results in three trips, which can cause a noticeable delay.
    - Data can become stale if it is updated in the database. This issue is mitigated by setting a time-to-live (TTL) which forces an update of the cache entry, or by using write-through.
    - When a node fails, it is replaced by a new, empty node, increasing latency.
- Write-through
  - Image:
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/0vBc0hN.png)
  - The application uses the cache as the main data store, reading and writing data to it, while the cache is responsible for reading and writing to the database:
    - Application adds/updates entry in cache
    - Cache synchronously writes entry to data store
    - Return
  - Write-through is a slow overall operation due to the write operation, but subsequent reads of just written data are fast
  - Users are generally more tolerant of latency when updating data than reading data. Data in the cache is not stale.
  - Disadvantage(s): write through
    - When a new node is created due to failure or scaling, the new node will not cache entries until the entry is updated in the database. - Cache-aside in conjunction with write through can mitigate this issue.
    - Most data written might never be read, which can be minimized with a TTL.
- Write-behind (write-back)
  - Image
    ![](https://github.com/donnemartin/system-design-primer/raw/master/images/rgSrvjG.png)
  - In write-behind, the application does the following:
    - Add/update entry in cache
    - Asynchronously write entry to the data store, improving write performance
  - Disadvantage(s): write-behind
    - There could be data loss if the cache goes down prior to its contents hitting the data store.
    - It is more complex to implement write-behind than it is to implement cache-aside or write-through.
- Refresh-ahead
  - ![](https://github.com/donnemartin/system-design-primer/raw/master/images/kxtjqgE.png)
  - You can configure the cache to automatically refresh any recently accessed cache entry prior to its expiration.
  - Refresh-ahead can result in reduced latency vs read-through if the cache can accurately predict which items are likely to be needed in the future.
  - Disadvantage(s): refresh-ahead
    - Not accurately predicting which items are likely to be needed in the future can result in reduced performance than without refresh-ahead.

### Disadvantage(s): cache
- Need to maintain consistency between caches and the source of truth such as the database through cache invalidation.
- Cache invalidation is a difficult problem, there is additional complexity associated with when to update the cache.
- Need to make application changes such as adding Redis or memcached.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Asynchronism](https://github.com/donnemartin/system-design-primer#asynchronism) <span id="content-11"></span>

### Intro
- Image
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/54GYsSx.png)
- Asynchronous workflows help reduce request times for expensive operations that would otherwise be performed in-line. They can also help by doing time-consuming work in advance, such as periodic aggregation of data.

### Message queues
- Message queues receive, hold, and deliver messages. If an operation is too slow to perform inline, you can use a message queue with the following workflow:
  - An application publishes a job to the queue, then notifies the user of job status
  - A worker picks up the job from the queue, processes it, then signals the job is complete
- The user is not blocked and the job is processed in the background. During this time, the client might optionally do a small amount of processing to make it seem like the task has completed.
- For example, if posting a tweet, the tweet could be instantly posted to your timeline, but it could take some time before your tweet is actually delivered to all of your followers.
- Redis is useful as a simple message broker but messages can be lost.
- RabbitMQ is popular but requires you to adapt to the 'AMQP' protocol and manage your own nodes.
- Amazon SQS is hosted but can have high latency and has the possibility of messages being delivered twice.

### Task queues
- Tasks queues receive tasks and their related data, runs them, then delivers their results. They can support scheduling and can be used to run computationally-intensive jobs in the background.
- Celery has support for scheduling and primarily has python support.

### Back pressure
- If queues start to grow significantly, the queue size can become larger than memory, resulting in cache misses, disk reads, and even slower performance.
- Back pressure can help by limiting the queue size, thereby maintaining a high throughput rate and good response times for jobs already in the queue. Once the queue fills up, clients get a server busy or HTTP 503 status code to try again later. Clients can retry the request at a later time, perhaps with exponential backoff.

### Disadvantage(s): asynchronism
- Use cases such as inexpensive calculations and realtime workflows might be better suited for synchronous operations, as introducing queues can add delays and complexity.
  

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Communication](https://github.com/donnemartin/system-design-primer#communication) <span id="content-12"></span>

### Introduction
- Image:
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/5KeocQs.jpg)

### Hypertext transfer protocol (HTTP)
- HTTP is a method for encoding and transporting data between a client and a server. It is a request/response protocol: clients issue requests and servers issue responses with relevant content and completion status info about the request.
- HTTP is self-contained, allowing requests and responses to flow through many intermediate routers and servers that perform load balancing, caching, encryption, and compression.
- A basic HTTP request consists of a verb (method) and a resource (endpoint)
- HTTP is an application layer protocol relying on lower-level protocols such as TCP and UDP.

### Transmission control protocol (TCP)
- Image:
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/JdAsdvG.jpg)
- TCP is a connection-oriented protocol over an IP network. Connection is established and terminated using a handshake. All packets sent are guaranteed to reach the destination in the original order and without corruption through:
  - Sequence numbers and checksum fields for each packet
  - Acknowledgement packets and automatic retransmission
- If the sender does not receive a correct response, it will resend the packets. If there are multiple timeouts, the connection is dropped. TCP also implements flow control and congestion control. These guarantees cause delays and generally result in less efficient transmission than UDP.
- TCP is useful for applications that require high reliability but are less time critical. Some examples include web servers, database info, SMTP, FTP, and SSH
- Use TCP over UDP when:
  - You need all of the data to arrive intact
  - You want to automatically make a best estimate use of the network throughput

### User datagram protocol (UDP)
- Image:
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/yzDrJtA.jpg)
- UDP is connectionless. Datagrams (analogous to packets) are guaranteed only at the datagram level. Datagrams might reach their destination out of order or not at all. UDP does not support congestion control. Without the guarantees that TCP support, UDP is generally more efficient.
- UDP is less reliable but works well in real time use cases such as VoIP, video chat, streaming, and realtime multiplayer games.
- Use UDP over TCP when:
  - You need the lowest latency
  - Late data is worse than loss of data
  - You want to implement your own error correction

### Remote procedure call (RPC)
- Image:
  ![](https://github.com/donnemartin/system-design-primer/raw/master/images/iF4Mkb5.png)
- In an RPC, a client causes a procedure to execute on a different address space, usually a remote server.
- The procedure is coded as if it were a local procedure call, abstracting away the details of how to communicate with the server from the client program. Remote calls are usually slower and less reliable than local calls so it is helpful to distinguish RPC calls from local calls.
- RPC is a request-response protocol:
  - Client program - Calls the client stub procedure. The parameters are pushed onto the stack like a local procedure call.
  - Client stub procedure - Marshals (packs) procedure id and arguments into a request message.
  - Client communication module - OS sends the message from the client to the server.
  - Server communication module - OS passes the incoming packets to the server stub procedure.
  - Server stub procedure - Unmarshalls the results, calls the server procedure matching the procedure id and passes the given arguments.
  - The server response repeats the steps above in reverse order.
- Disadvantage(s): RPC
  - RPC clients become tightly coupled to the service implementation.
  - A new API must be defined for every new operation or use case.
  - It can be difficult to debug RPC.
  - You might not be able to leverage existing technologies out of the box. For example, it might require additional effort to ensure RPC calls are properly cached on caching servers such as Squid.

### Representational state transfer (REST)
- REST is an architectural style enforcing a client/server model where the client acts on a set of resources managed by the server.
- REST is an architectural style enforcing a client/server model where the client acts on a set of resources managed by the server. The server provides a representation of resources and actions that can either manipulate or get a new representation of resources. All communication must be stateless and cacheable.
- There are four qualities of a RESTful interface:
  - Identify resources (URI in HTTP) - use the same URI regardless of any operation.
  - Change with representations (Verbs in HTTP) - use verbs, headers, and body.
  - Self-descriptive error message (status response in HTTP) - Use status codes, don't reinvent the wheel.
  - HATEOAS (HTML interface for HTTP) - your web service should be fully accessible in a browser.
- REST is focused on exposing data. It minimizes the coupling between client/server and is often used for public HTTP APIs. REST uses a more generic and uniform method of exposing resources through URIs, representation through headers, and actions through verbs such as GET, POST, PUT, DELETE, and PATCH. Being stateless, REST is great for horizontal scaling and partitioning.
- Disadvantage(s): REST
  - With REST being focused on exposing data, it might not be a good fit if resources are not naturally organized or accessed in a simple hierarchy. For example, returning all updated records from the past hour matching a particular set of events is not easily expressed as a path. With REST, it is likely to be implemented with a combination of URI path, query parameters, and possibly the request body.
  - REST typically relies on a few verbs (GET, POST, PUT, DELETE, and PATCH) which sometimes doesn't fit your use case. For example, moving expired documents to the archive folder might not cleanly fit within these verbs.
  - Fetching complicated resources with nested hierarchies requires multiple round trips between the client and server to render single views, e.g. fetching content of a blog entry and the comments on that entry. For mobile applications operating in variable network conditions, these multiple roundtrips are highly undesirable.
  - Over time, more fields might be added to an API response and older clients will receive all new data fields, even those that they do not need, as a result, it bloats the payload size and leads to larger latencies.


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://github.com/donnemartin/system-design-primer#next-steps