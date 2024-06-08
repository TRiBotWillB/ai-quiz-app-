import React from 'react'
import ReactDOM from 'react-dom/client'
import {
    createBrowserRouter, Outlet,
    RouterProvider,
} from "react-router-dom";
import {getSuperTokensRoutesForReactRouterDom} from "supertokens-auth-react/ui";
import {EmailPasswordPreBuiltUI} from 'supertokens-auth-react/recipe/emailpassword/prebuiltui';
import * as reactRouterDom from "react-router-dom";
import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import Session, {SessionAuth} from "supertokens-auth-react/recipe/session";
import {init, SuperTokensWrapper} from "supertokens-auth-react";

init({
    appInfo: {
        // learn more about this on https://supertokens.com/docs/thirdpartyemailpassword/appinfo
        appName: "MicroservicesExample",
        apiDomain: "http://127.0.0.1:3000",
        websiteDomain: "http://localhost:5173",
        apiBasePath: "/auth",
        websiteBasePath: "/auth"
    },
    recipeList: [
        EmailPassword.init(),
        Session.init()
    ]
});

const router = createBrowserRouter([
    ...getSuperTokensRoutesForReactRouterDom(reactRouterDom, [EmailPasswordPreBuiltUI]).map((r) => r.props),
    {
        element: <SessionAuth><Outlet/></SessionAuth>,
        children: [
            {
                path: '/',
                element: <p>This route is protected</p>
            },
            {
                path: '/test',
                element: <p>This route is also protected</p>
            }
        ]
    },
    {
        path: '/unprotected',
        element: <p>Unprotected route example</p>
    }
]);


ReactDOM.createRoot(document.getElementById("root")).render(
    <React.StrictMode>
        <SuperTokensWrapper>
            <RouterProvider router={router}/>
        </SuperTokensWrapper>

    </React.StrictMode>
);
