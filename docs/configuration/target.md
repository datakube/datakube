# Target Configuration

## Storage
```toml
# Target Array
[[target]]
    # Name of the Target - needs to be unique!
    name = "test_target"
    [target.db]
        ## Database Type - currently only mysql is supproted
        type = "mysql"
        host = "mariadb"
        name = "test"
        user = "user"
        password = "password"
        port = "3306"
    #schedule for the target. Currently supported:
    # monthly
    # weekly
    # daily
    # every_minute
    [target.schedule]
        interval = "every_minute"
```
