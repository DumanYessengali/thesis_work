package config

type Config struct {
	Data  []Data `json:"data"`
	Links Links  `json:"links"`
}
type Data struct {
	Attributes Attributes `json:"attributes"`
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Links      Links      `json:"links"`
}

type Attributes struct {
	RegionalInternetRegistry string               `json:"regional_internet_registry"`
	Jarm                     string               `json:"jarm"`
	Network                  string               `json:"network"`
	LastHttpsCertificateDate int64                `json:"last_https_certificate_date"`
	Tags                     []string             `json:"tags"`
	Country                  string               `json:"country"`
	AsOwner                  string               `json:"as_owner"`
	LastAnalysisStats        LastAnalysisStats    `json:"last_analysis_stats"`
	ASN                      int64                `json:"asn"`
	WhoisDate                int64                `json:"whois_date"`
	LastAnalysisResults      LastAnalysisResults  `json:"last_analysis_results"`
	Reputation               int64                `json:"reputation"`
	LastModificationDate     int64                `json:"last_modification_date"`
	TotalVotes               TotalVotes           `json:"total_votes"`
	LastHttpsCertificate     LastHttpsCertificate `json:"last_https_certificate"`
	Continent                string               `json:"continent"`
	Whois                    string               `json:"whois"`
}
type Links struct {
	Self string `json:"self"`
}

type LastAnalysisStats struct {
	Harmless   int64 `json:"harmless"`
	Malicious  int64 `json:"malicious"`
	Suspicious int64 `json:"suspicious"`
	Undetected int64 `json:"undetected"`
	Timeout    int64 `json:"timeout"`
}

type TotalVotes struct {
	Harmless  int64 `json:"harmless"`
	Malicious int64 `json:"malicious"`
}

type LastAnalysisResults struct {
	CMCThreatIntelligence    Scanner `json:"CMC Threat Intelligence"`
	SnortIPSampleList        Scanner `json:"Snort IP sample list"`
	Scanner0xSI_f33d         Scanner `json:"0xSI_f33d"`
	Armis                    Scanner `json:"Armis"`
	ViriBack                 Scanner `json:"ViriBack"`
	ComodoValkyrieVerdict    Scanner `json:"Comodo Valkyrie Verdict"`
	PhishLabs                Scanner `json:"PhishLabs"`
	K7AntiVirus              Scanner `json:"K7AntiVirus"`
	CINSArmy                 Scanner `json:"CINS Army"`
	Quttera                  Scanner `json:"Quttera"`
	OpenPhish                Scanner `json:"OpenPhish"`
	VXVault                  Scanner `json:"VX Vault"`
	WebSecurityGuard         Scanner `json:"Web Security Guard"`
	Scantitan                Scanner `json:"Scantitan"`
	AlienVault               Scanner `json:"AlienVault"`
	Sophos                   Scanner `json:"Sophos"`
	Phishtank                Scanner `json:"Phishtank"`
	EonScope                 Scanner `json:"EonScope"`
	Cyan                     Scanner `json:"Cyan"`
	Spam404                  Scanner `json:"Spam404"`
	SecureBrain              Scanner `json:"SecureBrain"`
	HopliteIndustries        Scanner `json:"Hoplite Industries"`
	CRDF                     Scanner `json:"CRDF"`
	Fortinet                 Scanner `json:"Fortinet"`
	AlphaMountainAI          Scanner `json:"alphaMountain.ai"`
	Lionic                   Scanner `json:"Lionic"`
	VirusdieExternalSiteScan Scanner `json:"Virusdie External Site Scan"`
	GoogleSafebrowsing       Scanner `json:"Google Safebrowsing"`
	SafeToOpen               Scanner `json:"SafeToOpen"`
	ADMINUSLabs              Scanner `json:"ADMINUSLabs"`
	CyberCrime               Scanner `json:"CyberCrime"`
	HeimdalSecurity          Scanner `json:"Heimdal Security"`
	AutoShun                 Scanner `json:"AutoShun"`
	Trustwave                Scanner `json:"Trustwave"`
	AICCMONITORAPP           Scanner `json:"AICC (MONITORAPP)"`
	CyRadar                  Scanner `json:"CyRadar"`
	DrWeb                    Scanner `json:"Dr.Web"`
	Emsisoft                 Scanner `json:"Emsisoft"`
	Abusix                   Scanner `json:"Abusix"`
	Webroot                  Scanner `json:"Webroot"`
	Avira                    Scanner `json:"Avira"`
	Securolytics             Scanner `json:"securolytics"`
	AntiyAVL                 Scanner `json:"Antiy-AVL"`
	Acronis                  Scanner `json:"Acronis"`
	QuickHeal                Scanner `json:"Quick Heal"`
	ESTsecurityThreatInside  Scanner `json:"ESTsecurity-Threat Inside"`
	DNS8                     Scanner `json:"DNS8"`
	BenkowCC                 Scanner `json:"benkow.cc""`
	EmergingThreats          Scanner `json:"EmergingThreats"`
	ChongLuaDao              Scanner `json:"Chong Lua Dao"`
	YandexSafebrowsing       Scanner `json:"Yandex Safebrowsing"`
	MalwareDomainList        Scanner `json:"MalwareDomainList"`
	Lumu                     Scanner `json:"Lumu"`
	Zvelo                    Scanner `json:"zvelo"`
	Kaspersky                Scanner `json:"Kaspersky"`
	Segasec                  Scanner `json:"Segasec"`
	SucuriSiteCheck          Scanner `json:"Sucuri SiteCheck"`
	DesenmascaraME           Scanner `json:"desenmascara.me"`
	URLhaus                  Scanner `json:"URLhaus"`
	PREBYTES                 Scanner `json:"PREBYTES"`
	StopForumSpam            Scanner `json:"StopForumSpam"`
	Blueliv                  Scanner `json:"Blueliv"`
	Netcraft                 Scanner `json:"Netcraft"`
	ZeroCERT                 Scanner `json:"ZeroCERT"`
	PhishingDatabase         Scanner `json:"Phishing Database"`
	MalwarePatrol            Scanner `json:"MalwarePatrol"`
	MalBeacon                Scanner `json:"MalBeacon"`
	IPsum                    Scanner `json:"IPsum"`
	Spamhaus                 Scanner `json:"Spamhaus"`
	Malwared                 Scanner `json:"Malwared"`
	BitDefender              Scanner `json:"BitDefender"`
	GreenSnow                Scanner `json:"GreenSnow"`
	GData                    Scanner `json:"G-Data"`
	StopBadware              Scanner `json:"StopBadware"`
	SCUMWAREORG              Scanner `json:"SCUMWARE.org"`
	MalwaresCOMURLChecker    Scanner `json:"malwares.com URL checker"`
	NotMining                Scanner `json:"NotMining"`
	ForcepointThreatSeeker   Scanner `json:"Forcepoint ThreatSeeker"`
	Certego                  Scanner `json:"Certego"`
	ESET                     Scanner `json:"ESET"`
	Threatsourcing           Scanner `json:"Threatsourcing"`
	MalSilo                  Scanner `json:"MalSilo"`
	Nucleon                  Scanner `json:"Nucleon"`
	BADWAREINFO              Scanner `json:"BADWARE.INFO"`
	ThreatHive               Scanner `json:"ThreatHive"`
	FraudScore               Scanner `json:"FraudScore"`
	Tencent                  Scanner `json:"Tencent"`
	BforeAiPreCrime          Scanner `json:"Bfore.Ai PreCrime"`
	BaiduInternational       Scanner `json:"Baidu-International"`
}

type Scanner struct {
	Category   string `json:"category"`
	Result     string `json:"result"`
	Method     string `json:"method"`
	EngineName string `json:"engine_name"`
}

type LastHttpsCertificate struct {
	Size               int64         `json:"size"`
	PublicKey          PublicKey     `json:"public_key"`
	ThumbprintSha256   string        `json:"thumbprint_sha256"`
	Tags               []string      `json:"tags"`
	CertSignature      CertSignature `json:"cert_signature"`
	Validity           Validity      `json:"validity"`
	Version            string        `json:"version"`
	Extensions         Extensions    `json:"extensions"`
	SignatureAlgorithm string        `json:"signature_algorithm"`
	SerialNumber       string        `json:"serial_number"`
	Thumbprint         string        `json:"thumbprint"`
	Issuer             Issuer        `json:"issuer"`
	Subject            Subject       `json:"subject"`
}

type PublicKey struct {
	EC        EC     `json:"ec"`
	Algorithm string `json:"algorithm"`
}

type EC struct {
	Oid string `json:"oid"`
	Pub string `json:"pub"`
}

type CertSignature struct {
	Signature          string `json:"signature"`
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type Validity struct {
	NotAfter  string `json:"not_after"`
	NotBefore string `json:"not_before"`
}

type Extensions struct {
	CertificatePolicies    []string               `json:"certificate_policies"`
	ExtendedKeyUsage       []string               `json:"extended_key_usage"`
	AuthorityKeyIdentifier AuthorityKeyIdentifier `json:"authority_key_identifier"`
	SubjectAlternativeName []string               `json:"subject_alternative_name"`
	Tags                   []string               `json:"tags"`
	SubjectKeyIdentifier   string                 `json:"subject_key_identifier"`
	CrlDistributionPoints  []string               `json:"crl_distribution_points"`
	KeyUsage               []string               `json:"key_usage"`
	Numbers                string                 `json:"1.3.6.1.4.1.11129.2.4.2"`
	CA                     bool                   `json:"CA"`
	CAInformationAccess    CAInformationAccess    `json:"ca_information_access"`
}

type AuthorityKeyIdentifier struct {
	KeyID string `json:"keyid"`
}

type CAInformationAccess struct {
	CAIssuers string `json:"CA Issuers"`
	OCSP      string `json:"OCSP"`
}

type Issuer struct {
	C  string `json:"C"`
	CN string `json:"CN"`
	O  string `json:"O"`
}

type Subject struct {
	C  string `json:"C"`
	ST string `json:"ST"`
	L  string `json:"L"`
	O  string `json:"O"`
	CN string `json:"CN"`
}
