package fake

import (
	harvesterhciv1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	k8scnicncfiov1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	snapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	upgradev1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubevirtv1 "kubevirt.io/client-go/api/v1"
	clusterv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	clusterv1alpha4.AddToScheme,
	harvesterhciv1beta1.AddToScheme,
	k8scnicncfiov1.AddToScheme,
	kubevirtv1.AddToScheme,
	longhornv1beta1.AddToScheme,
	networkingv1.AddToScheme,
	snapshotv1beta1.AddToScheme,
	upgradev1.AddToScheme,
}

var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	__traceStack()

	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
