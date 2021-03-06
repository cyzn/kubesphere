/*
Copyright 2018 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"time"

	"k8s.io/client-go/informers"
)

func (ctl *ControllerRevisionCtl) Name() string {
	return ctl.CommonAttribute.Name
}

func (ctl *ControllerRevisionCtl) sync(stopChan chan struct{}) {

	ctl.initListerAndInformer()
	ctl.informer.Run(stopChan)
}

func (ctl *ControllerRevisionCtl) total() int {

	return 0
}

func (ctl *ControllerRevisionCtl) initListerAndInformer() {

	informerFactory := informers.NewSharedInformerFactory(ctl.K8sClient, time.Second*resyncCircle)

	ctl.lister = informerFactory.Apps().V1().ControllerRevisions().Lister()

	informer := informerFactory.Apps().V1().ControllerRevisions().Informer()

	ctl.informer = informer
}

func (ctl *ControllerRevisionCtl) CountWithConditions(conditions string) int {

	return 0
}

func (ctl *ControllerRevisionCtl) ListWithConditions(conditions string, paging *Paging, order string) (int, interface{}, error) {

	return 0, nil, nil
}

func (ctl *ControllerRevisionCtl) Lister() interface{} {

	return ctl.lister
}
