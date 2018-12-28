// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
    "github.com/tinystack/goweb"
    "github.com/tinystack/syncd"
    userModule "github.com/tinystack/syncd/module/user"
    serverModule "github.com/tinystack/syncd/module/server"
    projectModule "github.com/tinystack/syncd/module/project"
    deployModule "github.com/tinystack/syncd/module/deploy"
)

func init() {
    handler()
}

func handler() {
    h := map[string]goweb.HandlerFunc{
        syncd.API_USER_GROUP_UPDATE: userModule.GroupUpdate,
        syncd.API_USER_GROUP_LIST: userModule.GroupList,
        syncd.API_USER_GROUP_DETAIL: userModule.GroupDetail,
        syncd.API_USER_GROUP_DELETE: userModule.GroupDelete,
        syncd.API_USER_LOGIN: userModule.Login,
        syncd.API_USER_LOGIN_STATUS: userModule.LoginStatus,
        syncd.API_USER_LOGOUT: userModule.Logout,
        syncd.API_USER_PRIV_LIST: userModule.PrivList,
        syncd.API_USER_UPDATE: userModule.UserUpdate,
        syncd.API_USER_LIST: userModule.UserList,
        syncd.API_USER_DETAIL: userModule.UserDetail,
        syncd.API_USER_EXISTS: userModule.UserExists,
        syncd.API_USER_DELETE: userModule.UserDelete,
        syncd.API_USER_SEARCH: userModule.UserSearch,

        syncd.API_SERVER_GROUP_UPDATE: serverModule.GroupUpdate,
        syncd.API_SERVER_GROUP_LIST: serverModule.GroupList,
        syncd.API_SERVER_GROUP_DETAIL: serverModule.GroupDetail,
        syncd.API_SERVER_GROUP_DELETE: serverModule.GroupDelete,
        syncd.API_SERVER_GROUP_MULTI: serverModule.GroupMulti,
        syncd.API_SERVER_UPDATE: serverModule.ServerUpdate,
        syncd.API_SERVER_LIST: serverModule.ServerList,
        syncd.API_SERVER_DETAIL: serverModule.ServerDetail,
        syncd.API_SERVER_DELETE: serverModule.ServerDelete,

        syncd.API_PROJECT_UPDATE: projectModule.ProjectUpdate,
        syncd.API_PROJECT_LIST: projectModule.ProjectList,
        syncd.API_PROJECT_DETAIL: projectModule.ProjectDetail,
        syncd.API_PROJECT_DELETE: projectModule.ProjectDelete,
        syncd.API_PROJECT_EXISTS: projectModule.ProjectExists,
        syncd.API_PROJECT_STATUS_CHANGE: projectModule.ProjectChangeStatus,
        syncd.API_PROJECT_SPACE_UPDATE: projectModule.SpaceUpdate,
        syncd.API_PROJECT_SPACE_LIST: projectModule.SpaceList,
        syncd.API_PROJECT_SPACE_DETAIL: projectModule.SpaceDetail,
        syncd.API_PROJECT_SPACE_DELETE: projectModule.SpaceDelete,
        syncd.API_PROJECT_SPACE_EXISTS: projectModule.SpaceExists,
        syncd.API_PROJECT_SPACE_USER_ADD: projectModule.UserAdd2Space,
        syncd.API_PROJECT_SPACE_USER_LIST: projectModule.UserList,
        syncd.API_PROJECT_SPACE_USER_REMOVE: projectModule.UserRemoveFromSpace,

        syncd.API_DEPLOY_APPLY_SUBMIT: deployModule.ApplySubmit,
        syncd.API_DEPLOY_APPLY_LIST: deployModule.ApplyList,
        syncd.API_DEPLOY_REPO_TAGLIST: deployModule.RepoTagList,
        syncd.API_DEPLOY_REPO_RESET: deployModule.RepoReset,
        syncd.API_DEPLOY_REPO_COMMITLIST: deployModule.RepoCommitList,
    }

    for k, v := range h {
        register(k, v)
    }
}