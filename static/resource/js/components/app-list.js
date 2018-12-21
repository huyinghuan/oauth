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
                            <button class="button is-small">删除</button>
                            <router-link class="button is-small" :to="{name: 'app-edit', params: {id: app.application.id}}" >编辑</router-link>
                            <a class="button is-small">用户管理</a>
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
            deleteApp: (id, name)=>{
                if(!confirm("是否删除应用:"+name)){
                    return
                }
                GetData("/api/app/"+id,{
                    method: "DELETE"
                }).then(()=>{
                    this.loadData()
                })
            },
            loadData: function(){
                GetData("/app/", {method:"GET"}).then((u)=>{
                    this.appList = u || []
                })
            }
        },
        beforeCreate: function(){
            GetData("/app/", {method:"GET"}).then((u)=>{
                this.appList = u || []
            })
        },
    })
})()