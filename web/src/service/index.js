import {notification} from 'antd'

export const get = function(url, options){
    url = "/api" + url
    return fetch(url, options).then((resp)=>{
        switch(resp.status){
            case 200:
                break
            case 401:
                resp.text().then((body)=>{
                    notification.warning({
                        message: '401',
                        description:body||"用户未登录",
                        placement: 'bottomRight',
                        duration: 3,
                    });
                })
                throw new Error("错误:"+resp.statusText)
            case 403:
                resp.text().then((body)=>{
                    notification.warning({
                        message: '403',
                        description:body||"此操作无权限",
                        placement: 'bottomRight',
                        duration: 3,
                    });
                })
                throw new Error("错误:"+resp.statusText)
            case 406:
                resp.text().then((body)=>{
                    notification.error({
                        message: '406',
                        description:body||"提交数据错误",
                        placement: 'bottomRight',
                        duration: 3,
                    });
                })
                throw new Error("错误:"+resp.statusText)
            default:
                resp.text().then((body)=>{
                    notification.error({
                        message: resp.status,
                        description:body|| ("未知错误:"+resp.statusText),
                        placement: 'bottomRight',
                        duration: 3,
                    });
                })
                throw new Error("错误:"+resp.statusText)
        }
        let contentType  = resp.headers.get("Content-Type")
        contentType = contentType && contentType.split(";").shift()
        switch(contentType){
            case "application/json":
                return resp.json()
            default:
                return resp.text()
        }
    })
}