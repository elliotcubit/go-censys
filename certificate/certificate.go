package certificate

import "time"

type Certificate struct {
	AddedAt                     time.Time          `json:"added_at"`
	Labels                      []string           `json:"labels"`
	ModifiedAt                  time.Time          `json:"modified_at"`
	ParseStatus                 string             `json:"parse_status"`
	Ct                          map[string]CtValue `json:"ct"`
	EverSeenInScan              bool               `json:"ever_seen_in_scan"`
	Precert                     bool               `json:"precert"`
	ValidationLevel             string             `json:"validation_level"`
	FingerprintMd5              string             `json:"fingerprint_md5"`
	FingerprintSha1             string             `json:"fingerprint_sha1"`
	FingerprintSha256           string             `json:"fingerprint_sha256"`
	ParentSpkiFingerprintSha256 string             `json:"parent_spki_fingerprint_sha256"`
	SpkiFingerprintSha256       string             `json:"spki_fingerprint_sha256"`
	TbsFingerprintSha256        string             `json:"tbs_fingerprint_sha256"`
	TbsNoCtFingerprintSha256    string             `json:"tbs_no_ct_fingerprint_sha256"`
	Names                       []string           `json:"names"`
	ValidatedAt                 time.Time          `json:"validated_at"`
	Parsed                      Parsed             `json:"parsed"`
	Revoked                     bool               `json:"revoked"`
	Revocation                  Revocation         `json:"revocation"`
	Validation                  Validation         `json:"validation"`
	Zlint                       Zlint              `json:"zlint"`
}

type Zlint struct {
	ErrorsPresent   bool      `json:"errors_present"`
	FailedLints     []string  `json:"failed_lints"`
	FatalsPresent   bool      `json:"fatals_present"`
	NoticesPresent  bool      `json:"notices_present"`
	Timestamp       time.Time `json:"timestamp"`
	Version         int64     `json:"version"`
	WarningsPresent bool      `json:"warnings_present"`
}

type Validation struct {
	Apple     Validation_Result `json:"apple"`
	Chrome    Validation_Result `json:"chrome"`
	Microsoft Validation_Result `json:"microsoft"`
	Nss       Validation_Result `json:"nss"`
}

type Validation_Result struct {
	Chains          []Validation_Result_Chain `json:"chains"`
	EverValid       bool                      `json:"ever_valid"`
	HadTrustedPath  bool                      `json:"had_trusted_path"`
	HasTrustedPath  bool                      `json:"has_trusted_path"`
	InRevocationSet bool                      `json:"in_revocation_set"`
	IsValid         bool                      `json:"is_valid"`
	Parents         []string                  `json:"parents"`
}

type Validation_Result_Chain struct {
	Sha256fp []string `json:"sha256fp"`
}

type Revocation struct {
	Crl  Revocation_Result `json:"crl"`
	Ocsp Revocation_Result `json:"ocsp"`
}

type Revocation_Result struct {
	NextUpdate     time.Time `json:"next_update"`
	Reason         string    `json:"reason"`
	RevocationTime time.Time `json:"revocation_time"`
	Revoked        bool      `json:"revoked"`
}

type Parsed struct {
	Redacted          bool               `json:"redacted"`
	Version           int                `json:"version"`
	IssuerDn          string             `json:"issuer_dn"`
	SerialNumber      string             `json:"serial_number"`
	SerialNumberHex   string             `json:"serial_number_hex"`
	SubjectDn         string             `json:"subject_dn"`
	Extensions        Extensions         `json:"extensions"`
	Issuer            DistinguishedName  `json:"issuer"`
	Signature         Signature          `json:"signature"`
	Subject           DistinguishedName  `json:"subject"`
	SubjectKeyInfo    SubjectKeyInfo     `json:"subject_key_info"`
	UnknownExtensions []UnknownExtension `json:"unknown_extensions"`
	ValidityPeriod    ValidityPeriod     `json:"validity_period"`
}

type ValidityPeriod struct {
	LengthSeconds int64     `json:"length_seconds"`
	NotAfter      time.Time `json:"not_after"`
	NotBefore     time.Time `json:"not_before"`
}

type UnknownExtension struct {
	Critical bool   `json:"critical"`
	Id       string `json:"id"`
	Value    string `json:"value"`
}

type SubjectKeyInfo struct {
	Dsa               DsaKey           `json:"dsa"`
	Ecdsa             EcdsaKey         `json:"ecdsa"`
	FingerprintSha256 string           `json:"fingerprint_sha256"`
	KeyAlgorithm      ObjectIdentifier `json:"key_algorithm"`
	Rsa               RsaKey           `json:"rsa"`
	Unrecognized      UnrecognizedKey  `json:"unrecognized"`
}

type UnrecognizedKey struct {
	Raw string `json:"raw"`
}

type RsaKey struct {
	Exponent int64  `json:"exponent"`
	Length   int64  `json:"length"`
	Modulus  string `json:"modulus"`
}

type EcdsaKey struct {
	B      string `json:"b"`
	Curve  string `json:"curve"`
	Gx     string `json:"gx"`
	Gy     string `json:"gy"`
	Length int64  `json:"length"`
	N      string `json:"n"`
	P      string `json:"p"`
	Pub    string `json:"pub"`
	X      string `json:"x"`
	Y      string `json:"y"`
}

type DsaKey struct {
	G string `json:"g"`
	P string `json:"p"`
	Q string `json:"q"`
	Y string `json:"y"`
}

type DistinguishedName struct {
	CommonName           []string `json:"common_name"`
	Country              []string `json:"country"`
	DomainComponent      []string `json:"domain_component"`
	EmailAddress         []string `json:"email_address"`
	GivenName            []string `json:"given_name"`
	JurisdictionCountry  []string `json:"jurisdiction_country"`
	JurisdictionLocality []string `json:"jurisdiction_locality"`
	JurisdictionProvince []string `json:"jurisdiction_province"`
	Locality             []string `json:"locality"`
	Organization         []string `json:"organization"`
	OrganizationId       []string `json:"organization_id"`
	OrganizationalUnit   []string `json:"organizational_unit"`
	PostalCode           []string `json:"postal_code"`
	Province             []string `json:"province"`
	SerialNumber         []string `json:"serial_number"`
	StreetAddress        []string `json:"street_address"`
	Surname              []string `json:"surname"`
}

type Signature struct {
	SelfSigned         bool             `json:"self_signed"`
	SignatureAlgorithm ObjectIdentifier `json:"signature_algorithm"`
	Valid              bool             `json:"valid"`
	Value              string           `json:"value"`
}

type ObjectIdentifier struct {
	Name string `json:"name"`
	Oid  string `json:"oid"`
}

type Extensions struct {
	AuthorityInfoAccess         AuthorityInfoAccess          `json:"authority_info_access"`
	AuthorityKeyId              string                       `json:"authority_key_id"`
	BasicConstraints            BasicConstraints             `json:"basic_constraints"`
	CabfOrganizationId          CabfOrganizationId           `json:"cabf_organization_id"`
	CertificatePolicies         []CertificatePolicy          `json:"certificate_policies"`
	CrlDistributionPoints       []string                     `json:"crl_distribution_points"`
	CtPoison                    bool                         `json:"ct_poison"`
	SignedCertificateTimestamps []SignedCertificateTimestamp `json:"signed_certificate_timestamps"`
	ExtendedKeyUsage            ExtendedKeyUsage             `json:"extended_key_usage"`
	IssuerAltName               AltName                      `json:"issuer_alt_name"`
	KeyUsage                    KeyUsage                     `json:"key_usage"`
	NameConstraints             NameConstraints              `json:"name_constraints"`
	QcStatements                QcStatements                 `json:"qc_statements"`
	SubjectAltName              AltName                      `json:"subject_alt_name"`
	SubjectKeyId                string                       `json:"subject_key_id"`
	TorServiceDescriptors       TorServiceDescriptors        `json:"tor_service_descriptors"`
}

type TorServiceDescriptors struct {
	AlgorithmName string `json:"algorithm_name"`
	Hash          string `json:"hash"`
	HashBits      int    `json:"hash_bits"`
	Onion         string `json:"onion"`
}

type QcStatements struct {
	Ids    []string            `json:"ids"`
	Parsed QcStatements_Parsed `json:"parsed"`
}

type QcStatements_Parsed struct {
	EtsiCompliance  bool                              `json:"etsi_compliance"`
	Legislation     QcStatements_Parsed_Legislation   `json:"legislation"`
	Limit           QcStatements_Parsed_Limit         `json:"limit"`
	PdsLocations    []QcStatements_Parsed_PdsLocation `json:"pds_locations"`
	RetentionPeriod int64                             `json:"retention_period"`
	Sscd            bool                              `json:"sscd"`
	Types           []QcStatements_Parsed_Type        `json:"types"`
}

type QcStatements_Parsed_Type struct {
	Ids []string `json:"ids"`
}

type QcStatements_Parsed_PdsLocation struct {
	Language string `json:"language"`
	Url      string `json:"url"`
}

type QcStatements_Parsed_Limit struct {
	Amount         int64  `json:"amount"`
	Currency       string `json:"currency"`
	CurrencyNumber int64  `json:"currency_number"`
	Exponent       int64  `json:"exponent"`
}

type QcStatements_Parsed_Legislation struct {
	CountryCodes []string `json:"country_codes"`
}

type NameConstraints struct {
	Critical                bool                                    `json:"critical"`
	ExcludedDirectoryNames  []DistinguishedName                     `json:"excluded_directory_names"`
	ExcludedEdiPartyNames   []NameConstraints_ExcludedEdiPartyName  `json:"excluded_edi_party_names"`
	ExcludedEmailAddresses  []string                                `json:"excluded_email_addresses"`
	ExcludedIpAddresses     []NameConstraints_ExcludedIpAddress     `json:"excluded_ip_addresses"`
	ExcludedNames           []string                                `json:"excluded_names"`
	ExcludedRegisteredIds   []string                                `json:"excluded_registered_ids"`
	ExcludedUris            []string                                `json:"excluded_uris"`
	PermittedDirectoryNames []DistinguishedName                     `json:"permitted_directory_names"`
	PermittedEdiPartyNames  []NameConstraints_PermittedEdiPartyName `json:"permitted_edi_party_names"`
	PermittedEmailAddresses []string                                `json:"permitted_email_addresses"`
	PermittedIpAddresses    []NameConstraints_PermittedIpAddress    `json:"permitted_ip_addresses"`
	PermittedNames          []string                                `json:"permitted_names"`
	PermittedRegisteredIds  []string                                `json:"permitted_registered_ids"`
	PermittedUris           []string                                `json:"permitted_uris"`
}

type NameConstraints_PermittedIpAddress struct {
	Begin string `json:"begin"`
	Cidr  string `json:"cidr"`
	End   string `json:"end"`
	Mask  string `json:"mask"`
}

type NameConstraints_PermittedEdiPartyName struct {
	NameAssigner string `json:"name_assigner"`
	PartyName    string `json:"party_name"`
}

type NameConstraints_ExcludedIpAddress struct {
	Begin string `json:"begin"`
	Cidr  string `json:"cidr"`
	End   string `json:"end"`
	Mask  string `json:"mask"`
}

type NameConstraints_ExcludedEdiPartyName struct {
	NameAssigner string `json:"name_assigner"`
	PartyName    string `json:"party_name"`
}

type KeyUsage struct {
	CertificateSign   bool   `json:"certificate_sign"`
	ContentCommitment bool   `json:"content_commitment"`
	CrlSign           bool   `json:"crl_sign"`
	DataEncipherment  bool   `json:"data_encipherment"`
	DecipherOnly      bool   `json:"decipher_only"`
	DigitalSignature  bool   `json:"digital_signature"`
	EncipherOnly      bool   `json:"encipher_only"`
	KeyAgreement      bool   `json:"key_agreement"`
	KeyEncipherment   bool   `json:"key_encipherment"`
	Value             uint64 `json:"value"`
}

type AltName struct {
	DirectoryNames             []DistinguishedName    `json:"directory_names"`
	DnsNames                   []string               `json:"dns_names"`
	EdiPartyNames              []AltName_EdiPartyName `json:"edi_party_names"`
	EmailAddresses             []string               `json:"email_addresses"`
	IpAddresses                []string               `json:"ip_addresses"`
	OtherNames                 []AltName_OtherName    `json:"other_names"`
	RegisteredIds              []string               `json:"registered_ids"`
	UniformResourceIdentifiers []string               `json:"uniform_resource_identifiers"`
}

type AltName_OtherName struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type AltName_EdiPartyName struct {
	NameAssigner string `json:"name_assigner"`
	PartyName    string `json:"party_name"`
}

type ExtendedKeyUsage struct {
	Any                            bool   `json:"any"`
	AppleCodeSigning               bool   `json:"apple_code_signing"`
	AppleCodeSigningDevelopment    bool   `json:"apple_code_signing_development"`
	AppleCodeSigningThirdParty     bool   `json:"apple_code_signing_third_party"`
	AppleCryptoDevelopmentEnv      bool   `json:"apple_crypto_development_env"`
	AppleCryptoEnv                 bool   `json:"apple_crypto_env"`
	AppleCryptoMaintenanceEnv      bool   `json:"apple_crypto_maintenance_env"`
	AppleCryptoProductionEnv       bool   `json:"apple_crypto_production_env"`
	AppleCryptoQos                 bool   `json:"apple_crypto_qos"`
	AppleCryptoTestEnv             bool   `json:"apple_crypto_test_env"`
	AppleCryptoTier0Qos            bool   `json:"apple_crypto_tier0_qos"`
	AppleCryptoTier1Qos            bool   `json:"apple_crypto_tier1_qos"`
	AppleCryptoTier2Qos            bool   `json:"apple_crypto_tier2_qos"`
	AppleCryptoTier3Qos            bool   `json:"apple_crypto_tier3_qos"`
	AppleIchatEncryption           bool   `json:"apple_ichat_encryption"`
	AppleIchatSigning              bool   `json:"apple_ichat_signing"`
	AppleResourceSigning           bool   `json:"apple_resource_signing"`
	AppleSoftwareUpdateSigning     bool   `json:"apple_software_update_signing"`
	AppleSystemIdentity            bool   `json:"apple_system_identity"`
	ClientAuth                     bool   `json:"client_auth"`
	CodeSigning                    bool   `json:"code_signing"`
	Dvcs                           bool   `json:"dvcs"`
	EapOverLan                     bool   `json:"eap_over_lan"`
	EapOverPpp                     bool   `json:"eap_over_ppp"`
	EmailProtection                bool   `json:"email_protection"`
	IpsecEndSystem                 bool   `json:"ipsec_end_system"`
	IpsecIntermediateSystemUsage   bool   `json:"ipsec_intermediate_system_usage"`
	IpsecTunnel                    bool   `json:"ipsec_tunnel"`
	IpsecUser                      bool   `json:"ipsec_user"`
	MicrosoftCaExchange            bool   `json:"microsoft_ca_exchange"`
	MicrosoftCertTrustListSigning  bool   `json:"microsoft_cert_trust_list_signing"`
	MicrosoftCspSignature          bool   `json:"microsoft_csp_signature"`
	MicrosoftDocumentSigning       bool   `json:"microsoft_document_signing"`
	MicrosoftDrm                   bool   `json:"microsoft_drm"`
	MicrosoftDrmIndividualization  bool   `json:"microsoft_drm_individualization"`
	MicrosoftEfsRecovery           bool   `json:"microsoft_efs_recovery"`
	MicrosoftEmbeddedNtCrypto      bool   `json:"microsoft_embedded_nt_crypto"`
	MicrosoftEncryptedFileSystem   bool   `json:"microsoft_encrypted_file_system"`
	MicrosoftEnrollmentAgent       bool   `json:"microsoft_enrollment_agent"`
	MicrosoftKernelModeCodeSigning bool   `json:"microsoft_kernel_mode_code_signing"`
	MicrosoftKeyRecovery21         bool   `json:"microsoft_key_recovery_21"`
	MicrosoftKeyRecovery3          bool   `json:"microsoft_key_recovery_3"`
	MicrosoftLicenseServer         bool   `json:"microsoft_license_server"`
	MicrosoftLicenses              bool   `json:"microsoft_licenses"`
	MicrosoftLifetimeSigning       bool   `json:"microsoft_lifetime_signing"`
	MicrosoftMobileDeviceSoftware  bool   `json:"microsoft_mobile_device_software"`
	MicrosoftNt5Crypto             bool   `json:"microsoft_nt5_crypto"`
	MicrosoftOemWhqlCrypto         bool   `json:"microsoft_oem_whql_crypto"`
	MicrosoftQualifiedSubordinate  bool   `json:"microsoft_qualified_subordinate"`
	MicrosoftRootListSigner        bool   `json:"microsoft_root_list_signer"`
	MicrosoftServerGatedCrypto     bool   `json:"microsoft_server_gated_crypto"`
	MicrosoftSgcSerialized         bool   `json:"microsoft_sgc_serialized"`
	MicrosoftSmartDisplay          bool   `json:"microsoft_smart_display"`
	MicrosoftSmartcardLogon        bool   `json:"microsoft_smartcard_logon"`
	MicrosoftSystemHealth          bool   `json:"microsoft_system_health"`
	MicrosoftSystemHealthLoophole  bool   `json:"microsoft_system_health_loophole"`
	MicrosoftTimestampSigning      bool   `json:"microsoft_timestamp_signing"`
	MicrosoftWhqlCrypto            bool   `json:"microsoft_whql_crypto"`
	NetscapeServerGatedCrypto      bool   `json:"netscape_server_gated_crypto"`
	OcspSigning                    bool   `json:"ocsp_signing"`
	SbgpCertAaServiceAuth          bool   `json:"sbgp_cert_aa_service_auth"`
	ServerAuth                     bool   `json:"server_auth"`
	TimeStamping                   bool   `json:"time_stamping"`
	Unknown                        string `json:"unknown"`
}

type SignedCertificateTimestamp struct {
	LogId     string                                `json:"log_id"`
	Signature SignedCertificateTimestamps_Signature `json:"signature"`
	Timestamp time.Time                             `json:"timestamp"`
	Version   int                                   `json:"version"`
}

type SignedCertificateTimestamps_Signature struct {
	HashAlgorithm      string `json:"hash_algorithm"`
	Signature          string `json:"signature"`
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type CertificatePolicy struct {
	Cps        []string                       `json:"cps"`
	Id         string                         `json:"id"`
	UserNotice []CertificatePolicy_UserNotice `json:"user_notice"`
}

type CertificatePolicy_UserNotice struct {
	ExplicitText    string                                       `json:"explicit_text"`
	NoticeReference CertificatePolicy_UserNotice_NoticeReference `json:"notice_reference"`
}

type CertificatePolicy_UserNotice_NoticeReference struct {
	NoticeNumbers []int  `json:"notice_numbers"`
	Organization  string `json:"organization"`
}

type CabfOrganizationId struct {
	Country   string `json:"country"`
	Reference string `json:"reference"`
	Scheme    string `json:"scheme"`
	State     string `json:"state"`
}

type BasicConstraints struct {
	IsCa       bool `json:"is_ca"`
	MaxPathLen int  `json:"max_path_len"`
}

type AuthorityInfoAccess struct {
	IssuerUrls []string `json:"issuer_urls"`
	OcspUrls   []string `json:"ocsp_urls"`
}

type CtValue struct {
	AddedToCtAt  time.Time `json:"added_to_ct_at"`
	CtToCensysAt time.Time `json:"ct_to_censys_at"`
	Index        int64     `json:"index"`
}
