# System Design - Software Architecture

</br>

## List of Contents:

### 1. [Design Netflix -- Streaming Platform](#content-1)

</br>

---

## Contents:
## [Design Netflix -- Streaming Platform](https://www.youtube.com/watch?v=psQzyFfsUGU) <span id="content-1"></span>

- Global infrastructure
- OpenConnect is a CDN by netflix
- Netflix load balancer will balance traffic based on region then instances (two-tier)
- Using round robin to balance traffic
- Converting videos to different formats and resolutions (transcoding)
- Raw video would be 50 GB
- Netword speed will decide video resolution
- Netflix create multiple copies of videos for different resolutions (1200 copies for single movie)
- Source --> video chunks (|||||) --> processing --> merge upload --> Amazon S3
- Client will find the openconnect server close to the client
- ZULU:
  - Netty proxy -> inbound filter -> endpoint filter (static response) -> application server -> outbound filter -> netty proxy
- Hystrix:
  - Used for microservices architecture
  - Tree like structure
  - One failure will cascase to subsequent 
  - Limiting time for response from a service. E.g. if it's more that one second, the call will be cancelled
  - If threadpool is full, reject the call
  - Give fallback default response
- RPC for internal communication
- Critical endpoints will be highly available
- Server should be stateless
- EV based on memcache
- Cache using SSD to store data
- Netflix use MySQL and Cassandra
  - MYSQL users data
  - Cassandra for big data
- Using master-master pattern (with sync-replication)
- One master could have multiple read replicas
- Cassandra: user history. R/W ratio = 9/1
- Database in Cassandra:
  - Live viewing history -- recent
  - Compressed viewing history -> old
- Data
  - 500 Billion evernts
  - 13 PB everyday
  - Video activity
  - UI activity
  - Error logs
  - Performance events
  - Troubleshooting events
- Logging and analytics for Chukwa
- Using ElasticSearch
  - 150 clusters
  - 3500 instances
  - For customer support
- Spark
  - Recommendatation
  - Sorting and relevance ranking
- Traffic data: 10 TB / sec
- Caching content:
  - Historical viewing content -> netx videos
  - Popular content
- Caching strategy:
  - Using consistent hashing

**[⬆ back to top](#list-of-contents)**

</br>

---

## References:

- https://www.youtube.com/watch?v=psQzyFfsUGU