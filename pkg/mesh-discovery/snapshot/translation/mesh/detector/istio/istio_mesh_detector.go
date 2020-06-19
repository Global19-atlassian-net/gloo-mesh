package istio

import (
	"github.com/rotisserie/eris"
	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/smh/pkg/mesh-discovery/snapshot/translation/mesh/detector"
	"github.com/solo-io/smh/pkg/mesh-discovery/snapshot/translation/utils"
	istiov1alpha1 "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/util/gogoprotomarshal"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

const (
	legacyPilotDeploymentName = "istio-pilot"
	istiodDeploymentName      = "istiod"
	istioContainerKeyword     = "istio"
	pilotContainerKeyword     = "pilot"
	istioConfigMapName        = "istio"
	istioConfigMapMeshDataKey = "mesh"
)

// detects Istio if a deployment contains the istiod container.
type meshDetector struct {
	configMaps corev1sets.ConfigMapSet
}

func NewMeshDetector(configMaps corev1sets.ConfigMapSet) detector.MeshDetector {
	return &meshDetector{configMaps: configMaps}
}

// returns nil, nil of the deployment does not contain the istiod image
func (d *meshDetector) DetectMesh(deployment *appsv1.Deployment) (*v1alpha1.Mesh, error) {
	version, err := d.getIstiodVersion(deployment)
	if err != nil {
		return nil, err
	}

	if version == "" {
		return nil, nil
	}

	trustDomain, err := getTrustDomain(d.configMaps, deployment.ClusterName, deployment.Namespace)
	if err != nil {
		return nil, err
	}

	mesh := &v1alpha1.Mesh{
		ObjectMeta: utils.DiscoveredObjectMeta(deployment),
		Spec: v1alpha1.MeshSpec{
			MeshType: &v1alpha1.MeshSpec_Istio_{
				Istio: &v1alpha1.MeshSpec_Istio{
					Installation: &v1alpha1.MeshSpec_MeshInstallation{
						Namespace: deployment.Namespace,
						Cluster:   deployment.ClusterName,
						Version:   version,
					},
					CitadelInfo: &v1alpha1.MeshSpec_Istio_CitadelInfo{
						TrustDomain: trustDomain,
						// This assumes that the istiod deployment is the cert provider
						CitadelServiceAccount: deployment.Spec.Template.Spec.ServiceAccountName,
					},
				},
			},
		},
	}

	return mesh, nil
}

func (d *meshDetector) getIstiodVersion(deployment *appsv1.Deployment) (string, error) {
	for _, container := range deployment.Spec.Template.Spec.Containers {
		if isIstiod(deployment, &container) {
			parsedImage, err := docker.ParseImageName(container.Image)
			if err != nil {
				return "", eris.Wrapf(err, "failed to parse istiod image tag: %s", container.Image)
			}
			version := parsedImage.Tag
			if parsedImage.Digest != "" {
				version = parsedImage.Digest
			}
			return version, nil
		}
	}
	return "", nil
}

// Return true if deployment is inferred to be an Istiod deployment
func isIstiod(deployment *appsv1.Deployment, container *corev1.Container) bool {
	return (deployment.GetName() == istiodDeploymentName || deployment.GetName() == legacyPilotDeploymentName) &&
		strings.Contains(container.Image, istioContainerKeyword) &&
		strings.Contains(container.Image, pilotContainerKeyword)
}

func getTrustDomain(
	configMaps corev1sets.ConfigMapSet,
	cluster,
	namespace string,
) (string, error) {
	istioConfigMap, err := configMaps.Find(&v1.ClusterObjectRef{
		Name:        istioConfigMapName,
		Namespace:   namespace,
		ClusterName: cluster,
	})
	if err != nil {
		return "", err
	}

	meshConfigString, ok := istioConfigMap.Data[istioConfigMapMeshDataKey]
	if !ok {
		return "", eris.Errorf("Failed to find 'mesh' entry in ConfigMap with name/namespace/cluster %s/%s/%s", istioConfigMapName, namespace, cluster)
	}
	var meshConfig istiov1alpha1.MeshConfig
	err = gogoprotomarshal.ApplyYAML(meshConfigString, &meshConfig)
	if err != nil {
		return "", eris.Errorf("Failed to find 'mesh' entry in ConfigMap with name/namespace/cluster %s/%s/%s", istioConfigMapName, namespace, cluster)
	}
	return meshConfig.TrustDomain, nil
}
