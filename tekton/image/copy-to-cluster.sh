#!/bin/sh
set +x -e
oc apply -f pvc-image.yaml
oc apply -f dummy.yaml
oc wait --for=condition=ready pod/dummy
oc rsync ./files/ dummy:/mnt/ --delete=true --strategy=tar
oc delete pod dummy
