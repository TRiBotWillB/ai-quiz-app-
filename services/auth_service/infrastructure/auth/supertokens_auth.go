package auth

import (
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func InitSupertokens() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "",
			APIKey:        "",
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "MicroservicesExample",
			APIDomain:       "http://localhost:54547",
			WebsiteDomain:   "http://localhost:5173",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			//thirdparty.Init(&tpmodels.TypeInput{ }), Don't want this for now :)
			emailpassword.Init(nil),
			session.Init(nil), // initializes session features
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
