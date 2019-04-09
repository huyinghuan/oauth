var Home = (function(){
    var template = `
        <div class="container is-fluid">
            <o-nav></o-nav>
            <div class="columns" style="margin-top: 0px">
                <div class="column is-full" id="content-body">
                    <router-view></router-view>
                </div>
            </div>
        </div>       
    `
    return {
        template: template,
    }
})()