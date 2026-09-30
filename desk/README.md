# Pontifex desk (proto)

This desk packages the static Pontifex frontend bundle and a Gall agent that can proxy to the roller with permissive CORS headers so the browser UI can call it from the desk origin.

## Contents
- `web/` — static assets (index/app/style/bundle binary/wasm_exec/planet.gif)
- `app/pontifex-proxy.hoon` — Gall agent exposing `/proxy/*` as a CORS-allowed proxy to the configured roller URL
- `desk.docket-0`, `desk.bill`, `sys.kelvin` — desk metadata

## Building the bundle

From the repo root:
```bash
GOOS=js GOARCH=wasm go build -o web/perigee.wasm ./wasm
cp web/* desk/web/
```

## Installing on your ship

```
|mount %pontifex
# copy desk/ contents into the mounted %pontifex directory
|commit %pontifex
|install our %pontifex
```

Visit `http://<your-ship>/~pontifex/web/index.html` and set the Roller URL in the UI (or via the agent scry below).

## Configuring the proxy (Gall agent)

The agent defaults to `https://roller.urbit.org`. You can change it at runtime:
```
:pontifex-proxy &roller-url 'https://your.roller.example'
```

It exposes a same-origin path: `http://<ship>/~pontifex/proxy/<rest>` which forwards to `<roller-url>/<rest>` and adds `Access-Control-Allow-Origin: *` (and related headers). Point the UI’s Roller URL field to `/~pontifex/proxy/` to avoid CORS issues.
