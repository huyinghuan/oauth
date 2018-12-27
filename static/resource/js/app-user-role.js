var AppUserRolePage = (function(){
    let template = `
    <div class="columns">
            <div class="column"></div>
            <div class="column is-one-three">

                    <h3 style="text-align: center;"  class="title is-3">角色修改</h3>

                    <div class="buttons">
                        <go-back></go-back>
                        <router-link class="button is-info" :to="{name: 'app-users', params: {id: $route.params.id}}" >应用: {{appName}}</router-link>
                        <button class="button is-warning"> 用户名: {{username}}</button>
                    </div>
                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">选择:</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                                <div class="select">
                                    <select v-model.number="newRole">
                                        <option value=0 >默认用户(全部权限)</option>
                                        <option v-for="role in roleList" :value="role.id">{{role.name}}</option>
                                    </select>
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
                                <button class="button is-info"  v-on:click="save">保存</button>
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
                appName: "",
                username: "",
                newRole: 0,
                roleList:[]
            }
        },
        methods: {
            save: function(){
                GetData(`/app/${this.$route.params.id}/user/${this.$route.params.uid}/role`,{
                    method: 'PUT',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({roleID: this.newRole})
                }).then(()=>{
                    alertify.success("修改成功")
                });
            }
        },
        beforeCreate() {
            GetData(`/app/${this.$route.params.id}`, {method:"GET"}).then((data)=>{
               this.appName = data.name
            })
            GetData("/user", {method:"GET"}).then((u)=>{
                this.username = u && u.username
            })
            GetData(`/app/${this.$route.params.id}/role`, {method: "GET"}).then((data)=>{
                this.roleList = data || []
                return
            })
        }
    }

})()