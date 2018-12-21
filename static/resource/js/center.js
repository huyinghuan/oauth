var Center = (function(){
    let template  = `
    <div class="columns" style="flex-direction: column">
        <div class="column is-full">
            <h3 class="title is-3">第三方应用列表:</h3>
            <app-list></app-list>
        </div>
        <div class="column is-full" v-if="isAdmin">
            <h3 class="title is-3">用户列表:</h3>
            <user-list></user-list>
        </div>
    </div>
    `
    return {
        template,
        data: function(){
            return {
                isAdmin: false
            }
        },
        beforeCreate() {
            GetData("/user", {method:"GET"}).then((u)=>{
               if(u.uid == 0){
                   this.isAdmin = true
               }
            })
        },
    }
})()