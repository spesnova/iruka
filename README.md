# iruka
A lightweight container platform.

<img src="iruka.png" height=250px>

## CLI

* [Managing Apps](docs/cli.md#managing-apps)
 * [Creating an App](docs/cli.md#creating-an-app)
 * [Listing Apps](docs/cli.md#listing-apps)
 * [Showing an App](docs/cli.md#showing-an-app-v030) (v0.3.0~)
 * [Destroying Apps](docs/cli.md#destroying-apps)
* [Deploying an App](docs/cli.md#deploying-an-app)
* [Managing Processes](docs/cli.md#managing-processes-containers)
 * [Listing Processes](docs/cli.md#listing-processes-of-an-app)
 * [Restarting Processes](docs/cli.md#restarting-processes-of-an-app-v030) (v0.3.0~)
 * [Scaling Processes of an App](docs/cli.md#scaling-processes-of-an-app-v030) (v0.3.0~)
* [Managing Config Vars](docs/cli.md#managing-config-vars)
 * [Setting Config Vars](docs/cli.md#setting-config-vars)
 * [Showing Config Vars](docs/cli.md#showing-config-vars)
 * [Unsetting Config Vars](docs/cli.md#unsetting-config-vars)
* [Running One-Off Process](docs/cli.md#running-one-off-process-v030) (v0.3.0~)

## HTTP API

* [Errors](docs/api-v1-alpha.md#errors)
* [App](docs/api-v1-alpha.md#app)
* [Container](docs/api-v1-alpha.md#container)
* [Config Vars](docs/api-v1-alpha.md#config-vars)

## Configuration
### Server

```bash
# development
$ export IRUKA_MACHINE=172.17.8.101
$ export IRUKA_ETCD_MACHINES=http://172.17.8.101:4001
$ export IRUKA_FLEET_API_URL=http://172.17.8.101:4002
$ export IRUKA_DOCKER_HOST=http://172.17.8.101:2375
$ gin --appPort 8080
```

```bash
# deploying as docker container
$ docker run \
    -d \
    --name irukad \
    -e IRUKA_MACHINE=172.17.8.101 \
    -e IRUKA_ETCD_MACHINES=http://172.17.8.101:4001 \
    -e IRUKA_FLEET_API_URL=http://172.17.8.101:4002 \
    -e IRUKA_DOCKER_HOST=http://172.17.8.101:2375 \
    -p 8080:8080 \
    quay.io/spesnova/iruka:latest
```

#### IRUKA_ETCD_MACHINES
Provide a custome set of etcd endpoints.

#### IRUKA_FLEET_API_URL
Provide a fleet HTTP API URL.

### CLI

```bash
$ export IRUKA_API_URL=http://172.17.8.101:8080
$ iruka apps
```

#### IRUKA_API_URL
Provide a iruka HTTP API URL.

## License
This project is releases under the [MIT license](http://opensource.org/licenses/MIT).
