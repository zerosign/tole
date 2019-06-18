from alpine:edge

env PATH=${PATH}:/opt/tole/bin

run mkdir -p /opt/tole/bin
copy tole /opt/tole/bin
