package npnweb

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnservice/user"
	"net/http"
	"net/url"
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"logur.dev/logur"
)

type RequestContext struct {
	App         AppInfo
	Logger      logur.Logger
	Profile     *npnuser.UserProfile
	Routes      *mux.Router
	Request     *url.URL
	Title       string
	Breadcrumbs Breadcrumbs
	Flashes     []string
	Session     *sessions.Session
}

func (r *RequestContext) Route(act string, pairs ...string) string {
	return Route(r.Routes, r.Logger, act, pairs...)
}

func ExtractContext(w http.ResponseWriter, r *http.Request, addIfMissing bool) *RequestContext {
	ai, ok := r.Context().Value(npncore.InfoKey).(AppInfo)
	if !ok {
		ai.Logger().Warn("cannot load AppInfo")
	}
	routes, ok := r.Context().Value(npncore.RoutesKey).(*mux.Router)
	if !ok {
		ai.Logger().Warn("cannot load Router")
	}
	session, err := Store.Get(r, npncore.AppName)
	if err != nil {
		session = sessions.NewSession(Store, npncore.AppName)
	}

	var userID uuid.UUID
	userIDValue, ok := session.Values[npncore.KeyUser]
	if ok && len(userIDValue.(string)) == 36 {
		userID, err = uuid.FromString(userIDValue.(string))
		if err != nil {
			ai.Logger().Warn(fmt.Sprintf("cannot parse uuid [%v]: %+v", userIDValue, err))
			userID = SetSessionUser(npncore.UUID(), session, r, w, ai.Logger())
		}
	} else {
		userID = SetSessionUser(npncore.UUID(), session, r, w, ai.Logger())
	}

	var u *user.SystemUser
	if ai.User() == nil {
		content, err := ai.Files().ReadFile("profile.json")
		if err != nil {
			return nil
		}
		tgt := &npnuser.UserProfile{}
		err = npncore.FromJSON([]byte(content), tgt)
		if err != nil {
			ai.Logger().Warn(fmt.Sprintf("can't load profile: %+v", err))
			return nil
		}
		u = &user.SystemUser{
			UserID:    tgt.UserID,
			Name:      tgt.Name,
			Role:      tgt.Role.String(),
			Theme:     tgt.Theme.String(),
			NavColor:  tgt.NavColor,
			LinkColor: tgt.LinkColor,
			Picture:   tgt.Picture,
			Locale:    tgt.Locale.String(),
			Created:   time.Now(),
		}
	} else {
		u = ai.User().GetByID(userID, addIfMissing)
	}
	var prof *npnuser.UserProfile
	if u == nil {
		prof = npnuser.NewUserProfile(userID, "")
	} else {
		prof = u.ToProfile()
	}

	flashes := make([]string, 0)
	for _, f := range session.Flashes() {
		flashes = append(flashes, fmt.Sprint(f))
	}

	logger := logur.WithFields(ai.Logger(), map[string]interface{}{"path": r.URL.Path, "method": r.Method})

	return &RequestContext{
		App:         ai,
		Logger:      logger,
		Profile:     prof,
		Routes:      routes,
		Request:     r.URL,
		Title:       npncore.AppName,
		Breadcrumbs: nil,
		Flashes:     flashes,
		Session:     session,
	}
}

func SetSessionUser(userID uuid.UUID, session *sessions.Session, r *http.Request, w http.ResponseWriter, logger logur.Logger) uuid.UUID {
	session.Values[npncore.KeyUser] = userID.String()
	session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteStrictMode}
	err := session.Save(r, w)
	if err != nil {
		logger.Warn(fmt.Sprintf("cannot save session: %+v", err))
	}
	return userID
}
