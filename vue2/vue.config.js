module.exports = {
    outputDir: "../assets/vue",    //指定输出目录
    publicPath: "./",               //指定相对于"输出目录"的子目录
    assetsDir: "static",            //指定最终静态文件的存放目录，放置生成的静态资源 (js、css、img、fonts) 的 (相对于 outputDir 的) 目录
    devServer: {
        open: true,
        port: 7070,
        // history模式下的url会请求到服务器端，但是服务器端并没有这一个资源文件，就会返回404，所以需要配置这一项
        // historyApiFallback: {
        //     index: '/index.html' //与output的publicPath
        // },
    },
}
