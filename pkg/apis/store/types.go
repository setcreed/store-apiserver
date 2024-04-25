package store

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type StoreService struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   StoreServiceSpec
	Status StoreServiceStatus
}

type StoreServiceStatus struct {
	Replicas int32
	Ready    string
}

type StoreServiceSpec struct {
	DbConfig  DBConfig
	SqlConfig []*SQLConfig
}

type DBConfig struct {
	Replicas    int
	Dsn         string
	MaxOpenConn int
	MaxLifeTime int
	MaxIdleConn int
}

type SQLConfig struct {
	Name   string
	Sql    string
	Select *Select
}

type Select struct {
	Sql string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type StoreServiceList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []StoreService
}
