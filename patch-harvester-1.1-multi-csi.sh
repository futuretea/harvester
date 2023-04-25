#!/usr/bin/env bash
[[ -n $DEBUG ]] && set -x
set -eou pipefail

usage() {
    cat <<HELP
USAGE:
    patch-harvester-1.1-multi-csi.sh KUBECONFIG TARGET_REPO TARGET_TAG
HELP
}

if [ $# -lt 3 ]; then
    usage
    exit 1
fi

export KUBECONFIG=$1
TARGET_REPO=$2
TARGET_TAG=$3

tmp_file="$(mktemp)"
cat > "${tmp_file}" <<EOF
spec:
  values:
    containers:
      apiserver:
        image:
          repository: ${TARGET_REPO}/harvester
          tag: ${TARGET_TAG}
    webhook:
      image:
        repository: ${TARGET_REPO}/harvester-webhook
        tag: ${TARGET_TAG}
EOF
kubectl -n fleet-local patch managedchart harvester --patch-file="${tmp_file}" --type merge
set +e
for _ in $(seq 1 10); do
    kubectl -n harvester-system get rs --no-headers -l app.kubernetes.io/name=harvester -owide | grep -v ${TARGET_TAG} | awk '{print $1}' | while read -r rs; do
      kubectl -n harvester-system delete rs "${rs}" --ignore-not-found=true
      sleep 3
    done
done
set -e
kubectl -n harvester-system wait --for=condition=Available deploy harvester --timeout 10m
kubectl -n harvester-system wait --for=condition=Available deploy harvester-webhook --timeout 10m
kubectl -n fleet-local wait --for=condition=Ready managedchart harvester --timeout 10m

cat > "${tmp_file}" <<EOF
value: 'bundled'
EOF
kubectl -n harvester-system patch setting ui-source --patch-file="${tmp_file}" --type merge
