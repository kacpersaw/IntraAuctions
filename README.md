# intra-auctions
Intranet auctions for organization. 

The purpose of this application is to be able to buy items that are for sale within an organization through auctions.

The application has an auction page where every user can bid on it.
On the auction page you will find things like:
* Title
* Short description
* Description
* Actual price
* Pictures
* Input for you offer and big "Bid" button

There is also a simple administration panel where you can add new auction, remove it, set to active or see details with bid history.

Things to do you can find in issues.

## Tech stack
* LDAP for auth
* MySQL as DB for auctions, bid history and images linked to an auction
* Server-sent-events (SSE) for real time price
* Vue.js + Vuetify as frontend
* Go + gorilla/mux + gorm as REST API
* Docker-compose as dev env
* i18n made with vue-i18n


## Screenshots

![Auction page](/screenshots/auction.png?raw=true "Auction page")

![Auction list](/screenshots/auctionList.png?raw=true "Auction list")

![Auction details](/screenshots/auctionDetails.png?raw=true "Auction details")

