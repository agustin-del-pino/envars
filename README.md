# envars

Utility for map OS env-vars into config json file.

## Usage

Use `-h` or `--help` for display the help of the command.

```shell
./envars -h
```

```
envars command allows to map env vars into a config.json file

Positional Arguments
jsonpath  the path of the config.json file to map the env vars

Flags
-h  --help  Show the command help
```

## How it works?

Using a json file as schema and default values of the env-vars,
it will cast them to the respective type. _All json types are supported_.
The env-var names **must be in UPPER CASE**.

Given this json:

```json lines
// config.json

{
  // Will cast the env-var to object
  "alias": {},
  
  // Will cast the env-var to number
  "port": 8080,

  // Will cast the env-var to array
  "tags": [],

  // Will cast the env-var to boolean
  "use_ssl": false,

  // Will cast the env-var to string
  "user_admin": "admin"
}
```

And given the next variables:

```shell
export TAGS='["dev", "test", "prod"]'
export PORT='7000'
```

When executes the command:

```shell
./envars ./config.json
```

Then the `config.json` will be overwritten with the new values.

```json lines
// config.json

{
  "alias": {},
  "port": 7000,
  "tags": [
    "dev",
    "test",
    "prod"
  ],
  "use_ssl": false,
  "user_admin": "admin"
}
```

## Supported boolean values
```yaml
truthy:
    - "true"
    - "yes" 
    - "1"   
    - "si"  
    - "hai" 
    - "ja"  
    - "da"  
    - "sim" 
    - "ok"
```

Any other not described value, will be interpreted as `false`.