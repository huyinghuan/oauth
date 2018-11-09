var Login = (function(){
    let template = `
<div class="columns login-body">
    <div class="column"></div>
    <div class="column is-one-fifth">
        <h2 style="text-align: center;"  class="title is-3">Open Auth User Register</h2>
        <div class="field is-horizontal">
            <div class="field-label is-normal">
                <label class="label">账户:</label>
            </div>
            <div class="field-body">
                <div class="field">
                <p class="control is-expanded">
                    <input class="input" type="text" id="username">
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
                    <input class="input" type="password" id="password">
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
                    <button class="button is-primary"  id="register">
                            注册
                    </button>
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
                    body: JSON.stringify({username: username, password: password})
                }).then((resp)=>{
                    if(resp.status == 200){
                        confirm("注册成功!")
                        window.location.href = "/"
                    }else{
                        confirm("账户重复或密码不符合")
                    }
                });
            }
        }
    }
})()