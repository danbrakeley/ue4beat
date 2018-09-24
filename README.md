# ue4beat

Translate Unreal Engine log lines into Filebeat-compatible json.

## Features

- Removes ANSI escape sequences
- `@timestamp` field with ISO8601 timestamp (assumes server is UTC)
- `frame` field from Unreal frame number (defaults to `-1`)
- `category` field from Unreal log category (ie `LogTemp`)
- `level` field with `info` (default), `warning`, or `error`, as appropriate
- `message` field with remaining text (whatever is left after parsing the above fields)

The `@timestamp` and `category` fields only appear in the output if they are not empty.

## Usage

```text
# pipe the output of your game into ue4beat, which will then print the json-ified log lines to stdout
/game/Binaries/Linux/gameServer /Game/Maps/Lobby | ue4beat

# same as above, but also ensures stderr is included (the "2>&1" bit), and outputs to a file instead of stdout
/game/Binaries/Linux/gameServer /Game/Maps/Lobby 2>&1 | ue4beat >> /logs/server1.json

# same as above, but also writes the original Unreal output to a second file
/game/Binaries/Linux/gameServer /Game/Maps/Lobby 2>&1 | tee >(/bin/ue4beat >> /logs/server1.json) >> /logs/server1.log
```

You can run multiple servers on the same machine by just changing the json/log file names (ie server2.json/log, etc).

Then you just need a Filebeat config that specifies json input. For example:

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
