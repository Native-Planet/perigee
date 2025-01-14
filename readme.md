## `perigee`

An HTTP endpoint turns master tickets into breaches (now with CLI mode!)


You can override the roller URL by setting a `ROLLER_URL` env var

### `get-wallet`
- generate a json wallet with key information
##### server

`curl http://localhost:8080/v1/gen/wallet\?ship=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&life\=2`

##### cli
```bash
perigee generate-wallet --point=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```

(providing the `life` is optional)
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

`curl -H 'Content-Type: application/json' \
    -d '{"point":"~dabdet-linnel","ticket":"~sampel-ticket-sampel-ticket","revision":2}' \
    http://localhost:8080/v1/mod/breach`

##### cli
```bash
perigee breach --point=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```
note you can also use the `--wait` flag with a length of time (eg `60m`, `2h`) to watch the roller until it clears the queue
---

### `escape`
- escape to a new sponsor
##### server

`curl -H 'Content-Type: application/json' \
    http://localhost:8080/v1/mod/escape?ship=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&sponsor=sampel`

##### cli
```bash
perigee escape --point=sampel-palnet --sponsor=sampel --master-ticket=sampel-palnet-sampel-palnet
```
---

### `cancel-escape`
- cancel an escape request
##### server

`curl -H 'Content-Type: application/json' \
    http://localhost:8080/v1/mod/cancel-escape?ship=\~satmun-wacnup\&ticket=\~sampel-ticket-sampel-ticket\&sponsor=sampel`

##### cli
```bash
perigee cancel-escape --point=sampel-palnet adoptee=sampel --master-ticket=sampel-palnet-sampel-palnet
```
---

### `adopt`
- accept an escape request as a sponsor
##### server

`curl -H 'Content-Type: application/json' \
    http://localhost:8080/v1/mod/escape?ship=\~satmun\&ticket=\~sampel-ticket-sampel-ticket\&adoptee=sampel-palnet`

##### cli
```bash
perigee adopt --point=sampel adoptee=sampel-palnet --master-ticket=sampel-palnet-sampel-palnet
```

> Note that you can use the `privkey` url parameter or `--private-key` cli arg instead of a master ticket and provide an ethereum wallet private key for an ownership or management address

Set the `ROLLER_URL` env var for custom rollers and `API_URL` with the desired solaris endpoint for automatic master ticket retrieval.

### Todo

- L1 breaches
