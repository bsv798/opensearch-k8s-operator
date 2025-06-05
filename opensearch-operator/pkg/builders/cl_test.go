package builders

import (
	"testing"

	opsterv1 "github.com/Opster/opensearch-k8s-operator/opensearch-operator/api/v1"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func TestAdd1(t *testing.T) {
	clusterObject := ClusterDescWithVersion("2.2.1")
	result := NewSTSForNodePool("foobar", &clusterObject, opsterv1.NodePool{}, "foobar", nil, nil, nil)
	Expect(len(result.Spec.Template.Spec.InitContainers)).To(Equal(1))
}

func TestAdd2(t *testing.T) {
	clusterObject := ClusterDescWithVersion("1.3.0")
	nodePool := opsterv1.NodePool{
		Component:         "masters",
		Roles:             []string{"cluster_manager"},
		PriorityClassName: "default",
	}
	result := NewSTSForNodePool("foobar", &clusterObject, nodePool, "foobar", nil, nil, nil)
	Expect(result.Spec.Template.Spec.PriorityClassName).To(Equal("default"))
}

func TestAdd3(t *testing.T) {
	user := int64(1000)
	podSecurityContext := &corev1.PodSecurityContext{
		RunAsUser:    &user,
		RunAsGroup:   &user,
		RunAsNonRoot: pointer.Bool(true),
	}
	securityContext := &corev1.SecurityContext{
		Privileged:               pointer.Bool(false),
		AllowPrivilegeEscalation: pointer.Bool(false),
	}
	clusterObject := ClusterDescWithVersion("2.2.1")
	clusterObject.Spec.General.PodSecurityContext = podSecurityContext
	clusterObject.Spec.General.SecurityContext = securityContext
	nodePool := opsterv1.NodePool{
		Replicas:  3,
		Component: "masters",
		Roles:     []string{"cluster_manager", "data"},
	}
	clusterObject.Spec.NodePools = append(clusterObject.Spec.NodePools, nodePool)
	result := NewSTSForNodePool("foobar", &clusterObject, opsterv1.NodePool{}, "foobar", nil, nil, nil)
	Expect(result.Spec.Template.Spec.SecurityContext).To(Equal(podSecurityContext))
	Expect(result.Spec.Template.Spec.Containers[0].SecurityContext).To(Equal(securityContext))
}

func TestAdd4(t *testing.T) {
	user := int64(1000)
	podSecurityContext := &corev1.PodSecurityContext{
		RunAsUser:    &user,
		RunAsGroup:   &user,
		RunAsNonRoot: pointer.Bool(true),
	}
	securityContext := &corev1.SecurityContext{
		Privileged:               pointer.Bool(false),
		AllowPrivilegeEscalation: pointer.Bool(false),
	}
	clusterObject := ClusterDescWithVersion("2.2.1")
	clusterObject.Spec.General.PodSecurityContext = podSecurityContext
	clusterObject.Spec.General.SecurityContext = securityContext
	nodePool := opsterv1.NodePool{
		Replicas:  1,
		Component: "masters",
		Roles:     []string{"cluster_manager", "data"},
	}
	clusterObject.Spec.NodePools = append(clusterObject.Spec.NodePools, nodePool)
	result := NewSTSForNodePool("foobar", &clusterObject, opsterv1.NodePool{}, "foobar", nil, nil, nil)
	Expect(result.Spec.Template.Spec.SecurityContext).To(Equal(podSecurityContext))
	Expect(result.Spec.Template.Spec.Containers[0].SecurityContext).To(Equal(securityContext))
}

func TestAdd5(t *testing.T) {
	user := int64(1000)
	podSecurityContext := &corev1.PodSecurityContext{
		RunAsUser:    &user,
		RunAsGroup:   &user,
		RunAsNonRoot: pointer.Bool(true),
	}
	securityContext := &corev1.SecurityContext{
		Privileged:               pointer.Bool(false),
		AllowPrivilegeEscalation: pointer.Bool(false),
	}
	clusterObject := ClusterDescWithVersion("2.2.1")
	clusterObject.Spec.General.PodSecurityContext = podSecurityContext
	clusterObject.Spec.General.SecurityContext = securityContext
	nodePool := opsterv1.NodePool{
		Replicas:  1,
		Component: "datas",
		Roles:     []string{"data"},
	}
	clusterObject.Spec.NodePools = append(clusterObject.Spec.NodePools, nodePool)
	result := NewSTSForNodePool("foobar", &clusterObject, opsterv1.NodePool{}, "foobar", nil, nil, nil)
	Expect(result.Spec.Template.Spec.SecurityContext).To(Equal(podSecurityContext))
	Expect(result.Spec.Template.Spec.Containers[0].SecurityContext).To(Equal(securityContext))
}

func TestAdd6(t *testing.T) {
	user := int64(1000)
	podSecurityContext := &corev1.PodSecurityContext{
		RunAsUser:    &user,
		RunAsGroup:   &user,
		RunAsNonRoot: pointer.Bool(true),
	}
	securityContext := &corev1.SecurityContext{
		Privileged:               pointer.Bool(false),
		AllowPrivilegeEscalation: pointer.Bool(false),
	}
	clusterObject := ClusterDescWithVersion("2.2.1")
	clusterObject.Spec.General.PodSecurityContext = podSecurityContext
	clusterObject.Spec.General.SecurityContext = securityContext
	nodePool := opsterv1.NodePool{
		Replicas:  3,
		Component: "datas",
		Roles:     []string{"data"},
	}
	clusterObject.Spec.NodePools = append(clusterObject.Spec.NodePools, nodePool)
	result := NewSTSForNodePool("foobar", &clusterObject, opsterv1.NodePool{}, "foobar", nil, nil, nil)
	Expect(result.Spec.Template.Spec.SecurityContext).To(Equal(podSecurityContext))
	Expect(result.Spec.Template.Spec.Containers[0].SecurityContext).To(Equal(securityContext))
}

func TestAdd7(t *testing.T) {
	clusterObject := ClusterDescWithVersion("2.2.1")
	customRepository := "mycustomrepo.cr"
	clusterObject.Spec.General.DefaultRepo = &customRepository
	result := NewBootstrapPod(&clusterObject, nil, nil)
	Expect(result.Spec.InitContainers[0].Image).To(Equal("mycustomrepo.cr/busybox:latest"))
}
