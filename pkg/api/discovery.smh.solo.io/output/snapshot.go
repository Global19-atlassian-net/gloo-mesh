// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package output

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	discovery_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	discovery_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of TrafficTargets with a given set of labels
	TrafficTargets() []LabeledTrafficTargetSet
	// return the set of Workloads with a given set of labels
	Workloads() []LabeledWorkloadSet
	// return the set of Meshes with a given set of labels
	Meshes() []LabeledMeshSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	trafficTargets []LabeledTrafficTargetSet
	workloads      []LabeledWorkloadSet
	meshes         []LabeledMeshSet
	clusters       []string
}

func NewSnapshot(
	name string,

	trafficTargets []LabeledTrafficTargetSet,
	workloads []LabeledWorkloadSet,
	meshes []LabeledMeshSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) Snapshot {
	return &snapshot{
		name: name,

		trafficTargets: trafficTargets,
		workloads:      workloads,
		meshes:         meshes,
		clusters:       clusters,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	trafficTargets discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet,
	workloads discovery_smh_solo_io_v1alpha2_sets.WorkloadSet,
	meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	partitionedTrafficTargets, err := partitionTrafficTargetsByLabel(labelKey, trafficTargets)
	if err != nil {
		return nil, err
	}
	partitionedWorkloads, err := partitionWorkloadsByLabel(labelKey, workloads)
	if err != nil {
		return nil, err
	}
	partitionedMeshes, err := partitionMeshesByLabel(labelKey, meshes)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedTrafficTargets,
		partitionedWorkloads,
		partitionedMeshes,
		clusters...,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	trafficTargets discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet,
	workloads discovery_smh_solo_io_v1alpha2_sets.WorkloadSet,
	meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	labeledTrafficTargets, err := NewLabeledTrafficTargetSet(trafficTargets, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledWorkloads, err := NewLabeledWorkloadSet(workloads, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledMeshes, err := NewLabeledMeshSet(meshes, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledTrafficTargetSet{labeledTrafficTargets},
		[]LabeledWorkloadSet{labeledWorkloads},
		[]LabeledMeshSet{labeledMeshes},
		clusters...,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, cli client.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.trafficTargets {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.workloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncLocalCluster(ctx, cli, errHandler)
}

// apply the desired resources to multiple cluster states; remove stale resources where necessary
func (s *snapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.trafficTargets {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.workloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		Clusters:    s.clusters,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, errHandler)
}

func partitionTrafficTargetsByLabel(labelKey string, set discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet) ([]LabeledTrafficTargetSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "TrafficTarget", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "TrafficTarget", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha2_sets.NewTrafficTargetSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedTrafficTargets []LabeledTrafficTargetSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledTrafficTargetSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedTrafficTargets = append(partitionedTrafficTargets, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedTrafficTargets, func(i, j int) bool {
		leftLabelValue := partitionedTrafficTargets[i].Labels()[labelKey]
		rightLabelValue := partitionedTrafficTargets[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedTrafficTargets, nil
}

func partitionWorkloadsByLabel(labelKey string, set discovery_smh_solo_io_v1alpha2_sets.WorkloadSet) ([]LabeledWorkloadSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha2_sets.WorkloadSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Workload", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Workload", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha2_sets.NewWorkloadSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedWorkloads []LabeledWorkloadSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledWorkloadSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedWorkloads = append(partitionedWorkloads, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedWorkloads, func(i, j int) bool {
		leftLabelValue := partitionedWorkloads[i].Labels()[labelKey]
		rightLabelValue := partitionedWorkloads[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedWorkloads, nil
}

func partitionMeshesByLabel(labelKey string, set discovery_smh_solo_io_v1alpha2_sets.MeshSet) ([]LabeledMeshSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha2_sets.MeshSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha2_sets.NewMeshSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshes []LabeledMeshSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshes = append(partitionedMeshes, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshes, func(i, j int) bool {
		leftLabelValue := partitionedMeshes[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshes[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshes, nil
}

func (s snapshot) TrafficTargets() []LabeledTrafficTargetSet {
	return s.trafficTargets
}

func (s snapshot) Workloads() []LabeledWorkloadSet {
	return s.workloads
}

func (s snapshot) Meshes() []LabeledMeshSet {
	return s.meshes
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	trafficTargetSet := discovery_smh_solo_io_v1alpha2_sets.NewTrafficTargetSet()
	for _, set := range s.trafficTargets {
		trafficTargetSet = trafficTargetSet.Union(set.Set())
	}
	snapshotMap["trafficTargets"] = trafficTargetSet.List()
	workloadSet := discovery_smh_solo_io_v1alpha2_sets.NewWorkloadSet()
	for _, set := range s.workloads {
		workloadSet = workloadSet.Union(set.Set())
	}
	snapshotMap["workloads"] = workloadSet.List()
	meshSet := discovery_smh_solo_io_v1alpha2_sets.NewMeshSet()
	for _, set := range s.meshes {
		meshSet = meshSet.Union(set.Set())
	}
	snapshotMap["meshes"] = meshSet.List()

	return json.Marshal(snapshotMap)
}

// LabeledTrafficTargetSet represents a set of trafficTargets
// which share a common set of labels.
// These labels are used to find diffs between TrafficTargetSets.
type LabeledTrafficTargetSet interface {
	// returns the set of Labels shared by this TrafficTargetSet
	Labels() map[string]string

	// returns the set of TrafficTargetes with the given labels
	Set() discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledTrafficTargetSet struct {
	set    discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet
	labels map[string]string
}

func NewLabeledTrafficTargetSet(set discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet, labels map[string]string) (LabeledTrafficTargetSet, error) {
	// validate that each TrafficTarget contains the labels, else this is not a valid LabeledTrafficTargetSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on TrafficTarget %v", k, v, item.Name)
			}
		}
	}

	return &labeledTrafficTargetSet{set: set, labels: labels}, nil
}

func (l *labeledTrafficTargetSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledTrafficTargetSet) Set() discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet {
	return l.set
}

func (l labeledTrafficTargetSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha2.TrafficTargetList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "TrafficTarget",
	}
}

// LabeledWorkloadSet represents a set of workloads
// which share a common set of labels.
// These labels are used to find diffs between WorkloadSets.
type LabeledWorkloadSet interface {
	// returns the set of Labels shared by this WorkloadSet
	Labels() map[string]string

	// returns the set of Workloades with the given labels
	Set() discovery_smh_solo_io_v1alpha2_sets.WorkloadSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledWorkloadSet struct {
	set    discovery_smh_solo_io_v1alpha2_sets.WorkloadSet
	labels map[string]string
}

func NewLabeledWorkloadSet(set discovery_smh_solo_io_v1alpha2_sets.WorkloadSet, labels map[string]string) (LabeledWorkloadSet, error) {
	// validate that each Workload contains the labels, else this is not a valid LabeledWorkloadSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Workload %v", k, v, item.Name)
			}
		}
	}

	return &labeledWorkloadSet{set: set, labels: labels}, nil
}

func (l *labeledWorkloadSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledWorkloadSet) Set() discovery_smh_solo_io_v1alpha2_sets.WorkloadSet {
	return l.set
}

func (l labeledWorkloadSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha2.WorkloadList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "Workload",
	}
}

// LabeledMeshSet represents a set of meshes
// which share a common set of labels.
// These labels are used to find diffs between MeshSets.
type LabeledMeshSet interface {
	// returns the set of Labels shared by this MeshSet
	Labels() map[string]string

	// returns the set of Meshes with the given labels
	Set() discovery_smh_solo_io_v1alpha2_sets.MeshSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshSet struct {
	set    discovery_smh_solo_io_v1alpha2_sets.MeshSet
	labels map[string]string
}

func NewLabeledMeshSet(set discovery_smh_solo_io_v1alpha2_sets.MeshSet, labels map[string]string) (LabeledMeshSet, error) {
	// validate that each Mesh contains the labels, else this is not a valid LabeledMeshSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Mesh %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshSet{set: set, labels: labels}, nil
}

func (l *labeledMeshSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshSet) Set() discovery_smh_solo_io_v1alpha2_sets.MeshSet {
	return l.set
}

func (l labeledMeshSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha2.MeshList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "Mesh",
	}
}

type builder struct {
	ctx      context.Context
	name     string
	clusters []string

	trafficTargets discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet
	workloads      discovery_smh_solo_io_v1alpha2_sets.WorkloadSet
	meshes         discovery_smh_solo_io_v1alpha2_sets.MeshSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		trafficTargets: discovery_smh_solo_io_v1alpha2_sets.NewTrafficTargetSet(),
		workloads:      discovery_smh_solo_io_v1alpha2_sets.NewWorkloadSet(),
		meshes:         discovery_smh_solo_io_v1alpha2_sets.NewMeshSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add TrafficTargets to the collected outputs
	AddTrafficTargets(trafficTargets ...*discovery_smh_solo_io_v1alpha2.TrafficTarget)

	// get the collected TrafficTargets
	GetTrafficTargets() discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet

	// add Workloads to the collected outputs
	AddWorkloads(workloads ...*discovery_smh_solo_io_v1alpha2.Workload)

	// get the collected Workloads
	GetWorkloads() discovery_smh_solo_io_v1alpha2_sets.WorkloadSet

	// add Meshes to the collected outputs
	AddMeshes(meshes ...*discovery_smh_solo_io_v1alpha2.Mesh)

	// get the collected Meshes
	GetMeshes() discovery_smh_solo_io_v1alpha2_sets.MeshSet

	// build the collected outputs into a label-partitioned snapshot
	BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error)

	// build the collected outputs into a snapshot with a single partition
	BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error)

	// add a cluster to the collected clusters.
	// this can be used to collect clusters for use with MultiCluster snapshots.
	AddCluster(cluster string)
}

func (b *builder) AddTrafficTargets(trafficTargets ...*discovery_smh_solo_io_v1alpha2.TrafficTarget) {
	for _, obj := range trafficTargets {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output TrafficTarget %v", sets.Key(obj))
		b.trafficTargets.Insert(obj)
	}
}
func (b *builder) AddWorkloads(workloads ...*discovery_smh_solo_io_v1alpha2.Workload) {
	for _, obj := range workloads {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output Workload %v", sets.Key(obj))
		b.workloads.Insert(obj)
	}
}
func (b *builder) AddMeshes(meshes ...*discovery_smh_solo_io_v1alpha2.Mesh) {
	for _, obj := range meshes {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output Mesh %v", sets.Key(obj))
		b.meshes.Insert(obj)
	}
}

func (b *builder) GetTrafficTargets() discovery_smh_solo_io_v1alpha2_sets.TrafficTargetSet {
	return b.trafficTargets
}
func (b *builder) GetWorkloads() discovery_smh_solo_io_v1alpha2_sets.WorkloadSet {
	return b.workloads
}
func (b *builder) GetMeshes() discovery_smh_solo_io_v1alpha2_sets.MeshSet {
	return b.meshes
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.trafficTargets,
		b.workloads,
		b.meshes,
		b.clusters...,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.trafficTargets,
		b.workloads,
		b.meshes,
		b.clusters...,
	)
}

func (b *builder) AddCluster(cluster string) {
	b.clusters = append(b.clusters, cluster)
}