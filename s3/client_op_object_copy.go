package s3

import (
	"context"
	"strings"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CopyObject executes `CopyObject` operation.
func (svc *S3) CopyObject(ctx context.Context, r CopyObjectRequest) (*CopyObjectResult, error) {
	out, err := svc.RawCopyObject(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CopyObject",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCopyObjectResult(out), nil
}

// CopyObjectRequest has parameters for `CopyObject` operation.
type CopyObjectRequest struct {
	Bucket     string
	Key        string
	CopySource string

	// optional
	ACL                            ObjectCannedACL
	CacheControl                   string
	ContentDisposition             string
	ContentEncoding                string
	ContentLanguage                string
	ContentType                    string
	CopySourceIfMatch              string
	CopySourceIfModifiedSince      time.Time
	CopySourceIfNoneMatch          string
	CopySourceIfUnmodifiedSince    time.Time
	CopySourceSSECustomerAlgorithm string
	CopySourceSSECustomerKey       string
	CopySourceSSECustomerKeyMD5    string
	Expires                        time.Time
	GrantFullControl               string
	GrantRead                      string
	GrantReadACP                   string
	GrantWriteACP                  string
	Metadata                       map[string]string
	MetadataDirective              MetadataDirective
	ObjectLockLegalHoldStatus      ObjectLockLegalHoldStatus
	ObjectLockMode                 ObjectLockMode
	ObjectLockRetainUntilDate      time.Time
	RequestPayer                   RequestPayer
	SSECustomerAlgorithm           string
	SSECustomerKey                 string
	SSECustomerKeyMD5              string
	SSEKMSEncryptionContext        string
	SSEKMSKeyID                    string
	ServerSideEncryption           ServerSideEncryption
	StorageClass                   StorageClass
	Tagging                        string
	TaggingDirective               TaggingDirective
	WebsiteRedirectLocation        string

	// TaggingTagSet is used for tagging. this won't override `Tagging` and can be used together.
	TaggingTagSet []Tag
}

func (r CopyObjectRequest) ToInput() *SDK.CopyObjectInput {
	in := &SDK.CopyObjectInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}
	if r.CopySource != "" {
		in.CopySource = pointers.String(r.CopySource)
	}

	in.ACL = SDK.ObjectCannedACL(r.ACL)

	if r.CacheControl != "" {
		in.CacheControl = pointers.String(r.CacheControl)
	}
	if r.ContentDisposition != "" {
		in.ContentDisposition = pointers.String(r.ContentDisposition)
	}
	if r.ContentEncoding != "" {
		in.ContentEncoding = pointers.String(r.ContentEncoding)
	}
	if r.ContentLanguage != "" {
		in.ContentLanguage = pointers.String(r.ContentLanguage)
	}
	if r.ContentType != "" {
		in.ContentType = pointers.String(r.ContentType)
	}
	if r.CopySourceIfMatch != "" {
		in.CopySourceIfMatch = pointers.String(r.CopySourceIfMatch)
	}
	if !r.CopySourceIfModifiedSince.IsZero() {
		in.CopySourceIfModifiedSince = &r.CopySourceIfModifiedSince
	}
	if r.CopySourceIfNoneMatch != "" {
		in.CopySourceIfNoneMatch = pointers.String(r.CopySourceIfNoneMatch)
	}
	if !r.CopySourceIfUnmodifiedSince.IsZero() {
		in.CopySourceIfUnmodifiedSince = &r.CopySourceIfUnmodifiedSince
	}
	if r.CopySourceSSECustomerAlgorithm != "" {
		in.CopySourceSSECustomerAlgorithm = pointers.String(r.CopySourceSSECustomerAlgorithm)
	}
	if r.CopySourceSSECustomerKey != "" {
		in.CopySourceSSECustomerKey = pointers.String(r.CopySourceSSECustomerKey)
	}
	if r.CopySourceSSECustomerKeyMD5 != "" {
		in.CopySourceSSECustomerKeyMD5 = pointers.String(r.CopySourceSSECustomerKeyMD5)
	}
	if !r.Expires.IsZero() {
		in.Expires = &r.Expires
	}
	if r.GrantFullControl != "" {
		in.GrantFullControl = pointers.String(r.GrantFullControl)
	}
	if r.GrantRead != "" {
		in.GrantRead = pointers.String(r.GrantRead)
	}
	if r.GrantReadACP != "" {
		in.GrantReadACP = pointers.String(r.GrantReadACP)
	}
	if r.GrantWriteACP != "" {
		in.GrantWriteACP = pointers.String(r.GrantWriteACP)
	}

	in.Metadata = r.Metadata
	in.MetadataDirective = SDK.MetadataDirective(r.MetadataDirective)
	in.ObjectLockLegalHoldStatus = SDK.ObjectLockLegalHoldStatus(r.ObjectLockLegalHoldStatus)
	in.ObjectLockMode = SDK.ObjectLockMode(r.ObjectLockMode)

	if !r.ObjectLockRetainUntilDate.IsZero() {
		in.ObjectLockRetainUntilDate = &r.ObjectLockRetainUntilDate
	}

	in.RequestPayer = SDK.RequestPayer(r.RequestPayer)

	if r.SSECustomerAlgorithm != "" {
		in.SSECustomerAlgorithm = pointers.String(r.SSECustomerAlgorithm)
	}
	if r.SSECustomerKey != "" {
		in.SSECustomerKey = pointers.String(r.SSECustomerKey)
	}
	if r.SSECustomerKeyMD5 != "" {
		in.SSECustomerKeyMD5 = pointers.String(r.SSECustomerKeyMD5)
	}
	if r.SSEKMSEncryptionContext != "" {
		in.SSEKMSEncryptionContext = pointers.String(r.SSEKMSEncryptionContext)
	}
	if r.SSEKMSKeyID != "" {
		in.SSEKMSKeyId = pointers.String(r.SSEKMSKeyID)
	}

	in.ServerSideEncryption = SDK.ServerSideEncryption(r.ServerSideEncryption)
	in.StorageClass = SDK.StorageClass(r.StorageClass)

	tags := r.Tagging
	if len(r.TaggingTagSet) != 0 {
		tt := make([]string, 0, len(r.TaggingTagSet))
		if tags != "" {
			tt = append(tt, tags)
		}
		for _, v := range r.TaggingTagSet {
			tt = append(tt, v.Key+"="+v.Value)
		}
		tags = strings.Join(tt, "&")
	}
	if tags != "" {
		in.Tagging = pointers.String(tags)
	}

	in.TaggingDirective = SDK.TaggingDirective(r.TaggingDirective)

	if r.WebsiteRedirectLocation != "" {
		in.WebsiteRedirectLocation = pointers.String(r.WebsiteRedirectLocation)
	}
	return in
}

// CopyObjectResult contains results from `CopyObject` operation.
type CopyObjectResult struct {
	CopySourceVersionID     string
	Expiration              string
	RequestCharged          RequestCharged
	SSECustomerAlgorithm    string
	SSECustomerKeyMD5       string
	SSEKMSEncryptionContext string
	SSEKMSKeyID             string
	ServerSideEncryption    ServerSideEncryption
	VersionID               string

	CopyResultETag         string
	CopyResultLastModified time.Time
}

func NewCopyObjectResult(output *SDK.CopyObjectResponse) *CopyObjectResult {
	r := &CopyObjectResult{}
	if output == nil {
		return r
	}

	if output.CopySourceVersionId != nil {
		r.CopySourceVersionID = *output.CopySourceVersionId
	}
	if output.Expiration != nil {
		r.Expiration = *output.Expiration
	}
	if output.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	if output.SSECustomerAlgorithm != nil {
		r.SSECustomerAlgorithm = *output.SSECustomerAlgorithm
	}
	if output.SSECustomerKeyMD5 != nil {
		r.SSECustomerKeyMD5 = *output.SSECustomerKeyMD5
	}
	if output.SSEKMSEncryptionContext != nil {
		r.SSEKMSEncryptionContext = *output.SSEKMSEncryptionContext
	}
	if output.SSEKMSKeyId != nil {
		r.SSEKMSKeyID = *output.SSEKMSKeyId
	}
	if output.ServerSideEncryption != "" {
		r.ServerSideEncryption = ServerSideEncryption(output.ServerSideEncryption)
	}
	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}

	if output.CopyObjectResult != nil {
		v := output.CopyObjectResult
		if v.ETag != nil {
			r.CopyResultETag = *v.ETag
		}
		if v.LastModified != nil {
			r.CopyResultLastModified = *v.LastModified
		}
	}
	return r
}
