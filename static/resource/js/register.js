var Register = (function(){
    let template = `
<div class="columns login-body">
    <div class="column"></div>
    <div class="column is-one-three">
        <h2 style="text-align: center;"  class="title is-3">Open Auth User Register</h2>
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
            <div class="field-label">
                <!-- Left empty for spacing -->
            </div>
            <div class="field-body">
                <div class="field">
                <div class="control">
                    <button class="button is-primary"   v-on:click="register">注册</button>
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
            register: function(){
                GetData('/user/register', {
                    method: 'POST',
                    headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({username: this.username, password: this.password})
                }).then((resp)=>{
                    alertify.success("注册成功!")
                });
            }
        }
    }
})()