const containers = [
  { name: "coredns", image: "registry.k8s.io/coredns/coredns:v1.9.3" },
  { name: "etcd", image: "registry.k8s.io/etcd:3.5.6-0" },
  { name: "kube-apiserver", image: "registry.k8s.io/kube-apiserver:v1.26.3" },
  {
    name: "kube-controller-manager",
    image: "registry.k8s.io/kube-controller-manager:v1.26.3",
  },
  { name: "kube-proxy", image: "registry.k8s.io/kube-proxy:v1.26.3" },
  { name: "kube-scheduler", image: "registry.k8s.io/kube-scheduler:v1.26.3" },
  {
    name: "kubernetes-bootcamp",
    image: "gcr.io/google-samples/kubernetes-bootcamp:v1",
  },
  {
    name: "metrics-server",
    image: "registry.k8s.io/metrics-server/metrics-server:v0.6.3",
  },
  { name: "nginx", image: "nginx" },
  {
    name: "storage-provisioner",
    image: "gcr.io/k8s-minikube/storage-provisioner:v5",
  },
  {
    name: "storage-provisionervery-veryveryveryvery-super-long-name",
    image: "gcr.io/k8s-made-up-image:v3",
  },
];

export default containers;
