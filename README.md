# dbm
> 💼 manage your database login easily

`dbm` is a command-line tool that allows you
to easily store and manage your database credentials
in one secure location. It provides a secure ways to connect
to your databases without typing credentials.

## 📦 Installation

`dbm` is free and open source. You can install it by following one of given ways.

### Linux

```bash
curl -sL https://github.com/ermos/dbm/releases/latest/download/dbm_Linux_$(uname -m).tar.gz | tar -xvz --wildcards 'dbm' \
&& sudo mv dbm /usr/local/bin/
```

### Mac

```bash
curl -sL https://github.com/ermos/dbm/releases/latest/download/dbm_macOS_all.tar.gz | tar -xvz --wildcards 'dbm' \
&& sudo mv dbm /usr/local/bin/
```

[//]: # (### Windows)

[//]: # ()
[//]: # (Download the right archive from [the latest release page]&#40;https://github.com/ermos/dbm/releases/latest&#41;.)

### Alternative

From `go install` :
```bash
go install github.com/ermos/dbm@latest
```

From `source` :
```bash
git clone git@github.com:ermos/dbm.git
make build/bin
```

## 📚 Usage

To get started with `dbm`, you will first need to create a master password.
This password will be used to encrypt and decrypt your sensitive information,
so be sure to choose a strong and secure password. Above all, don't forget it!

Run any command, the master password set-up will show up :
```bash
> dbm add

Welcome to dbm!
Please define your master password before continuing.
Use a strong and unique master password,
it will be asking before each command.

Your master password :

```

When it's done, you can add your first database :
```bash
> dbm add

[dbm] master password: 

Adding a new database to dbm
This command is designed to guide you through
the process of adding a new database to dbm.

? What is the alias name ? mydb
? What is the protocol ? mysql
? What is the hostname ? 127.0.0.1
? What is the port ? 3306
? What is your database username ? root
? What is your database password ? 
? What is the default database ? mydb
```

And for finish, let's connect to our database :
```bash
> dbm open mydb

[dbm] master password: 

Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MySQL connection id is 123456789
Server version: 0.0.00 MySQL Community Server (GPL)

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MySQL [mydb]> 
```

And voila ! 🎉

## 🔒 Security

Currently, `dbm` encrypt only passwords.
`dbm` use a secure ways used by the major part of password manager tools.
`dbm` encrypts passwords using the Advanced Encryption Standard (AES-256)
in Cipher Block Chaining (CBC) mode, along with a salt
and a password-derived key (PBKDF2).

According to [OWASP recommendations](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html),
the minimal required iteration for PBKDF2 is `600.000` but we have increased it to `1.200.000`.
If you want change this value, you can download source code and update it in `/internal/pkg/goliath/config.go`.

Since `dbm` is open-source, you can see the implementation here: `/internal/pkg/goliath`.

## 🤝 Contributing

Contributions to `dbm` are always welcome!
If you find a bug or have a feature request, please open an issue on GitHub.
If you want to contribute code, please fork the repository and