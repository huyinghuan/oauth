var Home = (function(){
    var template = `
    <div class="columns">
        <div class="column is-10 is-offset-1">
            <o-nav></o-nav>
            <router-view></router-view>
        </div>
    </div>
       
    `
    return {
        template: template,
    }
})()