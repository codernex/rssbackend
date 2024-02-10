### RSS Backend

This project is created for fun and to practice Golang. It serves as an opportunity to delve into Golang development by
building a project from start to finish.

### Motivation

The motivation behind this project is to deepen my understanding of Golang. Inspired by a YouTube video on freecodecamp
by [Lane Wagner](https://github.com/wagslane), I decided to create a similar kind of project. Having learned Golang from
Boot.Dev, I aimed to implement additional functionality beyond the original project.

### Note

    This marks my inaugural project in Golang.
    Expect more Golang projects in the future!

### Endpoints:

#### /healthz

- **GET**
    - Returns:
        - 200 -- Server Running

#### /users

- **POST**
    - Creates New Users
    - Body:
      ```json
      {
          "name": "Your Name"
      }
      ```

- **GET**
    - Requires Authentication
    - Headers:
      ```json
      {
          "Authorization": "ApiKey **********"
      }
      ```

#### /feeds

- **POST**
    - Requires Authentication
    - Headers:
      ```json
      {
          "Authorization": "ApiKey **********"
      }
      ```
    - Body:
      ```json
      {
          "name": "Feed Name",
          "url": "RSS Feed URL"
      }
      ```

#### /feeds/posts

- **GET**
    - Requires Authentication
    - Headers:
      ```json
      {
          "Authorization": "ApiKey **********"
      }
      ```
    - Returns Followed Feed Posts

#### /feed_follow

- **POST**
    - Requires Authentication
    - Headers:
      ```json
      {
          "Authorization": "ApiKey **********"
      }
      ```
    - Body:
      ```json
      {
          "feed_id": "Feed ID Wants to Follow"
      }
      ```

- **GET**
    - Requires Authentication
    - Headers:
      ```json
      {
          "Authorization": "ApiKey **********"
      }
      ```
    - Returns Following Feed For The User

#### /feed_follow/{feedFollowId}

- **DELETE**
    - Requires Authentication
    - Remove Feeds From Following

