/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ServerConfiguration Represents the server configuration.
type ServerConfiguration struct {
	// Gets or sets the number of days we should retain log files.
	LogFileRetentionDays int32 `json:"LogFileRetentionDays,omitempty"`
	// Gets or sets a value indicating whether this instance is first run.
	IsStartupWizardCompleted bool `json:"IsStartupWizardCompleted,omitempty"`
	// Gets or sets the cache path.
	CachePath *string `json:"CachePath,omitempty"`
	PreviousVersion Version `json:"PreviousVersion,omitempty"`
	// Gets or sets the stringified PreviousVersion to be stored/loaded,  because System.Version itself isn't xml-serializable.
	PreviousVersionStr *string `json:"PreviousVersionStr,omitempty"`
	// Gets or sets a value indicating whether to enable automatic port forwarding.
	EnableUPnP bool `json:"EnableUPnP,omitempty"`
	// Gets or sets a value indicating whether to enable prometheus metrics exporting.
	EnableMetrics bool `json:"EnableMetrics,omitempty"`
	// Gets or sets the public mapped port.
	PublicPort int32 `json:"PublicPort,omitempty"`
	// Gets or sets the public HTTPS port.
	PublicHttpsPort int32 `json:"PublicHttpsPort,omitempty"`
	// Gets or sets the HTTP server port number.
	HttpServerPortNumber int32 `json:"HttpServerPortNumber,omitempty"`
	// Gets or sets the HTTPS server port number.
	HttpsPortNumber int32 `json:"HttpsPortNumber,omitempty"`
	// Gets or sets a value indicating whether to use HTTPS.
	EnableHttps bool `json:"EnableHttps,omitempty"`
	EnableNormalizedItemByNameIds bool `json:"EnableNormalizedItemByNameIds,omitempty"`
	// Gets or sets the filesystem path of an X.509 certificate to use for SSL.
	CertificatePath *string `json:"CertificatePath,omitempty"`
	// Gets or sets the password required to access the X.509 certificate data in the file specified by MediaBrowser.Model.Configuration.ServerConfiguration.CertificatePath.
	CertificatePassword *string `json:"CertificatePassword,omitempty"`
	// Gets or sets a value indicating whether this instance is port authorized.
	IsPortAuthorized bool `json:"IsPortAuthorized,omitempty"`
	// Gets or sets if quick connect is available for use on this server.
	QuickConnectAvailable bool `json:"QuickConnectAvailable,omitempty"`
	EnableRemoteAccess bool `json:"EnableRemoteAccess,omitempty"`
	// Gets or sets a value indicating whether [enable case sensitive item ids].
	EnableCaseSensitiveItemIds bool `json:"EnableCaseSensitiveItemIds,omitempty"`
	DisableLiveTvChannelUserDataName bool `json:"DisableLiveTvChannelUserDataName,omitempty"`
	// Gets or sets the metadata path.
	MetadataPath *string `json:"MetadataPath,omitempty"`
	MetadataNetworkPath *string `json:"MetadataNetworkPath,omitempty"`
	// Gets or sets the preferred metadata language.
	PreferredMetadataLanguage *string `json:"PreferredMetadataLanguage,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode *string `json:"MetadataCountryCode,omitempty"`
	// Characters to be replaced with a ' ' in strings to create a sort name.
	SortReplaceCharacters *[]string `json:"SortReplaceCharacters,omitempty"`
	// Characters to be removed from strings to create a sort name.
	SortRemoveCharacters *[]string `json:"SortRemoveCharacters,omitempty"`
	// Words to be removed from strings to create a sort name.
	SortRemoveWords *[]string `json:"SortRemoveWords,omitempty"`
	// Gets or sets the minimum percentage of an item that must be played in order for playstate to be updated.
	MinResumePct int32 `json:"MinResumePct,omitempty"`
	// Gets or sets the maximum percentage of an item that can be played while still saving playstate. If this percentage is crossed playstate will be reset to the beginning and the item will be marked watched.
	MaxResumePct int32 `json:"MaxResumePct,omitempty"`
	// Gets or sets the minimum duration that an item must have in order to be eligible for playstate updates..
	MinResumeDurationSeconds int32 `json:"MinResumeDurationSeconds,omitempty"`
	// The delay in seconds that we will wait after a file system change to try and discover what has been added/removed  Some delay is necessary with some items because their creation is not atomic.  It involves the creation of several  different directories and files.
	LibraryMonitorDelay int32 `json:"LibraryMonitorDelay,omitempty"`
	// Gets or sets a value indicating whether [enable dashboard response caching].  Allows potential contributors without visual studio to modify production dashboard code and test changes.
	EnableDashboardResponseCaching bool `json:"EnableDashboardResponseCaching,omitempty"`
	ImageSavingConvention ImageSavingConvention `json:"ImageSavingConvention,omitempty"`
	MetadataOptions *[]MetadataOptions `json:"MetadataOptions,omitempty"`
	SkipDeserializationForBasicTypes bool `json:"SkipDeserializationForBasicTypes,omitempty"`
	ServerName *string `json:"ServerName,omitempty"`
	BaseUrl *string `json:"BaseUrl,omitempty"`
	UICulture *string `json:"UICulture,omitempty"`
	SaveMetadataHidden bool `json:"SaveMetadataHidden,omitempty"`
	ContentTypes *[]NameValuePair `json:"ContentTypes,omitempty"`
	RemoteClientBitrateLimit int32 `json:"RemoteClientBitrateLimit,omitempty"`
	EnableFolderView bool `json:"EnableFolderView,omitempty"`
	EnableGroupingIntoCollections bool `json:"EnableGroupingIntoCollections,omitempty"`
	DisplaySpecialsWithinSeasons bool `json:"DisplaySpecialsWithinSeasons,omitempty"`
	LocalNetworkSubnets *[]string `json:"LocalNetworkSubnets,omitempty"`
	LocalNetworkAddresses *[]string `json:"LocalNetworkAddresses,omitempty"`
	CodecsUsed *[]string `json:"CodecsUsed,omitempty"`
	PluginRepositories *[]RepositoryInfo `json:"PluginRepositories,omitempty"`
	IgnoreVirtualInterfaces bool `json:"IgnoreVirtualInterfaces,omitempty"`
	EnableExternalContentInSuggestions bool `json:"EnableExternalContentInSuggestions,omitempty"`
	// Gets or sets a value indicating whether the server should force connections over HTTPS.
	RequireHttps bool `json:"RequireHttps,omitempty"`
	EnableNewOmdbSupport bool `json:"EnableNewOmdbSupport,omitempty"`
	RemoteIPFilter *[]string `json:"RemoteIPFilter,omitempty"`
	IsRemoteIPFilterBlacklist bool `json:"IsRemoteIPFilterBlacklist,omitempty"`
	ImageExtractionTimeoutMs int32 `json:"ImageExtractionTimeoutMs,omitempty"`
	PathSubstitutions *[]PathSubstitution `json:"PathSubstitutions,omitempty"`
	EnableSimpleArtistDetection bool `json:"EnableSimpleArtistDetection,omitempty"`
	UninstalledPlugins *[]string `json:"UninstalledPlugins,omitempty"`
	// Gets or sets a value indicating whether slow server responses should be logged as a warning.
	EnableSlowResponseWarning bool `json:"EnableSlowResponseWarning,omitempty"`
	// Gets or sets the threshold for the slow response time warning in ms.
	SlowResponseThresholdMs int64 `json:"SlowResponseThresholdMs,omitempty"`
	// Gets or sets the cors hosts.
	CorsHosts *[]string `json:"CorsHosts,omitempty"`
	// Gets or sets the known proxies.
	KnownProxies *[]string `json:"KnownProxies,omitempty"`
}