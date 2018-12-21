var PasswordResetPage = (function(){
    let template = `
    <div class="columns">
            <div class="column"></div>
            <div class="column is-one-three">
                    <h2 style="text-align: center;"  class="title is-3">密码重置</h2>
                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">Old:</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                            <p class="control is-expanded">
                                <input class="input" type="password" v-model="oldPassword">
                            </p>
                            </div>
                        </div>
                    </div>
                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">New:</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                            <p class="control is-expanded">
                                <input class="input" type="password" v-model="newPassword">
                            </p>
                            </div>
                        </div>
                    </div>
                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">Repeat:</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                            <p class="control is-expanded">
                                <input class="input" type="password" v-model="newPassword2">
                            </p>
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
                                <button class="button is-primary"  v-on:click="save">
                                        保存
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
                oldPassword:"",
                newPassword:"",
                newPassword2:"",
            }
        },
        methods: {
            save: function(){
                if(this.newPassword != this.newPassword2){
                    alertify.error("两次密码不一致")
                }
                GetData('/user/password', {
                    method: 'PUT',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({oldPassword: this.oldPassword, newPassword: this.newPassword})
                }).then(()=>{
                    alertify.alert("密码修改成功!" , "请重新登陆", ()=>{
                        GetData("/user/logout",{
                            method: "DELETE"
                        }).then(()=>{
                            location.reload()
                        })
                    })
                   
                }).catch(()=>{});
            }
        }
    }

})()