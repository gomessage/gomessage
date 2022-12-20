module.exports = {
    publicPath: "./",               //指定相对子路径
    assetsDir: "static",            //放置生成的静态资源 (js、css、img、fonts) 的 (相对于 outputDir 的) 目录
    outputDir: "../assets/vue2",    //当运行 vue-cli-service build 时生成的生产环境构建文件的目录。注意目标目录的内容在构建之前会被清除 (构建时传入 --no-clean 可关闭该行为)
    devServer: {
        open: true,
        port: 7070,
        // history模式下的url会请求到服务器端，但是服务器端并没有这一个资源文件，就会返回404，所以需要配置这一项
        // historyApiFallback: {
        //     index: '/index.html' //与output的publicPath
        // },
    },
}
