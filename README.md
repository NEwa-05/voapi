# Voice Over API :D

Will be based on Alexa Skills and swapi.dev (Star Wars API)

## what we need

Get request from Alexa concerning a character from Star Wars and reply to Alexa with the information found in swapi.dev

Message from and to Alexa will be JSON.

swapi.dev parsing will be done via get request

## tools to test api

- ngrok (https://ngrok.com/) to create a redirect from an external https url to our local service

usage example:
ngrok http 8080

- fresh (https://github.com/gravityblast/fresh) to build go code at each changes

usage example:
fresh

- postman (https://www.postman.com/) to interact with api easily

usage example: ./Postman
