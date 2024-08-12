#!/bin/sh
set -x -e
oc new-project jumpstarter-pipelines

oc apply -f jumpstarter-pipelines-scc.yaml
oc adm policy add-scc-to-user jumpstarter-pipelines-scc -z pipeline

cd image
./copy-to-cluster.sh
cd ..
oc apply -f pipelines/task-git-clone.yaml
oc apply -f pipelines/task-prepare-image.yaml
oc apply -f pipelines/task-jumpstarter-script.yaml
oc apply -f pipelines/pipeline-jumpstarter-orin-nx.yaml

echo you can now run the pipeline with:
echo ./run-pipeline.sh


