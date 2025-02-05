// Package validation provides default Validator's that can be run with a list
// of arbitrary objects. The defaults exposed here consist of all Validator's
// implemented by this validation library.
//
// Each default Validator runs an independent set of validation functions on
// a set of objects. To run all implemented Validator's, use AllValidators.
// The Validator will not be run on objects of an inappropriate type.

package validation

import (
	interfaces "github.com/operator-framework/api/pkg/validation/interfaces"
	"github.com/operator-framework/api/pkg/validation/internal"
)

// PackageManifestValidator implements Validator to validate package manifests.
var PackageManifestValidator = internal.PackageManifestValidator

// ClusterServiceVersionValidator implements Validator to validate
// ClusterServiceVersions.
var ClusterServiceVersionValidator = internal.CSVValidator

// CustomResourceDefinitionValidator implements Validator to validate
// CustomResourceDefinitions.
var CustomResourceDefinitionValidator = internal.CRDValidator

// BundleValidator implements Validator to validate Bundles.
var BundleValidator = internal.BundleValidator

// OperatorHubValidator implements Validator to validate bundle objects
// for OperatorHub.io requirements.
var OperatorHubValidator = internal.OperatorHubValidator

// Object Validator validates various custom objects in the bundle like PDBs and SCCs.
// Object validation is optional and not a default-level validation.
var ObjectValidator = internal.ObjectValidator

// OperatorGroupValidator implements Validator to validate OperatorGroup manifests
var OperatorGroupValidator = internal.OperatorGroupValidator

// CommunityOperatorValidator implements Validator to validate bundle objects
// for the Community Operator requirements.
var CommunityOperatorValidator = internal.CommunityOperatorValidator

// AlphaDeprecatedAPIsValidator implements Validator to validate bundle objects
// for API deprecation requirements.
var AlphaDeprecatedAPIsValidator = internal.AlphaDeprecatedAPIsValidator

// AllValidators implements Validator to validate all Operator manifest types.
var AllValidators = interfaces.Validators{
	PackageManifestValidator,
	ClusterServiceVersionValidator,
	CustomResourceDefinitionValidator,
	BundleValidator,
	OperatorHubValidator,
	ObjectValidator,
	OperatorGroupValidator,
	CommunityOperatorValidator,
	AlphaDeprecatedAPIsValidator,
}

var DefaultBundleValidators = interfaces.Validators{
	ClusterServiceVersionValidator,
	CustomResourceDefinitionValidator,
	BundleValidator,
}
