# iruka

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
