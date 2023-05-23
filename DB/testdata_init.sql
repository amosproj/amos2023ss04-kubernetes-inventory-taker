-- Test data for pods table
INSERT INTO pods (pod_event_id, pod_id, timestamp, cluster_id, node_id, name, status) VALUES
(1, 1, '2023-05-23 15:08:25.506', 1, 1, 'hello-node-cddb6ccd5-tbjk9', 'Running'),
(2, 2, '2023-05-23 15:10:25.206', 1, 1, 'kubernetes-bootcamp-5485cc6795-fqnrs', 'Pending'),
(3, 3, '2023-05-23 15:10:25.306', 1, 1, 'nginx-depl-56cb8b6d7-2jh4n', 'Running'),
(4, 4, NOW(), 1, 1, 'coredns-787d4945fb-hkfk2', 'Succeeded'),
(5, 5, NOW(), 1, 1, 'etcd-minikube', 'Running'),
(6, 6, NOW(), 1, 1, 'kube-apiserver-minikube', 'Failed'),
(10, 7, NOW(), 1, 1, 'kube-scheduler-minikube', 'Running'),
(8, 8, NOW(), 1, 2, 'pod with very veryveryveryveryveryveryveryveryveryveryveryvery very super long name-6f6cdbf67d', 'Unknown'),
(9, 9, NOW(), 1, 2, 'storage-provisioner', 'Running'),
(7, 10, NOW(), 1, 2, 'metrics-server-6f6cdbf67d-hw5nx', 'Running');

-- Test data for containers table
INSERT INTO containers (container_event_id, container_id, timestamp, pod_id, name, image, status, ports) VALUES
(1, 1, '2023-05-23 15:08:25.506', 1, 'container1', 'image1', 'Terminated', '8080, 8443'),
(2, 2, NOW(), 1, '2023-05-23 15:10:25.206', 'image2', 'Waiting', '9000'),
(3, 3, NOW(), 2, '2023-05-23 15:10:25.306', 'image3', 'Waiting', '8080'),
(4, 4, NOW(), 2, 'container4', 'image4', 'Running', ''),
(5, 5, NOW(), 3, 'container5', 'image5', 'Running', '8080'),
(6, 6, NOW(), 3, 'container6', 'image6', 'Running', '9000'),
(7, 7, NOW(), 4, 'container7', 'image7', 'Running', ''),
(8, 8, NOW(), 4, 'container8', 'image8', 'Running', '8080'),
(9, 9, NOW(), 5, 'container9', 'image9', 'Terminated', '9000'),
(10, 10, NOW(), 5, 'container10', 'image10', 'Running', '8080, 8443');
