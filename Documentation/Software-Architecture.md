# Planned software architecture

## Runtime Diagram

```mermaid
erDiagram
    Backend ||..o{ Frontend : serves
    Backend }o--|| Database : "reads from"
    Proxy }o--|| Database : "writes to"
    Proxy ||--|| "K8s Api Server" : "receives updates from"
```

## Code Components

### Proxy

The Proxy is responsible for reflecting the state of the K8s cluster into the relational DB.
It uses the [K8s client libraries](https://github.com/kubernetes/client-go/) to interact with the API server
and [the Bun ORM framework](https://bun.uptrace.dev/) in connection with the [Golang Postgres Driver](https://github.com/lib/pq) to insert data into the DB.

### Web service/Explorer

The user facing web service consists of a [next.js](https://github.com/vercel/next.js) frontend application and a next.js router,
which is responsible for loading the current state from the DB and sending updates to the existing clients.

The frontend is then responsible for displaying the data in a variety of views, allowing the user to explore the current state of a cluster.

## Tech Stack summary

To increase reproducibility, we have commited to using Docker and Docker-Compose to set up our software. By building in the container, developers are not
required to install any of the development tools (with the exception of a K8s API provider) locally.

Additionally, we are very mindful of polling vs event driven updates and are evaluating each part of our stack by its ability to receive, transform and emits
events efficiently.

## Explanation

The use of a relational DB as the only interface between the proxy and the web service was client requirement, as was the choice of Golang for the backend.

For the frontend we were looking for a modern framework that wasn't burden with too much complexity so we could focus
on the data modeling and the best way to present such complex information.
