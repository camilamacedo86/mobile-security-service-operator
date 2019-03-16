package mobilesecurityservice

import (
	mobilesecurityservicev1alpha1 "github.com/aerogear/mobile-security-service-operator/pkg/apis/mobilesecurityservice/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// getServiceForMobileSecurityService returns a MobileSecurityService Service resource
func (r *ReconcileMobileSecurityService) getConfigMapForMobileSecurityService(m *mobilesecurityservicev1alpha1.MobileSecurityService) *corev1.ConfigMap {
	ls := labelsForMobileSecurityService(m.Name)
	ser := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
			Labels:    ls,
		},
		Data: getConfigMapForMobileSecurityService(),
	}
	// Set MobileSecurityService instance as the owner and controller
	controllerutil.SetControllerReference(m, ser, r.scheme)
	return ser
}
