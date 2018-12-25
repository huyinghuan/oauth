var AppRoleAndPermissionPage = (function(){

    let template = `
    <div class="columns">
        <div class="column is-full">
            <h3  class="title is-3"> 当前应用: {{appName}}  角色: {{roleName}}</h3>
            <div class="buttons">
               <go-back></go-back>
            </div>
            <hr>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">权限:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" v-model="rule.pattern" placeholder="url正则 or 字符串">
                        </p>
                    </div>
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" v-model="rule.method" placeholder="HTTP Method 多个用逗号 ',' 隔开 ">
                        </p>
                    </div>
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" v-model="rule.name" placeholder="备注">
                        </p>
                    </div>
                    <div class="field">
                        <button v-if="isAdd" class="button is-success" @click="addRule()">添加</button>
                        <button v-if="isEdit" class="button is-success" @click="saveRule()">保存</button>
                    </div>
                </div>
            </div>
            <hr>
            <h3 class="title is-3">权限列表:</h3>
            <table class="table is-striped is-hoverable  is-fullwidth">
                <thead>
                    <tr>
                        <th>权限</th>
                        <th>HTTP Method</th>
                        <th>备注</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="p in permissionList">
                        <td>{{p.pattern}}</td>
                        <td>{{p.method}}</td>
                        <td>{{p.name}}</td>
                        <td>
                            <div class="buttons">
                                <button class="button is-small" @click="delPermission(p.id,p.name)">删除</button>
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
                appName:"",
                roleName:"",
                rule:{
                    name: "",
                    pattern:"",
                    method:"",
                },
                isAdd: true,
                isEdit: false,
                permissionList:[]
            }
        },
        methods: {
            loadList(){
                GetData(`/app/${this.$route.params.id}/role/${this.$route.params.roleID}/permission`,{method: 'GET'}).then((data)=>{
                    this.permissionList = data
                });
            },
            saveRule(){
                GetData(`/app/${this.$route.params.id}/role/${this.$route.params.roleID}/permission/${this.rule.id}`,{
                    method: 'PUT',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(this.rule)
                }).then((resp)=>{
                    if(resp.status == 200){
                        confirm("添加成功!")
                        location.reload()
                    }else{
                        confirm(resp.statusText)
                    }
                });
            },
            addRule(){
                GetData(`/app/${this.$route.params.id}/role/${this.$route.params.roleID}/permission`,{
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(this.rule)
                }).then(()=>{
                    alertify.success("添加成功")
                    this.loadList()
                });
            },
            delPermission(id, name){
                alertify.confirm('是否删除权限:', name, ()=>{
                    GetData(`/app/${this.$route.params.id}/role/${this.$route.params.roleID}/permission/${id}`,{ method: "DELETE" }).then(()=>{
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
            GetData(`/app/${this.$route.params.id}/role/${this.$route.params.roleID}`, {method:"GET"}).then((data)=>{
               this.roleName = data.name
            })
        },
    }


})();