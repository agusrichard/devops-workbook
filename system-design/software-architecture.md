# System Design - Software Architecture

</br>

## List of Contents:

### 1. [Design Netflix -- Streaming Platform](#content-1)
### 2. [Design Twitter](#content-2)

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
  - One failure will cascade to subsequent 
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
  - 500 Billion events
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

## [Design Twitter](https://www.youtube.com/watch?v=wYk0xPP_P_8&list=PLkQkbY7JNJuC99VDJcpQdww-4aT3QhdJv&index=19) <span id="content-2"></span>
- Requirements:
  - User can tweet -- eventual consistency
  - Timelime
    - Home
    - User
    - Search
  - Trends
- Before architecting we need to make assumptions
- Calcs:
  - 300M+ users
  - 600 tweets / sec -- write
  - 600,000 tweets / sec -- reads
- Desc:
  - Read heavy
  - Eventual consistency
  - Don't need to worry too much for storage
- Combining Redis and DB
- Tables: User, Tweet, and Followers
- DB:
  - user
    - tweet
    - follower
- Redis
  - user_id -> tweets
  - user_id -> followers
- User timeline
  - DB --  Shard by time
    - Tweets
    - Followers
  - Redis:
    - Get tweets by user (tweet id)
    - Real tweet data
- Home timeline
  - Get followers
  - Get latest tweets
  - Merge and display
- Efficient way to load home timeline
  - Using fanout
  - General solution: If a user tweets something, then that tweet will be stored inside the database and that user own timeline. But other than that, this event will also distributed to his followers, so his followers will have the updated home timelime. All of this is happened in Redis cache.
- Trends
  - 1000 tweets in 5 min
  - 10,000 tweets in 1 month
  - Tweet
    - filter (removing violation check like adult content)
    - Parse (is there any hastag in here or removing stop words)
      - Count hashtag
        - Rank
      - Geolocation
        - Count location
    - All goes to Redis and internal API
- Search:
  - Inverted full text index:
  - Tweet is broken down into words
  - Look up to the table including that word
  - Scatter and gather:
    - Query -> nodes / data centers
- Twitter system architecture:
  - High level design:
    ![](./images/Screenshot%202022-04-21%20231947.png)
  - Http push websocket for streaming data
  - Thousands nodes of Redis -- cordinated using zookeeper



**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://www.youtube.com/watch?v=psQzyFfsUGU
- https://www.youtube.com/watch?v=wYk0xPP_P_8&list=PLkQkbY7JNJuC99VDJcpQdww-4aT3QhdJv&index=19