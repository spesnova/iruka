# iruka CLI (prototype)

```bash
# Apps
$ iruka create hello
$ iruka apps
$ iruka info hello
$ iruka destroy hello

# Deploy
$ iruka deploy
$ iruka deploy --app <APP>

# Containers
$ iruka ps
$ iruka scale web=3:2X
$ iruka restart example.web.1

# Config vars
$ iruka config
$ iruka config --app <APP>
$ iruka set HELLO=world
$ iruka unset HELLO

# Releases
$ iruka releases

# Admin task
$ iruka run bash

# Machines
$ iruka machines
```

