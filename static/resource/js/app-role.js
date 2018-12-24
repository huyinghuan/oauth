var AppRolesPage = (function(){
    let template = `
        <div class="columns">
            <div class="column is-full">
                <h3 class="title is-3"> 当前应用:{{appName}}</h3>
                <div class="buttons">
                   <go-back></go-back>
                </div>
                <hr>
                <div class="field is-horizontal special-form">
                    <div class="field-label is-normal">
                        <label class="label">角色:</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <p class="control">
                                <input class="input" type="text" v-model="newRole">
                            </p>
                        </div>
                        <div class="field">
                            <button class="button is-success" @click="addRole()">添加</button>
                        </div>
                    </div>
                </div>
                <hr>
                <h3 class="title is-3">角色列表:</h3>
                <table class="table is-striped is-hoverable  is-fullwidth">
                    <thead>
                        <tr>
                            <th>角色</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                            <tr v-for="role in roleList">
                                <td>{{role.name}}</td>
                                <td>
                                    <div class="buttons">
                                        <button class="button is-small" @click="delRole(role.id, role.name)">删除</button>
                                        <a class="button is-small is-info">权限分配</a>
                                    </div>
                                </td>
                            </tr>
                    </tbody>
                </table>
            </div>
        </div>
    `
    return {
        template: template,
        data: function(){
            return {
                appName: "",
                roleList: [],
                newRole: ""
            }
        },
        methods: {
            loadList: function(){
                GetData(`/app/${this.$route.params.id}/role`).then((data)=>{
                    this.roleList = data
                })
            },
            addRole: function(){
                GetData(`/app/${this.$route.params.id}/role`,{
                    method: 'POST',
                    headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({role: this.newRole})
                }).then(()=>{
                    alertify.success("添加成功!")
                    this.loadList()
                });
            },
            delRole: function(id, name){
                alertify.confirm('是否删除角色:', name, ()=>{
                    GetData(`/app/${this.$route.params.id}/role/${id}`,{ method: "DELETE" }).then(()=>{
                        alertify.success("删除成功")
                        this.loadList()
                    })
                }, ()=>{});
            }
        },
        created() {
            this.loadList()
        },
        beforeCreate() {
            GetData(`/app/${this.$route.params.id}`, {method:"GET"}).then((data)=>{
               this.appName = data.name
            })
        },
    }
})()