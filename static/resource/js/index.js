const router = new VueRouter({
    routes:[
        {
            path:"/", redirect:"/home/center"
        },
        {
            path: "/login", component: Login, name:"login"
        },{
            path: "/register", component: Register, name:"register"
        },{
            name: "home",
            path: "/home",
            component: Home,
            children:[
                {
                    name: "center",
                    path: "center",
                    component: Center
                },{
                    name: "app-register",
                    path: "app-register",
                    component: AppRegister
                },{
                    name: "app-edit",
                    path: "app/:id",
                    component: AppEditPage
                },{
                    name: "app-users",
                    path: "app/:id/users",
                    component: AppUsersPage
                },{
                    name: "app-roles",
                    path: "app/:id/roles",
                    component: AppRolesPage
                },{
                    name: "app-role-permission",
                    path: "app/:id/roles/:roleID/permission",
                    component: AppRoleAndPermissionPage
                },{
                    name:"password-reset",
                    path:"password-reset",
                    component: PasswordResetPage
                },{
                    name: "admin-reset-anyone",
                    path: "user/:id/password",
                    component: AdminResetAnyonePage
                }
            ]
        }
    ]
})

const GetData = function(url, options){
    url = "/api" + url
    return fetch(url, options).then((resp)=>{
        switch(resp.status){
            case 200:
                break
            case 401:
                router.push("/login")
                return
            case 403:
                resp.text().then((body)=>{
                    alertify.error(body||"此操作无权限")
                })
                return
            case 406:
                resp.text().then((body)=>{
                    alertify.error(body||"提交数据错误")
                })
                throw new Error()
            default:
                resp.text().then((body)=>{
                    alertify.error(body||resp.statusText)
                })
                return
        }
        let contentType  = resp.headers.get("Content-Type")
        contentType = contentType && contentType.split(";").shift()
        switch(contentType){
            case "application/json":
                return resp.json()
            default:
                return resp.text()
        }
        
    })
}

new Vue({
    router
}).$mount("#app")