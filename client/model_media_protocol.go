/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// MediaProtocol the model 'MediaProtocol'
type MediaProtocol string

// List of MediaProtocol
const (
	MEDIAPROTOCOL_FILE MediaProtocol = "File"
	MEDIAPROTOCOL_HTTP MediaProtocol = "Http"
	MEDIAPROTOCOL_RTMP MediaProtocol = "Rtmp"
	MEDIAPROTOCOL_RTSP MediaProtocol = "Rtsp"
	MEDIAPROTOCOL_UDP MediaProtocol = "Udp"
	MEDIAPROTOCOL_RTP MediaProtocol = "Rtp"
	MEDIAPROTOCOL_FTP MediaProtocol = "Ftp"
)
