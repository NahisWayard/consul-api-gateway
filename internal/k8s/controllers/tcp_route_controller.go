package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	gwv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul-api-gateway/internal/k8s/gatewayclient"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/reconciler"
)

// TCPRouteReconciler reconciles a TCPRoute object
type TCPRouteReconciler struct {
	Context        context.Context
	Client         gatewayclient.Client
	Log            hclog.Logger
	ControllerName string
	Manager        reconciler.ReconcileManager
}

//+kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=tcproutes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=tcproutes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=tcproutes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *TCPRouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.With("tcp-route", req.NamespacedName)

	route, err := r.Client.GetTCPRoute(ctx, req.NamespacedName)
	if err != nil {
		logger.Error("failed to get tcp route", "error", err)
		return ctrl.Result{}, err
	}

	if route == nil {
		// clean up cached resources
		err := r.Manager.DeleteTCPRoute(ctx, req.NamespacedName)
		return ctrl.Result{}, err
	}

	// let the route get upserted so long as there's a single gateway we control
	// that it's managed by -- the underlying reconciliation code will handle the
	// validation of gateway attachment
	err = r.Manager.UpsertTCPRoute(ctx, route)
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *TCPRouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gwv1alpha2.TCPRoute{}).
		Watches(
			&source.Kind{Type: &gwv1alpha2.ReferenceGrant{}},
			handler.EnqueueRequestsFromMapFunc(r.referenceGrantToRouteRequests),
		).
		Watches(
			&source.Kind{Type: &gwv1alpha2.ReferencePolicy{}},
			handler.EnqueueRequestsFromMapFunc(r.referencePolicyToRouteRequests),
		).
		Complete(gatewayclient.NewRequeueingMiddleware(r.Log, r))
}

func (r *TCPRouteReconciler) referenceGrantToRouteRequests(object client.Object) []reconcile.Request {
	return r.getRouteRequestsFromReferenceGrant(object.(*gwv1alpha2.ReferenceGrant))
}

func (r *TCPRouteReconciler) referencePolicyToRouteRequests(object client.Object) []reconcile.Request {
	refPolicy := object.(*gwv1alpha2.ReferencePolicy)
	refGrant := gwv1alpha2.ReferenceGrant{Spec: refPolicy.Spec}
	return r.getRouteRequestsFromReferenceGrant(&refGrant)
}

// For UpdateEvents which contain both a new and old object, this transformation
// function is run on both objects and both sets of Requests are enqueued.
//
// This is needed to reconcile any objects matched by both current and prior
// state in case a ReferenceGrant has been modified to revoke permission from a
// namespace or to a service
//
// It may be possible to improve performance here by filtering Routes by
// BackendRefs selectable by the To fields, but currently we just revalidate
// all Routes allowed in the From Namespaces
func (r *TCPRouteReconciler) getRouteRequestsFromReferenceGrant(refGrant *gwv1alpha2.ReferenceGrant) []reconcile.Request {
	routes := r.getRoutesAffectedByReferenceGrant(refGrant)
	requests := []reconcile.Request{}

	for _, route := range routes {
		requests = append(requests, reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      route.Name,
				Namespace: route.Namespace,
			},
		})
	}

	return requests
}

// getRoutesAffectedByReferencePolicy retrieves all TCPRoutes potentially impacted
// by the ReferencePolicy being modified. Currently, this is unfiltered and so returns
// all TCPRoutes in the namespace referenced by the ReferencePolicy.
func (r *TCPRouteReconciler) getRoutesAffectedByReferenceGrant(refGrant *gwv1alpha2.ReferenceGrant) []gwv1alpha2.TCPRoute {
	var matches []gwv1alpha2.TCPRoute

	for _, from := range refGrant.Spec.From {
		// TODO: search by from.Group and from.Kind instead of assuming this ReferenceGrant references a TCPRoute
		routes, err := r.Client.GetTCPRoutesInNamespace(r.Context, string(from.Namespace))
		if err != nil {
			r.Log.Error("error fetching routes", err)
			return matches
		}

		matches = append(matches, routes...)
	}

	return matches
}
