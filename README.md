# dun

![](https://github.com/xchacha20-poly1305/dun/actions/workflows/lint.yml/badge.svg)

sing-box and some additional functionality or API.

Name from: [赵盾](https://zh.wikipedia.org/wiki/赵盾)

## Version name

`{sing-box-version}-dun-{dun-version}`

## Components

see distro/all

### dunapi

- V2Ray service API (An implementation that replaces the sing-box V2Ray option, the caller can 
  set it after obtaining the router).
- Golang http.Client API.

### dunbox

Custom Box.

Use this instead of `github.com/sagernet/sing-box`.

`dunbox.New(Options,PlatformOptions)` is used in `libgojni.so`.

### dunmain

Custom CLI tools.

`dunmain.Create(jsonContent)` creates `dunbox.Box` instead of `box.Box` (used in `nekobox_core`).

# Credits

- [SagerNet/sing-box](https://github.com/SagerNet/sing-box)
- [MatsuriDayo/sing-box-extra](https://github.com/MatsuriDayo/sing-box-extra)