#!/bin/bash

curl 10.29.102.173:9200/metadata -XDELETE

curl 10.203.209.28:9200/metadata -XPUT -d'{"mappings":{"_doc":{"properties":{"name":{"type":"string","index":"not_analyzed"},"version":{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"string"}}}}}'
