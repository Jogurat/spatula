# Golang web scraper

## Usage & installation

### Installation

⚠ You must have the Docker CLI (and/or docker-compose) tool installed, and have Docker running!

First off, get the `docker-compose.yml` file from this repo. Either by doing a `git clone` or just downloading it.

Now, when in the path where the file is located, run:

`POSTGRES_PASSWORD=example NODE_URL=http://node:4000/api DATABASE_URL=postgres://postgres:${POSTGRES_PASSWORD}@db/postgres docker compose up`

Only thing you can/should change in the command is the POSTGRES_PASSWORD env var.

⚠ Note: The env var POSTGRES_PASSWORD that you use while first booting up the Postgres container will stay as the main user password for the entire duration that the container uses the same VOLUME, meaning, if you want to change the main user password for Postgres, you must first remove the docker volume attached to the container, by using `docker volume rm <name-of-volume>`

⚠ Note: If you have an older version of the Docker CLI, you will have to use `docker-compose up` instead of `docker compose up`, notice the "-"!

### Usage

### Frontend

After all of the containers have booted up & are running, you can visit [localhost:8080](http://localhost:8080) for the frontend of the scraper. Here, you can find the simple UI to navigate around scraping user info.

You can find the frontend Github repository [HERE](https://github.com/Jogurat/spatula-vue).

![Spatula UI](https://i.imgur.com/teQAVZX.png)

Start by entering the username of the user you want to check the stats of, and press the button corresponding to the social network you want to search. A nice card will appear! 😊 The avatars are randomly generated, based on the username (same username will always have the same avatar).

![Spatula Search](https://i.imgur.com/1LMT10U.png)

### Swagger

On [localhost:1414](http://localhost:1414) you can find the Swagger API explorer

Here, if you wish to explore your locally hosted version of the app, please select the appropriate server in the "Servers" dropdown, after clicking the "Try it out" button (located where the "Cancel" button is on the image):

![Swagger Server tab 1](https://i.imgur.com/XCb1xDl.png)

### Adminer

On [localhost:1111](http://localhost:1111) you can find the Adminer DB tool. This is a simple, phpmyadmin-like tool for exploring your DB. If you used the Swagger API explorer for your locally hosted app, here, you will be able to see your data stored in the DB.

![Adminer Login](https://i.imgur.com/WUVZy33.png)

In the "Password" field, fill in the POSTGRES_PASSWORD you used when starting the docker-compose command.

![Adminer Dashboard](https://i.imgur.com/jebD4fW.png)

By clicking the "select" next to the table name, you will see all of the data stored in that table. Most noteable columns are: username, followers_count, posts_count & updated_at (all except created at 😛, also please ignore the "completed_at" field in the image 🤪)

The "updated_at" field is checked when we already have a row in the DB, to see if the cache needs to be refreshed - the time frame is 10 minutes, meaning, if the row was updated longer than 10 minutes ago, we will scrape the data again (this value is lower for purpose of presentation, it would probably be much higher in a prod environment).

## Brief overview of the architecture

Whenever you search for a user, it is first looked for in the Postgres DB. If no results are returned, the user will be searched using the appropriate scraper: for Twitter, it's on the Golang server, and for TikTok it's on the Node server. Once the user is returned from the scraper, it is returned to the user, as well as being inserted into the DB.

If, however, a user with the given username already exists in the DB, its' "updated_at" field is looked at, and being checked if it's older than 10 minutes. If not, that user's info is returned from the DB without the need to be scraped again. However, if the cache is "too stale", the flow of getting the data from the scrapers is started. Once the scraping is done, the user gets updated in the DB, and, with it, of course, its' "updated_at" field, which is being done with a DB trigger.

I tried creating a swimchart diagram for this, but the tool used was being very hard on me 😠
![Diagram attempt](https://i.imgur.com/9bzfaL5.png)

## All repositories tied to this project:

- [Go Server - you are here](https://github.com/Jogurat/spatula)
- [Node Server - TikTok microservice](https://github.com/Jogurat/spatula-tiktok)
- [Vue Frontend](https://github.com/Jogurat/spatula-vue)

## Noteable resources:

- [Timestamps in Postgres](https://x-team.com/blog/automatic-timestamps-with-postgresql/)
- [TikTok Scraper](https://github.com/drawrowfly/tiktok-scraper)
- [Tabler Icons](https://tablericons.com/)
- [Official Docker Image of Postgres](https://hub.docker.com/_/postgres)
- [Random Avatars](https://avatars.dicebear.com/)
- [pgx API Overview at pkg.go.dev](https://pkg.go.dev/github.com/jackc/pgx)
