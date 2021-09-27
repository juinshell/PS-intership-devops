# 目的
创建 Kubernetes 对象时，必须提供对象的规约，用来描述该对象的期望状态， 以及关于对象的一些基本信息（例如名称）。 当使用 Kubernetes API 创建对象时（或者直接创建，或者基于kubectl）， API 请求必须在请求体中包含 JSON 格式的信息。 大多数情况下，需要在 .yaml 文件中为 kubectl 提供这些信息。 kubectl 在发起 API 请求时，将这些信息转换成 JSON 格式。
# 官方例子
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  # Unique key of the Deployment instance
  name: deployment-example
spec:
  # 3 Pods should exist at all times.
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        # Apply this label to pods and default
        # the Deployment label selector to this value
        app: nginx
    spec:
      containers:
      - name: nginx
        # Run this image
        image: nginx:1.14
```
# 编写
## 必须字段
* apiVersion - 创建该对象所使用的 Kubernetes API 的版本
* kind - 想要创建的对象的类别
* metadata - 帮助唯一性标识对象的一些数据，包括一个 name 字符串、UID 和可选的 namespace
## Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80

```
1. .spec.replicas字段指示Pod副本数目
2. .spec.selector字段定义了 Deployment 如何找到要管理的 Pod
* template PodTemplateSpec	Template describes the pods that will be created.
* 参考文档https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#deploymentspec-v1-apps
## Server类型
### Type 的取值以及行为如下：
* ClusterIP：通过集群的内部 IP 暴露服务，选择该值时服务只能够在集群内部访问。 这也是默认的 ServiceType。
* NodePort：通过每个节点上的 IP 和静态端口（NodePort）暴露服务。 NodePort 服务会路由到自动创建的 ClusterIP 服务。 通过请求 <节点 IP>:<节点端口>，你可以从集群的外部访问一个 NodePort 服务。
https://www.qikqiak.com/k8s-book/docs/18.YAML%20%E6%96%87%E4%BB%B6.html
* targetport:要代理的集群内的pod端口（可以用）
* port：集群内其他部分访问service时的目标端口
* 