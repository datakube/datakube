[![Build Status](https://semaphoreci.com/api/v1/santode/datahamster/branches/master/badge.svg)](https://semaphoreci.com/santode/datahamster)
[![codecov](https://codecov.io/gh/SantoDE/datahamster/branch/master/graph/badge.svg)](https://codecov.io/gh/SantoDE/datahamster)

Datakube is a modern method to backup / save your application data which are based in the cloud easily. Currently it's
early in development and definitely **NOT** production ready.

##Overview
You might have already faced that situation: You happily deployed your set of services to a cloud of your choice 
(e.g. Amazon RDS or Google Cloud SQL) and your nightly snapshots are running just fine. However, you might also want to
have backups as simple SQL File Dumps to archive them somewhere else and have them accessible with your normal
ecosytem. **This is when Datakube comes into play**!

Datakube's agent creates dumbs of it's known targets (currently by [a .toml file](https://stackoverflow.com/questions/tagged/traefik))
and a given schedule. Once a particular dump is made, its transferred back to the server and stored safely there. From
there, you can grab this dumb with a simple HTTP call and reuse the file whereever you want to for whatver
you want to.

##Features
- Continuously updates to target configuration (No restarts!)
- Save files to a storage of your choise (currently only file storage is supported)
- Set different schedules for different targets
- Easy to configure
- Small footprint as its made with go in :heart:

##Roadmap
As mentioned, datakube is very early in development and therefore its missing **a lot** of functionality just
know. Features planned already:

- User Management
- Restrict Targets / Dumps per User
- Web UI
- More File Storages (S3. Cloud Storage...)
- More Target providers
- Metrics


##Quickstart
No version is released yet. Will ne updated soon!
##Examples
Please see the [examples](https://stackoverflow.com/questions/tagged/traefik) directory for configuration examples.