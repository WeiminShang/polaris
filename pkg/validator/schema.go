package validator

import (
	"bytes"
	"fmt"
	"io"

	packr "github.com/gobuffalo/packr/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/fairwindsops/polaris/pkg/config"
)

var (
	schemaBox = (*packr.Box)(nil)
	checks    = map[config.TargetKind][]config.SchemaCheck{
		config.TargetContainer: []config.SchemaCheck{},
		config.TargetPod:       []config.SchemaCheck{},
	}
	// We explicitly set the order to avoid thrash in the
	// tests as we migrate toward JSON schema
	checkOrder = []string{
		// Pod checks
		"hostIPC",
		"hostPID",
		"hostNetwork",
		// Container checks
		"memoryLimitsMissing",
		"memoryRequestsMissing",
		"cpuLimitsMissing",
		"cpuRequestsMissing",
		"readinessProbe",
		"livenessProbe",
		"pullPolicyNotAlways",
		"tagNotSpecified",
		"hostPortSet",
		"runAsRootAllowed",
		"runAsPrivileged",
		"notReadOnlyRootFileSystem",
		"privilegeEscalationAllowed",
		"dangerousCapabilities",
		"insecureCapabilities",
	}
)

func init() {
	schemaBox = packr.New("Schemas", "../../checks")
	for _, file := range checkOrder {
		contents, err := schemaBox.Find(file + ".yaml")
		if err != nil {
			panic(err)
		}
		check, err := parseCheck(contents)
		if err != nil {
			panic(err)
		}
		checks[check.Target] = append(checks[check.Target], check)
	}
}

func parseCheck(rawBytes []byte) (config.SchemaCheck, error) {
	reader := bytes.NewReader(rawBytes)
	check := config.SchemaCheck{}
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)
	for {
		if err := d.Decode(&check); err != nil {
			if err == io.EOF {
				return check, nil
			}
			return check, fmt.Errorf("Decoding schema check failed: %v", err)
		}
	}
}

func applyPodSchemaChecks(conf *config.Configuration, pod *corev1.PodSpec, controllerName string, controllerType config.SupportedController, pv *PodValidation) error {
	for _, check := range checks[config.TargetPod] {
		if !conf.IsActionable(check.Category, check.Name, controllerName) {
			continue
		}
		if !check.IsActionable(config.TargetPod, controllerType, false) {
			continue
		}
		severity := conf.GetSeverity(check.Category, check.Name)
		passes, err := check.CheckPod(pod)
		if err != nil {
			return err
		}
		if passes {
			pv.addSuccess(check.SuccessMessage, check.Category, check.ID)
		} else {
			pv.addFailure(check.FailureMessage, severity, check.Category, check.ID)
		}
	}
	return nil
}

func applyContainerSchemaChecks(conf *config.Configuration, controllerName string, controllerType config.SupportedController, cv *ContainerValidation) error {
	for _, check := range checks[config.TargetContainer] {
		if !conf.IsActionable(check.Category, check.Name, controllerName) {
			continue
		}
		if !check.IsActionable(config.TargetContainer, controllerType, cv.IsInitContainer) {
			continue
		}
		severity := conf.GetSeverity(check.Category, check.Name)
		var passes bool
		var err error
		if check.SchemaTarget == config.TargetPod {
			cv.parentPodSpec.Containers = []corev1.Container{*cv.Container}
			passes, err = check.CheckPod(&cv.parentPodSpec)
			cv.parentPodSpec.Containers = []corev1.Container{}
		} else {
			passes, err = check.CheckContainer(cv.Container)
		}
		if err != nil {
			return err
		}
		if passes {
			cv.addSuccess(check.SuccessMessage, check.Category, check.ID)
		} else {
			cv.addFailure(check.FailureMessage, severity, check.Category, check.ID)
		}
	}
	for checkName, severity := range conf.Checks {
		check, ok := conf.CustomChecks[checkName]
		if !ok {
			return fmt.Errorf("Custom check %s not found", checkName)
		}
		// FIXME: check actionability here
		/*
			if !conf.IsActionable(check.Category, check.Name, controllerName) {
				continue
			}
		*/
		if !check.IsActionable(config.TargetContainer, controllerType, cv.IsInitContainer) {
			continue
		}
		var passes bool
		var err error
		if check.SchemaTarget == config.TargetPod {
			cv.parentPodSpec.Containers = []corev1.Container{*cv.Container}
			passes, err = check.CheckPod(&cv.parentPodSpec)
			cv.parentPodSpec.Containers = []corev1.Container{}
		} else {
			passes, err = check.CheckContainer(cv.Container)
		}
		if err != nil {
			return err
		}
		if passes {
			cv.addSuccess(check.SuccessMessage, check.Category, check.ID)
		} else {
			cv.addFailure(check.FailureMessage, severity, check.Category, check.ID)
		}
	}
	return nil
}
