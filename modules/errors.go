package modules

const (
	APIErrorCodeUnknown = iota
	APIErrorCodeUnauthorized
	APIErrorCodeDigestInvalid
	APIErrorCodeSizeInvalid
	APIErrorCodeNameInvalid
	APIErrorCodeTagInvalid
	APIErrorCodeNameUnknown
	APIErrorCodeManifestUnknown
	APIErrorCodeManifestInvalid
	APIErrorCodeManifestUnverified
	APIErrorCodeBlobUnknown
	APIErrorCodeBlobUploadUnknown
)

type ErrorDescriptor struct {
	Code        int
	Value       string
	Message     string
	Description string
}

var ErrorDescriptors = []ErrorDescriptor{
	{
		Code:        APIErrorCodeUnknown,
		Value:       "UNKNOWN",
		Message:     "unknown error",
		Description: `Generic error returned when the error does not have an API classification.`,
	},
	{
		Code:        APIErrorCodeUnauthorized,
		Value:       "UNAUTHORIZED",
		Message:     "access to the requested resource is not authorized",
		Description: `The access controller denied access for the operation on a resource. Often this will be accompanied by a 401 Unauthorized response status.`,
	},
	{
		Code:        APIErrorCodeDigestInvalid,
		Value:       "DIGEST_INVALID",
		Message:     "provided digest did not match uploaded content",
		Description: `When a blob is uploaded, the registry will check that the content matches the digest provided by the client. The error may include a detail structure with the key "digest", including the invalid digest string. This error may also be returned when a manifest includes an invalid layer digest.`,
	},
	{
		Code:        APIErrorCodeSizeInvalid,
		Value:       "SIZE_INVALID",
		Message:     "provided length did not match content length",
		Description: `When a layer is uploaded, the provided size will be checked against the uploaded content. If they do not match, this error will be returned.`,
	},
	{
		Code:        APIErrorCodeNameInvalid,
		Value:       "NAME_INVALID",
		Message:     "manifest name did not match URI",
		Description: `During a manifest upload, if the name in the manifest does not match the uri name, this error will be returned.`,
	},
	{
		Code:        APIErrorCodeTagInvalid,
		Value:       "TAG_INVALID",
		Message:     "manifest tag did not match URI",
		Description: `During a manifest upload, if the tag in the manifest does not match the uri tag, this error will be returned.`,
	},
	{
		Code:        APIErrorCodeNameUnknown,
		Value:       "NAME_UNKNOWN",
		Message:     "repository name not known to registry",
		Description: `This is returned if the name used during an operation is unknown to the registry.`,
	},
	{
		Code:        APIErrorCodeManifestUnknown,
		Value:       "MANIFEST_UNKNOWN",
		Message:     "manifest unknown",
		Description: `This error is returned when the manifest, identified by name and tag is unknown to the repository.`,
	},
	{
		Code:        APIErrorCodeManifestInvalid,
		Value:       "MANIFEST_INVALID",
		Message:     "manifest invalid",
		Description: `During upload, manifests undergo several checks ensuring validity. If those checks fail, this error may be returned, unless a more specific error is included. The detail will contain information the failed validation.`,
	},
	{
		Code:        APIErrorCodeManifestUnverified,
		Value:       "MANIFEST_UNVERIFIED",
		Message:     "manifest failed signature verification",
		Description: `During manifest upload, if the manifest fails signature verification, this error will be returned.`,
	},
	{
		Code:        APIErrorCodeBlobUnknown,
		Value:       "BLOB_UNKNOWN",
		Message:     "blob unknown to registry",
		Description: `This error may be returned when a blob is unknown to the registry in a specified repository. This can be returned with a standard get or if a manifest references an unknown layer during upload.`,
	},
	{
		Code:        APIErrorCodeBlobUploadUnknown,
		Value:       "BLOB_UPLOAD_UNKNOWN",
		Message:     "blob upload unknown to registry",
		Description: `If a blob upload has been cancelled or was never started, this error code may be returned.`,
	},
}