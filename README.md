[![wercker status](https://app.wercker.com/status/59d57ea64a657bfeae067b54dad17823/s/master "wercker status")](https://app.wercker.com/project/byKey/59d57ea64a657bfeae067b54dad17823)
# Giffy
This repository cotains the documentation and source for a coding assesment.   

## Overview
Giffy is a GIF explorer, which allows your to curate through GIFs from Giphy's API and rank your favourite GIFs.

## Objective
Create a web app that allows you to search with a list of tag words to display a list of GIFs. Clicking one of the GIFs in the list should bring that item into focus, providing it's details. There also should be way to rank/rate items and sort them by their ranks.

## Requirements
- [x] Use Giphy’s API to search for tags and display a list of GIFs. The API is paged, so I would anticipate some way of handling multi-paged items.
- [ ] Clicking on a GIF should bring that item into focus and display more information
- [x] Provide the ability to rank / rate the GIFs and sort by highest rank.
- [x] The specific UI is up to you. Bonus points for awesome stuff you bake into the app.
- [ ] Provide specific instructions for installing. For example, if you are using a specific
library, please call that out, so we know to install them.
- [ ] Testing is something SoapBox takes seriously. You should too.

## Getting Started
To run the application you need to run both the Giffy service and UI.   

Run with Docker:   
1. `docker pull danewilson/giffy`
2. `docker run -p 3000:3000 danewilson/giffy`
3. Visit [localhost:3000](http://localhost:3000)

If you have go and node installed you can follow these instructions.

1. `cd assets`
2. `yarn`
3. `yarn build`
4. `cd ..`
5. `go build`
6. `./giffy`
7. Visit [localhost:3000](http://localhost:3000)