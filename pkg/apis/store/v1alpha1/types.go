package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type StoreService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,name=metadata"`

	Spec   StoreServiceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status StoreServiceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type StoreServiceStatus struct {
	Replicas int32  `json:"replicas,omitempty"`
	Ready    string `json:"ready,omitempty"`
}

type StoreServiceSpec struct {
	DbConfig  DBConfig     `json:"dbConfig,omitempty"`
	SqlConfig []*SQLConfig `json:"sqlConfig,omitempty"`
}

type DBConfig struct {
	Replicas    int    `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	Dsn         string `json:"dsn,omitempty"`
	MaxOpenConn int    `json:"maxOpenConn,omitempty"`
	MaxLifeTime int    `json:"maxLifeTime,omitempty"`
	MaxIdleConn int    `json:"maxIdleConn,omitempty"`
}

type SQLConfig struct {
	Name   string  `json:"name,omitempty"`
	Sql    string  `json:"sql,omitempty"`
	Select *Select `json:"select,omitempty"`
}

type Select struct {
	Sql string `json:"sql,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type StoreServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []StoreService `json:"items" protobuf:"bytes,2,rep,name=items"`
}
