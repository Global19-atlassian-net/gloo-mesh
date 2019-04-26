package kubernetes_test

import (
	"context"
	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/test/helpers"
	"k8s.io/client-go/kubernetes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	. "github.com/solo-io/supergloo/pkg/api/custom/clients/kubernetes"
	"github.com/solo-io/supergloo/test/testutils"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Namespace base client", func() {

	var (
		kube         kubernetes.Interface
		kcache       cache.KubeCoreCache
		namespace    string
		namespaceObj *kubev1.Namespace
	)

	BeforeEach(func() {
		var err error
		kube = testutils.MustKubeClient()
		kcache, err = cache.NewKubeCoreCache(context.TODO(), kube)
		Expect(err).NotTo(HaveOccurred())
		namespace = "namespaceclient-" + helpers.RandString(8)

		namespaceObj, err = kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      namespace,
			},
		})
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		err := kube.CoreV1().Namespaces().Delete(namespace, &metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())
	})
	It("converts a kubernetes namespace to solo-kit resource", func() {

		rc := NewnamespaceResourceClient(kube, kcache)

		Eventually(func() (bool, error) {
			var err error
			namespaces, err := rc.List(namespace, clients.ListOpts{})
			if err != nil {
				return false, err
			}
			for _, v := range namespaces {
				if v.GetMetadata().Name == namespaceObj.Name &&
					v.GetMetadata().Namespace == namespaceObj.Namespace {
					return true, nil
				}
			}
			return false, nil
		}, time.Second*10).Should(BeTrue())
	})
})