package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	util "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/setcreed/store-apiserver/pkg/apis/store"
	"github.com/setcreed/store-apiserver/pkg/apis/store/v1alpha1"
)

// 当我们有了scheme实例时，就可以来调这个install来把这个api server支持的object信息注册进来了
func Install(scheme *runtime.Scheme) {
	util.Must(store.AddToScheme(scheme))
	util.Must(v1alpha1.AddToScheme(scheme))
	util.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
