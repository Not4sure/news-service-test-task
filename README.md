# News Service
News service written in Golang with htmx ui and mongodb database.

## Startup instructions
You need `docker` and `docker-compose` installed
Just run
```
make start
```
or 
```
docker-compose up -d 
```

To access application, open 
`http://localhost:8080`

To stop application, run 
```
make stop
```

## Architecture and general structure
This application's is heavily inspired by Domain Driven Design principles 
described in book by Eric Evans, and also excellent book by Miłosz Smółka and Robert Laszczak named "Go with domain".

Application is divided into couple layers, namely: infrastructure (ports and adapters), application, and domain logic.

The only port for application (for now) is htmx view, which provides access to all application main functions true 
htmx web application.

Mongodb repository implemetation is used for storage of Articles, which are main(and only) Entities in application.
There's also a in-memory Article repository implemetation which is used primarily for unit test in domain layer.

Although domain layer in this application is quite thin and, as some may say, stupid, it's separation from application
and infrastructure logic provides ability to expand it, and still contain critical business logic in one place.

I must confess that this project was my first experience with htmx and i probably done something terribly wrong.

## What should (and must) be done 
FE lacks validation, as it happens only on BE in domain layer. Obviously some some FE validation should be added together 
with user-friendly validation errors.

Validation logic is hard-coded into domain layer, and probably needs be configurable by application config. 
This calls for implemetation of configurable ArticleFactory which accepts validation parameters.

Some form of pagination should be added to List() Articles.

Db integrations tests are triggered with unit tests. This is not a problem for this project, but as integration 
tests got bigger and more costly to run, they should be separated from unit tests.

## What I'm proud of in this application
Firstly, general project structure.

Secondly, with my minimal experience with frontend and no prior experience with htmx, quite not bad frontend ui and ux,
especially with article edit form which appears seamlessly after "Edit" button is pushed.
