package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var flatedData = []byte{93, 82, 91, 142, 131, 48, 12, 188, 75, 190, 17, 114,
	156, 23, 225, 50, 85, 10, 9, 27, 45, 37, 40, 132, 237, 174, 170, 222, 125, 29,
	218, 125, 126, 128, 98, 103, 60, 99, 143, 115, 99, 91, 113, 37, 110, 37, 14,
	27, 235, 111, 108, 73, 163, 63, 197, 145, 245, 76, 75, 173, 207, 98, 244, 65,
	104, 193, 26, 54, 204, 209, 47, 229, 192, 92, 99, 136, 172, 135, 134, 149, 84,
	220, 76, 167, 123, 195, 114, 74, 37, 108, 167, 125, 115, 147, 167, 76, 11,
	162, 51, 18, 181, 5, 80, 157, 236, 26, 54, 185, 226, 175, 238, 131, 120, 1,
	251, 193, 247, 62, 244, 131, 235, 131, 239, 209, 17, 121, 201, 46, 132, 56,
	84, 242, 242, 94, 255, 171, 27, 94, 125, 85, 67, 32, 14, 99, 26, 54, 230, 180,
	174, 158, 26, 227, 160, 177, 97, 231, 143, 226, 235, 181, 232, 120, 7, 168,
	68, 109, 225, 168, 252, 186, 224, 86, 26, 174, 80, 170, 230, 23, 153, 177, 74,
	162, 208, 4, 14, 41, 95, 93, 38, 186, 239, 10, 97, 17, 140, 37, 165, 111, 184,
	64, 129, 21, 123, 153, 46, 229, 244, 232, 235, 137, 53, 202, 106, 73, 159,
	146, 191, 217, 37, 8, 206, 53, 255, 170, 248, 211, 143, 18, 202, 116, 146, 26,
	32, 219, 126, 4, 12, 8, 129, 202, 220, 107, 137, 191, 164, 76, 254, 220, 216,
	224, 134, 151, 58, 168, 177, 150, 6, 125, 122, 204, 81, 43, 36, 31, 207, 123,
	8, 62, 87, 53, 146, 163, 49, 178, 39, 187, 59, 80, 88, 151, 176, 175, 37, 94,
	40, 214, 150, 230, 238, 90, 160, 97, 226, 56, 251, 103, 82, 1, 128, 105, 165,
	174, 90, 219, 203, 233, 109, 93, 170, 218, 148, 211, 190, 30, 107, 29, 19, 224,
	225, 189, 63, 4, 30, 9, 14, 192, 250, 101, 159, 103, 218, 0, 133, 88, 195, 27,
	243, 244, 104, 206, 115, 220, 142, 62, 53, 7, 11, 216, 146, 228, 189, 14, 66,
	48, 254, 159, 134, 31, 52, 255, 235, 80, 107, 108, 65, 226, 179, 232, 224,
	174, 82, 68, 67, 169, 53, 167, 193, 111, 91, 53, 143, 94, 197, 195, 5, 41,
	104, 205, 251, 178, 196, 101, 34, 75, 8, 52, 39, 55, 186, 55, 10, 232, 197,
	145, 141, 159}

func TestDecompressFlate(t *testing.T) {
	//t.Skip("Currently no realistic test data")
	assert := assert.New(t)
	out, err := Deflate(flatedData)
	assert.Nil(err)
	assert.NotNil(out)
	assert.True(len(out) > 1)
}

func TestCompressFlate(t *testing.T) {
	assert := assert.New(t)
	out, err := DeflateCompress([]byte("Test string"))
	assert.Nil(err)
	assert.NotNil(out)
	assert.True(len(out) > 0)

	decompressedBytes, err := Deflate(out)
	assert.Nil(err)
	assert.Equal("Test string", string(decompressedBytes))
}
