// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ContainerCPUTime:                     MetricConfig{Enabled: true},
					ContainerCPUUsage:                    MetricConfig{Enabled: true},
					ContainerCPUUtilization:              MetricConfig{Enabled: true},
					ContainerFilesystemAvailable:         MetricConfig{Enabled: true},
					ContainerFilesystemCapacity:          MetricConfig{Enabled: true},
					ContainerFilesystemUsage:             MetricConfig{Enabled: true},
					ContainerMemoryAvailable:             MetricConfig{Enabled: true},
					ContainerMemoryMajorPageFaults:       MetricConfig{Enabled: true},
					ContainerMemoryPageFaults:            MetricConfig{Enabled: true},
					ContainerMemoryRss:                   MetricConfig{Enabled: true},
					ContainerMemoryUsage:                 MetricConfig{Enabled: true},
					ContainerMemoryWorkingSet:            MetricConfig{Enabled: true},
					ContainerUptime:                      MetricConfig{Enabled: true},
					K8sContainerCPUNodeUtilization:       MetricConfig{Enabled: true},
					K8sContainerCPULimitUtilization:      MetricConfig{Enabled: true},
					K8sContainerCPURequestUtilization:    MetricConfig{Enabled: true},
					K8sContainerMemoryNodeUtilization:    MetricConfig{Enabled: true},
					K8sContainerMemoryLimitUtilization:   MetricConfig{Enabled: true},
					K8sContainerMemoryRequestUtilization: MetricConfig{Enabled: true},
					K8sNodeCPUTime:                       MetricConfig{Enabled: true},
					K8sNodeCPUUsage:                      MetricConfig{Enabled: true},
					K8sNodeCPUUtilization:                MetricConfig{Enabled: true},
					K8sNodeFilesystemAvailable:           MetricConfig{Enabled: true},
					K8sNodeFilesystemCapacity:            MetricConfig{Enabled: true},
					K8sNodeFilesystemUsage:               MetricConfig{Enabled: true},
					K8sNodeMemoryAvailable:               MetricConfig{Enabled: true},
					K8sNodeMemoryMajorPageFaults:         MetricConfig{Enabled: true},
					K8sNodeMemoryPageFaults:              MetricConfig{Enabled: true},
					K8sNodeMemoryRss:                     MetricConfig{Enabled: true},
					K8sNodeMemoryUsage:                   MetricConfig{Enabled: true},
					K8sNodeMemoryWorkingSet:              MetricConfig{Enabled: true},
					K8sNodeNetworkErrors:                 MetricConfig{Enabled: true},
					K8sNodeNetworkIo:                     MetricConfig{Enabled: true},
					K8sNodeUptime:                        MetricConfig{Enabled: true},
					K8sPodCPUNodeUtilization:             MetricConfig{Enabled: true},
					K8sPodCPUTime:                        MetricConfig{Enabled: true},
					K8sPodCPUUsage:                       MetricConfig{Enabled: true},
					K8sPodCPUUtilization:                 MetricConfig{Enabled: true},
					K8sPodCPULimitUtilization:            MetricConfig{Enabled: true},
					K8sPodCPURequestUtilization:          MetricConfig{Enabled: true},
					K8sPodFilesystemAvailable:            MetricConfig{Enabled: true},
					K8sPodFilesystemCapacity:             MetricConfig{Enabled: true},
					K8sPodFilesystemUsage:                MetricConfig{Enabled: true},
					K8sPodMemoryAvailable:                MetricConfig{Enabled: true},
					K8sPodMemoryMajorPageFaults:          MetricConfig{Enabled: true},
					K8sPodMemoryNodeUtilization:          MetricConfig{Enabled: true},
					K8sPodMemoryPageFaults:               MetricConfig{Enabled: true},
					K8sPodMemoryRss:                      MetricConfig{Enabled: true},
					K8sPodMemoryUsage:                    MetricConfig{Enabled: true},
					K8sPodMemoryWorkingSet:               MetricConfig{Enabled: true},
					K8sPodMemoryLimitUtilization:         MetricConfig{Enabled: true},
					K8sPodMemoryRequestUtilization:       MetricConfig{Enabled: true},
					K8sPodNetworkErrors:                  MetricConfig{Enabled: true},
					K8sPodNetworkIo:                      MetricConfig{Enabled: true},
					K8sPodUptime:                         MetricConfig{Enabled: true},
					K8sVolumeAvailable:                   MetricConfig{Enabled: true},
					K8sVolumeCapacity:                    MetricConfig{Enabled: true},
					K8sVolumeInodes:                      MetricConfig{Enabled: true},
					K8sVolumeInodesFree:                  MetricConfig{Enabled: true},
					K8sVolumeInodesUsed:                  MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					AwsVolumeID:                  ResourceAttributeConfig{Enabled: true},
					ContainerID:                  ResourceAttributeConfig{Enabled: true},
					FsType:                       ResourceAttributeConfig{Enabled: true},
					GcePdName:                    ResourceAttributeConfig{Enabled: true},
					GlusterfsEndpointsName:       ResourceAttributeConfig{Enabled: true},
					GlusterfsPath:                ResourceAttributeConfig{Enabled: true},
					K8sContainerName:             ResourceAttributeConfig{Enabled: true},
					K8sNamespaceName:             ResourceAttributeConfig{Enabled: true},
					K8sNodeName:                  ResourceAttributeConfig{Enabled: true},
					K8sPersistentvolumeclaimName: ResourceAttributeConfig{Enabled: true},
					K8sPodName:                   ResourceAttributeConfig{Enabled: true},
					K8sPodUID:                    ResourceAttributeConfig{Enabled: true},
					K8sVolumeName:                ResourceAttributeConfig{Enabled: true},
					K8sVolumeType:                ResourceAttributeConfig{Enabled: true},
					Partition:                    ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ContainerCPUTime:                     MetricConfig{Enabled: false},
					ContainerCPUUsage:                    MetricConfig{Enabled: false},
					ContainerCPUUtilization:              MetricConfig{Enabled: false},
					ContainerFilesystemAvailable:         MetricConfig{Enabled: false},
					ContainerFilesystemCapacity:          MetricConfig{Enabled: false},
					ContainerFilesystemUsage:             MetricConfig{Enabled: false},
					ContainerMemoryAvailable:             MetricConfig{Enabled: false},
					ContainerMemoryMajorPageFaults:       MetricConfig{Enabled: false},
					ContainerMemoryPageFaults:            MetricConfig{Enabled: false},
					ContainerMemoryRss:                   MetricConfig{Enabled: false},
					ContainerMemoryUsage:                 MetricConfig{Enabled: false},
					ContainerMemoryWorkingSet:            MetricConfig{Enabled: false},
					ContainerUptime:                      MetricConfig{Enabled: false},
					K8sContainerCPUNodeUtilization:       MetricConfig{Enabled: false},
					K8sContainerCPULimitUtilization:      MetricConfig{Enabled: false},
					K8sContainerCPURequestUtilization:    MetricConfig{Enabled: false},
					K8sContainerMemoryNodeUtilization:    MetricConfig{Enabled: false},
					K8sContainerMemoryLimitUtilization:   MetricConfig{Enabled: false},
					K8sContainerMemoryRequestUtilization: MetricConfig{Enabled: false},
					K8sNodeCPUTime:                       MetricConfig{Enabled: false},
					K8sNodeCPUUsage:                      MetricConfig{Enabled: false},
					K8sNodeCPUUtilization:                MetricConfig{Enabled: false},
					K8sNodeFilesystemAvailable:           MetricConfig{Enabled: false},
					K8sNodeFilesystemCapacity:            MetricConfig{Enabled: false},
					K8sNodeFilesystemUsage:               MetricConfig{Enabled: false},
					K8sNodeMemoryAvailable:               MetricConfig{Enabled: false},
					K8sNodeMemoryMajorPageFaults:         MetricConfig{Enabled: false},
					K8sNodeMemoryPageFaults:              MetricConfig{Enabled: false},
					K8sNodeMemoryRss:                     MetricConfig{Enabled: false},
					K8sNodeMemoryUsage:                   MetricConfig{Enabled: false},
					K8sNodeMemoryWorkingSet:              MetricConfig{Enabled: false},
					K8sNodeNetworkErrors:                 MetricConfig{Enabled: false},
					K8sNodeNetworkIo:                     MetricConfig{Enabled: false},
					K8sNodeUptime:                        MetricConfig{Enabled: false},
					K8sPodCPUNodeUtilization:             MetricConfig{Enabled: false},
					K8sPodCPUTime:                        MetricConfig{Enabled: false},
					K8sPodCPUUsage:                       MetricConfig{Enabled: false},
					K8sPodCPUUtilization:                 MetricConfig{Enabled: false},
					K8sPodCPULimitUtilization:            MetricConfig{Enabled: false},
					K8sPodCPURequestUtilization:          MetricConfig{Enabled: false},
					K8sPodFilesystemAvailable:            MetricConfig{Enabled: false},
					K8sPodFilesystemCapacity:             MetricConfig{Enabled: false},
					K8sPodFilesystemUsage:                MetricConfig{Enabled: false},
					K8sPodMemoryAvailable:                MetricConfig{Enabled: false},
					K8sPodMemoryMajorPageFaults:          MetricConfig{Enabled: false},
					K8sPodMemoryNodeUtilization:          MetricConfig{Enabled: false},
					K8sPodMemoryPageFaults:               MetricConfig{Enabled: false},
					K8sPodMemoryRss:                      MetricConfig{Enabled: false},
					K8sPodMemoryUsage:                    MetricConfig{Enabled: false},
					K8sPodMemoryWorkingSet:               MetricConfig{Enabled: false},
					K8sPodMemoryLimitUtilization:         MetricConfig{Enabled: false},
					K8sPodMemoryRequestUtilization:       MetricConfig{Enabled: false},
					K8sPodNetworkErrors:                  MetricConfig{Enabled: false},
					K8sPodNetworkIo:                      MetricConfig{Enabled: false},
					K8sPodUptime:                         MetricConfig{Enabled: false},
					K8sVolumeAvailable:                   MetricConfig{Enabled: false},
					K8sVolumeCapacity:                    MetricConfig{Enabled: false},
					K8sVolumeInodes:                      MetricConfig{Enabled: false},
					K8sVolumeInodesFree:                  MetricConfig{Enabled: false},
					K8sVolumeInodesUsed:                  MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					AwsVolumeID:                  ResourceAttributeConfig{Enabled: false},
					ContainerID:                  ResourceAttributeConfig{Enabled: false},
					FsType:                       ResourceAttributeConfig{Enabled: false},
					GcePdName:                    ResourceAttributeConfig{Enabled: false},
					GlusterfsEndpointsName:       ResourceAttributeConfig{Enabled: false},
					GlusterfsPath:                ResourceAttributeConfig{Enabled: false},
					K8sContainerName:             ResourceAttributeConfig{Enabled: false},
					K8sNamespaceName:             ResourceAttributeConfig{Enabled: false},
					K8sNodeName:                  ResourceAttributeConfig{Enabled: false},
					K8sPersistentvolumeclaimName: ResourceAttributeConfig{Enabled: false},
					K8sPodName:                   ResourceAttributeConfig{Enabled: false},
					K8sPodUID:                    ResourceAttributeConfig{Enabled: false},
					K8sVolumeName:                ResourceAttributeConfig{Enabled: false},
					K8sVolumeType:                ResourceAttributeConfig{Enabled: false},
					Partition:                    ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				AwsVolumeID:                  ResourceAttributeConfig{Enabled: true},
				ContainerID:                  ResourceAttributeConfig{Enabled: true},
				FsType:                       ResourceAttributeConfig{Enabled: true},
				GcePdName:                    ResourceAttributeConfig{Enabled: true},
				GlusterfsEndpointsName:       ResourceAttributeConfig{Enabled: true},
				GlusterfsPath:                ResourceAttributeConfig{Enabled: true},
				K8sContainerName:             ResourceAttributeConfig{Enabled: true},
				K8sNamespaceName:             ResourceAttributeConfig{Enabled: true},
				K8sNodeName:                  ResourceAttributeConfig{Enabled: true},
				K8sPersistentvolumeclaimName: ResourceAttributeConfig{Enabled: true},
				K8sPodName:                   ResourceAttributeConfig{Enabled: true},
				K8sPodUID:                    ResourceAttributeConfig{Enabled: true},
				K8sVolumeName:                ResourceAttributeConfig{Enabled: true},
				K8sVolumeType:                ResourceAttributeConfig{Enabled: true},
				Partition:                    ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				AwsVolumeID:                  ResourceAttributeConfig{Enabled: false},
				ContainerID:                  ResourceAttributeConfig{Enabled: false},
				FsType:                       ResourceAttributeConfig{Enabled: false},
				GcePdName:                    ResourceAttributeConfig{Enabled: false},
				GlusterfsEndpointsName:       ResourceAttributeConfig{Enabled: false},
				GlusterfsPath:                ResourceAttributeConfig{Enabled: false},
				K8sContainerName:             ResourceAttributeConfig{Enabled: false},
				K8sNamespaceName:             ResourceAttributeConfig{Enabled: false},
				K8sNodeName:                  ResourceAttributeConfig{Enabled: false},
				K8sPersistentvolumeclaimName: ResourceAttributeConfig{Enabled: false},
				K8sPodName:                   ResourceAttributeConfig{Enabled: false},
				K8sPodUID:                    ResourceAttributeConfig{Enabled: false},
				K8sVolumeName:                ResourceAttributeConfig{Enabled: false},
				K8sVolumeType:                ResourceAttributeConfig{Enabled: false},
				Partition:                    ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
