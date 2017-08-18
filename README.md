# makedo

## Installation

```
go get github.com/r0fls/makedo
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
