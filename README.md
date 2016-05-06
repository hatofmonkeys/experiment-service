# A/B Experimentation Route Service

## About
A simple system for distributing traffic between two Cloud Foundry applications in a set ratio.

If you would like more information about Route Services in Cloud Foundry, please refer to [CF dev docs](http://docs.cloudfoundry.org/services/index.html#route-services).

## Prerequisites
- A Cloud Foundry and Diego deployment
- CF CLI v6.16+
- Two versions of an application deployed for experimentation:
- One version called myapp-a
- Another version called myapp-b
```
$ git clone https://github.com/hatofmonkeys/experiment-service.git
```

## Install

## Set up Redis

Create a Redis service to store A/B ratios and configure the connection in main.go

```

Addr:     "xxxxxxxxxxxxx",
Password: "xxxxxxxxxxxxx", // no password set
DB:       0,               // use default DB

```

Set a value in Redis with the key 'ratio' between 0 and 100 to dictate the percentage of traffic to be directed to app A. All other traffic will go to app B.

```
myredis.com:17966> set ratio 100
OK
myredis.com:17966> get ratio
"100"
```

### Deploy Experiment App
```
$ cd experiment-service
$ cf push experiment-service
```

### Create Route Service
The following will create a route service instance using a user-provided service and specifies the route service url (see step above).

```
$ cf create-user-provided-service experiment-service -r https://experiment-service.cfapps.io
Creating user provided service experiment-service in org my-org / space as admin...
OK
```

### Map both versions to true hostname

```
$ cf map-route myapp-a cfapps.io --hostname myapp
$ cf map-route myapp-b cfapps.io --hostname myapp
```

### Bind Route to Service Instance
The following will create bind the application's route to the route service instance.

```
$ cf bind-route-service cfapps.io experiment-service --hostname myapp
Binding route myapp.cfapps.io to service instance experiment-service in org my-org / my-space as admin...
OK
```


## Try it out

Adjust the ratio in real time to direct varying levels of traffic to versions A and B.
