const router = new VueRouter({
    routes:[
        {
            path: "/login", component: Login, name:"login"
        },{
            name: "home",
            path: "/",
            component: Home,
            children:[
                {
                    name: "center",
                    path: "home",
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
                router.push("login")
                return
            default:
                resp.text().then((body)=>{
                    alertify.error(body||resp.statusText)
                })
                return
        }
        let contentType  = resp.headers.get("Content-Type").split(";").shift()
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