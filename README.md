# ue4beat

Translate Unreal Engine log lines into Filebeat-compatible json.

---
---
---

## DEPRECATED

I wrote a ue4 filebeat module, and am in the process of switching to that instead of running this. It is available to all here: [proletariatgames/filebeat-ue4-module](https://github.com/proletariatgames/filebeat-ue4-module).

---
---
---

## Features

- Removes ANSI escape sequences
- `@timestamp` field with ISO8601 timestamp (assumes server is UTC)
- `frame` field from Unreal frame number (defaults to `-1`)
- `category` field from Unreal log category (ie `LogTemp`)
- `level` field with `info` (default), `warning`, or `error`, as appropriate
- `message` field with remaining text (whatever is left after parsing the above fields)
- additional static fields can be passed on the command line, and will be added to every output line

The `@timestamp` and `category` fields only appear in the output if they are not empty.

## Command Line Parameters

```text
-f, --field <name> <value>   Add name:value field to each output line
-h, --help                   Print this information
-v, --version                Print just the app name and version
```

### Usage Examples

```text
$ ./ue4beat -v
ue4beat v0.0.3

$ echo "LogWhat: This is a thing" | ./ue4beat
{"fields.category":"LogWhat","fields.level":"info","message":"This is a thing"}

$ echo "LogTemp: Warning: This is important" | ./ue4beat --field server_num 1
{"fields.category":"LogTemp","fields.level":"warning","fields.server_num":1,"message":"This is important"}

$ echo "[2018.01.01-11.11.00:993][  3]LogFoo: Error: Bar" | ./ue4beat -f ip 192.168.0.0 -f port 7000 --field name "server 1"
{"@timestamp":"2018-01-01T11:11:00.993Z","fields.category":"LogFoo","fields.frame":3,"fields.ip":"192.168.0.0","fields.level":"error","fields.name":"server 1","fields.port":7000,"message":"Bar"}
```

## Server Setup

### Piping the output of Unreal Server

```text
# pipe the output of a game server into ue4beat, which will then print the json-ified log lines to stdout
/game/Binaries/Linux/gameServer /Game/Maps/Lobby | ue4beat

# same as above, but also ensures stderr is included (the "2>&1" bit), and outputs to a file instead of stdout
/game/Binaries/Linux/gameServer /Game/Maps/Lobby 2>&1 | ue4beat >> /logs/server1.json

# same as above, but also writes the original Unreal output to a second file
/game/Binaries/Linux/gameServer /Game/Maps/Lobby 2>&1 | tee >(/bin/ue4beat >> /logs/server1.json) >> /logs/server1.log
```

Multiple servers can be run on the same machine by just changing the json/log file names (ie server2.json/log, etc).

### Filebeat Config Example

Filebeat can be configured to monitor for one or more files with wildcards, for example:

```yaml
filebeat.inputs:
- type: log
  enabled: true
  paths:
  - /logs/server*.json
  json:
    keys_under_root: true
    overwrite_keys: true
    add_error_key: true
    message_key: message
...
```
