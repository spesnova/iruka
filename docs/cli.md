# iruka CLI (prototype)
## Managing Apps
### Creating an App

```bash
$ iruka create <APP_NAME>

# Example
$ iruka create hello
Creating hello... done
```

### Listing Apps

```bash
$ iruka apps

# Example
$ iruka apps
foo
bar
baz
```

### Showing an App (v0.3.0~)

```bash
$ iruka info <APP_NAME>

# Example
$ iruka info hello
Name: hello
Maintenance Mode: false
```

### Destroying an App

```bash
$ iruka destroy <APP_NAME>

# Example
$ iruka deploy hello
Destroying hello... done
```


## Deploying an App

```bash
$ iruka deploy <APP_NAME>

# Example
$ iruka deploy hello
Deploying hello... done
```


## Managing Processes (Containers)
### Listing processes of an App

```
$ iruka ps

# Example
$ iruka ps
NAME                        IMAGE           SIZE    STATUS  CREATED     COMMAND
hello.web.bb1b59fd-14ac-11e5-8bd7-5cf93896cc38  coreos/example:1.0.0    2X  Up 50 seconds  1 minutes   /bin/sh -c nginx
```

### Restarting processes of an App (v0.3.0~)

```bash
$ iruka restart <PROCESS or TYPE>

# Example: restart web type processes
$ iruka restart web
Restarting web processes... done

# Example: restart hello.web.bb1b59fd-14ac-11e5-8bd7-5cf93896cc38
$ iruka restart hello.web.bb1b59fd-14ac-11e5-8bd7-5cf93896cc38
Restarting hello.web.bb1b59fd-14ac-11e5-8bd7-5cf93896cc38... done
```

### Scaling processes of an App (v0.3.0~)

```bash
$ iruka scale <TYPE>=<QUANTITY>:<SIZE>

# Example: scale web process's quantity to 2
$ iruka scale web=2
Scaling processes... done

# Example: scale web processes's size to 3X
$ iruka scale web=:3X
Scaling processes... done

# Example: scale web processes's quantity to 3 and size to 2X
$ iruka scale web=3:2X
Scaling processes... done

# Example: scale web processes's quantity to 3 and size to 2X
#          and scale worker processes's quantity to 10
$ iruka scale web=3:2X worker=10
Scaling processes... done
```


## Managing Config Vars
### Setting Config Vars

```bash
$ iruka set <KEY>=<VALUE>...

# Example: set a key-value
$ iruka set FOO=bar
FOO:   bar

# Example: set multiple key-value
$ iruka set FOO=bar BAZ=qux
FOO:   bar
BAZ:   qux
```

### Showing Config Vars

```bash
$ iruka config

# Example
$ iruka config
FOO:   bar
BAR:   baz
BAZ:   qux
```

### Unsetting Config Vars

```bash
$ iruka unset <KEY>...

# Example: unset a key
$ iruka unsert FOO
BAR:   baz
BAZ:   qux

# Example: unset multiple key
$ iruka unsert FOO,BAR
BAZ:   qux
```


## Running One-Off Process (v0.3.0~)

```bash
$ iruka run <COMMAND>

# Example: launch rails console
$ iruka run rails console

# Example: login to a container
$ iruka run bash

# Example: kick batch job
$ iruka run rake some:task
```


---

## CLI Ideas

```
# Managing Releases
$ iruka releases
$ iruka rollback

# Managing Machines
$ iruka machines
```

