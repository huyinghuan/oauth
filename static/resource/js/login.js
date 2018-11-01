var Login = (function(){
    let template = `
<div class="columns">
    <div class="column"></div>
    <div class="column is-one-fifth">
            <h2 style="text-align: center;"  class="title is-3">Open Auth</h2>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">账户:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                    <p class="control is-expanded">
                        <input class="input" type="text" id="username">
                    </p>
                    </div>
                </div>
            </div>
            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">密码:</label>
                </div>
                <div class="field-body">
                    <div class="field">
                    <div class="control">
                        <input class="input" type="password" id="password">
                    </div>
                    </div>
                </div>
            </div>
            <div class="field is-horizontal">
                <div class="field-label"></div>
                <div class="field-body">
                    <div class="field">
                    <div class="control">
                        <button class="button is-primary" id="login">登陆</button>
                    </div>
                    </div>
                </div>
            </div>
        </div>
    <div class="column"></div>     
</div>
    `

    return {
        template: template
    }
})()