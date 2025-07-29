//go:build js

package web_transport

import (
	"encoding/hex"
	"io"
	"net/http"
	"testing"

	"github.com/sourcenetwork/goji"
	"github.com/sourcenetwork/goji/streams"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebTransportBidirectionalStream(t *testing.T) {
	certHash, err := fetchServerCertificateHash()
	require.NoError(t, err, "failed to fetch server certificate hash")

	certHashValue := CertificateHash(CertificateHashAlgorithmSHA256, certHash)
	webTransport := WebTransport.New("https://127.0.0.1:4443", WebTransportOptions.WithServerCertificateHashes(certHashValue))

	_, err = goji.Await(webTransport.Ready())
	require.NoError(t, err, "web transport failed to connect")

	streamRes, err := goji.Await(webTransport.CreateBidirectionalStream())
	require.NoError(t, err, "failed to create bidirectional stream")

	stream := WebTransportBidirectionalStreamValue(streamRes[0])
	writer := streams.NewWriter(stream.Writable().GetWriter())
	reader := streams.NewReader(stream.Readable().GetBYOBReader())

	writeBytes := []byte("hello")
	readBytes := make([]byte, len(writeBytes))

	_, err = writer.Write(writeBytes)
	require.NoError(t, err, "failed to write to stream")

	_, err = io.ReadFull(reader, readBytes)
	require.NoError(t, err, "failed to read from stream")

	assert.Equal(t, writeBytes, readBytes)

	err = writer.Close()
	assert.NoError(t, err)

	err = reader.Close()
	assert.NoError(t, err)
}

// fetchServerCertificateHash fetches the certificate hash from
// the local webtransport-test server.
func fetchServerCertificateHash() ([]byte, error) {
	req, err := http.NewRequest("GET", "http://localhost:8000", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("js.fetch:mode", "cors")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(string(bytes))
}
