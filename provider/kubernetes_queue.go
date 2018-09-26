package provider

import (
	"github.com/datakube/datakube/pkg/apis/backuptarget/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func NewBackupEventHandler() backupEventHandler {
	return backupEventHandler{
		queue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
	}
}

type backupEventHandler struct {
	queue     workqueue.RateLimitingInterface
	Notify 	  chan string
}

func (b backupEventHandler) OnAdd(obj interface{}) {
	target := obj.(*v1.BackupTarget)
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	b.addToQueue(key, err)
	b.Notify <- target.Namespace
}

func (b backupEventHandler) OnUpdate(oldObj, newObj interface{}) {
	target := newObj.(*v1.BackupTarget)
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(newObj)
	b.addToQueue(key, err)
	b.Notify <- target.Namespace
}

func (b backupEventHandler) OnDelete(obj interface{}) {
	target := obj.(*v1.BackupTarget)
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	b.addToQueue(key, err)
	b.Notify <- target.Namespace
}

func (b backupEventHandler) addToQueue(key string, err error) {
	if err == nil {
		b.queue.Add(key)
	}
}

func (b backupEventHandler) ShutDown() {
	b.queue.ShutDown()
}

func (b backupEventHandler) Get() (interface{}, bool) {
	return b.queue.Get()
}

func (b backupEventHandler) Done(item interface{}) {
	b.queue.Done(item)
}
