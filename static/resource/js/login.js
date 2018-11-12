var Login = (function(){
    let template = `
<div class="columns login-body">
    <div class="column"></div>
    <div class="column is-one-fifth">
            <h2 style="text-align: center;"  class="title is-3">Open Auth</h2>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">账户:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                    <p class="control is-expanded">
                        <input class="input" type="text" v-model="username">
                    </p>
                    </div>
                </div>
            </div>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">密码:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                    <div class="control">
                        <input class="input" type="password" v-model="password">
                    </div>
                    </div>
                </div>
            </div>
            <div class="field is-horizontal">
                <div class="field-label"></div>
                <div class="field-body">
                    <div class="field">
                        <div class="control">
                            <button class="button is-primary" v-on:click="login">登陆</button>
                        </div>
                    </div>
                    <div class="field">
                        <div class="control">
                            <router-link to="/register">用户注册</router-link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    <div class="column"></div>     
</div>
    `

    return {
        template: template,
        data: ()=>{
            return {
                username:"",
                password:""
            }
        },
        methods: {
            login: function(){
                GetData('/user/login', {
                    method: 'POST',
                    headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({username: this.username, password: this.password})
                }).then((resp)=>{
                    router.push({name:"center"})
                })
            }
        }
    }
})()