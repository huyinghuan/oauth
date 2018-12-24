
;(function(){
    let template = `
        <a class="button is-success" @click="goBack()">返回上级</a>
    `
    Vue.component("go-back", {
        template: template,
        methods: {
            goBack: function(){
                window.history.length > 1
                ? this.$router.go(-1)
                : this.$router.push('/')
            }
        }
    })
})()