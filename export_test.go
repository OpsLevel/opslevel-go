package opslevel

// Workaround for testing unexported functions.
//
// Running `go help build` displays:
// When compiling packages, build ignores files that end in '_test.go'.
var (
	ExtractAliases           = extractAliases
	ExtractTagIdsToDelete    = extractTagIdsToDelete
	ExtractTagInputsToCreate = extractTagInputsToCreate
)
