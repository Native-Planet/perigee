/-  http-utils
/+  logs, dbug, verb
|%
+  card  card:agent:gall
--

::  pontifex-proxy: stateless HTTP proxy for roller calls
::
::  Bind path: /~pontifex/proxy/[<@uw-encoded-url>]
::  Decodes the target URL from the path, forwards the request with a
::  Forwarded header, and returns the upstream response with permissive
::  CORS headers so the web UI can call it same-origin.
::
|_  =bowl:gall
+*  this  .
    l     log(our our.bowl url ~)

++  fetch
  |=  $:  [for=@ta secure=?]
          =request:http
      ==
  ^-  card
  =.  header-list.request
    =-  (set-header:http 'forwarded' - header-list.request)
    %+  rap  3
    :~  'for="pontifex-proxy";'
        'proto='  ?:(secure 'https' 'http')
    ==
  [%pass /fetch/[for]/(scot %t url.request) %arvo %i %request request *outbound-config:iris]

++  l
  |_  [our=@p url=(unit @t)]
  ++  fail
    |=  [desc=term trace=tang]
    %-  link
    (~(fail logs our /logs) desc trace deez)
  ++  tell
    |=  [=volume:logs =echo:logs]
    %-  link
    (~(tell logs our /logs) volume echo deez)
  ++  deez
    ^-  (list [@t json])
    :-  %flow^s+'http proxy'
    =;  l=(list (unit [@t json]))
      (murn l same)
    :~  ?~(url ~ `[%url s+u.url])
    ==
  ++  link
    |=  cad=card
    |*  [caz=(list card) etc=*]
    [[cad caz] etc]
  --

++  on-save  !>(~)

++  on-init
  ^-  (quip card _this)
  :_  this
  [%pass /eyre/bind %arvo %e %connect [~ /~pontifex/proxy] dap.bowl]~

++  on-load
  |=  ole=vase
  ^-  (quip card _this)
  [~ this]

++  on-poke
  |=  [=mark =vase]
  ^-  (quip card _this)
  ?>  =(our src):bowl
  ?+  mark  ~|([%strange-mark mark=mark] !!)
      %handle-http-request
    =+  !<(order:http vase)
    =+  (purse:http url.request)
    ?.  ?=([%~pontifex %proxy *] site)
      :_  this
      %^  spout:http  id
        [404 ~]
      `(as-octs:mimes:html (cat 3 'bad route into ' dap.bowl))
    =/  tail  t.t.site
    ?+  tail
      [(spout:http id [404 ~] `(as-octs:mimes:html 'bad path')) this]
        [@ ~]
      =|  msg=@t
      =*  bad-req
        %-  (tell:l %warn msg url.request ~)
        [(spout:http id [400 ~] `(as-octs:mimes:html msg)) this]
      ?~  target=(slaw %uw i.tail)
        =.  msg  'target not @uw'
        bad-req
      ?~  (de-purl:html u.target)
        =.  msg  'target not parseable'
        bad-req
      =.  url.request  u.target
      =.  header-list.request
        %+  skip  header-list.request
        |=([k=@t @t] =('cookie' k))
      [[(fetch [id secure] request)]~ this]
    ==
  ==

++  on-arvo
  |=  [=wire sign=sign-arvo]
  =-  -(log ^l)
  ^-  (quip card _this)
  ?+  wire  ~|(%strange-sign-arvo !!)
      [%eyre %bind ~]
    ?>  ?=(%bound +<.sign)
    ?:  accepted.sign  [~ this]
    %-  (tell:l %crit 'failed to eyre-bind' ~)
    %-  (slog dap.bowl 'failed to eyre-bind' ~)
    [~ this]
      [%fetch @ @ ~]
    =/  eid=@ta  i.t.wire
    =/  url=@t   (slav %t i.t.t.wire)
    =.  url.log  `url
    ?>  ?=([%iris %http-response *] sign)
    =*  res  client-response.sign
    %-  ?.  ?=(%progress -.res)  same
        (tell:l %warn 'strange iris %progress response' ~)
    =?  res  ?=(%progress -.res)
      ~&  [dap.bowl %strange-iris-progress-response]
      [%cancel ~]
    ?:  ?=(%cancel -.res)
      :_  this
      %+  spout:http  eid
      :-  [502 'x-pontifex-proxy'^'cancelled' ~]
      ~
    ?>  ?=(%finished -.res)
    =*  cod  status-code.response-header.res
    :_  this
    %+  spout:http  eid
    :-  =,  response-header.res
        :-  status-code
        (snoc headers 'x-pontifex-proxy'^'finished')
    ?~  full-file.res  ~
    `data.u.full-file.res
  ==

++  on-watch
  |=  =path
  ^-  (quip card _this)
  ?>  =(our src):bowl
  ?>  ?=([%http-response @ ~] path)
  [~ this]

++  on-leave
  |=  =path
  ^-  (quip card _this)
  [~ this]

++  on-agent
  |=  [=wire =sign:agent:gall]
  ?:  =(/logs wire)  [~ this]
  %-  (tell:l %crit 'unexpected on-agent' (spat wire) -.sign ~)
  ~&  [dap.bowl %unexpected-on-agent wire=wire]
  [~ this]

++  on-peek
  |=  =path
  ^-  (unit (unit cage))
  ~

++  on-fail
  |=  [=term =tang]
  ^-  (quip card _this)
  %-  (fail:l term tang)
  %-  (slog dap.bowl '+on-fail' term tang)
  [~ this]
--
