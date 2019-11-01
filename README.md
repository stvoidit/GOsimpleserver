# GOsimpleserver

config.ini exemple (for postgres database)

__config.ini__

    [server]
    host = localhost
    port = 11111
    user = USER
    password = PASSWORD
    dbname = DBNAME
    
    [application]
    secret = MySecretKey

# Start

##
make build-go

its create files:

1. app - web application (api router (JWT), privat router (Cookies))
2. monitor - for crontab application
3. /static - its frontend files

## Dependences
1. install node.js and install packages -> npm i --save in __/src/frontend__ folder
2. install golang

and end, you can run command __make go-build__. Its create __/build__ folder in your directory with compiled project