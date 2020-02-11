# challengesite

To get rolling on a dev environment:

Make sure you have docker-compose installed: [Installation Link](https://docs.docker.com/compose/install/)
1.  Build the packages with docker-compose : `docker-compose build`
2.  Start the service: `docker-compose up`
3.  Make a request to the service:

`curl http://HOSTNAME/` : root Handler

`curl http://HOSTNAME/redisPing` : ping the redis Docker container

`curl http://HOSTNAME/postgresPing` : ping the postgres Docker container
