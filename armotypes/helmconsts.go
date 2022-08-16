package armotypes

const (
	// In-cluster namespaces
	ArmoSystemNamespace = "armo-system" // deprecated, kept for backward compatibility
	KubescapeNamespace  = "kubescape"

	ArmoKollectorContainerName = "armo-collector"

	// registry scan
	LowestHelmVersionSupportedRegistryScan = "v1.7.14"
	RegistryInfoArgKey                     = "registryInfo-v1"
	RegistryScanSecretName                 = "kubescape-registry-scan"

	// vulnerability scan
	LowestHelmVersionSupportedVulnerabilityScan = "v1.7.17"

	// cronjob template annotation and labels
	CronJobTemplateAnnotationArmoJobIDKeyDeprecated      = "armo.jobid"       // deprecated
	CronJobTemplateAnnotationArmoCloudJobIDKeyDeprecated = "armo.cloud/jobid" // deprecated
	CronJobTemplateAnnotationJobIDKey                    = "app.kubescape/job-id"

	CronJobTemplateAnnotationUpdateJobIDDeprecated = "armo.updatejobid" // deprecated
	CronJobTemplateAnnotationUpdateJobID           = "app.kubescape/update-job-id"

	CronJobTemplateAnnotationNamespaceKeyDeprecated = "armo.namespace" // deprecated
	CronJobTemplateAnnotationNamespaceKey           = "app.kubescape/namespace"

	CronJobTemplateAnnotationRegistryNameKey = "armo.cloud/registryname"
	CronJobTemplateAnnotationHostScannerKey  = "armo.host-scanner"
	CronJobTemplateAnnotationFrameworkKey    = "armo.framework"

	CronJobTemplateLabelKey               = "armo.tier"
	CronJobTemplateLabelValueKubescape    = "kubescape-scan"
	CronJobTemplateLabelValueVulnScan     = "vuln-scan"
	CronJobTemplateLabelValueRegistryScan = "registry-scan"
)

func GetInClusterSupportedNamespaces() []string {
	return []string{KubescapeNamespace, ArmoSystemNamespace}
}
