package accg

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ecsaas/api-cookie-config/DEFINE_VARIABLES/accgf"
)

type CookieParam struct {
	CookieMaxAge            int
	CookieMaxAgeTimeout     int
	CookieDomain            string
	CookiePath              string
	CookieSecure            bool
	CookieHttponly          bool
	CookieSameSite          int
	CookieConfigAuthParams  string
	CookieConfigServerCache string
	CookieConfigNewPassword string
	CookieConfigShopId      string
	Unparsed                []string
}

func GetCookieDomain(domain string) string {
	var arrDomain = strings.Split(domain, ".")
	if len(arrDomain) > 2 {
		arrDomain = arrDomain[len(arrDomain)-2:]
	} else if len(arrDomain) < 2 {
		return ""
	}
	return fmt.Sprintf(".%s", strings.Join(arrDomain, "."))
}

func GetCookieEnvParam(request *http.Request) CookieParam {
	var (
		cookieMaxAge, _        = strconv.Atoi(os.Getenv(accgf.COOKIE_MAX_AGE))
		cookieMaxAgeTimeout, _ = strconv.Atoi(os.Getenv(accgf.COOKIE_MAX_AGE_TIMEOUT))
		cookieSameSite, _      = strconv.Atoi(os.Getenv(accgf.COOKIE_SAME_SITE))
		cookieDomain           = os.Getenv(accgf.COOKIE_DOMAIN)
	)
	if cookieDomain == "" {
		cookieDomain = GetCookieDomain(request.Host)
	}

	//http.SameSiteDefaultMode 1
	//http.SameSiteLaxMode 2
	//http.SameSiteStrictMode 3
	//http.SameSiteNoneMode 4
	return CookieParam{
		CookieMaxAge:            cookieMaxAge,
		CookieMaxAgeTimeout:     cookieMaxAgeTimeout,
		CookieDomain:            cookieDomain,
		CookiePath:              os.Getenv(accgf.COOKIE_PATH),
		CookieSecure:            strings.ToLower(os.Getenv(accgf.COOKIE_SECURE)) == "true",
		CookieHttponly:          strings.ToLower(os.Getenv(accgf.COOKIE_HTTPONLY)) == "true",
		CookieSameSite:          cookieSameSite,
		CookieConfigAuthParams:  os.Getenv(accgf.COOKIE_CONFIG_AUTH_PARAMS),
		CookieConfigServerCache: os.Getenv(accgf.COOKIE_CONFIG_SERVER_CACHE),
		CookieConfigNewPassword: os.Getenv(accgf.COOKIE_CONFIG_NEW_PASSWORD),
		CookieConfigShopId:      os.Getenv(accgf.COOKIE_CONFIG_SHOP_ID),
		Unparsed:                []string{},
	}
}
