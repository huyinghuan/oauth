var AdminResetAnyonePage = (function(){
    let template = `
<div class="columns">
    <div class="column"></div>
    <div class="column is-one-three">
        <h2 style="text-align: center;"  class="title is-3">密码重置</h2>
        <div class="field is-horizontal">
            <div class="field-label is-normal">
                <label class="label">账户:</label>
            </div>
            <div class="field-body">
                <div class="field">
                <p class="control is-expanded">
                    <input class="input" type="text" v-model="u.username" readonly>
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
                    <input class="input" type="password" v-model="u.password">
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
                    <button class="button is-primary" v-on:click="save">重置</button>
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
                u:{
                    username:"",
                    password:""
                }
            }
        },
        methods: {
            save: function(){
                GetData(`/user/password/${this.$route.params.id}`, {
                    method: 'PUT',
                    headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({password: this.u.password})
                }).then((resp)=>{
                    alertify.success("修改成功!")
                });
            }
        },
        beforeCreate() {
            GetData(`/user/info/${this.$route.params.id}`, {method:"GET"}).then((data)=>{
                this.u.username = data.name
            })
        },
    }
})()