<h1><img width="125" src="assets/logo.png"></h1>

Flare is a simple and lightweight migration tool for Go projects. You can manage databases and migrations while keeping it up to date with Go structs.

It also helps you to keep you schemas in sync across multiple developers and servers.

#### Works on
macOS, Linux and Docker

#### Supported databases
PostgreSQL and MySQL

## Quick usage
You can start using Flare CLI by creating a **flare.toml** file in the root path of your project with the database connection info

```toml
Driver="postgres"
User="flare"
Name="flare"
Host="localhost"
Port="5432"
Password="12345678"
```

### Database
The `database` or shortcut `db` command allows you to manage your database directly. With it you can:

```bash
// creates the specified database
flare database create

// drops the existing database
flare database drop
```

## Documentation
You can find our getting started guides and general CLI usage here.

## Contributing
Want to help develop Flare CLI? Check out our contributing documentation.

If you find an issue, please report it on the issue tracker.

## License
This project is under [Apache License 2.0](https://github.com/ecorreiax/flare/blob/main/LICENSE) license.