package main

import (
	"context"
	"github.com/VajiraPrabuddhaka/k8s-informer-for-crd/internal/api/v1alpha1"
	"github.com/VajiraPrabuddhaka/k8s-informer-for-crd/internal/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"log"
	"time"
)

func main() {

	stop := make(chan struct{})
	defer close(stop)
	//err = controller.Run(stop)
	c, _ := WatchResources(client.GetOutClusterClientSetV1alpha1())

	c.Run(stop)

	select {}

}

func WatchResources(clientSet v1alpha1.APIV1Alpha1Interface) (cache.Controller, cache.Store) {
	h := APIEventHandler{}
	apiStore, apiController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.APIs("default").List(context.TODO(), lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.APIs("default").Watch(context.TODO(), lo)
			},
		},
		&v1alpha1.API{},
		55*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    h.onAdd,
			UpdateFunc: h.OnUpdate,
			DeleteFunc: h.OnDelete,
		},
	)

	//go httpRouteController.Run(wait.NeverStop)
	return apiController, apiStore
}

// APIEventHandler is used to provide functions for resource event handler
type APIEventHandler struct {
}

func (h *APIEventHandler) onAdd(obj interface{}) {
	log.Printf("onAdd called : %v", obj)
}

func (h *APIEventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	log.Printf("onUpdate called, oldObj: %v newObj:%v", oldObj, newObj)
}

func (h *APIEventHandler) OnDelete(obj interface{}) {
	log.Printf("onAdd called : %v", obj)
}
