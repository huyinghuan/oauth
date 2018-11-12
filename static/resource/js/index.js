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
                alertify.error("此操作无权限")
                return
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