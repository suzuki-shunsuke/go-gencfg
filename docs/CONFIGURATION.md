# Congiguration

## dest

generated code file path.

## formatters

`formatters` is the list of command to format generated code file.
For example,

```yaml
formatters:
- gofmt -l -s -w
```

the target file path is added to each command.

## global

`global` defines the global configuration.
If each param's configuration overwrites the global configuration.
The `global` configuration is optional.

* global.env.bind
* global.env.prefix
* global.flag.bind

## params[].name

`param.name` is the viper key.

## params[].type

`param.type` is the paramter's type.
If it is not set, the type is inferred using the default value.
For example, if `param.default` is `10`, `param.type` is inferred as `"int"`.

## params[].description

## params[].default

`params[].default` is the default value of the parameter.

```golang
viper.SetDefault(portKey, 3000)
```

```golang
pflag.IntP("port", "p", 3000, "app port number")
```

## params[].flag

## params[].env

## bind

`env.bind`

```golang
viper.BindEnv(portKey, "SAMPLE_PORT")
```

`flag.bind`

```golang
pflag.IntP("port", "p", 3000, "app port number")
viper.BindPFlag(portKey, pflag.Lookup("port"))
pflag.Parse()
```

Each param's `end.bind` and `flag.bind` is true if each param's other configuration such as `env.prefix` defines explicitly.

## flag.description

`flag.description` is the description of the flag.
If it is not set, `params[].description` is used.

```golang
pflag.IntP("port", "p", 3000, "app port number")
```

## flag.name

`flag.name` is the flag name.

```golang
pflag.IntP("port", "p", 3000, "app port number")
```

## flag.short

Optional.

If `flag.short` is set,

```golang
pflag.IntP("port", "p", 3000, "app port number")
```

otherwise

```golang
pflag.Int("port", 3000, "app port number")
```

## env.name

`env.name` is the binded environment variable name.
The default value is `{{env.prefix}}{{param.name}}` .

## env.prefix

`env.prefix` is the prefix of environment variable name.
If `env.name` is set, `env.prefix` is ignored.
