# Page-Hit-Tracker

A simple static web page hit tracker using **Golang** and **SQLite**, Can tracking different subpages of a static web page.

### Features

- [x] Build a simple backend server yourself for tracking page hits
- [x] Use a lightweight database like SQLite
- [x] Fast and easy to use
- [x] Track static web subpages

### Requirements

- [x] dokcer with docker-compose

### How to use

1. Install docker and docker-compose
2. Clone this repository
3. Setting you publishing url in config.json
    - "trackerURL": "PUBLISHING_URL"
4. Run `bash RUN.sh` to start the server
    -   The server will run on port 8080

Use the following code to track page hits, and replace the PUBLISHING_URL with your publishing url.
```
<script async src="${ PUBLISHING_URL }/tracker.js"></script>
<span id="page-hit-tracker"></span>
```