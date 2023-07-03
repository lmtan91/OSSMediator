# OSS Mediator Collector

The OSS Mediator Collector is a command line API client. It connects with NDAC APIGW and gets the FM and PM data using REST interface at regular intervals based on the collector configuration for the customer’s managed networks only.

### Prerequisites

MediatorCollector is compatible with only Unix/Linux system.

### Project Structure

    .  
    ├── resources               # Resource files  
        └── conf.json
        └── alarm_notifier.yaml  
    ├── src                     # Source files  
    ├── Makefile  
    ├── Dockerfile  
    └── README.md  

### Installation steps

OSSMediatorCollector's binary should be built by running `make all` command followed by `make build_package` command.  
It will create binary named as `collector` inside `bin` directory and package containing the binary and resource file, named as `OSSMediatorCollector-<VERSION>.zip` inside `package` directory.

Please follow below procedure to install OSSMediatorCollector-<VERSION>.zip in your home directory:

````
$ mkdir collector
$ cp OSSMediatorCollector-<VERSION>.zip collector/
$ cd collector/
$ unzip OSSMediatorCollector-<VERSION>.zip
````

MediatorCollector directory structure after installation will be as shown below:

````
    .
    ├── OSSMediatorCollector.zip
    ├── bin
        └── collector
    ├── log
        └── collector.log
    └── resources
        └── conf.json
````

## Usage
```
Usage: ./collector [options]
Options:
        -h, --help
                Output a usage message and exit.
        -conf_file string
                Config file path (default "../resources/conf.json")
        -cert_file string
                Certificate file path (if cert_file is not passed then it will establish TLS auth using root certificates.)
        -log_dir string
                Log Directory (default "../log"), logs will be stored in collector.log file.
        -log_level int
                Log Level (default 4). Values: 0 (PANIC), 1 (FATAl), 2 (ERROR), 3 (WARNING), 4 (INFO), 5 (DEBUG)
        -skip_tls
                Skip TLS Authentication
        -enable_console_log
                Enable console logging, if true logs won't be written to file
        -v
                Prints OSSMediator's version
```

## Configuration

PM / FM data collection by collector is performed using REST interface at regular intervals based on configuration.  

* To collect PM / FM data, it is required to modify `conf.json` configuration file in `resource` directory as shown in the example.

````json
{
  "base_url": "https://api.dac.nokia.com/api/ndac/v2",
  "users": [
    {
      "email_id": "<USER EMAIL>",
      "response_dest": "<DIRECTORY PATH>",
      "auth_type": "<ADTOKEN/PASSWORD>"
    },
    {
      "email_id": "<USER EMAIL>",
      "response_dest": "<DIRECTORY PATH>",
      "auth_type": "<ADTOKEN/PASSWORD>"
    }
  ],
  "azure_session_api": {
    "refresh": "/azure-refresh-sessions"
  },
  "um_api": {
    "login": "/login-session",
    "refresh": "/refresh-session",
    "logout": "/logout-session"
  },
  "list_nhg_api": {
    "api": "/network-hardware-groups",
    "interval": 60
  },
  "sim_apis": [
    {
      "api": "/network-hardware-groups/{nhg_id}/sims",
      "interval": 60
    },
    {
      "api": "/account-sims",
      "interval": 60
    },
    {
      "api": "/access-point-sims",
      "interval": 60
    }
  ],
  "metric_apis": [
    {
      "api": "/network-hardware-groups/{nhg_id}/pmdata",
      "metric_type": "RADIO",
      "interval": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/pmdata",
      "metric_type": "EDGE",
      "interval": 5
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/pmdata",
      "metric_type": "CORE",
      "interval": 5
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "ACTIVE",
      "metric_type": "RADIO",
      "interval": 1,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "HISTORY",
      "metric_type": "RADIO",
      "interval": 1,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "ACTIVE",
      "metric_type": "DAC",
      "interval": 1,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "HISTORY",
      "metric_type": "DAC",
      "interval": 1,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "ACTIVE",
      "metric_type": "CORE",
      "interval": 1,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "HISTORY",
      "metric_type": "CORE",
      "interval": 1,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "ACTIVE",
      "metric_type": "APPLICATION",
      "interval": 5,
      "sync_duration": 15
    },
    {
      "api": "/network-hardware-groups/{nhg_id}/fmdata",
      "type": "HISTORY",
      "metric_type": "APPLICATION",
      "interval": 5,
      "sync_duration": 15
    }
  ],
  "limit": 10000,
  "delay": 5
}
````

| Field                     | Type                | Description                                                                                                                                                                   |
|---------------------------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| base_url                  | string              | APIGW base URL.                                                                                                                                                               |
| users                     | [object]            | Users details.                                                                                                                                                                |
| users.email_id            | string              | User's email ID.                                                                                                                                                              |
| users.auth_type           | string              | User's authorization type - either ADTOKEN or PASSWORD.                                                                                                                       |
| users.response_dest       | string              | Base directory to store the response from the REST APIs. Subdirectories will be created inside the base directory for storing each APIs response in their respective location |
| um_api                    | object              | User management APIs.                                                                                                                                                         |
| um_api.login              | string              | Customer portal login API.                                                                                                                                                    |
| um_api.refresh            | string              | Customer portal refresh session API.                                                                                                                                          |
| um_api.logout             | string              | Customer portal logout API.                                                                                                                                                   |
| list_nhg_api.api          | string              | API URl for getting user's network details. Collector uses the list of NHGs for each FM/PM data collection.                                                                   |
| list_nhg_api.interval     | integer             | Interval at which list_nhg_api will be called..                                                                                                                               |
| sim_apis                  | [object] (Optional) | Get SIM APIs.                                                                                                                                                                 |
| sim_apis.api              | string              | API URL for fetching SIM data.                                                                                                                                                |
| sim_apis.interval         | integer             | Interval at which SIM API should be called to collect data.                                                                                                                   |
| metric_apis               | [object]            | Get PM/FM APIs.                                                                                                                                                               |
| metric_apis.api           | string              | API URL of get PM/FM data.                                                                                                                                                    |
| metric_apis.interval      | integer             | Interval at which API should be called to collect data.                                                                                                                       |
| metric_apis.type          | string              | Type of FM request ("ACTIVE" or "HISTORY").                                                                                                                                   |
| metric_apis.metric_type   | string              | Type of metric for PM("RADIO" or "CORE" or "EDGE") or FM("DAC" or "RADIO" or "CORE" or "APPLICATION").                                                                        |
| metric_apis.sync_duration | integer             | Time duration in minutes, for syncing FM for the given duration.                                                                                                              |
| limit                     | integer             | Number of records to be fetched from the API, should be within 1-10000.                                                                                                       |
| delay                     | integer             | Time duration in minutes, for adding delay in API calls.                                                                                                                      |

## Starting OSSMediatorCollector  

Configure the passwords for the users configured in `resources/conf.json` file for API access by executing `storesecret` inside `bin` directory.  

#### storesecret Usage

```
Usage: ./storesecret [options]
Options:
        -h, --help
                Output a usage message.
        -c string
                Config file path (default "../resources/conf.json")
```

Check if execute permissions/ are there for the `storesecret` binary, if not set it as `chmod 777 storesecret`, then execute `sudo ./storesecret` command to store the user passwords.
Enter the password for each customer having the right permission.  

NOTE:
* For login details (email ID and password) contact Nokia DAC support/operations team.
* In case the user’s password is updated, execute `sudo ./storesecret` and input the updated password, then restart the OSSMediatorCollector module.  

To start collector, go to the installed path of the collector bin directory and start by calling the following command:

````
./collector
````

Once the login is successful for all users, the collector will periodically start collecting the data by calling the configured APIs for the customer’s managed network.

Collector logs can be checked in $cd $collector_basepath/log/collector.log file.

### Alarm notification

User can enable alarm notification feature to receive details of specific alarm raised from the network.  
This feature is optional and disabled by default.
* To enable alarm notification, it is required to add `alarm_notifier.yaml` file in `resource` directory as shown in the example.  

```yaml
  webhook_url: <WEBHOOK URL>
  radio_alarm_filters:
    - specific_problem: <ALARM SPECIFIC PROBLEM>
    - specific_problem: <ALARM SPECIFIC PROBLEM>
      fault_ids:
        - <FAULT ID>
        - <FAULT ID>
    - specific_problem: <ALARM SPECIFIC PROBLEM>
      fault_ids:
        - <FAULT ID>
        - <FAULT ID>
  dac_alarm_filters:
    - alarm_id: <ALARM ID>
    - alarm_id: <ALARM ID>
  core_alarm_filters:
    - alarm_id: <ALARM ID>
    - alarm_id: <ALARM ID>
  alarm_sync_duration: 60
  group_events: <true/false>
  notify_clear_event: <true/false>
  message_format: <ms_teams/json>
```

| Field                                 | Type        | Description                                                                                            |
|---------------------------------------|-------------|--------------------------------------------------------------------------------------------------------|
| webhook_url                           | string      | Webhook url.                                                                                           |
| radio_alarm_filters.specific_problem  | integer     | Specific problem of the alarm of radio module for which notification should be sent.                   |
| radio_alarm_filters.fault_ids         | integer     | Fault id of the radio alarm (can be found in Alarm text' second part).                                 |
| dac_alarm_filters.alarm_id            | integer     | Alarm ID of the DAC alarm for which notification should be sent.                                       |
| core_alarm_filters.alarm_id           | integer     | Alarm ID of the CORE alarm for which notification should be sent.                                      |
| alarm_sync_duration                   | integer     | Duration in minutes after which notification for the already notified active alarms wil be sent again. |
| group_events                          | boolean     | To group notification events based on Network Hardware level. Default: False                           |
| notify_clear_event                    | boolean     | To enable clear alarm notifications. Default: False                                                    |
| message_format                        | string      | Message format (ms_teams or json)                                                                      |
