var AppRoleAndPermissionPage = (function(){

    let template = `
    <div class="columns">
        <div class="column is-full">
            <h3  class="title is-3"> 当前应用:</h3>
            <div class="buttons">
                <a class="button is-white" href="/">首页</a>
                <a class="button is-success">角色与权限</a>
            </div>
            <hr>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">权限:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" id="pattern" placeholder="url正则">
                        </p>
                    </div>
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" id="method" placeholder="HTTP Method">
                        </p>
                    </div>
                    <div class="field">
                        <p class="control">
                            <input class="input" type="text" id="name" placeholder="备注">
                        </p>
                    </div>
                    <div class="field">
                        <button class="button is-success" onclick="addRole()">添加</button>
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
                                <button class="button is-small" onclick="delPermission({{p.id}}, {{p.name}})">删除</button>
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
    }


})();