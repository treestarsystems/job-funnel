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
```bash
git clone https://github.com/treestarsystems/job-funnel.git
cd ./job-funnel
go build -o <destination dire>job-funnel
```

This will build the binary in the destination directory.
It requires a .env file in the same directory your are executing the compiled binary with the following variables:
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
```bash
./<destination>/job-funnel 
#
```

<p align="center">
<img src="./job-search.drawio.svg" width="512"/>
</p>