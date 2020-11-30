/*
Copyright 2020 Baechul.

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
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mygroupv1 "github.com/ibestgo/hello-operator/api/v1"
)

// MyKindReconciler reconciles a MyKind object
type MyKindReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=mygroup.hello.k8s.io,resources=mykinds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mygroup.hello.k8s.io,resources=mykinds/status,verbs=get;update;patch

func (r *MyKindReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("mykind", req.NamespacedName)

	// your logic here - start
	// 1. Load the Named MyKind
	var msg mygroupv1.MyKind
	if err := r.Get(ctx, req.NamespacedName, &msg); err != nil {
		logger.Error(err, "Unable to fetch MyKind")
		return ctrl.Result{}, nil
	}
	logger.Info("A new MyKind posted: ", "Message", msg.Spec.Message, "Image", msg.Spec.Image)

	// 2.  Pod Creation Business Logic TODO

	// 3. Update the CRD instance status
	msg.Status.Printed = true
	msg.Status.PrintedDate = time.Now().Format(time.RFC3339)
	if err := r.Update(ctx, &msg); err != nil {
		logger.Error(err, "Unable to update MyKind")
		return ctrl.Result{}, err
	}

	// your logic here - end

	return ctrl.Result{}, nil
}

func (r *MyKindReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mygroupv1.MyKind{}).
		Complete(r)
}
