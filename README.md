# Taplist

This a simple self hosted talist. It is intended to run on a Raspberry, Onion Omega etc.
and serves a page which shows which homebrewed beers you have on tap.

## Building

### Prerequisites

You need to have yarn and vue cli installed to build the frontend and go to build the backend.
Additionally you need to have [packr2](https://github.com/gobuffalo/packr/tree/master/v2) since
the frontend is bundled directly with the binary

## Building

A simple `make build` will build everything and create a binary called `taplist` in the root directory.

# Running

Taplist is configured via env variables. All variables need to be prefixed with `TAPLIST_`.

| Variable | Default | Required | Description |
|----------|---------|----------|-------------|
| KEGERATOR_NAME | My Smart Kegerator | no | The name of your kegerator/brewery to be displayed on the taplist |
| NUM_TAPS | 2 | no | The number of Taps your kegerator has |
| HTTP_ADDR | :8080 | no |The address the server listens on |
| LOG_LEVEL | info | no | How detailed you want the log to be |
| NO_AUTH | false | no | Disable authentication for mutations and queries which might be sensitive |
| HTTP_USER_HEADER | X-Auth-User | no | In which header to look for the admin user name |
| ADMIN_USER | | no | Name of the admin user |
| DATA_PATH | ./data/data.json | no | Where to persist the data about the taplist |
| BREWFATHER_USERID | | no | If you want to add batches from brewfather to taps, specify your brewfather user id |
| BREWFATHER_SECRET | | no | Brewfather API secret |
| FRONTEND_PATH | | no | Use your own frontend instead of the embedded one |
| BREWFATHER_BREWERY| | no | Brewery name to use if tap data comes from brewfather | 