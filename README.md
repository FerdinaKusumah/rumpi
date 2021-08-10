# rumpi
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub license](https://img.shields.io/github/license/narutoxxx/rumpi.svg)](https://github.com/FerdinaKusumah/rumpi/blob/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/narutoxxx/rumpi.svg)](https://GitHub.com/narutoxxx/rumpi/issues/)
[![GitHub issues-closed](https://img.shields.io/github/issues-closed/narutoxxx/rumpi.svg)](https://GitHub.com/narutoxxx/rumpi/issues?q=is%3Aissue+is%3Aclosed)
[![GitHub pull-requests](https://img.shields.io/github/issues-pr/Naereen/StrapDown.js.svg)](https://GitHub.com/Naereen/StrapDown.js/pull/)

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Resource</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- GETTING STARTED -->
## Getting Started
`rümpï` is engine to watch changes or health api based on scheduler.
This will periodic checking api at given specified range interval in second.
You can watch multiple api changes and notify to another api when there's changes in data.

### Watch multiple api
![Watch Api](/example/watch-api.png)

### Notify other api when data source is changes
![Notify Api](/example/notify.png)

`rümpï` will post message to notify url with payload 
```json
{
    "date_time": "2021-01-02 08:46:24", # this will be UTC datetime
    "message": "OK", #  http status response text when success fetch watch api 
    "status": {
        "Data": "[1]", # what is changes from api
        "Error": null, # if any error when fetch api
        "StatusCode": 200 #  http status response when success fetch watch api
    },
    "url": "http://localhost:8090/watch1" # information url watch url
}
```  

### Prerequisites

Need go `1.12+` or later.

## Installation

### from GitHub and make executable

```bash
▶ git clone https://github.com/narutoxxx/rumpi
▶ cd rumpi
▶ go build .
▶ (sudo) mv rumpi /usr/local/bin
```

<!-- USAGE EXAMPLES -->
## Usage

| **Flag**          	| **Description**                                                 	                                |
|-------------------	|-----------------------------------------------------------------------------------------------	|
| -s, --source         	| source JSON file                 	                                                                |
| -h, --help        	| Display its helps                                               	                                |


### Examples

#### Example Json Config
```json
[
    {
        "watch_api": "http://localhost:8090/watch1",
        "notify_api": "http://localhost:8090/report1",
        "interval": 1,
        "verbose": true
    },
    {
        "watch_api": "http://localhost:8090/watch2",
        "notify_api": "http://localhost:8090/report2",
        "interval": 2,
        "verbose": true
    },
    {
        "watch_api": "http://localhost:8090/watch3",
        "notify_api": "http://localhost:8090/report3",
        "interval": 3,
        "verbose": true
    }
]
```

#### Watch From File Config Path

```bash
▶ rumpi --source "./config.json"
```

#### Watch From File Config Url

```bash
▶ rumpi --source "https://raw.githubusercontent.com/narutoxxx/rumpi/master/example/config.json"
```

## TODOs

- [ ] Supporting header or token
- [ ] Notify to many source like, email, telegram, etc ...

## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/narutoxxx/rumpi/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/narutoxxx/rumpi/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.


## Pronunciation
`id_ID` • **/rümpï/** — suka liat emak - emak belanja sayur `ngegibah` kan ??, nah dari situ awal mulanya!

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**rumpi** released under MIT. See `LICENSE` for more details.
