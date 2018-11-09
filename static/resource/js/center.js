var Center = (function(){
    let template  = `
    <div class="columns">
        <div class="column">
            <h3 class="title is-3">第三方应用列表:</h3>
            <app-list></app-list>
            <h3 class="title is-3">用户列表:</h3>
            <user-list></user-list>
        </div>
    </div>
    `
    return {
        template
    }
})()