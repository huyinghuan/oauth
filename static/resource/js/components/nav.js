(function(){
    let template = `
        <nav class="navbar is-info" role="navigation" aria-label="main navigation">
        <!--<div class="navbar-brand" style="margin-left: 0px">
            <a class="navbar-item" href="https://github.com/huyinghuan">
                <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
            </a>
        </div>
        -->
        <div class="navbar-menu" style="margin-right:0px">
        <div class="navbar-start">
            <router-link class="navbar-item" to="/">OAUTH</router-link>
            <router-link class="navbar-item" :to="{name:'open-apps'}" active-class="is-active">应用入口</router-link>
            <router-link class="navbar-item" :to="{name:'apps'}" active-class="is-active">我的应用</router-link>
            <router-link class="navbar-item" :to="{name:'users'}"  v-if="isAdmin" active-class="is-active">用户列表</router-link>
        </div>
        <div class="navbar-end">
            <router-link class="navbar-item" to="/register">注册用户</router-link>
            <router-link class="navbar-item" to="/home/app-register">注册应用</router-link>
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
                isAdmin: false,
                username:"未登陆",
                isDropdownUserinfo: false
            }
        },
        beforeCreate: function(){
            GetData("/user", {method:"GET"}).then((u)=>{
                this.username = u && u.username
                if(u.uid == 0){
                    this.isAdmin = true
                }
            })
        },
        methods:{
            logout: function(){
                GetData("/user/logout",{
                    method: "DELETE"
                }).then(()=>{
                    router.push({name:"login"})
                })
            }
        }
    })
})()