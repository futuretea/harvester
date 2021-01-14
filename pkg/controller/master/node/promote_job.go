package node

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/name"
)

const (
	promoteJobNamespaceName = "harvester-system"

	promoteImage           = "busybox:1.32.0"
	promoteImagePullPolicy = corev1.PullIfNotPresent
	promoteRootMountPath   = "/host"

	// restart after modify the k3s service profile
	promoteCommand = `echo start promote && \
if [ ! -f /var/lib/rancher/k3os/config.yaml ]; then \
	sudo cp /k3os/system/config.yaml /var/lib/rancher/k3os/config.yaml && \
	echo clone config; \
fi && \
echo update config && \
sudo sed -i 's/agent/server/g' /var/lib/rancher/k3os/config.yaml && \
sudo sed -i 's/command_args=\"agent/command_args=\"server/g' /etc/init.d/k3s-service && \
echo restart k3s && \
cat /var/run/k3s-restarter-trap.pid | xargs -r kill -HUP && \
echo finish promote
`
)

var (
	promotePrivileged   = true
	promoteBackoffLimit = int32(2)

	ConditionJobComplete = condition.Cond(batchv1.JobComplete)
	ConditionJobFailed   = condition.Cond(batchv1.JobFailed)
)

func buildPromoteJobName(nodeName string) string {
	return name.SafeConcatName("harvester", "promote", nodeName)
}

func (h *Handler) createPromoteJob(node *corev1.Node) (*batchv1.Job, error) {
	job := buildPromoteJob(node.Name, promoteImage, []string{"sh"}, []string{"-c",
		fmt.Sprintf(`chroot %s bash -c "%s"`, promoteRootMountPath, promoteCommand)})
	j, err := h.jobs.Create(job)
	return j, checkError(fmt.Sprintf("create job %s on node %s", job.Name, node.Name), err)
}

func (h *Handler) deleteJob(job *batchv1.Job, deletionPropagation metav1.DeletionPropagation) error {
	err := h.jobs.Delete(job.Namespace, job.Name, &metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return checkError(fmt.Sprintf("delete job %s", job.Name), err)
}

func buildPromoteJob(nodeName string, image string, command, args []string) *batchv1.Job {
	hostPathDirectory := corev1.HostPathDirectory
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      buildPromoteJobName(nodeName),
			Namespace: promoteJobNamespaceName,
			Labels: labels.Set{
				HarvesterPromoteNodeLabelKey: nodeName,
			},
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &promoteBackoffLimit,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels.Set{
						HarvesterPromoteNodeLabelKey: nodeName,
					},
				},
				Spec: corev1.PodSpec{
					HostIPC:     true,
					HostPID:     true,
					HostNetwork: true,
					DNSPolicy:   corev1.DNSClusterFirstWithHostNet,
					Affinity: &corev1.Affinity{
						NodeAffinity: &corev1.NodeAffinity{
							RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
								NodeSelectorTerms: []corev1.NodeSelectorTerm{{
									MatchExpressions: []corev1.NodeSelectorRequirement{{
										Key:      corev1.LabelHostname,
										Operator: corev1.NodeSelectorOpIn,
										Values: []string{
											nodeName,
										},
									}},
								}},
							},
						},
						PodAntiAffinity: &corev1.PodAntiAffinity{
							RequiredDuringSchedulingIgnoredDuringExecution: []corev1.PodAffinityTerm{
								{
									LabelSelector: &metav1.LabelSelector{
										MatchExpressions: []metav1.LabelSelectorRequirement{
											{
												Key:      HarvesterPromoteNodeLabelKey,
												Operator: metav1.LabelSelectorOpIn,
												Values: []string{
													nodeName,
												},
											},
										},
									},
									TopologyKey: corev1.LabelHostname,
								},
							},
						},
					},
					Tolerations: []corev1.Toleration{
						{
							Key:      corev1.TaintNodeUnschedulable,
							Operator: corev1.TolerationOpExists,
							Effect:   corev1.TaintEffectNoSchedule,
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{{
						Name: `host-root`,
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/", Type: &hostPathDirectory,
							},
						},
					}},
				},
			},
		},
	}
	podTemplate := &job.Spec.Template

	podTemplate.Spec.Containers = []corev1.Container{
		{
			Name:      "promote",
			Image:     image,
			Command:   command,
			Args:      args,
			Resources: corev1.ResourceRequirements{},
			VolumeMounts: []corev1.VolumeMount{
				{Name: "host-root", MountPath: promoteRootMountPath},
			},
			ImagePullPolicy: promoteImagePullPolicy,
			SecurityContext: &corev1.SecurityContext{
				Capabilities: &corev1.Capabilities{
					Add: []corev1.Capability{
						"CAP_SYS_BOOT",
					},
				},
				Privileged: &promotePrivileged,
			},
		},
	}

	return job
}
