[![Build Status](https://semaphoreci.com/api/v1/datakube/datakube/branches/master/badge.svg)](https://semaphoreci.com/datakube/datakube)
[![codecov](https://codecov.io/gh/datakube/datakube/branch/master/graph/badge.svg)](https://codecov.io/gh/datakube/datakube)

Datakube is a modern method to backup / save your application data which are based in the cloud easily.

## Overview
You might have already faced that situation: You happily deployed your set of services to a cloud of your choice 
(e.g. Amazon RDS or Google Cloud SQL) and your nightly snapshots are running just fine. However, you might also want to
have backups as simple SQL File Dumps to archive them somewhere else and have them accessible with your normal
ecosytem. **This is when Datakube comes into play**!

Datakube's agent creates dumbs of it's known targets (currently by [a .toml file](./examples/targets.toml))
and a given schedule. Once a particular dump is made, its transferred back to the server and stored safely there. From
there, you can grab this dumb with a simple HTTP call and reuse the file whereever you want to for whatver
you want to.

## Disclaimer
Currently it's very early in development and definitely **NOT** production ready. **USE AT YOUR OWN RISK**

## Features
- Continuously updates to target configuration (No restarts!)
- Save files to a storage of your choise (currently only file storage is supported)
- Set different schedules for different targets
- Easy to configure
- Small footprint as its made with go in :heart:

## Roadmap
As mentioned, datakube is very early in development and therefore its missing **a lot** of functionality just
know. Features planned already:

- User Management
- Restrict Targets / Dumps per User
- Web UI
- More File Storages (S3. Cloud Storage...)
- More Target providers
- Metrics

## Quickstart
Have a look at the [quickstart docs](./docs/quickstart.md)

## Configuration
Best place is, to have a look at the [configuration docs](./docs/configuration)

## Examples
Please  see the [examples](./examples) directory for configuration examples and a docker-compose setup.