package paypayopa

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
