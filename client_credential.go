package paypayopa

// Environment is a type used to specify the destination of the request.
//
// Environment はリクエストの送信先を指定するための型.
type Environment string

const (
	// EnvProduction is a value that specifies that the request should be
	// sent to the production environment.
	//
	// EnvProduction は本番環境にリクエストを送ることを指定する値.
	EnvProduction Environment = "https://api.paypay.ne.jp"

	// EnvStaging is a value that specifies that the request should be
	// sent to the staging environment.
	//
	// EnvStaging はステージング環境にリクエストを送ることを指定する値.
	EnvStaging Environment = "https://stg-api.paypay.ne.jp"

	// EnvSandbox is a value that specifies that the request should be
	// sent to the sandbox environment.
	//
	// EnvSandbox はサンドボックス環境にリクエストを送ることを指定する値.
	EnvSandbox Environment = "https://stg-api.sandbox.paypay.ne.jp"
)

// Credential is a structure that holds credentials.
//
// Credential は資格情報を保持する構造体.
type Credential struct {
	env          Environment
	apiKey       string
	apiKeySecret string
	merchantID   string
}

// NewCredential creates a structure to hold the credentials.
//
// NewCredential は資格情報を保持する構造体を生成する.
//
// About Arguments
//
// - env
//   Where to send the request. Specify one of the following
//   environments: production, staging, or sandbox.
//
//   リクエストの送信先. 本番環境, ステージング環境, サンドボックス環境の中から指定する.
//
// - apiKey
//   API key created in PayPay for Developers.
//
//   PayPay for Developers で作成した API キー.
//
// - apiKeySecret
//   Secret key created in PayPay for Developers.
//
//   PayPay for Developers で作成したシークレットキー.
//
// - merchantID
//   The merchant ID registered with PayPay for Developers.
//
//   PayPay for Developers で登録した加盟店の ID.
func NewCredential(
	env Environment,
	apiKey string,
	apiKeySecret string,
	merchantID string,
) *Credential {
	return &Credential{
		env:          env,
		apiKey:       apiKey,
		apiKeySecret: apiKeySecret,
		merchantID:   merchantID,
	}
}
