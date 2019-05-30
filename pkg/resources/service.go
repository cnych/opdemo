package resources

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"github.com/cnych/opdemo/pkg/apis/app/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewService(app *v1.AppService) *corev1.Service {
	return &corev1.Service {
		TypeMeta: metav1.TypeMeta {
			Kind: "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, schema.GroupVersionKind{
					Group: v1.SchemeGroupVersion.Group,
					Version: v1.SchemeGroupVersion.Version,
					Kind: "AppService",
				}),
			},
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Ports: app.Spec.Ports,
			Selector: map[string]string{
				"app": app.Name,
			},
		},
	}
}
