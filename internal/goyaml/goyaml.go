package goyaml

import (
	"cuelang.org/go/cue/load"
)

func goyaml() {

	// create a context
	v := load.FromString()(YML)

}

var YML = `
apiVersion: infra.tce.io/v1
kind: ServicePlan
metadata:
    name: zk-gogo-test
spec:
    serviceClass: zookeeper
    metadata:
      cpu: 1
      storage: 50G
      memory: 1G

---
apiVersion: infra.tce.io/v1
kind: ServiceInstance
metadata:
    name: zk-gogo-test
    namespace: sso
spec:
    serviceClass: zookeeper
    servicePlan: zk-gogo-test
    parameters:
        version: 3.4.14

---
apiVersion: infra.tce.io/v1
kind: ServiceBinding
metadata:
    name: zk-gogo-test
    namespace: sso
spec:
    instanceRef:
        name: zk-gogo-test
        namespace: sso
`
