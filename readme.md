## `perigee`

An HTTP server and CLI tool for performing Azimuth PKI operations for Urbit IDs

![image](https://github.com/user-attachments/assets/4c252a1c-72d5-497d-aca3-8dd3dd958a47)

This is a pure Go utility that you can run as an HTTP server or command line tool. It is a wrapper around the [L2 roller RPC](https://urbit.github.io/roller-rpc-client/) client spec and go bindings for the Azimuth/Ecliptic L1 contracts, but it also has macros and conveniences for deriving the data required for transactions. This allows you to e.g. breach a ship on L1 or L2 with a single command, knowing only the `@p` and the master ticket (or eth wallet private key). 

Additionally, it contains a library (`github.com/Native-Planet/perigee/libprg`) with a simple interface that can be imported by other projects, and a library (`github.com/Native-Planet/perigee/aura`) for casting to `@uw` in golang, which allows you to generate valid keyfiles to boot your ship -- this removes the dependency on [Bridge](https://bridge.urbit.org) and allows you to automate PKI updates. You can also generate a ship's `+code`.

Big thanks to [stephenlacy](https://github.com/stephenlacy/go-urbit), [nathanlever](https://github.com/nathanlever/keygen) and everyone who worked on [Bridge](https://github.com/urbit/bridge) and [urbit-key-generation](https://github.com/urbit/urbit-key-generation) for doing the hard parts.

~~For now this library can only perform L2 transactions.~~ *Now with L1!* ✅

Set the `ROLLER_URL` env var for custom roller. Set the `ADMIN_TOKEN` env var if you want authentication in server mode.

To run:
- download latest release from sidebar
- `chmod +x perigee-amd64 && mv perigee-amd64 perigee`
- `./perigee`
  
To verify binary provenance:
- download 
- Use [slsa3-verifier](https://github.com/slsa-framework/slsa-verifier): `/slsa-verifier-linux-amd64 verify-artifact perigee-amd64 --provenance-path perigee-amd64.intoto.jsonl --source-uri=git+https://github.com/Native-Planet/perigee`

To build: 
- [install go](https://go.dev/doc/install) >=1.23.2 
- `git clone https://github.com/Native-Planet/perigee && cd perigee`
- `go build -o perigee .`

To run docker container:

```bash
docker build -t perigee
docker run -v $(pwd)/out:/out -p 8080:8080 perigee
```

> Note that you can use the `privkey` url parameter or `--private-key` cli arg instead of a master ticket and provide an ethereum wallet private key for an ownership or management address


### `get-wallet`
- generate a json wallet with key information
##### server

`curl http://localhost:8080/v1/gen/wallet\?ship=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&life\=2`

##### cli
```bash
perigee generate-wallet --point=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```

(optional flags: `--life`, `--output-dir`; also writes to `./out/sampel-palnet-1-wallet.json` unless output path is overriden)

---


### `get-keyfile`
- generate a `@uw`-encoded keyfile to boot a ship
##### server

`curl http://localhost:8080/v1/gen/wallet\?ship=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&life\=2`

##### cli
```bash
perigee generate-wallet --point=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```

(optional flags: `--life`, `--output-dir`; also writes to `./out/sampel-palnet-1.key` unless output path is overriden)

---

### `get-point` 
- get the azimuth state of a point
##### server

`curl http://localhost:8080/v1/get/point\?point=\~satmun-wacnup`

##### cli
```bash
perigee get-point --point=sampel-palnet
```

---

### `get-code`
- get the `+code` of a point
##### server

`curl http://localhost:8080/v1/get/code\?point=sampel-palnet&ticket=sampel-palnet-sampel-palnet`

##### cli

```bash
perigee get-code --point=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```

(optional flags: `--life`, `--output-dir`, `step` (integer you can increment if the +code has been reset); also writes to `./out/sampel-palnet.code` unless output path is overriden)

### `get-pending`
- get all pending rollup txos
##### server

`curl http://localhost:8080/v1/get/pending`

##### cli
```bash
perigee get-pending
```

---



### `breach`
- continuity breach
##### server

`curl http://localhost:8080/v1/mod/breach?point=sampel-palnet\&ticket=~sampel-palnet-sampel-palnet`

##### cli
```bash
perigee breach --point=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```
note you can also use the `--wait` flag with a length of time (eg `60m`, `2h`) to watch the roller until it clears the queue

---



### `escape`
- escape to a new sponsor
##### server

`curl http://localhost:8080/v1/mod/escape?point=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&sponsor=sampel`

##### cli
```bash
perigee escape --point=sampel-palnet --sponsor=sampel --master-ticket=sampel-palnet-sampel-palnet
```

---



### `cancel-escape`
- cancel an escape request
##### server

`curl http://localhost:8080/v1/mod/cancel-escape?ship=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&sponsor=sampel`

##### cli
```bash
perigee cancel-escape --point=sampel-palnet adoptee=sampel --master-ticket=sampel-palnet-sampel-palnet
```

---



### `adopt`
- accept an escape request as a sponsor
##### server

`curl http://localhost:8080/v1/mod/escape?ship=\~satmun\&ticket=\~sampel-ticket-sampel-ticket\&adoptee=sampel-palnet`

##### cli
```bash
perigee adopt --point=sampel adoptee=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```


### Todo

- L1 breaches
