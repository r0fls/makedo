# makedo

## Installation

```
go install github.com/r0fls/makedo
```

## Example


**makedo.yaml**
```yaml
shoes:
  depends:
    - socks
  commands:
    - echo putting on shoes

socks:
  commands:
    - echo putting on socks
```

Then run the `shoes` step with the command `makedo shoes`:

```shell
 $ makedo shoes
putting on socks
putting on shoes
```

## Status

This is new and should by no means replace `make` in your workflow. Unless you *really* disklike makefiles and like to live on the bleeding edge. If you find bugs or have feature requests, please create an issue.
