<p align="center">
<img src="./job-funnel-logo.jpg" width="256"/>
</p>

# Job-Funnel
An aggregator used to collect job postings and funnel them into communication channels. Why? Because searching for jobs sucks and I also wanted to write this in GoLang.

## Requirements
- [GoLang 1.23.1 or higher](https://go.dev/doc/install)
- [MongoDB 7.0.14 or higher (optional)](https://www.mongodb.com/docs/manual/administration/install-community/)
- [Discord Bot Token (optional)](https://discord.com/developers/applications)
- Properly formatted .env file (see below)

## Installation
This will build the binary in the destination directory. I suggest you run this app in it's own directory as it will create a SQLite database file in the same directory. 

##### Note: This may change in the future to allow for a specified path for the SQLite database file.
```bash
git clone https://github.com/treestarsystems/job-funnel.git
cd ./job-funnel/
go build -o <destination dir>job-funnel
cd <destination dir>
```
To run this app it requires an .env file in the same directory or you must specify a path (Ex: -e=/path/to/.env) to a .env file. 

The .env file should be formatted as follows (template included repository):
```bash
# .env
PORT="8080"
GIN_MODE="release"
DB_NAME="jobFunnel"
DB_TABLE_NAME="jobPosts"
DB_SQLITE_ENABLE="true"
DB_MONGODB_ENABLE="false"
DB_SQLITE_FILENAME="job-funnel.sqlite.db"
DB_MONGODB_URI="mongodb://localhost:27017"
COMMUNICATION_DISCORD_ENABLE="true"
COMMUNICATION_DISCORD_BOT_TOKEN="<token>"
```

## Usage

Without -e flag, the program will use the .env file in the same directory
```bash
./<destination>/job-funnel
```

With -e flag, the program will use the .env file at the specified path
```bash
./<destination>/job-funnel -e=/path/to/.env
```

## Progress
#### *Green = Completed

<p align="center">
<img src="./job-search.drawio.svg" width="512"/>
</p>