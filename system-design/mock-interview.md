# System Design Mock Interview

</br>

## List of Contents:

### 1. [Design Netflix](#content-1)

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

## References:

- https://www.youtube.com/watch?v=VvZf7lISfgs
