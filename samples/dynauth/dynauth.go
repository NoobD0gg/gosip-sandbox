package dynauth

import (
	"fmt"

	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/auth/addin"
	"github.com/koltyakov/gosip/auth/adfs"
	"github.com/koltyakov/gosip/auth/anon"
	"github.com/koltyakov/gosip/auth/fba"
	"github.com/koltyakov/gosip/auth/ntlm"
	"github.com/koltyakov/gosip/auth/saml"
	"github.com/koltyakov/gosip/auth/tmg"

	ntlm2 "github.com/NoobD0gg/gosip-sandbox/strategies/ntlm"
	"github.com/NoobD0gg/gosip-sandbox/strategies/device"
	"github.com/NoobD0gg/gosip-sandbox/strategies/ondemand"
)

// NewAuthCnfg creates AuthCnfg object based on strategy and config path
func NewAuthCnfg(strategy string, configPath string) (gosip.AuthCnfg, error) {
	var auth gosip.AuthCnfg

	switch strategy {
	case "addin":
		auth = &addin.AuthCnfg{}
		break
	case "adfs":
		auth = &adfs.AuthCnfg{}
		break
	case "fba":
		auth = &fba.AuthCnfg{}
		break
	case "ntlm":
		auth = &ntlm.AuthCnfg{}
		break
	case "ntlm2":
		auth = &ntlm2.AuthCnfg{}
		break
	case "saml":
		auth = &saml.AuthCnfg{}
		break
	case "tmg":
		auth = &tmg.AuthCnfg{}
		break
	case "device":
		auth = &device.AuthCnfg{}
		break
	case "ondemand":
		auth = &ondemand.AuthCnfg{}
		break
	case "anon":
		auth = &anon.AuthCnfg{}
		break
	default:
		return nil, fmt.Errorf("can't resolve the strategy: %s", strategy)
	}

	if configPath == "" {
		return nil, fmt.Errorf("config path must be provided")
	}

	if err := auth.ReadConfig(configPath); err != nil {
		return nil, fmt.Errorf("unable to get config: %v", err)
	}

	return auth, nil
}
