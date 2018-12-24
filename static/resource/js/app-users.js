var AppUsersPage = (function(){
    var template = `
    <div class="columns">
        <div class="column is-full">
            <h3 class="title is-3"> 当前应用: {{appName}}</h3>
            <div class="buttons">
                <button class="button is-info" onclick="runMode('black')">黑名单模式运行</button>
                <button class="button is-warning" onclick="runMode('white')">白名单模式运行</button>
                <router-link class="button is-success" :to="{name: 'app-roles', params: {id: $route.params.id}}" >角色与权限</router-link>
            </div>
            <hr>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">用户:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" id="username">
                        </p>
                    </div>
                    <div class="field">
                        <div class="select">
                            <select id="category">
                                <option value="white">白名单</option>
                                <option value="black">黑名单</option>
                            </select>
                        </div>
                        <div  class="select">
                            <select id="roleID">
                                <option value=0 >默认用户(全部权限)</option>
                                <option v-for="role in roleList">{{role.name}}</option>
                            </select>
                        </div>
                    </div>
                    <div class="field">
                        <button class="button is-success" onclick="addUser2CategoryList()">添加</button>
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
                <tbody id="WhiteList">
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
                <tbody id="BlackList">
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
                roleList: []
            }
        },
        methods: {
            loadList: function(){
                GetData(`/app/${this.$route.params.id}/user`).then((data)=>{
                    console.log(data)
                    data = data || []
                    
                })
            },
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