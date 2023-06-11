#!/bin/bash

echo -n "this object will be separate to 4+2 shards" | openssl dgst -sha256 -binary | base64

curl -v 127.0.0.1:12350/objects/test5 -XPUT -d "this object will be separate to 4+2 shards" -H "Digest: SHA-256=MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8="

ls -ltr /Users/cjm/tmp/?/objects
echo
curl 127.0.0.1:12350/objects/test5
echo
curl 127.0.0.1:12350/locate/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=

rm /Users/cjm/tmp/1/objects/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=.*
echo some_garbage > /tmp/2/objects/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=.*
ls -ltr /tmp/?/objects
echo
curl 10.29.2.1:12345/objects/test5
echo
ls -ltr /tmp/?/objects

echo -n "this object is test6" | openssl dgst -sha256 -binary | base64
S0fDdFXdWBW3wb9KEETRiDiwFJuCaYawULY/sBxuqF4=

curl -v 127.0.0.1:12350/objects/test6 -XPUT -d "this object is test6" -H "Digest: SHA-256=S0fDdFXdWBW3wb9KEETRiDiwFJuCaYawULY/sBxuqF4="

curl 127.0.0.1:12350/objects/test6