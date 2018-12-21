var AppEditPage = (function(){
    let template = `
    <div class="columns">
            <div class="column"></div>
            <div class="column is-one-three">
                    <h2 style="text-align: center;"  class="title is-3">App 修改</h2>
                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">应用名称:</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                            <p class="control is-expanded">
                                <input class="input" type="text" v-model="app.name">
                            </p>
                            </div>
                        </div>
                    </div>
                    <div class="field is-horizontal">
                            <div class="field-label is-normal">
                                <label class="label">回调地址:</label>
                            </div>
                            <div class="field-body">
                                <div class="field">
                                <p class="control is-expanded">
                                    <input class="input" type="text" v-model="app.callback">
                                </p>
                                </div>
                            </div>
                        </div>


                    <div class="field is-horizontal">
                        <div class="field-label">
                            <!-- Left empty for spacing -->
                        </div>
                        <div class="field-body">
                            <div class="field">
                            <div class="control">
                                <button class="button is-primary"  v-on:click="save">
                                        保存
                                </button>
                            </div>
                            </div>
                        </div>
                    </div>
                </div>
            <div class="column"></div>
            </div>
    `

    return {
        template: template,
        data: ()=>{
            return {
                app:{
                    name:"",
                    callback:""
                }
            }
        },
        beforeCreate() {
            GetData(`/app/${this.$route.params.id}`).then((app)=>{
                this.app = app
            })
        },
        methods: {
            save: function(){
                GetData(`/app/${this.$route.params.id}`, {
                    method: 'PUT',
                    headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(this.app)
                }).then(()=>{
                    alertify.success("报存成功!")
                })
            }
        }
    }

})()