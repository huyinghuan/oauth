const router = new VueRouter({
    routes:[
        {
            path: "/login", component: Login,
            path: "/", component: Home
        }
    ]
})

new Vue({
    router
}).$mount("#app")