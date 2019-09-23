package provenance

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/build"
	"google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/provenance"

	"github.com/Shopify/voucher"
)

var (
	builderIdentityTestData = "trusted-person@email.com"
	imageSHA256TestData     = "sha256:1234c923e00e0fd2ba78041bfb64a105e1ecb7678916d1f7776311e45bf57890"
	imageURLTestData        = "gcr.io/" + projectTestData + "/name@" + imageSHA256TestData
	projectTestData         = "test"
)

var buildDetailsTestData = &build.Details{
	Provenance: &provenance.BuildProvenance{
		Id:        "foo",
		ProjectId: projectTestData,
		Creator:   builderIdentityTestData,
		BuiltArtifacts: []*provenance.Artifact{
			{
				Id:       imageURLTestData,
				Checksum: imageSHA256TestData,
				Names:    []string{"foo", "bar"},
			},
		},
	},
	ProvenanceBytes: "base64blob",
}

func TestArtifactIsImage(t *testing.T) {
	imageDataTestData, err := voucher.NewImageData(imageURLTestData)
	require.NoError(t, err)

	assert := assert.New(t)
	result := validateArtifacts(imageDataTestData, buildDetailsTestData)
	assert.True(result)
}
