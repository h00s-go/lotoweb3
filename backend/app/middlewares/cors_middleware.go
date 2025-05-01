package middlewares

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-raptor/raptor/v4/core"
)

// CORSConfig defines the configuration for CORS middleware.
type CORSConfig struct {
	// AllowOrigins specifies the allowed origins for CORS requests.
	// Supports wildcards '*' and '?' (converted to regex '.*' and '.').
	// Default: ["*"].
	AllowOrigins []string `yaml:"allow_origins"`

	// AllowOriginFunc is a custom function to validate origins.
	// If set, AllowOrigins is ignored.
	AllowOriginFunc func(origin string) (bool, error) `yaml:"-"`

	// AllowMethods specifies the allowed HTTP methods.
	// Default: [GET, HEAD, PUT, PATCH, POST, DELETE].
	AllowMethods []string `yaml:"allow_methods"`

	// AllowHeaders specifies the allowed request headers.
	// Default: [].
	AllowHeaders []string `yaml:"allow_headers"`

	// AllowCredentials indicates whether credentials are allowed.
	// Default: false.
	AllowCredentials bool `yaml:"allow_credentials"`

	// UnsafeWildcardOriginWithAllowCredentials allows wildcard '*' with AllowCredentials.
	// WARNING: This is insecure and should be used cautiously.
	// Default: false.
	UnsafeWildcardOriginWithAllowCredentials bool `yaml:"unsafe_wildcard_origin_with_allow_credentials"`

	// ExposeHeaders specifies headers exposed to clients.
	// Default: [].
	ExposeHeaders []string `yaml:"expose_headers"`

	// MaxAge specifies the cache duration (in seconds) for preflight responses.
	// If 0, the header is not set. Negative values send "0".
	// Default: 0.
	MaxAge int `yaml:"max_age"`
}

// DefaultCORSConfig provides default CORS settings.
var DefaultCORSConfig = CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		"GET",
		"HEAD",
		"PUT",
		"PATCH",
		"POST",
		"DELETE",
	},
}

// CORSMiddleware implements CORS handling for Raptor.
type CORSMiddleware struct {
	core.Middleware
	config              CORSConfig
	allowOriginPatterns []*regexp.Regexp
	allowMethods        string
	allowHeaders        string
	exposeHeaders       string
	maxAge              string
}

// NewCORSMiddleware creates a new CORSMiddleware with the given config.
func NewCORSMiddleware(config CORSConfig) *CORSMiddleware {
	return &CORSMiddleware{
		config: config,
	}
}

// Init initializes the middleware with Resources and compiles origin patterns.
func (m *CORSMiddleware) Init(r *core.Resources) {
	m.Middleware.Init(r)

	// Apply defaults
	if len(m.config.AllowOrigins) == 0 {
		m.config.AllowOrigins = DefaultCORSConfig.AllowOrigins
	}
	if len(m.config.AllowMethods) == 0 {
		m.config.AllowMethods = DefaultCORSConfig.AllowMethods
	}

	// Compile origin patterns
	m.allowOriginPatterns = make([]*regexp.Regexp, 0, len(m.config.AllowOrigins))
	for _, origin := range m.config.AllowOrigins {
		if origin == "*" {
			continue // Handled separately
		}
		pattern := regexp.QuoteMeta(origin)
		pattern = strings.ReplaceAll(pattern, "\\*", ".*")
		pattern = strings.ReplaceAll(pattern, "\\?", ".")
		pattern = "^" + pattern + "$"

		re, err := regexp.Compile(pattern)
		if err != nil {
			r.Log.Warn("Invalid origin pattern, skipping", "origin", origin, "error", err)
			continue
		}
		m.allowOriginPatterns = append(m.allowOriginPatterns, re)
	}

	// Precompute header values
	m.allowMethods = strings.Join(m.config.AllowMethods, ",")
	m.allowHeaders = strings.Join(m.config.AllowHeaders, ",")
	m.exposeHeaders = strings.Join(m.config.ExposeHeaders, ",")
	if m.config.MaxAge == 0 {
		m.maxAge = "0"
	} else {
		m.maxAge = strconv.Itoa(m.config.MaxAge)
	}

	// Warn about insecure configuration
	if m.config.AllowCredentials && contains(m.config.AllowOrigins, "*") && !m.config.UnsafeWildcardOriginWithAllowCredentials {
		r.Log.Warn("Insecure CORS configuration: AllowCredentials with AllowOrigins [*] requires UnsafeWildcardOriginWithAllowCredentials")
	}

	r.Log.Info("CORSMiddleware initialized")
}

// New handles CORS for a request, setting headers and processing preflight requests.
func (m *CORSMiddleware) New(c *core.Context, next func(*core.Context) error) error {
	req := c.Request()
	res := c.Response()

	origin := req.Header.Get(core.HeaderOrigin)
	allowOrigin := ""

	res.Header().Add(core.HeaderVary, core.HeaderOrigin)

	// Preflight request (OPTIONS method)
	preflight := req.Method == "OPTIONS"

	// No Origin header: proceed with next handler for non-preflight, or return 204 for preflight
	if origin == "" {
		if !preflight {
			return next(c)
		}
		return c.NoContent(204)
	}

	// Validate origin
	if m.config.AllowOriginFunc != nil {
		allowed, err := m.config.AllowOriginFunc(origin)
		if err != nil {
			m.Resources.Log.Error("AllowOriginFunc error", "origin", origin, "error", err)
			return err
		}
		if allowed {
			allowOrigin = origin
		}
	} else {
		for _, o := range m.config.AllowOrigins {
			if o == "*" && m.config.AllowCredentials && m.config.UnsafeWildcardOriginWithAllowCredentials {
				allowOrigin = origin
				break
			}
			if o == "*" || o == origin {
				allowOrigin = o
				break
			}
			if matchSubdomain(origin, o) {
				allowOrigin = origin
				break
			}
		}

		if allowOrigin == "" && len(origin) <= (253+3+5) && strings.Contains(origin, "://") {
			for _, re := range m.allowOriginPatterns {
				if re.MatchString(origin) {
					allowOrigin = origin
					break
				}
			}
		}
	}

	// Origin not allowed
	if allowOrigin == "" {
		if !preflight {
			m.Resources.Log.Warn("Unauthorized origin", "origin", origin)
			return c.JSON(401, map[string]string{"error": "Unauthorized origin"})
		}
		return c.NoContent(204)
	}

	// Set CORS headers
	res.Header().Set(core.HeaderAccessControlAllowOrigin, allowOrigin)
	if m.config.AllowCredentials {
		res.Header().Set(core.HeaderAccessControlAllowCredentials, "true")
	}

	// Simple request
	if !preflight {
		if m.exposeHeaders != "" {
			res.Header().Set(core.HeaderAccessControlExposeHeaders, m.exposeHeaders)
		}
		return next(c)
	}

	// Preflight request
	res.Header().Add(core.HeaderVary, core.HeaderAccessControlRequestMethod)
	res.Header().Add(core.HeaderVary, core.HeaderAccessControlRequestHeaders)

	res.Header().Set(core.HeaderAccessControlAllowMethods, m.allowMethods)

	if m.allowHeaders != "" {
		res.Header().Set(core.HeaderAccessControlAllowHeaders, m.allowHeaders)
	} else {
		h := req.Header.Get(core.HeaderAccessControlRequestHeaders)
		if h != "" {
			res.Header().Set(core.HeaderAccessControlAllowHeaders, h)
		}
	}

	if m.config.MaxAge != 0 {
		res.Header().Set(core.HeaderAccessControlMaxAge, m.maxAge)
	}

	return c.NoContent(204)
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// matchSubdomain checks if the origin matches a subdomain pattern.
func matchSubdomain(origin, pattern string) bool {
	if !strings.Contains(pattern, "*") {
		return false
	}
	pattern = regexp.QuoteMeta(pattern)
	pattern = strings.ReplaceAll(pattern, "\\*", ".*")
	re, err := regexp.Compile("^" + pattern + "$")
	if err != nil {
		return false
	}
	return re.MatchString(origin)
}
