# makedo

## Motivation

This was inspired by the blog post [Afraid of Makefiles? Don't be!](https://matthias-endler.de/2017/makefiles/)

Personally, I'm still scared of makefiles! Just kidding, they're fine. Anyway, if you want to explore this interactively, I suggest converting the rest of the example from that blog to a makedo.yaml file. It's really quite simple. :ghost:

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
