# ApiHelpers

Basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.
# Controllers

Contains handler functions for particular route to be called when an api is called.
#  Helpers

Contains helper functions used in all apis
# Middlewares

Middleware to be used for the project
# Models

Database tables to be used as models struct
# Resources

Resources contains all structures other than models which can be used as responses
# Routers

Resources define the routes for your project
# Seeder

It is optional, but if you want to insert lots of dummy records in your database, then you can use seeder.
# Services

All the core apis for your projects should be within services.
# Storage

It is generally for storage purpose.
# Templates

Contains the HTML templates used in your project
# .env

Contains environment variables.
Pre-requirements before starting your first go projects

. Install latest version of go i.e 1.12 (Released in Feb 2019) . Setup GOROOT and GOPATH
### RUN THE SERVER (Basic commands)

. For running the server you have to run following command go run main.go It will start your server at the port you have mentioned in .env file.

. If you want to run the server in port other than default, then run following command go run main.go <specific port>

. If you want to create a build for your project and upload in server, then you have to run following command go build
API with versioning
For using version 1 api

127.0.0.1:8099/api/v1/user-list
For using version 2 api

127.0.0.1:8099/api/v2/user-list