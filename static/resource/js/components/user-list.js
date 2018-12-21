(function(){
    let template = `
    <table class="table is-striped is-hoverable  is-fullwidth">
        <thead>
            <tr>
                <th>用户名</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="u in userList">
                <td>{{u.name}}</td>
                <td>
                    <div class="buttons">
                        <button class="button is-small">删除</button>
                        <router-link class="button is-small" :to="{name: 'admin-reset-anyone', params: {id: u.id}}" >修改密码</router-link>
                    </div>
                </td>
            </tr>
        </tbody>
    </table>
    `
    Vue.component("user-list", {
        template: template,
        data: ()=>{
            return {
                userList: []
            }
        },
        methods:{
            deleteUser: function(id, name){
                if(!confirm("是否删除用户:"+name)){
                    return
                }
                GetData("/user/"+id, {
                    method:"DELETE"
                }).then((resp)=>{
                    if(resp.status ==200){
                        location.reload()
                        return
                    }
                    resp.text().then((t)=>{
                        confirm(resp.status + ":"+t)
                    })
                })
            },
            
        },
        loadData: function(){
            GetData("/user/list", {method:"GET"}).then((u)=>{
                this.userList = u || []
            })
        },
        beforeCreate: function(){
            GetData("/user/list", {method:"GET"}).then((u)=>{
                this.userList = u || []
            })
        },
    })
})()