# Reddit to VK Auto Poster

## APIs
* Reddit API documentation: https://www.reddit.com/dev/api/
* VK API documentation: https://vk.com/dev/methods

## Getting tokens and other stuff
Make your new app here https://www.reddit.com/prefs/apps and save its secret key and user agent (app's title)

Generate VK ADMIN token from here: https://vkhost.github.io/

## Environment setup
Make .env file according to the example
```
CLIENTSECRET=<reddit secret key>
PASSWORD=<reddit account password>
USERAGENT=<title of your app>
USERNAME=<your reddit username>
SUBREDDIT=<subreddit title>
FREQUENCY=<posts/day, should be less than LIMIT>
VK_TOKEN=<your vk tokem>
OWNER_ID=<your vk group id with "-" before it>
SOURCE=<source of your pictures>
LIMIT=<amount of pictures to download from subreddit>
```

## Building project
1. Clone this repository
2. `$ cd go-reddit-to-vk-auto-poster`
3. `$ touch .env` then open it, for example, with nano `$ nano .env` and fill it with data according to template above
4. Install golang (read https://golang.org/dl)
5. `$ go build src/main.go -o app`
6. `$ ./app`

That's it!

## Run in docker
1. Clone this repository
2. `$ cd go-reddit-to-vk-auto-poster`
3. `$ touch .env` then open it, for example, with nano `$ nano .env` and fill it with data according to template above
4. `$ docker build --tag reddit-vk-image .`
5. `$ docker run --name reddit-vk reddit-vk-image`