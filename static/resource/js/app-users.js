var AppUsersPage = (function(){
    var template = `
    <div class="columns">
        <div class="column is-full">
            <h3 class="title is-3"> 当前应用: {{appName}}</h3>
            <div class="buttons">
                <go-back></go-back>
                <button class="button is-info" @click="runMode('black')">黑名单模式运行</button>
                <button class="button is-warning" @click="runMode('white')">白名单模式运行</button>
                <router-link class="button is-success" :to="{name: 'app-roles', params: {id: $route.params.id}}" >角色与权限</router-link>
                <span style="font-size:8px">黑名单模式: 禁止名单内用户访问应用；白名单模式: 仅允许名单用户访问应用</span>
            </div>
            <hr>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">用户:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" v-model="newUser.username">
                        </p>
                    </div>
                    <div class="field">
                        <div class="select">
                            <select v-model="newUser.category">
                                <option value="white">白名单</option>
                                <option value="black">黑名单</option>
                            </select>
                        </div>
                        <div  class="select">
                            <select v-model.number="newUser.role_id">
                                <option value=0 >默认用户(全部权限)</option>
                                <option v-for="role in roleList" :value="role.id">{{role.name}}</option>
                            </select>
                        </div>
                    </div>
                    <div class="field">
                        <button class="button is-success" @click="addUser2CategoryList()">添加</button>
                    </div>
                </div>
            </div>
            <hr>
            <h3 class="title is-3">白名单列表:</h3>
            <table class="table is-striped is-hoverable  is-fullwidth">
                <thead>
                    <tr>
                        <th>用户名</th>
                        <th>角色</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in whiteList">
                        <td>{{item.user.name}}</td>
                        <td>{{item.appUser.role_id}}</td>
                        <td>
                            <div class="buttons">
                                <button class="button is-small" @click="deleteFromUserList(item.appUser.id, item.user.name)">删除</button>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
            <hr>
            <h3 class="title is-3">黑名单列表:</h3>
            <table class="table is-striped is-hoverable  is-fullwidth">
                <thead>
                    <tr>
                        <th>用户名</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                <tr v-for="item in blackList">
                    <td>{{item.user.name}}</td>
                    <td>
                        <div class="buttons">
                            <button class="button is-small" @click="deleteFromUserList(item.appUser.id, item.user.name)">删除</button>
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
                roleList:[],
                whiteList: [],
                blackList: [],
                newUser:{
                    role_id: 0,
                    category:"white",
                    username: ""
                }
            }
        },
        methods: {
            loadList(){
                GetData(`/app/${this.$route.params.id}/user`).then((data)=>{
                    data = data || []
                    let whiteList = []
                    let blackList = []
                    data.forEach((item)=>{
                        if(item.appUser.category == "white"){
                            whiteList.push(item)
                        }else{
                            blackList.push(item)
                        }
                    })
                    this.whiteList = whiteList
                    this.blackList = blackList
                })
            },
            addUser2CategoryList(){
                GetData(`/app/${this.$route.params.id}/user`,{
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(this.newUser)
                }).then(()=>{
                    alertify.success("添加成功")
                    this.newUser = {
                        role_id: 0,
                        category:"white",
                        username: ""
                    }
                    this.loadList()
                });
            },
            deleteFromUserList(id, name){
                alertify.confirm('是否从名单中删除:', name +"的访问权限", ()=>{
                    GetData(`/app/${this.$route.params.id}/user/${id}`,{ method: "DELETE" }).then(()=>{
                        alertify.success("删除成功")
                        this.loadList()
                    })
                }, ()=>{});
            },
            runMode(mode){
                GetData(`/app/${this.$route.params.id}/user_mode/${mode}`,{
                    method:"PATCH"
                }).then((resp)=>{
                    alertify.success("保存成功")
                })
            }
        },
        created() {
            this.loadList()
        },
        beforeCreate() {
            GetData(`/app/${this.$route.params.id}`, {method:"GET"}).then((data)=>{
               this.appName = data.name
            })
            GetData(`/app/${this.$route.params.id}/role`, {method: "GET"}).then((data)=>{
                this.roleList = data || []
            })
        },
    }
})()