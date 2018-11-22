var AppRegister = (function(){
    let template = `
    <div class="columns">
            <div class="column"></div>
            <div class="column is-one-three">
                    <h2 style="text-align: center;"  class="title is-3">Open Auth App Register</h2>
                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">应用名称:</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                            <p class="control is-expanded">
                                <input class="input" type="text" v-model="name">
                            </p>
                            </div>
                        </div>
                    </div>
                    <div class="field is-horizontal">
                            <div class="field-label is-normal">
                                <label class="label">回调地址:</label>
                            </div>
                            <div class="field-body">
                                <div class="field">
                                <p class="control is-expanded">
                                    <input class="input" type="text" v-model="callback">
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
                                <button class="button is-primary"  v-on:click="register">
                                        注册
                                </button>
                                <a href="https://github.com/huyinghuan/oauth/blob/master/README.md" target="_blank">如何使用?</a>
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
                name:"",
                callback:""
            }
        },
        methods: {
            register: function(){
                GetData('/app/register', {
                    method: 'POST',
                    headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({name: this.name, callback: this.callback})
                }).then((resp)=>{
                    alertify.success("注册成功!")
                    router.push("/")
                }).catch(()=>{});
            }
        }
    }

})()