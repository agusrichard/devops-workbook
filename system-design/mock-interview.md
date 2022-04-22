# System Design Mock Interview

</br>

## List of Contents:

### 1. [Design Netflix](#content-1)
### 2. [Design Youtube](#content-2)
### 3. [Design Youtube](#content-3)
### 4. [Design Pastebin.com](#content-4)

</br>

---

## Contents:

## [Design Netflix](https://www.youtube.com/watch?v=VvZf7lISfgs) <span id="content-1"></span>

- Always start by asking or clarifying questions. Like focusing on what features important features.
- Need to focus on user activity data. Aggregate and process.
- Make sure that we're not diverging to another topic like recommendation engine. Because it's not important right now.
- Add non functional requirements: low latency and global
- Recommendation engine becomes background async process
- Talk about number of users will use the application. Like 200 M users
- Types of data:
  - Video content
  - Video static content: title, description, thumbnail, list of casts
  - User meta data -> is that video watched yet or last watched timestamp, likes
  - activity logs -> clicks, impressions, scrolls, fine tune -> for the purpose of recommendation engine
- After talk about requirements, assumptions, and data. Move to do some math and make estimations.
- Let's say that netflix has 10K videos
- 1 Hr long of videos
- 30 GB/hr
- 10K videos * 1 hr avg * 30 GB/hr = 300 TB
- The amount storage that is used only 300 TB
- We can use blog storage like S3.
- Move to high level design.
- Admin API is connected to blog store if the admin wants to add new video (movie or new series)
- Video static content:
  - Titles, description, cast lists
  - Where to store this data. Is that titles, description, and cast lists
  - Do we need cache for this? Cache top movies
  - Static content <== Service ==> Cache
- User metadata:
  - 200 M * 1 K * 100 Bytes = 20 TB
  - User metadata is stored in Postgres
  - The interviewee considers sharding
- Load balancer from the external users
- Add CDN and CDN populator. CDN populator will add new data to CDN for data to be cached.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Design Youtube](https://www.youtube.com/watch?v=1xV5WI0OFkg) <span id="content-2"></span>

- Functional requirements:
  - Upload videos from multiple devices
  - Viewing experience should be agnostic
  - Ask about what to focus
- Non-functional requiements:
  - High availability
  - Has low latency
- More people reading/streaming compared to write
- Client --> video upload service. WebRTC to be used for streaming.
- Video upload service (Video content) --> Transcoding queue --> blog store
- Video upload service (video's metadata) -> DB
- CDN <-- User --> Video streaming service


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Design Instagram](https://www.youtube.com/watch?v=VJpfO6KdyWE) <span id="content-3"></span>

- Start with requirements
- Requirements:
  - Upload images from a mobile client
  - Users follow other users
  - Generate a feed of images
  - Scale: 10 million users
- Assumptions
  - 10 mil users on monthly basis
  - 2 photos per month
  - 5 MB per photo
  - 10^7 * 2 * 5 MB = 1.2 PB per month
- Make tables of users, photos, followers
- User table:
  - id: primary key, serial
  - name: string
  - email: string
  - location: string
- Photo table:
  - id: primary key, serial
  - user_id: foreign key
  - caption: string
  - location: string
  - path: string
- Follower table (self-reference table):
  - from: foreign key referencing user_id
  - to: foreign key referencing user_id
- Make metadata DB
- Object storage like S3
- App server (read)
- App server (write)
- App server (read) --> cache <-- database
- App server (write) 
  - --> database
  - --> Object storage (upload)
- Load balancer is sitting in front of our server
- High level design:
  ![](./images/Screenshot%202022-04-22%20015754.png)


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Design Pastebin.com](https://github.com/donnemartin/system-design-primer/blob/master/solutions/system_design/pastebin/README.md) <span id="content-4"></span>

### Step 1: Outline use cases and constraints
#### Use cases
- User enters a block of text and gets a randomly generated link
  - Expiration
    - Default setting does not expire
    - Can optionally set a timed expiration
  - User enters a paste's url and views the contents
  - User is anonymous
  - Service tracks analytics of pages
  - Monthly visit stats
    - Service deletes expired pastes
  - Service has high availability

#### Constraints and assumptions
- State assumptions
  - Traffic is not evenly distributed
  -  Following a short link should be fast
  -  Pastes are text only
  -  Page view analytics do not need to be realtime
  -  10 million users
  -  10 million paste writes per month
  -  100 million paste reads per month
  -  10:1 read to write ratio
- Calculate usage
  - Size per paste
    - 1 KB content per paste
    - shortlink - 7 bytes
    - expiration_length_in_minutes - 4 bytes
    - created_at - 5 bytes
    - paste_path - 255 bytes
    - total = ~1.27 KB
  - 12.7 GB of new paste content per month
    - 1.27 KB per paste * 10 million pastes per month
    - ~450 GB of new paste content in 3 years
    - 360 million shortlinks in 3 years
    - Assume most are new pastes instead of updates to existing ones
  - 4 paste writes per second on average
  - 40 read requests per second on average

### Step 2: Create a high level design
- Image:
  ![](https://camo.githubusercontent.com/3cb97853adab74ef0abfe733685aab2a224702191cb696560be2b48a0d6c1d08/687474703a2f2f692e696d6775722e636f6d2f424b73426e6d472e706e67)

### Step 3: Design core components

#### Use case: User enters a block of text and gets a randomly generated link
- We could use a relational database as a large hash table, mapping the generated url to a file server and path containing the paste file.
- Instead of managing a file server, we could use a managed Object Store such as Amazon S3 or a NoSQL document store.
- Paths:
  - The Client sends a create paste request to the Web Server, running as a reverse proxy
  - The Web Server forwards the request to the Write API server
  - The Write API server does the following:
    - Generates a unique url
      - Checks if the url is unique by looking at the SQL Database for a duplicate
      - If the url is not unique, it generates another url
      - If we supported a custom url, we could use the user-supplied (also check for a duplicate)
    - Saves to the SQL Database pastes table
    - Saves the paste data to the Object Store
    - Returns the url
- Clarify with your interviewer how much code you are expected to write.
- The pastes table could have the following structure:
  ```text
  shortlink char(7) NOT NULL
  expiration_length_in_minutes int NOT NULL
  created_at datetime NOT NULL
  paste_path varchar(255) NOT NULL
  PRIMARY KEY(shortlink)
  ```
- Setting the primary key to be based on the shortlink column creates an index that the database uses to enforce uniqueness.
- We'll create an additional index on created_at to speed up lookups (log-time instead of scanning the entire table) and to keep the data in memory.
- Reading 1 MB sequentially from memory takes about 250 microseconds, while reading from SSD takes 4x and from disk takes 80x longer.
- To generate the unique url, we could:
  - Take the MD5 hash of the user's ip_address + timestamp
    - MD5 is a widely used hashing function that produces a 128-bit hash value
    - MD5 is uniformly distributed
    - Alternatively, we could also take the MD5 hash of randomly-generated data
  - Base 62 encode the MD5 hash
    - Base 62 encodes to [a-zA-Z0-9] which works well for urls, eliminating the need for escaping special characters
    - There is only one hash result for the original input and Base 62 is deterministic (no randomness involved)
    - Base 64 is another popular encoding but provides issues for urls because of the additional + and / characters
    - The following Base 62 pseudocode runs in O(k) time where k is the number of digits = 7:
  - Snippet:
    ```python
    def base_encode(num, base=62):
        digits = []
        while num > 0
            remainder = modulo(num, base)
            digits.push(remainder)
            num = divide(num, base)
        digits = digits.reverse
    ```
  - Take the first 7 characters of the output, which results in 62^7 possible values and should be sufficient to handle our constraint of 360 million shortlinks in 3 years:
  - We'll use a public REST API:
    ```text
    $ curl -X POST --data '{ "expiration_length_in_minutes": "60", \
        "paste_contents": "Hello World!" }' https://pastebin.com/api/v1/paste
    ```
  - Response:
    ```json
    {
        "shortlink": "foobar"
    }
    ```
  - For internal communications, we could use Remote Procedure Calls.

#### Use case: User enters a paste's url and views the contents
- The Client sends a get paste request to the Web Server
- The Web Server forwards the request to the Read API server
- The Read API server does the following:
  - Checks the SQL Database for the generated url
  - If the url is in the SQL Database, fetch the paste contents from the - Object Store
  - Else, return an error message for the user
- REST API:
  ```text
  $ curl https://pastebin.com/api/v1/paste?shortlink=foobar
  ```
- Response:
  ```json
  {
      "paste_contents": "Hello World",
      "created_at": "YYYY-MM-DD HH:MM:SS",
      "expiration_length_in_minutes": "60"
  }
  ```

#### Use case: Service tracks analytics of pages
- Since realtime analytics are not a requirement, we could simply MapReduce the Web Server logs to generate hit counts.
- Code:
  ```python
  class HitCounts(MRJob):
      def extract_url(self, line):
          """Extract the generated url from the log line."""
          ...

      def extract_year_month(self, line):
          """Return the year and month portions of the timestamp."""
          ...

      def mapper(self, _, line):
          """Parse each log line, extract and transform relevant lines.

          Emit key value pairs of the form:

          (2016-01, url0), 1
          (2016-01, url0), 1
          (2016-01, url1), 1
          """
          url = self.extract_url(line)
          period = self.extract_year_month(line)
          yield (period, url), 1

      def reducer(self, key, values):
          """Sum values for each key.

          (2016-01, url0), 2
          (2016-01, url1), 1
          """
          yield key, sum(values)
  ```

#### Use case: Service deletes expired pastes
- To delete expired pastes, we could just scan the SQL Database for all entries whose expiration timestamp are older than the current timestamp. All expired entries would then be deleted (or marked as expired) from the table.

### Step 4: Scale the design
- Image:
  ![](https://camo.githubusercontent.com/018f5e8780f65abf18eee58d0dfe8a0501213be42829a3da0d4d94118a5979c9/687474703a2f2f692e696d6775722e636f6d2f346564584730542e706e67)
- State you would do this iteratively: 1) Benchmark/Load Test, 2) Profile for bottlenecks 3) address bottlenecks while evaluating alternatives and trade-offs, and 4) repeat.
- We'll introduce some components to complete the design and to address scalability issues. Internal load balancers are not shown to reduce clutter.
- The Analytics Database could use a data warehousing solution such as Amazon Redshift or Google BigQuery.
- An Object Store such as Amazon S3 can comfortably handle the constraint of 12.7 GB of new content per month.
- To address the 40 average read requests per second (higher at peak), traffic for popular content should be handled by the Memory Cache instead of the database. The Memory Cache is also useful for handling the unevenly distributed traffic and traffic spikes.
- 4 average paste writes per second (with higher at peak) should be do-able for a single SQL Write Master-Slave.


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:

- https://www.youtube.com/watch?v=VvZf7lISfgs
- https://www.youtube.com/watch?v=1xV5WI0OFkg
- https://github.com/donnemartin/system-design-primer/blob/master/solutions/system_design/pastebin/README.md
- https://www.youtube.com/watch?v=VJpfO6KdyWE
