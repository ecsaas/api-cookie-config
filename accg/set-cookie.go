package accg

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ecsaas/api-cookie-config/DEFINE_VARIABLES/accgf"
)

func CookieSetAuthLogin(request *http.Request, writer http.ResponseWriter, token string) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(writer,
		k.CookieConfigAuthParams,
		token,
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAge),
		k.CookieMaxAge,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
	CookieSetLogoutNewPassword(request, writer)
}

func CookieSetShopId(request *http.Request, writer http.ResponseWriter, shopId string) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(writer,
		k.CookieConfigShopId,
		shopId,
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAge),
		k.CookieMaxAge,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
}

func CookieSetAuthNewPassword(
	request *http.Request,
	writer http.ResponseWriter,
	token string,
	timeout int,
) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(writer,
		k.CookieConfigNewPassword,
		token,
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(timeout),
		timeout,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
	AppSetCookie(writer,
		os.Getenv(accgf.COOKIE_CLIENT_WINDOW_NEW_PASSWORD),
		"true",
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(timeout),
		timeout,
		k.CookieSecure,
		false,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
	CookieSetLogout(request, writer)
}

func CookieServerDeleteCache(request *http.Request, writer http.ResponseWriter) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(writer,
		k.CookieConfigServerCache,
		"",
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAgeTimeout),
		k.CookieMaxAgeTimeout,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
}

func CookieServerSetCache(
	request *http.Request,
	writer http.ResponseWriter,
	serverCacheValue string,
) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(writer,
		k.CookieConfigServerCache,
		serverCacheValue,
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAge),
		k.CookieMaxAge,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
}

func CookieSetLogout(request *http.Request, writer http.ResponseWriter) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(
		writer,
		k.CookieConfigAuthParams,
		"",
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAgeTimeout),
		k.CookieMaxAgeTimeout,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
}

func CookieSetLogoutNewPassword(request *http.Request, writer http.ResponseWriter) {
	var k = GetCookieEnvParam(request)
	AppSetCookie(writer,
		k.CookieConfigNewPassword,
		"",
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAgeTimeout),
		k.CookieMaxAgeTimeout,
		k.CookieSecure,
		k.CookieHttponly,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
	AppSetCookie(writer,
		os.Getenv(accgf.COOKIE_CLIENT_WINDOW_NEW_PASSWORD),
		"",
		k.CookiePath,
		k.CookieDomain,
		GetExpireCookie(k.CookieMaxAgeTimeout),
		k.CookieMaxAgeTimeout,
		k.CookieSecure,
		false,
		http.SameSite(k.CookieSameSite),
		k.Unparsed,
		"",
	)
}

func AppSetCookie(
	writer http.ResponseWriter,
	name string,
	value string,
	path string,
	domain string,
	expire time.Time,
	cookieMaxAge int,
	cookieSecure bool,
	cookieHttponly bool,
	cookieSameSite http.SameSite,
	unparsed []string,
	raw string,
) {
	if raw == "" && value != "" {
		raw = fmt.Sprintf("%s=%s", name, value)
	} else if value == "" {
		raw = ""
	}
	http.SetCookie(writer, &http.Cookie{
		Name:       name,
		Value:      value,
		Path:       path,
		Domain:     domain,
		Expires:    expire,
		RawExpires: expire.Format(time.UnixDate),
		MaxAge:     cookieMaxAge,
		Secure:     cookieSecure,
		HttpOnly:   cookieHttponly,
		SameSite:   cookieSameSite,
		Raw:        raw,
		Unparsed:   unparsed,
	})
}

func GetExpireCookie(cookieMaxAge int) time.Time {
	return time.Now().Add(time.Second * time.Duration(cookieMaxAge))
}
