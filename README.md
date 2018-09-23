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

TODO: add usage example
TODO: add filebeat config example
