package bp2build

import (
	"android/soong/android"
	"fmt"
	"strings"
)

// Simple metrics struct to collect information about a Blueprint to BUILD
// conversion process.
type CodegenMetrics struct {
	// Total number of Soong/Blueprint modules
	TotalModuleCount int

	// Counts of generated Bazel targets per Bazel rule class
	RuleClassCount map[string]int

	// Total number of handcrafted targets
	handCraftedTargetCount int

	moduleWithUnconvertedDepsMsgs []string

	convertedModules []string
}

// Print the codegen metrics to stdout.
func (metrics CodegenMetrics) Print() {
	generatedTargetCount := 0
	for _, ruleClass := range android.SortedStringKeys(metrics.RuleClassCount) {
		count := metrics.RuleClassCount[ruleClass]
		fmt.Printf("[bp2build] %s: %d targets\n", ruleClass, count)
		generatedTargetCount += count
	}
	fmt.Printf(
		"[bp2build] Generated %d total BUILD targets and included %d handcrafted BUILD targets from %d Android.bp modules.\n With %d modules with unconverted deps \n\t%s",
		generatedTargetCount,
		metrics.handCraftedTargetCount,
		metrics.TotalModuleCount,
		len(metrics.moduleWithUnconvertedDepsMsgs),
		strings.Join(metrics.moduleWithUnconvertedDepsMsgs, "\n\t"))
}

func (metrics CodegenMetrics) AddConvertedModule(moduleName string) {
	// Undo prebuilt_ module name prefix modifications
	moduleName = android.RemoveOptionalPrebuiltPrefix(moduleName)
	metrics.convertedModules = append(metrics.convertedModules, moduleName)
}
