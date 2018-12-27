(function(){
    let template = `
    <table class="table is-striped is-hoverable  is-fullwidth">
        <thead>
            <tr>
                <th>所属</th>
                <th>应用名称</th>
                <th>ClientID</th>
                <th>PrivateKey</th>
                <th>回调地址</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
                <tr v-for="app in appList">
                    <td>{{app.user.name}}</td>
                    <td>{{app.application.name}}</td>
                    <td>{{app.application.client_id}}</td>
                    <td>{{app.application.private_key}}</td>
                    <td>{{app.application.callback}}</td>
                    <td>
                        <div class="buttons">
                            <button class="button is-small is-danger" @click="deleteApp(app.application.id, app.application.name)">删除</button>
                            <router-link class="button is-small" :to="{name: 'app-edit', params: {id: app.application.id}}" >编辑</router-link>
                            <router-link class="button is-small" :to="{name: 'app-users', params: {id: app.application.id}}">用户管理</router-link>
                        </div>
                    </td>
                </tr>
        </tbody>
    </table>
    `
    Vue.component("app-list", {
        template: template,
        data: ()=>{
            return {
                appList: []
            }
        },
        methods:{
            deleteApp: function(id, name){                
                alertify.confirm('是否删除应用:', name, ()=>{
                    GetData(`/app/${id}`,{ method: "DELETE" }).then(()=>{
                        alertify.success("删除成功")
                        this.loadData()
                    })
                }, ()=>{});
            },
            loadData: function(){
                GetData("/app", {method:"GET"}).then((u)=>{
                    this.appList = u || []
                })
            }
        },
        created() {
            this.loadData()
        }
    })
})()