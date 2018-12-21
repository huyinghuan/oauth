(function(){
    let template = `
        <nav class="navbar is-info" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
        <a class="navbar-item" href="https://bulma.io">
            <!-- <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">-->
        </a>
    
        <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
        </a>
        </div>
    
        <div class="navbar-menu" style="margin-right:0px">
        <div class="navbar-start">
            <router-link class="navbar-item" to="/">Home</router-link>
            <router-link class="navbar-item" to="/home/app-register">注册应用</router-link>
            <router-link class="navbar-item" to="/register">注册用户</router-link>
        </div>
        <div class="navbar-end">
            <div class="navbar-item has-dropdown" :class="{'is-active':isDropdownUserinfo}" v-on:click="isDropdownUserinfo=!isDropdownUserinfo" >
                <a class="navbar-link">{{username}}</a>
                <div class="navbar-dropdown is-right">
                    <router-link class="navbar-item" to="/home/password-reset">修改密码</router-link>
                    <hr class="navbar-divider">
                    <a class="navbar-item" v-on:click="logout">注销</a>
                </div>
            </div>
        </div>
        </div>
        </nav>
    `
    Vue.component('o-nav', {
        template: template,
        data:function(){
            return {
                username:"未登陆",
                isDropdownUserinfo: false
            }
        },
        beforeCreate: function(){
            GetData("/user", {method:"GET"}).then((u)=>{
                this.username = u && u.username
            })
        },
        methods:{
            logout: function(){
                GetData("/user/logout",{
                    method: "DELETE"
                }).then(()=>{
                    location.reload()
                })
            }
        }
    })
})()