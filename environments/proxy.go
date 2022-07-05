//go:build proxy

package environments

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"srvexec/common"
	"strings"
)

const (
	prefixEnv = "PROXY_ENV_"
)

var (
	MainEnvironments = common.Environment{
		Name:    "proxy",
		Handler: handleProxy,
	}
	executorsUrl map[string]string
)

type contextProxy struct {
	Env string `json:"env"`
}

func init() {
	executorsUrl = make(map[string]string)
	for _, k := range common.Config.Keys() {
		if strings.HasPrefix(k, prefixEnv) {
			// trim prefix
			woPrefix := strings.TrimPrefix(k, prefixEnv)
			// replace _ => -
			wDashes := strings.ReplaceAll(woPrefix, "_", "-")
			// lower case and add
			executorsUrl[strings.ToLower(wDashes)] = common.Config.Get(k)
		}
	}

	for k, v := range executorsUrl {
		common.LogDebug("Executor " + k + " is at " + v)
	}
}

func handleProxy(j common.ToHandle) (string, common.Status) {
	var ctx contextProxy

	if err := j.Exercice.UnmarshalContexte(&ctx); err != nil {
		common.LogError("error while unmarshalling context: " + err.Error())
		return "Error unmarshalling context", common.ErrorInternal
	}

	targetEnv := ctx.Env

	// check if env is defined
	if targetEnv == "" {
		return "No env defined", common.ErrorInternal
	}

	// check if env is in config keys
	if _, ok := executorsUrl[targetEnv]; !ok {

		// get keys
		envs := make([]string, len(executorsUrl))
		i := 0
		for k := range executorsUrl {
			envs[i] = k
			i++
		}

		// return error accordingly
		last := envs[len(envs)-1]
		exceptLast := envs[:len(envs)-1]

		common.LogDebug("envs: %v", envs)
		if len(envs) == 1 {
			return "Environment " + targetEnv + " don't exists. Available environment is " + last + ".", common.ErrorInternal
		}

		return "Environment " + targetEnv + " don't exists. Choose among " + strings.Join(exceptLast, ", ") + " or " + last + ".", common.ErrorInternal
	}

	//  Forword request to target env
	body, _ := json.Marshal(j)
	bodyBuffer := bytes.NewReader(body)

	common.LogDebug("forwarding request to " + targetEnv)
	httpResp, err := http.Post(executorsUrl[targetEnv]+"/exec", "application/json", bodyBuffer)

	// Can't reach target env/executor
	if err != nil {
		common.LogError("error while sending request to target env: " + err.Error())
		return "Error while sending request to target env", common.ErrorInternal
	}

	// Read response
	out, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		common.LogError("error while reading response from target env: " + err.Error())
		return "Error while reading response from target env", common.ErrorInternal
	}

	var resp map[string]string
	json.Unmarshal(out, &resp)

	// send
	return resp["output"], common.StatusFromString(resp["status"])
}
